# portfolio

A single-binary Go server that renders your project portfolio. No frameworks,
no build step — the server itself is one of the projects it shows off.

## Run it locally

```
go run .
```

Then open http://localhost:8080

## Before you deploy

- Add your photo as `static/img/avatar.jpg` (it's referenced in the header
  already — just drop the file in).
- Open `main.go` and fill in `Contact.LinkedIn` if you want that row to show.
- Edit the `projects` slice to add, remove, or update project entries.

## Deploy

Because it's a single static binary, it runs almost anywhere:

```
GOOS=linux GOARCH=amd64 go build -o portfolio .
```

Copy the `portfolio` binary to a server (a $5/mo VPS, Fly.io, Render, Railway
all work) and run it. It reads the `PORT` environment variable if you need
something other than 8080.

## Structure

```
main.go              — server, routes, project data
templates/index.html — page template
static/css/style.css — styling
```
