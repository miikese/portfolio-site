// Run before the stylesheet to avoid a flash when a theme has been saved.
(() => {
  document.documentElement.classList.add('js');
  try {
    const theme = localStorage.getItem('portfolio-theme');
    if (theme === 'light' || theme === 'dark') document.documentElement.dataset.theme = theme;
  } catch { /* Device theme works when browser storage is unavailable. */ }
})();
