# Michael Ikese Emmanuel — Portfolio

A professional portfolio for Michael Ikese Emmanuel: software engineer and programmer, project manager in training, and cybersecurity enthusiast. The curated collection highlights Global Browser, Smile Learning, Campus Resource Manager, TalentGrid (in development), and this portfolio. It runs as a single Go binary with embedded templates and assets, with no frontend framework, package manager, or build step.

## Run locally

Install Go 1.27.2 or newer, then run:

```sh
git clone https://github.com/miikese/portfolio-site.git
cd portfolio-site
go run .
```

Open http://localhost:8080. The site includes a homepage, a searchable project directory at `/projects`, and a detail page for each project. Search and theme controls use small vanilla JavaScript enhancements; navigation, content, contact links, and CV downloads also work without JavaScript.

## Customize

- **Profile, contact details, skills, education:** edit `portfolioData`, `skillGroups`, and `education` in `content.go`.
- **Projects:** edit the `projects` slice in `content.go`. Use a unique URL-safe `Slug`, a `Title`, a `Category`, a description, and a technology stack. `Featured: true` includes a project on the homepage.
- **Links:** supply the actual repository in `RepoURL`, a suitable `SourceLabel`, and an optional `DemoURL`. The public collection is curated; small exercises and unimplemented concepts are omitted.
- **Portrait and branding:** the homepage uses `static/img/portrait.png`, an AI-assisted professional portrait based on Michael's original photo, with a navy background coordinated with the site. Visitors can pause its subtle motion. The monogram and custom profile banners remain under `static/img/`.
- **CV:** edit `cv/content.json`, then run `python3 scripts/build-cv.py`. This generates a selectable-text, single-page A4 PDF using only the Python standard library.
- **Design:** edit `static/css/style.css`; shared navigation and metadata live in `templates/shared.html`.
- **Categories:** the project filter derives its categories from the content automatically. Private entries have a written overview and no public source link.

All files under `static/` and `templates/` are embedded at build time. Rebuild after editing them when running a compiled binary.

## Configuration

| Variable | Default | Purpose |
| --- | --- | --- |
| `PORT` | `8080` | HTTP listen port |
| `SITE_URL` | unset | Public origin, e.g. `https://portfolio.example`; enables `/sitemap.xml` and its entry in `/robots.txt` |

The server reads environment variables directly; it does not load `.env` files. `.env.example` is a configuration reference.

```sh
PORT=8081 SITE_URL=https://portfolio.example go run .
```

Use the scheme and hostname of your deployed site for `SITE_URL`, without a subpath, credentials, query, or fragment. Without it, the portfolio still works and `/sitemap.xml` returns 404 rather than advertising an incorrect domain.

## Verify

```sh
gofmt -w main.go content.go main_test.go
go vet ./...
go test -race -cover ./...
go build -trimpath -o portfolio .
```

GitHub Actions checks formatting, analysis, route and asset tests, a production binary, and the container build on pushes to `main` and pull requests. Tests cover project pages, 404s, unsupported methods, HEAD requests, the CV and image, security headers, health checks, and sitemap configuration.

## Deploy

Build and run the executable:

```sh
CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o portfolio .
PORT=8080 ./portfolio
```

No separate template or static directories are needed alongside the binary. For a different target, set `GOOS` and `GOARCH` at build time.

Or use the multi-stage container, which runs as a non-root user:

```sh
docker build -t portfolio .
docker run --rm -p 8080:8080 -e PORT=8080 portfolio
```

Configure your hosting platform's HTTP health check to `/healthz`, which returns `{"status":"ok"}`. Termination signals trigger graceful shutdown. For public hosting, put the HTTP server behind the platform's HTTPS endpoint or a reverse proxy that terminates TLS. Set `SITE_URL` to that public origin.

## Structure

```text
content.go                 Profile and project content
main.go                    HTTP routes, rendering, headers, server lifecycle
main_test.go               HTTP and deployment configuration tests
templates/shared.html      Metadata, navigation, footer, project cards
templates/index.html       Homepage
templates/projects.html    Project directory
templates/project.html     Project detail
templates/error.html       Accessible 404 page
static/css/style.css       Responsive light/dark design
static/js/theme.js         Theme initialization before first paint
static/js/main.js          Mobile menu, theme, filters, and motion controls
static/img/                Custom logo, banners, original portrait, and favicon
static/cv/resume.pdf        Downloadable CV
.github/workflows/ci.yml    Automated verification
Dockerfile                 Minimal non-root production image
```

The layout includes keyboard focus styles, a skip link, labeled search controls, reduced-motion support, print styles, page descriptions, and social preview metadata. Static directory listings are disabled. The server sets security headers and bounded request timeouts.

## Profile assets and presentation

`github-profile/README.md` is the prepared GitHub profile introduction, project showcase, training focus, and verified contact information. `github-profile/SETUP.md` records account fields, profile visibility, and suggested pins.

The original banner and logo share the site design. Run `python3 scripts/build-profile-graphics.py` to rebuild the GIF and PNG graphics; this optional script requires Pillow and the DejaVu fonts. The site has no runtime Python dependency. The SVG sources can be edited directly. The homepage uses lightweight entrance, portrait, and orbit motion, with an explicit pause control and a reduced-motion override. Social previews use the static PNG banner for compatibility.

Project search accepts shareable `q`, `category`, and `collection` query parameters. With JavaScript disabled, every project remains visible and navigation and contact links still work.

See [project selection](docs/project-selection.md) for the content review and [refresh report](docs/refresh-report.md) for validation and delivery details.
