package main

var projects = []Project{
	{
		Slug: "past-quest-solver", Title: "Smile Learning", Category: "Web applications", Featured: true,
		FuncName: "SmileLearning", Params: "", Returns: "StudyPlatform",
		Blurb:   "A responsive study platform for browsing past exam questions and answers, built with React and TypeScript.",
		Bullets: []string{"Organized the interface around question cards, search, and subject filters.", "Built reusable components with React, TypeScript, and shadcn/ui.", "Used Vite and Tailwind CSS for the development workflow and responsive styling."},
		Stack:   []string{"React", "TypeScript", "Vite", "Tailwind CSS"},
		RepoURL: "https://github.com/miikese/past-quest-solver", SourceLabel: "View source", DemoURL: "https://past-quest-solver.vercel.app",
	},
	{
		Slug: "portfolio-site", Title: "Go Portfolio", Category: "HTTP services", Featured: true,
		FuncName: "Portfolio", Params: "", Returns: "HTTPService",
		Blurb:   "This portfolio: a single Go binary with embedded templates, a project directory, and responsive pages.",
		Bullets: []string{"Rendered accessible pages with Go's html/template and embedded the assets in the binary.", "Added HTTP route tests, security headers, server timeouts, and graceful shutdown.", "Built mobile navigation, project search, and a persistent theme preference with vanilla JavaScript."},
		Stack:   []string{"Go", "net/http", "html/template", "CSS", "JavaScript"},
		RepoURL: "https://github.com/miikese/portfolio-site", SourceLabel: "View source",
	},
	{
		Slug:     "ascii-art-foundations",
		Title:    "ASCII Art Foundations",
		Category: "CLI tools",
		FuncName: "ValidateBanner",
		Params:   "m map[rune][]string",
		Returns:  "bool",
		Blurb:    "Foundational ASCII-art rendering engine — the banner-parsing core that later became ascii-art-web.",
		Bullets: []string{
			"Built ValidateBanner and MergeBanners around map[rune][]string banner maps.",
			"Established recurring patterns used across later projects: rune indexing, nil-map detection, slice-length validation.",
			"Worked directly with strings.Builder and rune-level ASCII arithmetic — no higher-level string libraries.",
		},
		Stack:       []string{"Go", "runes", "strings.Builder"},
		RepoURL:     "https://github.com/miikese",
		SourceLabel: "GitHub profile",
	},
	{
		Slug:     "ascii-art-web",
		Featured: true,
		Title:    "ASCII Art Web",
		Category: "Web applications",
		FuncName: "AsciiArtWeb",
		Params:   "text string",
		Returns:  "HTTPService",
		Blurb:    "A Go HTTP server with a web GUI for generating ASCII art, supporting three banner styles with live switching.",
		Bullets: []string{
			"Built a form-driven front end that auto-submits on style change, no client-side JS framework.",
			"Rendered banners server-side with html/template — no CSS build step.",
			"Handled invalid or empty input with correct HTTP status codes.",
		},
		Stack:       []string{"Go", "net/http", "html/template"},
		RepoURL:     "https://github.com/miikese",
		SourceLabel: "GitHub profile",
	},
	{
		Slug:     "echo-form-server",
		Title:    "Echo & Form Server",
		Category: "HTTP services",
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
		Stack:       []string{"Go", "net/http", "Bash"},
		RepoURL:     "https://github.com/miikese",
		SourceLabel: "GitHub profile",
	},
	{
		Slug:     "ascii-art-color",
		Title:    "ASCII Art Color",
		Category: "CLI tools",
		FuncName: "AsciiArtColor",
		Params:   "text string, color string",
		Returns:  "string",
		Blurb:    "Group project extending ascii-art with terminal color support via a --color flag.",
		Bullets: []string{
			"Used only the Go standard library — no external color packages.",
			"Collaborated on a shared codebase with another contributor.",
			"Code review caught an empty-substring crash risk and spec-wording mismatches before ship.",
		},
		Stack:       []string{"Go", "code review"},
		RepoURL:     "https://github.com/miikese",
		SourceLabel: "GitHub profile",
	},
}

var skillGroups = []SkillGroup{
	{Category: "GO", Skills: []string{"net/http", "html/template", "os", "bufio", "go test", "slices / maps / structs", "runes", "strings.Builder", "CLI flags", "file I/O"}},
	{Category: "PYTHON", Skills: []string{"Python 3.12", "scripting", "venv", "structured JSON output"}},
	{Category: "FRONTEND", Skills: []string{"HTML5 semantic markup", "server-side templating", "CSS basics"}},
	{Category: "TOOLING", Skills: []string{"Git branching / merging", "Bash", "Linux CLI", "unit testing", "project management"}},
}

var education = []EduEntry{
	{
		Period: "Present",
		Title:  "Software Development Programme — Go & Python Curriculum",
		Org:    "LEEF Centre, Otukpo",
		Detail: "Project-based curriculum built around Go fundamentals through HTTP services and CLI tools, migrated from an earlier 01edu-network Go program; the curriculum has since expanded to include Python.",
	},
}

func portfolioData() PageData {
	return PageData{
		Name: "Michael Ikese Emmanuel", RoleTag: "Backend developer · Go & Python",
		Headline: "Thoughtful code.", HeadlineAccent: "Useful software.",
		SubHeadline: "I build HTTP services, command-line tools, and web experiences. My work starts with solid foundations, clear interfaces, and a curiosity for how things work.",
		AvatarURL:   "/static/img/avatar.jpg", CVPath: "/static/cv/resume.pdf",
		Bio:         "I'm Michael, a backend developer based in Nigeria. I build with Go's standard library, study Python, and enjoy turning a problem into something people can use. I'm currently learning at LEEF Centre, Otukpo, through a project-based software development programme, and I'm open to remote opportunities.",
		SkillGroups: skillGroups, Projects: projects, Education: education,
		Contact: Contact{Email: "emmanlemichel2019@gmail.com", GitHub: "https://github.com/miikese", Phone: "08160345977", PhoneURL: "+2348160345977", Based: "Nigeria"},
	}
}
