import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import topLevelAwait from 'vite-plugin-top-level-await'
import fs from 'node:fs'
import path from 'node:path'
import https from 'node:https'

export default defineConfig({
  plugins: [
    react(), 
    topLevelAwait(),
    {
      name: 'tle-cache-proxy',
      configureServer(server) {
        server.middlewares.use(async (req, res, next) => {
          if (req.url && req.url.startsWith('/api-celestrak')) {
            const url = new URL(req.url, `http://${req.headers.host}`);
            const group = url.searchParams.get('GROUP') || 'active';
            
            const cacheDir = path.resolve(process.cwd(), 'cache');
            if (!fs.existsSync(cacheDir)) fs.mkdirSync(cacheDir, { recursive: true });
            const cachePath = path.join(cacheDir, `${group}.tle`);

            const CACHE_DURATION = 2 * 60 * 60 * 1000; // 2 hours
            let shouldFetch = true;

            if (fs.existsSync(cachePath)) {
              const stats = fs.statSync(cachePath);
              const age = Date.now() - stats.mtimeMs;
              if (age < CACHE_DURATION) {
                shouldFetch = false;
                console.log(`--- [CACHE HIT] Serving ${group} from local file (${Math.round(age/60000)}m old) ---`);
              } else {
                console.log(`--- [CACHE EXPIRED] ${group} is ${Math.round(age/60000)}m old, trying to refresh ---`);
              }
            }

            if (!shouldFetch) {
              const data = fs.readFileSync(cachePath, 'utf8');
              res.setHeader('Content-Type', 'text/plain');
              res.end(data);
              return;
            }

            // Fetch fresh data
            const targetUrl = `https://celestrak.org/NORAD/elements/gp.php?GROUP=${group}&FORMAT=tle`;
            console.log(`--- [FETCHING] Requesting fresh data from Celestrak: ${group} ---`);

            const options = {
              headers: {
                'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36'
              }
            };

            https.get(targetUrl, options, (remoteRes) => {
              if (remoteRes.statusCode === 200) {
                let body = '';
                remoteRes.on('data', chunk => body += chunk);
                remoteRes.on('end', () => {
                  if (body.length > 100) {
                    fs.writeFileSync(cachePath, body);
                    console.log(`--- [CACHE UPDATED] Successfully saved ${group} (${body.length} bytes) ---`);
                    res.setHeader('Content-Type', 'text/plain');
                    res.end(body);
                  } else {
                    console.warn(`--- [FETCH FAILED] Data too short, likely error page ---`);
                    serveOldCache(res, cachePath, next);
                  }
                });
              } else {
                console.error(`--- [FETCH FAILED] Celestrak returned ${remoteRes.statusCode} ---`);
                serveOldCache(res, cachePath, next);
              }
            }).on('error', (err) => {
              console.error(`--- [FETCH ERROR] ${err.message} ---`);
              serveOldCache(res, cachePath, next);
            });
            return;
          }
          next();
        });
      }
    }
  ],
  server: {
    // 既存の proxy 設定は不要になるので削除（ミドルウェアが処理するため）
  }
})

function serveOldCache(res: any, cachePath: string, next: any) {
  if (fs.existsSync(cachePath)) {
    console.log(`--- [FALLBACK] Serving old cache as fallback ---`);
    const data = fs.readFileSync(cachePath, 'utf8');
    res.setHeader('Content-Type', 'text/plain');
    res.end(data);
  } else {
    next();
  }
}
