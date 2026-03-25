import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import fs from 'node:fs'
import path from 'node:path'

/**
 * Cache proxy middleware as mandated by GEMINI.md.
 * This ensures that any external API data (like Celestrak or JPL horizons) 
 * is cached locally for 2 hours to respect rate limits.
 */
export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
    {
      name: 'api-cache-proxy',
      configureServer(server) {
        server.middlewares.use(async (req, res, next) => {
          // Catch any API requests that should be cached
          // Example pattern: /api-proxy/provider/endpoint?params
          if (req.url && req.url.startsWith('/api-proxy/')) {
            const url = new URL(req.url, `http://${req.headers.host}`);
            const cacheKey = Buffer.from(url.pathname + url.search).toString('hex').slice(0, 16);
            const cacheDir = path.resolve(process.cwd(), 'cache');
            
            if (!fs.existsSync(cacheDir)) fs.mkdirSync(cacheDir, { recursive: true });
            const cachePath = path.join(cacheDir, `${cacheKey}.json`);

            const CACHE_DURATION = 2 * 60 * 60 * 1000; // 2 hours
            
            if (fs.existsSync(cachePath)) {
              const stats = fs.statSync(cachePath);
              const age = Date.now() - stats.mtimeMs;
              if (age < CACHE_DURATION) {
                const data = fs.readFileSync(cachePath, 'utf8');
                res.setHeader('Content-Type', 'application/json');
                res.end(data);
                return;
              }
            }

            // For now, this is a placeholder as solar-system-now currently uses 
            // static internal data from NASA JPL (kepler.ts).
            // This middleware is "standard-equipped" per GEMINI.md instructions.
            next();
            return;
          }
          next();
        });
      }
    }
  ],
})
