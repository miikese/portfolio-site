# Portrait and presentation refresh — 9 October 2026

The homepage now introduces Michael by name and uses an edited professional portrait based on his original photo, with softer lighting and a navy backdrop. The original photo was removed from the served site. Visitors can pause motion and use either color theme. Reduced-motion preferences disable animation automatically. The bio focuses on implemented projects and current studies, with project management and cybersecurity described as developing interests.

Global Browser appears once as the project spotlight; the other three projects follow it. Social preview metadata uses a static PNG banner instead of SVG. The GitHub profile README includes the portrait, a clearer introduction, selected projects, a toolkit, and education. The one-page CV includes the National Diploma in Civil Engineering Technology from Federal Polytechnic, Nasarawa and the user-provided referee. Account-avatar instructions point to the same image.

Validation: Go race tests pass with 77.6% statement coverage; `go vet`, the production build, and `git diff --check` pass. Native Chrome checks cover 25 page/viewport combinations from 320 to 1440 pixels, absence of the original-photo control, motion pause, reduced-motion preferences, project filters, theme persistence, mobile-menu Escape handling, and navigation with JavaScript disabled.

The local preview uses port 8081 because port 8080 is occupied by Global Browser. No public hosting deployment was requested. Connected GitHub tools support README and asset changes, but do not expose account bio or avatar updates, repository visibility, or profile pins. The username repository is private, so its README does not currently appear on the public profile.

See `portrait-prompt.md` for the image-generation prompt and asset provenance.
