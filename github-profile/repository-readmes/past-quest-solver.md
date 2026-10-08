# Smile Learning

A responsive study interface for exploring past exam questions and answers. Built with React and TypeScript, with reusable question cards, subject views, and search and filtering controls.

## Project status

The repository includes static and mock study data. This is a study-interface project; the README does not imply a verified exam database, professional answer review, or production usage.

## Features

- Question cards and subject-focused navigation.
- Search and filters for finding relevant study material.
- Responsive layout for desktop and mobile screens.
- Reusable UI components built with shadcn/ui and Tailwind CSS.

## Technology

React 18, TypeScript, Vite, Tailwind CSS, shadcn/ui, and React Router.

## Run locally

Requires Node.js and npm.

```sh
git clone https://github.com/miikese/past-quest-solver.git
cd past-quest-solver
npm install
npm run dev
```

Open the local URL printed by Vite.

## Development commands

```sh
npm run build    # production build
npm run preview  # preview the build locally
npm run lint     # static code checks
```

The commands above are defined in `package.json`. This documentation refresh does not certify that they pass in every environment.

## Structure

- `src/components/`: reusable interface components, including question and search controls.
- `src/pages/`: route-level pages.
- `src/data/`: static and mock study material.
- `src/hooks/` and `src/lib/`: shared interface logic and utilities.

## Author

**Michael Ikese Emmanuel** — software engineer and programmer, project manager in training, and cybersecurity enthusiast.

[GitHub @miikese](https://github.com/miikese) · [Engineering portfolio](https://github.com/miikese/portfolio-site) · [Email](mailto:emmanlemichel2019@gmail.com)
