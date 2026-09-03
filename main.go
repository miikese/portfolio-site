package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed templates/*.html
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

type Project struct {
	Slug      string
	FuncName  string
	Params    string
	Returns   string
	Blurb     string
	Bullets   []string
	Stack     []string
	RepoURL   string
}

type SkillGroup struct {
	Category string
	Color    string
	Skills   []string
}

type EduEntry struct {
	Period string
	Title  string
	Org    string
	Detail string
}

type Meta struct {
	Stack          string
	Based          string
	ProjectsCount  int
}

type PageData struct {
	Name        string
	RoleTag     string // "BACKEND DEVELOPER — GO"
	Headline    string
	SubHeadline string
	Meta        Meta
	AvatarURL   string
	CVPath      string
	Bio         string
	SkillGroups []SkillGroup
	Projects    []Project
	Education   []EduEntry
	Contact     Contact
}

type ProjectPageData struct {
	Name      string
	AvatarURL string
	Project   Project
}

type ProjectsIndexData struct {
	Name     string
	Projects []Project
}

type Contact struct {
	Email  string
	GitHub string
	Phone  string
	Based  string
}

var projects = []Project{
	{
		Slug:     "ascii-art-foundations",
		FuncName: "ValidateBanner",
		Params:   "m map[rune][]string",
		Returns:  "bool",
		Blurb:    "Foundational ASCII-art rendering engine — the banner-parsing core that later became ascii-art-web.",
		Bullets: []string{
			"Built ValidateBanner and MergeBanners around map[rune][]string banner maps.",
			"Established recurring patterns used across later projects: rune indexing, nil-map detection, slice-length validation.",
			"Worked directly with strings.Builder and rune-level ASCII arithmetic — no higher-level string libraries.",
		},
		Stack:   []string{"Go", "runes", "strings.Builder"},
		RepoURL: "https://github.com/miikese",
	},
	{
		Slug:     "ascii-art-web",
		FuncName: "AsciiArtWeb",
		Params:   "text string",
		Returns:  "HTTPService",
		Blurb:    "A Go HTTP server with a web GUI for generating ASCII art, supporting three banner styles with live switching.",
		Bullets: []string{
			"Built a form-driven front end that auto-submits on style change, no client-side JS framework.",
			"Rendered banners server-side with html/template — no CSS build step.",
			"Handled invalid or empty input with correct HTTP status codes.",
		},
		Stack:   []string{"Go", "net/http", "html/template"},
		RepoURL: "https://github.com/miikese",
	},
	{
		Slug:     "echo-form-server",
		FuncName: "EchoFormServer",
		Params:   "r *http.Request",
		Returns:  "Response",
		Blurb:    "A Go HTTP server built from scratch handling /echo and /form endpoints.",
		Bullets: []string{
			"Enforced Content-Type headers and rejected empty request bodies.",
			"Validated form fields before they reached any business logic.",
			"Verified endpoint behavior with a Bash smoke-test script.",
			"Built entirely on net/http — no router library.",
		},
		Stack:   []string{"Go", "net/http", "Bash"},
		RepoURL: "https://github.com/miikese",
	},
	{
		Slug:     "ascii-art-color",
		FuncName: "AsciiArtColor",
		Params:   "text string, color string",
		Returns:  "string",
		Blurb:    "Group project extending ascii-art with terminal color support via a --color flag.",
		Bullets: []string{
			"Used only the Go standard library — no external color packages.",
			"Collaborated on a shared codebase with another contributor.",
			"Code review caught an empty-substring crash risk and spec-wording mismatches before ship.",
		},
		Stack:   []string{"Go", "code review"},
		RepoURL: "https://github.com/miikese",
	},
}

var skillGroups = []SkillGroup{
	{Category: "GO", Color: "#5FB4B0", Skills: []string{"net/http", "html/template", "os", "bufio", "go test", "slices / maps / structs", "runes", "strings.Builder", "CLI flags", "file I/O"}},
	{Category: "PYTHON", Color: "#D4A24C", Skills: []string{"Python 3.12", "scripting", "venv", "structured JSON output"}},
	{Category: "FRONTEND", Color: "#8FA9D6", Skills: []string{"HTML5 semantic markup", "server-side templating", "CSS basics"}},
	{Category: "TOOLING", Color: "#C77B5B", Skills: []string{"Git branching / merging", "Bash", "Linux CLI", "unit testing", "project management"}},
}

var education = []EduEntry{
	{
		Period: "Present",
		Title:  "Software Development Programme — Go & Python Curriculum",
		Org:    "LEEF Centre, Otukpo",
		Detail: "Project-based curriculum built around Go fundamentals through HTTP services and CLI tools, migrated from an earlier 01edu-network Go program; the curriculum has since expanded to include Python.",
	},
}

func main() {
	tmpl, err := template.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		log.Fatalf("parsing templates: %v", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(staticFS)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		data := PageData{
			Name:        "Michael Ikese Emmanuel",
			RoleTag:     "BACKEND DEVELOPER — GO & PYTHON",
			Headline:    "Backend systems built from the standard library up.",
			SubHeadline: "Michael builds HTTP services and CLI tools in Go, backed by a project-based curriculum in algorithms, data structures, and Git-based collaboration.",
			Meta: Meta{
				Stack:         "Go · Python",
				Based:         "Nigeria",
				ProjectsCount: len(projects),
			},
			AvatarURL: "/static/img/avatar.jpg",
			CVPath:    "/static/cv/resume.pdf",
			Bio:       "Backend developer focused on Go, building HTTP servers, APIs, and command-line tools from the ground up using the standard library. Currently studying software development at LEEF Centre, Otukpo, and picking up Python alongside Go. Comfortable working through problems from first principles rather than reaching for a framework, and open to remote opportunities.",
			SkillGroups: skillGroups,
			Projects:    projects,
			Education:   education,
			Contact: Contact{
				Email:  "emmanlemichel2019@gmail.com",
				GitHub: "https://github.com/miikese",
				Phone:  "08160345977",
				Based:  "Nigeria",
			},
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			log.Printf("template error: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/projects" {
			http.NotFound(w, r)
			return
		}
		data := ProjectsIndexData{
			Name:     "Michael Ikese Emmanuel",
			Projects: projects,
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "projects.html", data); err != nil {
			log.Printf("template error: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/projects/")
		for _, p := range projects {
			if p.Slug == slug {
				data := ProjectPageData{
					Name:      "Michael Ikese Emmanuel",
					AvatarURL: "/static/img/avatar.jpg",
					Project:   p,
				}
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				if err := tmpl.ExecuteTemplate(w, "project.html", data); err != nil {
					log.Printf("template error: %v", err)
					http.Error(w, "internal error", http.StatusInternalServerError)
				}
				return
			}
		}
		http.NotFound(w, r)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("serving on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
