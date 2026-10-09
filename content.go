package main

// Repository-backed entries were reviewed against GitHub on 8 October 2026.
// Private projects publish summaries only; RepoURL is deliberately empty.
// The public collection includes implemented work and the requested TalentGrid project in development.
var projects = []Project{
	{
		Slug: "global-browser", Title: "Global Browser", RepoName: "global-browser",
		Category: "Systems & infrastructure", Featured: true, Status: "Working MVP", Visibility: "Private source",
		Blurb: "A browser you control through the web. Go powers isolated Chromium sessions, account access, and a live browser viewer.",
		Focus: "Browser infrastructure · access control", Outcome: "A working local browser platform with account management and session controls.",
		Bullets: []string{"Streams browser frames and input over a protected WebSocket connection.", "Combines isolated Chromium sessions with authentication, account roles, quotas, and SQLite persistence.", "Includes destination checks, origin validation, security headers, and browser lifecycle cleanup.", "Supports configured regional gateways with outbound verification. Payment processing and email delivery remain future work."},
		Stack:   []string{"Go", "Chromium", "WebSocket", "SQLite", "Linux"},
	},
	{
		Slug: "past-quest-solver", Title: "Smile Learning", RepoName: "past-quest-solver",
		Category: "Web applications", Featured: true, Status: "Implemented", Visibility: "Public source",
		Blurb: "A study interface that helps students explore past exam questions, filter by subject, and read answers in one place.",
		Focus: "Learning experience · component design", Outcome: "A responsive question-browsing interface built with React and TypeScript.",
		Bullets: []string{"Organizes study material into reusable question cards and subject views.", "Includes search and filtering controls for discovering relevant questions.", "Uses React, TypeScript, shadcn/ui, and Tailwind CSS for a responsive interface. The repository includes static and mock study data."},
		Stack:   []string{"React", "TypeScript", "Vite", "Tailwind CSS"},
		RepoURL: "https://github.com/miikese/past-quest-solver", SourceLabel: "View source",
	},
	{
		Slug: "campus-resource-manager", Title: "Campus Resource Manager", RepoName: "assesment",
		Category: "CLI tools", Featured: true, Status: "Implemented", Visibility: "Private source",
		Blurb: "A Python application for managing campus resources: inventory, borrowing, returns, and reports, with saved JSON records.",
		Focus: "Data integrity · practical automation", Outcome: "A complete menu-driven resource-management application with regression tests.",
		Bullets: []string{"Supports resource creation, search, category filters, borrowing, returns, and inventory reports.", "Validates requests before changing inventory and saves successful operations to JSON.", "Includes demonstration output, design documentation, and standard-library regression tests.", "Designed for a single local operator; concurrent writes would require a different storage strategy."},
		Stack:   []string{"Python", "JSON", "unittest", "CLI"},
	},
	{
		Slug: "talentgrid", Title: "TalentGrid", RepoName: "talentgrid",
		Category: "Web applications", Featured: true, Status: "In development", Visibility: "Private source",
		Blurb: "A sector-based talent directory that connects professional profiles, skills, and work experience with searchable discovery.",
		Focus: "Talent discovery · relational data", Outcome: "An evolving Go and PostgreSQL application with profile, account, and talent-search handlers.",
		Bullets: []string{"Current source includes registration, login, JWT authentication, and account-management handlers.", "Models sectors, skills, and work experience in PostgreSQL, with profile editing and filtered, paginated talent search.", "Includes a web interface and routes for contact requests and support-message storage.", "Source reviewed on 9 October 2026. Full database-backed runtime validation and production deployment have not been verified in this portfolio review."},
		Stack:   []string{"Go", "PostgreSQL", "JWT", "HTML", "CSS", "JavaScript"},
	},
	{
		Slug: "portfolio-site", Title: "Engineering Portfolio", RepoName: "portfolio-site",
		Category: "Web applications", Featured: true, Status: "Implemented", Visibility: "Public source",
		Blurb: "The site you're exploring: a compact Go application with an editorial interface, a searchable project collection, and embedded assets.",
		Focus: "Accessible interfaces · Go delivery", Outcome: "A portfolio served from a single binary, with no frontend build pipeline.",
		Bullets: []string{"Uses Go's html/template for server-rendered pages and embeds templates, images, and styles in one binary.", "Provides project search, category filtering, shareable filtered URLs, and persistent light and dark themes.", "Includes HTTP route tests, security headers, timeouts, graceful shutdown, and a production container configuration."},
		Stack:   []string{"Go", "HTML", "CSS", "JavaScript"},
		RepoURL: "https://github.com/miikese/portfolio-site", SourceLabel: "View source",
	},
}

var skillGroups = []SkillGroup{
	{Category: "01 / Languages", Skills: []string{"Go", "Python", "JavaScript", "TypeScript", "HTML & CSS"}},
	{Category: "02 / Backend & data", Skills: []string{"net/http", "html/template", "SQLite", "JSON persistence", "WebSocket", "CLI tools"}},
	{Category: "03 / Interfaces & workflow", Skills: []string{"React", "Tailwind CSS", "Git", "Linux", "Bash", "Unit testing", "Documentation"}},
	{Category: "04 / Exploring next", Skills: []string{"Cybersecurity fundamentals", "Secure software design", "Networking", "PostgreSQL", "New languages & frameworks"}},
	{Category: "05 / Project management in training", Skills: []string{"Scope & requirements", "Work breakdown & scheduling", "Agile & Scrum fundamentals", "Prioritization", "Risk & issue tracking", "Stakeholder communication", "Progress reporting", "Delivery documentation"}},
}

var education = []EduEntry{
	{Period: "Present", Title: "Software Development Programme", Org: "LEEF Centre, Otukpo", Detail: "A project-based Go and Python curriculum: building command-line tools, HTTP services, and practical applications while developing stronger problem-solving habits."},
	{Period: "National Diploma", Title: "Civil Engineering Technology", Org: "Federal Polytechnic, Nasarawa", Detail: "An engineering qualification that forms part of my background alongside my current software development studies."},
}

func portfolioData() PageData {
	return PageData{
		Name: "Michael Ikese Emmanuel", RoleTag: "Software engineer & programmer",
		Headline: "Useful software.", HeadlineAccent: "Thoughtful engineering.",
		SubHeadline: "I'm Michael Ikese Emmanuel, a software engineer based in Nigeria. I build practical tools and web applications with Go, Python, and TypeScript — with care for the people using them and the systems behind them.",
		AvatarURL:   "/static/img/portrait.png", CVPath: "/static/cv/resume.pdf",
		Bio:         "I build software to solve everyday problems, from a browser platform with isolated sessions to a study interface and a campus resource manager. I enjoy connecting a clear interface to dependable backend logic, testing edge cases, and documenting how things work. Alongside my software development studies at LEEF Centre, Otukpo, I'm training in project management and exploring cybersecurity fundamentals. I'm looking for opportunities to contribute, learn from others, and grow through meaningful work.",
		SkillGroups: skillGroups, Projects: projects, Education: education,
		Contact: Contact{Email: "emmanlemichel2019@gmail.com", GitHub: "https://github.com/miikese", Phone: "08160345977", PhoneURL: "+2348160345977", Based: "Nigeria"},
	}
}
