# Smile Student

An early static study website with English and mathematics questions, sample answers, and simple filtering. This prototype documents the learning steps that preceded the React-based [Smile Learning](https://github.com/miikese/past-quest-solver) interface.

## Features

- Separate English and mathematics question pages.
- Keyword search and year filters.
- Answer controls that toggle their accessible expanded state.
- Sample study content held directly in the page scripts.

## Stack and structure

HTML, CSS, and vanilla JavaScript. `index.html` is the entry page; `english.html` and `math.html` contain the subject views, and `styles.css` provides shared styling.

## Preview locally

With Python 3 installed:

```sh
git clone https://github.com/miikese/Smile-student-.git
cd Smile-student-
python3 -m http.server 8000
```

Open http://localhost:8000. No frontend package installation is required.

## Status

Learning prototype with sample content. It is separate from the newer React application and is not presented as a production exam-content service.

## Author

**Michael Ikese Emmanuel** · [@miikese](https://github.com/miikese) · [Engineering portfolio](https://github.com/miikese/portfolio-site)
