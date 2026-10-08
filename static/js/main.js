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
  const stage = document.getElementById('project-stage');
  const params = new URLSearchParams(location.search);
  search.value = params.get('q') || '';
  category.value = params.get('category') || '';
  stage.value = params.get('collection') || '';
  const cards = [...document.querySelectorAll('[data-project-card]')];
  const count = document.getElementById('project-count');
  const empty = document.getElementById('no-projects');
  const filter = (updateURL = true) => {
    const query = search.value.trim().toLowerCase();
    let visible = 0;
    for (const card of cards) {
      const inCollection = !stage.value || (stage.value === 'featured' && card.dataset.featured === 'true') || (stage.value === 'public' && card.dataset.visibility === 'Public source') || (stage.value === 'private' && card.dataset.visibility === 'Private source');
      card.hidden = !(card.textContent.toLowerCase().includes(query) && (!category.value || card.dataset.category === category.value) && inCollection);
      if (!card.hidden) visible++;
    }
    count.textContent = `${visible} ${visible === 1 ? 'entry' : 'entries'}`;
    empty.hidden = visible !== 0;
    if (updateURL) {
      const next = new URL(location.href);
      for (const [key, value] of [['q', search.value.trim()], ['category', category.value], ['collection', stage.value]]) {
        if (value) next.searchParams.set(key, value); else next.searchParams.delete(key);
      }
      history.replaceState(null, '', next);
    }
  };
  search.addEventListener('input', () => filter());
  category.addEventListener('change', () => filter());
  stage.addEventListener('change', () => filter());
  window.addEventListener('popstate', () => {
    const restored = new URLSearchParams(location.search);
    search.value = restored.get('q') || '';
    category.value = restored.get('category') || '';
    stage.value = restored.get('collection') || '';
    filter(false);
  });
  document.querySelector('[data-clear-filters]').addEventListener('click', () => { search.value = ''; category.value = ''; stage.value = ''; filter(); search.focus(); });
  filter(false);
})();
