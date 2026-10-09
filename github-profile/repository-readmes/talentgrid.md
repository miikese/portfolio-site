# TalentGrid

A sector-based talent directory built with Go and PostgreSQL. Professionals maintain a profile, skills, and work experience; clients discover talent through sector, skill, location, availability, and text filters.

**Status: in development.** The current source contains application handlers and a browser interface, beyond the original scaffold. This documentation reflects source review on 9 October 2026; it does not claim a verified production deployment or a completed database-backed test suite.

## What the source includes

- Talent/client registration, bcrypt password hashing, login, and JWT authentication.
- Public talent profiles and an authenticated personal profile; public profile responses omit email and phone.
- Profile editing, sector/subsector selection, skills, work-experience records, and profile-completeness calculation.
- Paginated discovery with sector, skill, location, availability, and text filters, using parameterized PostgreSQL queries.
- Contact-request creation, inboxes, and response routes.
- Password changes, account deactivation/reactivation, and support-message storage.
- A browser interface served by the Go application, plus a JSON health endpoint.

Support messages are stored in the database. This repository does not establish that outbound support email delivery is implemented.

## Stack and structure

Go (`net/http`), PostgreSQL (`pgx`), JWT, bcrypt, and HTML/CSS/JavaScript.

```text
cmd/api/               Server entry point and route registration
internal/auth/         JWT creation and verification
internal/config/       Environment configuration
internal/db/           PostgreSQL connection pool
internal/handlers/     Accounts, profiles, discovery, contact and support
internal/middleware/   Authentication
internal/models/       Domain records
migrations/            Numbered SQL schema and seed files
web/index.html         Browser interface
```

The SQL migrations are the authoritative schema. Do not duplicate their definitions in documentation.

## Local setup

Install Go compatible with `go.mod`, PostgreSQL, and the `psql` client. Run commands from the repository root.

1. Copy `.env.example` to `.env`, then configure a real local `DATABASE_URL` and a strong `JWT_SECRET`. Keep `.env` out of Git.
2. Create the configured database, then apply all three migrations in order:

   ```sh
   psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f migrations/0001_init.sql
   psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f migrations/0002_seed_sectors.sql
   psql "$DATABASE_URL" -v ON_ERROR_STOP=1 -f migrations/0003_support.sql
   ```

   For these shell commands, export `DATABASE_URL` in the shell; configuring the application's `.env` alone does not export it for `psql`.

3. Resolve the dependencies and start the application:

   ```sh
   go mod tidy
   go run ./cmd/api
   ```

4. Open the configured port in the browser (the existing setup uses http://localhost:8080) and check `/health`.

Use a different `PORT` if another local application already uses 8080.

## API overview

| Area | Routes |
| --- | --- |
| Health | `GET /health` |
| Accounts | `POST /api/users/register`, `POST /api/users/login`, `POST /api/users/reactivate` |
| Discovery | `GET /api/talent`, `GET /api/sectors`, `GET /api/skills`, `GET /api/users/{id}` |
| Personal profile | `GET/PATCH /api/me`, `PUT /api/me/sector`, `PUT /api/me/skills` |
| Experience | `POST /api/me/experience`, `PUT/DELETE /api/me/experience/{id}` |
| Contact requests | `POST /api/talent/{id}/contact-request`, `GET /api/me/contact-requests`, `PATCH /api/me/contact-requests/{id}` |
| Account management | `PATCH /api/me/password`, `POST /api/me/deactivate` |
| Support | `POST /api/support` |

Protected routes require `Authorization: Bearer <token>`. Public registration accepts `talent` and `client`, with administrative accounts excluded from self-service registration.

## Validation and next steps

Before presenting the application as ready for public use, complete a reproducible database-backed build and test pass for registration/login, profile privacy and ownership, search/pagination, experience operations, contact permissions, and account lifecycle. Review authentication-token revocation, abuse controls, and operational configuration. Add deployment and backup documentation once validated.

Useful development commands:

```sh
go test ./...
go vet ./...
go build ./cmd/api
```

The current tree does not include automated test files; a test command reporting no tests is not proof that user flows work.
