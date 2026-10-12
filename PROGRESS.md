# Portfolio recovery checkpoint — 12 October 2026

- Last completed task: recovered the original portfolio, confirmed all nine cards
  and the three project source/demo links, and refreshed actual app screenshots.
- Current phase: tested recovery and local desktop handoff.
- Implemented: a Linux launcher builds missing previews, starts all four separate
  apps, opens their URLs, reuses healthy servers, and can stop its own processes.
  Benue's description now records reliable discovery through moderation backlogs.
- Latest verified source before this checkpoint:
  `4ab1d99fd3a4a004ed5698d959e3a88e586689ec`. Resolve the checkpoint with
  `git rev-parse HEAD`; final remote/CI outcomes are in the workspace recovery report.
- Tests: Go vet, race-enabled tests (77.6% statement coverage), production binary
  and Playwright passed. Nine cards, all six original projects, screenshots, links,
  search/filters, theme persistence, mobile menu and no-JavaScript content verified.
  Home, collection and three new detail pages pass at 320/390/768/1440px.
- Deployment: GitHub has no Pages site, deployment records or configured public
  homepage. No public portfolio deployment is claimed. All three linked apps have
  verified restricted Sites deployments. Local portfolio: http://127.0.0.1:8085.
- Remaining features: public portfolio hosting requires a chosen Go/container host.
- Next task: keep project details current as the independent product plans advance.
- Reopen: `python3 scripts/open-local-projects.py`; rebuild with `--rebuild`,
  stop owned previews with `--stop`. Logs and process records: `.local-state/`.
