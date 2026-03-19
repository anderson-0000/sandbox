import puppeteer from 'puppeteer';

(async () => {
  const browser = await puppeteer.launch({
    executablePath: '/Users/sho/.cache/puppeteer/chrome-headless-shell/mac-145.0.7632.77/chrome-headless-shell-mac-x64/chrome-headless-shell',
    args: ['--no-sandbox', '--disable-setuid-sandbox']
  });
  const page = await browser.newPage();

  // コンソールログをキャプチャ
  page.on('console', msg => console.log('BROWSER LOG:', msg.text()));
  page.on('pageerror', err => console.error('BROWSER ERROR:', err.message));

  try {
    console.log('Navigating to http://localhost:5173/ ...');
    await page.goto('http://localhost:5173/', { waitUntil: 'networkidle0', timeout: 5000 });
    
    // HTML構造を確認
    const content = await page.content();
    console.log('HTML Length:', content.length);
    
    const rootHasContent = await page.evaluate(() => {
      const root = document.getElementById('root');
      return root ? root.innerHTML.length > 0 : false;
    });
    console.log('Root has content:', rootHasContent);

    // 画面が真っ白か、何か表示されているかを確認
  } catch (e) {
    console.error('Navigation failed:', e.message);
  } finally {
    await browser.close();
  }
})();
