package main

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	handler, err := newHandler("https://portfolio.example")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		path, contentType, contains string
		status                      int
	}{
		{"/", "text/html", "Useful software.", 200},
		{"/projects", "text/html", "project-search", 200},
		{"/projects/portfolio-site", "text/html", "Inside the project", 200},
		{"/projects/not-a-project", "text/html", "off the path.", 404},
		{"/projects/portfolio-site/extra", "text/html", "off the path.", 404},
		{"/missing", "text/html", "off the path.", 404},
		{"/healthz", "application/json", `"status":"ok"`, 200},
		{"/robots.txt", "text/plain", "Sitemap: https://portfolio.example/sitemap.xml", 200},
		{"/static/css/style.css", "text/css", "prefers-reduced-motion", 200},
		{"/static/js/main.js", "javascript", "localStorage", 200},
		{"/static/img/portrait.png", "image/png", "", 200},
		{"/static/cv/resume.pdf", "application/pdf", "%PDF", 200},
		{"/static/", "text/html", "off the path.", 404},
		{"/static/img/", "text/html", "off the path.", 404},
		{"/static/no-such-file", "text/html", "off the path.", 404},
	}
	for _, tc := range tests {
		t.Run(tc.path, func(t *testing.T) {
			r := httptest.NewRecorder()
			handler.ServeHTTP(r, httptest.NewRequest(http.MethodGet, tc.path, nil))
			if r.Code != tc.status {
				t.Fatalf("status %d, want %d", r.Code, tc.status)
			}
			if !strings.Contains(r.Header().Get("Content-Type"), tc.contentType) {
				t.Errorf("content type %q", r.Header().Get("Content-Type"))
			}
			if !strings.Contains(r.Body.String(), tc.contains) {
				t.Errorf("missing expected content %q", tc.contains)
			}
			if r.Header().Get("X-Content-Type-Options") != "nosniff" || !strings.Contains(r.Header().Get("Content-Security-Policy"), "frame-ancestors 'none'") {
				t.Error("missing security headers")
			}
		})
	}
	// Exercise every project to catch bad data and template regressions.
	for _, p := range projects {
		r := httptest.NewRecorder()
		handler.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/projects/"+p.Slug, nil))
		if r.Code != 200 || !strings.Contains(r.Body.String(), template.HTMLEscapeString(p.Title)) {
			t.Errorf("project %s failed to render", p.Slug)
		}
	}
}

func TestMethodAndHead(t *testing.T) {
	handler, err := newHandler("")
	if err != nil {
		t.Fatal(err)
	}
	// Use a real server: net/http itself suppresses HEAD bodies for static files.
	server := httptest.NewServer(handler)
	defer server.Close()
	for _, path := range []string{"/", "/projects", "/projects/global-browser", "/healthz", "/static/cv/resume.pdf", "/missing"} {
		response, err := http.Head(server.URL + path)
		if err != nil {
			t.Fatal(err)
		}
		var body [1]byte
		n, _ := response.Body.Read(body[:])
		response.Body.Close()
		if n != 0 {
			t.Errorf("HEAD %s returned a body", path)
		}
	}
	for _, path := range []string{"/", "/projects", "/healthz", "/static/css/style.css"} {
		r := httptest.NewRecorder()
		handler.ServeHTTP(r, httptest.NewRequest(http.MethodPost, path, nil))
		if r.Code != http.StatusMethodNotAllowed || r.Header().Get("Allow") != "GET, HEAD" {
			t.Errorf("POST %s: status %d, Allow %q", path, r.Code, r.Header().Get("Allow"))
		}
	}
}

func TestDiscoveryAndConfiguration(t *testing.T) {
	for _, origin := range []string{"relative", "ftp://example.com", "https://user:pass@example.com", "https://example.com/path", "https://example.com?x=1", "https://example.com#fragment"} {
		if _, err := newHandler(origin); err == nil {
			t.Errorf("accepted invalid SITE_URL %q", origin)
		}
	}
	handler, err := newHandler("https://portfolio.example/")
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	handler.ServeHTTP(r, httptest.NewRequest("GET", "/sitemap.xml", nil))
	var sitemap struct {
		URLs []struct {
			Location string `xml:"loc"`
		} `xml:"url"`
	}
	if err := xml.Unmarshal(r.Body.Bytes(), &sitemap); err != nil {
		t.Fatal(err)
	}
	if len(sitemap.URLs) != len(projects)+2 {
		t.Errorf("got %d sitemap URLs", len(sitemap.URLs))
	}
	for _, entry := range sitemap.URLs {
		if !strings.HasPrefix(entry.Location, "https://portfolio.example/") {
			t.Errorf("invalid URL %q", entry.Location)
		}
	}
	r = httptest.NewRecorder()
	handler.ServeHTTP(r, httptest.NewRequest("GET", "/healthz", nil))
	var health map[string]string
	if err := json.Unmarshal(r.Body.Bytes(), &health); err != nil || health["status"] != "ok" {
		t.Error("invalid health response")
	}
	noOrigin, _ := newHandler("")
	r = httptest.NewRecorder()
	noOrigin.ServeHTTP(r, httptest.NewRequest("GET", "/sitemap.xml", nil))
	if r.Code != 404 {
		t.Error("sitemap must require a configured origin")
	}
}

// Guard the original collection while adding screenshot-backed project entries.
func TestProjectCollectionIntegrity(t *testing.T) {
	original := []string{"smartconvert", "global-browser", "past-quest-solver", "campus-resource-manager", "talentgrid", "portfolio-site"}
	handler, err := newHandler("")
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRecorder()
	handler.ServeHTTP(r, httptest.NewRequest("GET", "/projects", nil))
	seen := map[string]bool{}
	for _, p := range projects {
		if seen[p.Slug] {
			t.Fatalf("duplicate project slug %q", p.Slug)
		}
		seen[p.Slug] = true
		if !strings.Contains(r.Body.String(), template.HTMLEscapeString(p.Title)) {
			t.Errorf("missing card for %s", p.Slug)
		}
		if p.ScreenshotURL == "" {
			continue
		}
		if p.ScreenshotAlt == "" || p.ScreenshotWidth < 1 || p.ScreenshotHeight < 1 || p.RepoURL == "" || p.Status == "" || len(p.Bullets) == 0 || len(p.Stack) == 0 {
			t.Errorf("incomplete screenshot project %s", p.Slug)
		}
		image := httptest.NewRecorder()
		handler.ServeHTTP(image, httptest.NewRequest("GET", p.ScreenshotURL, nil))
		if image.Code != 200 || image.Header().Get("Content-Type") != "image/png" {
			t.Errorf("unavailable application screenshot %s", p.ScreenshotURL)
		}
		detail := httptest.NewRecorder()
		handler.ServeHTTP(detail, httptest.NewRequest("GET", "/projects/"+p.Slug, nil))
		for _, required := range []string{p.ScreenshotURL, p.RepoURL, "in development", "Actual application screenshot"} {
			if !strings.Contains(detail.Body.String(), required) {
				t.Errorf("%s detail missing %q", p.Slug, required)
			}
		}
	}
	for _, slug := range original {
		if !seen[slug] {
			t.Errorf("original project removed: %s", slug)
		}
	}
}

// Collaboration credit and source visibility must survive template changes.
func TestProjectAttributionAndPrivateSources(t *testing.T) {
	handler, err := newHandler("")
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		slug string
		want []string
	}{
		{"marvellous-motors", []string{"My role", "Development collaborator", "Collaboration", "Otiti Mordecai Marvellous", "project owner", "https://github.com/mordecai12-cool/Marvellous-motor2"}},
		{"campusflow", []string{"AI-assisted project development", "https://github.com/miikese/campusflow", "human collaboration and peer-review evidence remain outstanding"}},
	}
	for _, tc := range cases {
		t.Run(tc.slug, func(t *testing.T) {
			r := httptest.NewRecorder()
			handler.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/projects/"+tc.slug, nil))
			if r.Code != http.StatusOK {
				t.Fatalf("status %d", r.Code)
			}
			for _, want := range tc.want {
				if !strings.Contains(r.Body.String(), want) {
					t.Errorf("missing attribution %q", want)
				}
			}
		})
	}
	for _, project := range projects {
		if project.Visibility != "Private source" {
			continue
		}
		if project.RepoURL != "" {
			t.Errorf("private project %s exposes a repository link", project.Slug)
		}
		r := httptest.NewRecorder()
		handler.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/projects/"+project.Slug, nil))
		if !strings.Contains(r.Body.String(), "Discuss this project") || strings.Contains(r.Body.String(), "GitHub source for") {
			t.Errorf("private project %s does not offer an appropriate contact action", project.Slug)
		}
	}
}
