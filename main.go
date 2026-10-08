package main

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//go:embed templates/*.html static
var assets embed.FS

type Project struct {
	Slug, Title, Category, FuncName, Params, Returns, Blurb string
	Bullets, Stack                                          []string
	RepoURL, SourceLabel, DemoURL                           string
	Featured                                                bool
}
type SkillGroup struct {
	Category string
	Skills   []string
}
type EduEntry struct{ Period, Title, Org, Detail string }
type Contact struct{ Email, GitHub, Phone, PhoneURL, Based string }
type PageData struct {
	Name, RoleTag, Headline, HeadlineAccent, SubHeadline, AvatarURL, CVPath, Bio string
	Title, Description, Page, SiteURL, CanonicalURL                              string
	Year, Status                                                                 int
	SkillGroups                                                                  []SkillGroup
	Projects, FeaturedProjects                                                   []Project
	Education                                                                    []EduEntry
	Contact                                                                      Contact
	Project                                                                      Project
}

func newHandler(siteURL string) (http.Handler, error) {
	if siteURL != "" {
		u, err := url.Parse(siteURL)
		if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") {
			return nil, fmt.Errorf("SITE_URL must be an absolute http(s) origin, such as https://example.com")
		}
		siteURL = strings.TrimRight(siteURL, "/")
	}
	tmpl, err := template.ParseFS(assets, "templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}
	static, err := fs.Sub(assets, "static")
	if err != nil {
		return nil, err
	}
	base := portfolioData()
	base.Year = time.Now().Year()
	base.SiteURL = siteURL
	for _, p := range base.Projects {
		if p.Featured {
			base.FeaturedProjects = append(base.FeaturedProjects, p)
		}
	}
	render := func(w http.ResponseWriter, r *http.Request, name string, data PageData, status int) {
		if siteURL != "" && status == http.StatusOK {
			data.CanonicalURL = siteURL + r.URL.EscapedPath()
		}
		var body bytes.Buffer
		if err := tmpl.ExecuteTemplate(&body, name, data); err != nil {
			log.Printf("render %s: %v", name, err)
			http.Error(w, "Unable to load this page.", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(status)
		if r.Method != http.MethodHead {
			_, _ = body.WriteTo(w)
		}
	}
	notFound := func(w http.ResponseWriter, r *http.Request) {
		data := base
		data.Title, data.Description, data.Status = "Page not found", "Find your way back to Michael's portfolio.", http.StatusNotFound
		render(w, r, "error.html", data, http.StatusNotFound)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, "/static/")
		info, err := fs.Stat(static, name)
		if err != nil || info.IsDir() {
			notFound(w, r)
			return
		}
		w.Header().Set("Cache-Control", "public, max-age=3600, must-revalidate")
		http.StripPrefix("/static/", http.FileServer(http.FS(static))).ServeHTTP(w, r)
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		data := base
		data.Title, data.Page, data.Description = data.Name+" — Backend Developer", "home", "Michael Ikese Emmanuel's portfolio: Go HTTP services, CLI tools, and web applications. Based in Nigeria and open to remote opportunities."
		render(w, r, "index.html", data, http.StatusOK)
	})
	mux.HandleFunc("GET /projects", func(w http.ResponseWriter, r *http.Request) {
		data := base
		data.Title, data.Page, data.Description = "Projects — "+data.Name, "projects", "Explore Michael's Go services, command-line tools, and web applications."
		render(w, r, "projects.html", data, http.StatusOK)
	})
	mux.HandleFunc("GET /projects/{slug}", func(w http.ResponseWriter, r *http.Request) {
		for _, p := range base.Projects {
			if p.Slug == r.PathValue("slug") {
				data := base
				data.Project, data.Page, data.Title, data.Description = p, "projects", p.Title+" — "+data.Name, p.Blurb
				render(w, r, "project.html", data, http.StatusOK)
				return
			}
		}
		notFound(w, r)
	})
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if r.Method != http.MethodHead {
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		}
	})
	mux.HandleFunc("GET /robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if r.Method == http.MethodHead {
			return
		}
		_, _ = fmt.Fprint(w, "User-agent: *\nAllow: /\n")
		if siteURL != "" {
			_, _ = fmt.Fprintf(w, "Sitemap: %s/sitemap.xml\n", siteURL)
		}
	})
	mux.HandleFunc("GET /sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		if siteURL == "" {
			notFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		if r.Method == http.MethodHead {
			return
		}
		_, _ = fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
		paths := []string{"/", "/projects"}
		for _, p := range base.Projects {
			paths = append(paths, "/projects/"+p.Slug)
		}
		for _, path := range paths {
			_, _ = fmt.Fprintf(w, "<url><loc>%s</loc></url>", template.HTMLEscapeString(siteURL+path))
		}
		_, _ = fmt.Fprint(w, "</urlset>")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			w.Header().Set("Allow", "GET, HEAD")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		notFound(w, r)
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self'; font-src 'self'; object-src 'none'; base-uri 'none'; frame-ancestors 'none'; form-action 'self'")
		w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
		mux.ServeHTTP(w, r)
	}), nil
}

func main() {
	handler, err := newHandler(os.Getenv("SITE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{Addr: ":" + port, Handler: handler, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 15 * time.Second, IdleTimeout: 60 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	failures := make(chan error, 1)
	go func() { log.Printf("serving on %s", server.Addr); failures <- server.ListenAndServe() }()
	select {
	case err := <-failures:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("shutdown: %v", err)
			_ = server.Close()
		}
	}
}
