(() => {
  const menu = document.querySelector('[data-menu-toggle]');
  const nav = document.getElementById('nav-links');
  const closeMenu = () => { menu.setAttribute('aria-expanded', 'false'); nav.classList.remove('is-open'); };
  menu.hidden = false;
  menu.addEventListener('click', () => {
    const open = menu.getAttribute('aria-expanded') !== 'true';
    menu.setAttribute('aria-expanded', String(open));
    nav.classList.toggle('is-open', open);
  });
  nav.addEventListener('click', (event) => { if (event.target.closest('a')) closeMenu(); });
  document.addEventListener('keydown', (event) => {
    if (event.key === 'Escape' && menu.getAttribute('aria-expanded') === 'true') { closeMenu(); menu.focus(); }
  });
  document.addEventListener('click', (event) => { if (!event.target.closest('.topnav')) closeMenu(); });
  const narrow = matchMedia('(max-width: 760px)');
  narrow.addEventListener('change', closeMenu);

  const themeButton = document.querySelector('[data-theme-toggle]');
  const deviceTheme = matchMedia('(prefers-color-scheme: dark)');
  const currentTheme = () => document.documentElement.dataset.theme || (deviceTheme.matches ? 'dark' : 'light');
  const updateThemeLabel = () => {
    const next = currentTheme() === 'dark' ? 'light' : 'dark';
    themeButton.textContent = next === 'light' ? '☀' : '◐';
    themeButton.setAttribute('aria-label', `Switch to ${next} theme`);
    themeButton.title = `Switch to ${next} theme`;
  };
  themeButton.hidden = false;
  updateThemeLabel();
  deviceTheme.addEventListener('change', updateThemeLabel);
  themeButton.addEventListener('click', () => {
    const theme = currentTheme() === 'dark' ? 'light' : 'dark';
    document.documentElement.dataset.theme = theme;
    try { localStorage.setItem('portfolio-theme', theme); } catch { /* Current-page choice still works. */ }
    updateThemeLabel();
  });

  const filters = document.querySelector('[data-project-filters]');
  if (!filters) return;
  filters.hidden = false;
  const search = document.getElementById('project-search');
  const category = document.getElementById('project-category');
  const cards = [...document.querySelectorAll('[data-project-card]')];
  const count = document.getElementById('project-count');
  const empty = document.getElementById('no-projects');
  const filter = () => {
    const query = search.value.trim().toLowerCase();
    let visible = 0;
    for (const card of cards) {
      card.hidden = !(card.textContent.toLowerCase().includes(query) && (!category.value || card.dataset.category === category.value));
      if (!card.hidden) visible++;
    }
    count.textContent = `${visible} ${visible === 1 ? 'project' : 'projects'}`;
    empty.hidden = visible !== 0;
  };
  search.addEventListener('input', filter);
  category.addEventListener('change', filter);
  document.querySelector('[data-clear-filters]').addEventListener('click', () => { search.value = ''; category.value = ''; filter(); search.focus(); });
  filter();
})();
