"""Validate the production portfolio binary: pip install playwright."""
import os
from pathlib import Path
from playwright.sync_api import sync_playwright
base = os.environ.get('PORTFOLIO_URL', 'http://127.0.0.1:8085')
original = ['smartconvert', 'global-browser', 'past-quest-solver', 'campus-resource-manager', 'talentgrid', 'portfolio-site']
with sync_playwright() as p:
    executable = os.environ.get('PLAYWRIGHT_CHROMIUM_EXECUTABLE')
    browser = p.chromium.launch(headless=True, **({'executable_path': executable} if executable else {}), args=['--no-sandbox', '--disable-dev-shm-usage'])
    page = browser.new_page(viewport={'width': 1440, 'height': 1000}, reduced_motion='reduce')
    errors = []
    page.on('pageerror', lambda error: errors.append(str(error)))
    page.goto(base + '/projects')
    page.get_by_label('Find a project or technology').wait_for()
    cards = page.locator('[data-project-card]')
    total = cards.count()
    for slug in original:
        assert page.locator(f'h3 a[href="/projects/{slug}"]').count() == 1, slug
    screenshots = page.locator('.project-screenshot img')
    added = [page.locator('.project-screenshot').nth(i).get_attribute('href').split('/')[-1] for i in range(screenshots.count())]
    assert added, 'No new screenshot-backed projects'
    for slug in added:
        card = cards.filter(has=page.locator(f'h3 a[href="/projects/{slug}"]'))
        title = card.locator('h3').inner_text()
        assert 'in development' in card.inner_text()
        assert card.locator('a[href="https://github.com/miikese/' + slug + '"]').count() == 1
        image = card.locator('img')
        image.scroll_into_view_if_needed()
        image.evaluate('(img) => img.decode()')
        assert image.evaluate('(img) => img.naturalWidth > 0')
        search = page.get_by_label('Find a project or technology')
        search.fill(title)
        assert page.locator('[data-project-card]:visible').count() == 1
        assert 'q=' in page.url
        page.reload()
        assert page.locator('[data-project-card]:visible').count() == 1
        page.get_by_role('button', name='Reset', exact=True).click()
        assert page.locator('[data-project-card]:visible').count() == total
        card.get_by_role('link', name='Explore ' + title, exact=True).click()
        page.get_by_role('heading', name=title, exact=True).wait_for()
        assert page.locator('.detail-screenshot img').count() == 1
        page.goto(base + '/projects')
    page.get_by_label('Collection', exact=True).select_option('private')
    assert page.locator('[data-project-card]:visible').count() == 4
    page.get_by_role('button', name='Reset', exact=True).click()
    page.get_by_label('Find a project or technology').fill('no-matching-project-123')
    assert page.get_by_role('heading', name='No matching projects').is_visible()
    page.get_by_role('button', name='Reset', exact=True).click()
    page.get_by_role('button', name='Switch to dark theme', exact=True).click()
    page.reload()
    assert page.evaluate('document.documentElement.dataset.theme') == 'dark'
    page.get_by_role('button', name='Switch to light theme', exact=True).click()
    for width in [1440, 768, 390, 320]:
        page.set_viewport_size({'width': width, 'height': 844})
        for path in ['/', '/projects', *['/projects/' + slug for slug in added]]:
            page.goto(base + path)
            assert page.evaluate('document.documentElement.scrollWidth <= innerWidth'), (path, width)
    page.set_viewport_size({'width': 390, 'height': 844})
    page.goto(base + '/projects')
    menu = page.get_by_role('button', name='Menu', exact=False)
    menu.click()
    assert menu.get_attribute('aria-expanded') == 'true'
    page.keyboard.press('Escape')
    assert menu.get_attribute('aria-expanded') == 'false'
    Path('docs/screenshots').mkdir(exist_ok=True)
    page.screenshot(path='docs/screenshots/projects-mobile.png', full_page=True)
    page.set_viewport_size({'width': 1440, 'height': 1000})
    page.reload()
    page.screenshot(path='docs/screenshots/projects-desktop.png', full_page=True)
    no_js = browser.new_context(java_script_enabled=False)
    plain = no_js.new_page()
    plain.goto(base + '/projects')
    assert plain.locator('[data-project-card]').count() == total
    assert not errors, errors
    browser.close()
print(f'PASS: {total} project cards ({len(added)} new), original projects preserved, screenshots, source links, search/reset/shareable filters, detail pages, theme persistence, mobile menu, no-JS content, 1440/768/390/320px layouts; no browser errors.')
