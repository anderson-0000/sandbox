import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import fs from 'fs';
import path from 'path';
import axios from 'axios';

// 2時間キャッシュ
const CACHE_TTL = 2 * 60 * 60 * 1000;
const CACHE_DIR = path.resolve(__dirname, 'cache');

if (!fs.existsSync(CACHE_DIR)) {
  fs.mkdirSync(CACHE_DIR);
}

// CelesTrak API をキャッシュするミドルウェア
const celestrakCachePlugin = () => ({
  name: 'celestrak-cache',
  configureServer(server: any) {
    server.middlewares.use(async (req: any, res: any, next: any) => {
      if (req.url?.startsWith('/api/celestrak/')) {
        const query = req.url.replace('/api/celestrak/', '');
        const cachePath = path.join(CACHE_DIR, `${query.replace(/\//g, '_')}.txt`);

        // キャッシュチェック
        if (fs.existsSync(cachePath)) {
          const stats = fs.statSync(cachePath);
          if (Date.now() - stats.mtimeMs < CACHE_TTL) {
            const data = fs.readFileSync(cachePath, 'utf-8');
            res.setHeader('Content-Type', 'text/plain');
            res.end(data);
            return;
          }
        }

        // API リクエスト
        try {
          const targetUrl = `https://celestrak.org/NORAD/elements/gp.php?${query}`;
          const response = await axios.get(targetUrl);
          fs.writeFileSync(cachePath, response.data);
          res.setHeader('Content-Type', 'text/plain');
          res.end(response.data);
        } catch (error) {
          console.error('CelesTrak API Error:', error);
          res.statusCode = 500;
          res.end('Internal Server Error');
        }
        return;
      }
      next();
    });
  },
});

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), celestrakCachePlugin()],
});
