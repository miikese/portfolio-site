package main

// Current public project descriptions reviewed on 27 October 2026.
// Older private summaries retain their stated review boundaries.
// Private projects publish summaries only; RepoURL is deliberately empty.
// Collaboration credit records the user-confirmed project role, not project ownership.
var projects = []Project{
	{
		Slug: "marvellous-motors", Title: "Marvellous Motors", RepoName: "Marvellous-motor2",
		Category: "Web applications", Featured: true, Status: "Working application · in development", Visibility: "Public source",
		Role: "Development collaborator", Collaborators: "Otiti Mordecai Marvellous · project owner",
		Blurb: "A collaborative vehicle-sales application connecting a searchable showroom to customer accounts, itemised delivery pricing and an administrator-led order workflow.",
		Focus: "Commerce workflows · inventory · customer experience", Outcome: "Collaborated with Otiti Mordecai Marvellous on her Motors project, supporting practical customer and administrator workflows. Project ownership remains with Mordecai.",
		Bullets: []string{"Customer registration, editable profiles, password settings and account management backed by SQLite.", "Vehicle discovery, administrator-managed stock quantities and configurable delivery charges feed a saved order price breakdown.", "Bank-transfer evidence is submitted for administrator review; delivery tracking shows administrator-set arrival estimates.", "Local validation does not establish production deployment, real payment processing or inventory availability."},
		Stack:   []string{"Go", "SQLite", "HTML", "CSS", "JavaScript"},
		RepoURL: "https://github.com/mordecai12-cool/Marvellous-motor2", SourceLabel: "View collaborative source",
	},
	{
		Slug: "campusflow", Title: "CampusFlow", RepoName: "campusflow",
		Category: "CLI tools", Featured: false, Status: "Implemented · AI-assisted", Visibility: "Public source",
		Role:  "AI-assisted project development",
		Blurb: "A Python campus helpdesk for creating support tickets, prioritising incidents, assigning staff and following work through a validated status workflow.",
		Focus: "Workflow validation · safe persistence", Outcome: "A standard-library command-line application with priority rules, an open-ticket queue, reports and atomic JSON persistence. The implementation is AI-authored; human collaboration and peer-review evidence remain outstanding.",
		Bullets: []string{"Calculate urgency from impact and affected users, assign staff and enforce open → in progress → resolved transitions with explicit reopening.", "Preserve records on invalid input or corrupted storage, allocate stable numeric IDs and save using atomic file replacement.", "Repository includes regression tests, a temporary-data demo and workflow documentation. JSON storage is intended for one local operator."},
		Stack:   []string{"Python", "JSON", "unittest", "CLI"},
		RepoURL: "https://github.com/miikese/campusflow", SourceLabel: "View source",
	},

	{
		Slug: "community-connect", Title: "CommunityConnect", RepoName: "community-connect",
		Category: "Web applications", Featured: false, Status: "Working MVP · in development", Visibility: "Public source",
		Blurb: "A community development platform where residents report local issues, coordinate volunteers, discover opportunities and share skills to help improve their communities.",
		Focus: "Community reporting · accountable participation", Outcome: "A tested community-coordination MVP with moderated reporting and real activity counts. Public boards start empty and community review does not imply official government action.",
		Bullets: []string{"Submit private pending reports, follow reviewed issues from acknowledgement to resolution, and view status history and notifications.", "Register reviewed organizations, publish opportunities and enroll volunteers with ownership and deadline checks.", "Discover skills offers, jobs, training, scholarships, events and nonfinancial support through moderated community boards.", "Use approximate public locations, flag abuse, export aggregate counts, download event calendars and delete account data."},
		Stack:   []string{"React", "TypeScript", "Leaflet", "Cloudflare Workers", "D1 / SQLite", "R2"},
		RepoURL: "https://github.com/miikese/community-connect", SourceLabel: "View source",
		PreviewCaption:  "Actual application screenshot from an earlier local preview; current features may differ. Hosted demo access depends on sign-in and project permissions.",
		ScreenshotWidth: 1440, ScreenshotHeight: 1603, ScreenshotURL: "/static/img/projects/community-connect.png", ScreenshotAlt: "CommunityConnect dashboard with community issue reporting and opportunity discovery",
		DemoURL: "https://community-connect.readyloom3.chatgpt.site", DemoLabel: "Hosted MVP · sign-in required ↗",
	},

	{
		Slug: "mindgrow", Title: "MindGrow", RepoName: "mindgrow",
		Category: "Web applications", Featured: true, Status: "Working MVP · in development", Visibility: "Public source",
		Blurb: "An educational gaming and brain-training platform for children learning programming and problem-solving, with relaxed memory, logic and reasoning activities for teenagers and adults.",
		Focus: "Playable learning · private progress", Outcome: "Thirteen playable activities and 30 introductory lessons, with Blockly programming, sandboxed HTML/CSS previews and private guardian-owned learning history.",
		Bullets: []string{"Play thirteen activities, including memory, logic, visual coding and Falling blocks, with clear rules and keyboard/touch controls.", "Build Blockly movement programs and practice HTML/CSS in a sandbox. Thirty lessons cover foundations, websites, Python, Go and CSS; Python and Go examples are read-and-reason exercises.", "Choose age-appropriate paths, track practice through private nickname profiles, and receive rules-based recommendations.", "Resume device-local games, use calm mode and accessible controls, and delete learner profiles and progress. No public child profiles or intelligence claims."},
		Stack:   []string{"React", "TypeScript", "Blockly", "Cloudflare Workers", "D1 / SQLite"},
		RepoURL: "https://github.com/miikese/mindgrow", SourceLabel: "View source",
		PreviewCaption:  "Actual application screenshot from an earlier local preview; current features may differ. Hosted demo access depends on sign-in and project permissions.",
		ScreenshotWidth: 1440, ScreenshotHeight: 1439, ScreenshotURL: "/static/img/projects/mindgrow.png", ScreenshotAlt: "MindGrow learning dashboard with playable memory, logic and coding activities",
		DemoURL: "https://mindgrow.readyloom3.chatgpt.site", DemoLabel: "Hosted MVP · sign-in required ↗",
	},

	{
		Slug: "benue-street-locator", Title: "Benue Street Locator", RepoName: "benue-street-locator",
		Category: "Web applications", Featured: true, Status: "Working MVP · in development", Visibility: "Public source",
		Blurb: "An interactive mapping and location-discovery platform for Benue State, starting with Otukpo. Explore streets, landmarks and public places, share locations and contribute local knowledge.",
		Focus: "Mapping · search · reviewed contributions", Outcome: "A mapping MVP with street and satellite views, predictive local search, source-labelled references and reviewed contributions across Benue. Coverage remains incomplete; unpinned references are never presented as exact coordinates.",
		Bullets: []string{"Predict local names from the first letter, match aliases and spelling variations, and filter by LGA or category. Source-labelled street references remain distinct from map pins.", "Explore Leaflet maps, save places, and share private pins or directory locations through links, WhatsApp and QR codes.", "Submit places and corrections privately; approved places and your submissions remain discoverable through large moderation backlogs.", "Open external directions; in-app driving routes require a configured OSRM provider. Low-data mode pauses map tiles and cached public details work offline."},
		Stack:   []string{"React", "TypeScript", "Leaflet", "Cloudflare Workers", "D1 / SQLite", "R2"},
		RepoURL: "https://github.com/miikese/benue-street-locator", SourceLabel: "View source",
		PreviewCaption:  "Actual application screenshot from an earlier local preview; current features may differ. Hosted demo access depends on sign-in and project permissions.",
		ScreenshotWidth: 1440, ScreenshotHeight: 1110, ScreenshotURL: "/static/img/projects/benue-street-locator.png", ScreenshotAlt: "Benue Street Locator showing place search and the OpenStreetMap of Otukpo",
		DemoURL: "https://benue-street-locator.readyloom3.chatgpt.site", DemoLabel: "Hosted MVP · sign-in required ↗",
	},

	{
		Slug: "smartconvert", Title: "SmartConvert", RepoName: "smartconvert",
		Category: "Web applications", Featured: false, Status: "Tested demo MVP", Visibility: "Private source",
		Blurb: "A booking-slip translation workspace that compares selection meaning, explains uncertain matches, and asks for review before export.",
		Focus: "Go + Python orchestration · semantic matching", Outcome: "A locally runnable hybrid application with verified demo exports and persistent review receipts.",
		Bullets: []string{"Go owns authentication, roles, subscriptions, independent NGN/USD prices, notifications, history, and database writes; a private FastAPI service handles matching.", "Uses RapidFuzz with event identity checks and exact market lines. Different outcome sides, periods, or settlement rules cannot be accepted as equivalent.", "Three fictitious demo adapters support inspect, compare, confirm, export, and re-import verification. Demo batches never consume real conversion credits.", "Validated with Go race tests, Python tests, Go-to-Python integration, PostgreSQL credit concurrency, and desktop/mobile browser checks. Real bookmaker integrations and live payments are unavailable."},
		Stack:   []string{"Go", "Python", "FastAPI", "RapidFuzz", "PostgreSQL", "JavaScript"},
	},

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
		Category: "Web applications", Featured: false, Status: "Implemented", Visibility: "Public source",
		Blurb: "A study interface that helps students explore past exam questions, filter by subject, and read answers in one place.",
		Focus: "Learning experience · component design", Outcome: "A responsive question-browsing interface built with React and TypeScript.",
		Bullets: []string{"Organizes study material into reusable question cards and subject views.", "Includes search and filtering controls for discovering relevant questions.", "Uses React, TypeScript, shadcn/ui, and Tailwind CSS for a responsive interface. The repository includes static and mock study data."},
		Stack:   []string{"React", "TypeScript", "Vite", "Tailwind CSS"},
		RepoURL: "https://github.com/miikese/past-quest-solver", SourceLabel: "View source",
	},
	{
		Slug: "campus-resource-manager", Title: "Campus Resource Manager", RepoName: "assesment",
		Category: "CLI tools", Featured: false, Status: "Implemented", Visibility: "Private source",
		Blurb: "A Python application for managing campus resources: inventory, borrowing, returns, and reports, with saved JSON records.",
		Focus: "Data integrity · practical automation", Outcome: "A complete menu-driven resource-management application with regression tests.",
		Bullets: []string{"Supports resource creation, search, category filters, borrowing, returns, and inventory reports.", "Validates requests before changing inventory and saves successful operations to JSON.", "Includes demonstration output, design documentation, and standard-library regression tests.", "Designed for a single local operator; concurrent writes would require a different storage strategy."},
		Stack:   []string{"Python", "JSON", "unittest", "CLI"},
	},
	{
		Slug: "talentgrid", Title: "TalentGrid", RepoName: "talentgrid",
		Category: "Web applications", Featured: false, Status: "In development", Visibility: "Private source",
		Blurb: "A sector-based talent directory that connects professional profiles, skills, and work experience with searchable discovery.",
		Focus: "Talent discovery · relational data", Outcome: "An evolving Go and PostgreSQL application with profile, account, and talent-search handlers.",
		Bullets: []string{"Current source includes registration, login, JWT authentication, and account-management handlers.", "Models sectors, skills, and work experience in PostgreSQL, with profile editing and filtered, paginated talent search.", "Includes private contact requests, employer vacancies, applications and account-deactivation workflows. Search guards prevent delayed results from replacing a newer query.", "Search behavior was checked with isolated browser fixtures on 28 October 2026. Full database-backed runtime validation and production deployment were not repeated in this review."},
		Stack:   []string{"Go", "PostgreSQL", "JWT", "HTML", "CSS", "JavaScript"},
	},
	{
		Slug: "reachaid", Title: "ReachAid", RepoName: "ReachAid",
		Category: "Web applications", Featured: false, Status: "Closed pilot · in development", Visibility: "Private source",
		Blurb: "A location-aware directory prototype for discovering essential community services, with controlled account access and review workflows.",
		Focus: "Service discovery · responsible information", Outcome: "An implemented Go and Next.js pilot with PostgreSQL and PostGIS. The public overview describes the software, not verified service availability or an open public launch.",
		Bullets: []string{"Supports location-based service discovery, organization records and moderated information workflows.", "Keeps source code private and preserves the pilot's access controls.", "Real provider verification, production email delivery and operational release checks remain prerequisites for wider launch."},
		Stack:   []string{"Go", "Next.js", "TypeScript", "PostgreSQL", "PostGIS"},
	},
	{
		Slug: "portfolio-site", Title: "Engineering Portfolio", RepoName: "portfolio-site",
		Category: "Web applications", Featured: false, Status: "Implemented", Visibility: "Public source",
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
		Bio:         "I build software to solve everyday problems, from mapping and learning platforms to browser infrastructure and practical Python tools. I also collaborate with Otiti Mordecai Marvellous on her Marvellous Motors application. I enjoy connecting a clear interface to dependable backend logic, testing edge cases, and documenting how things work. Alongside my software development studies at LEEF Centre, Otukpo, I'm training in project management and exploring cybersecurity fundamentals. I'm looking for opportunities to contribute, learn from others, and grow through meaningful work.",
		SkillGroups: skillGroups, Projects: projects, Education: education,
		Contact: Contact{Email: "emmanlemichel2019@gmail.com", GitHub: "https://github.com/miikese", Phone: "08160345977", PhoneURL: "+2348160345977", Based: "Nigeria"},
	}
}
