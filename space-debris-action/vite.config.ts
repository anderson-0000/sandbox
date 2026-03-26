import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import fs from 'fs';
import path from 'path';
import axios from 'axios';

const CACHE_TTL = 2 * 60 * 60 * 1000;
const CACHE_DIR = path.resolve(__dirname, 'cache');

if (!fs.existsSync(CACHE_DIR)) {
  fs.mkdirSync(CACHE_DIR);
}

const celestrakCachePlugin = () => ({
  name: 'celestrak-cache',
  configureServer(server: any) {
    server.middlewares.use(async (req: any, res: any, next: any) => {
      if (req.url?.startsWith('/api/celestrak/')) {
        const query = req.url.replace('/api/celestrak/', '');
        const cachePath = path.join(CACHE_DIR, `${query.replace(/\//g, '_')}.txt`);

        if (fs.existsSync(cachePath)) {
          const stats = fs.statSync(cachePath);
          if (Date.now() - stats.mtimeMs < CACHE_TTL) {
            const data = fs.readFileSync(cachePath, 'utf-8');
            res.setHeader('Content-Type', 'text/plain; charset=utf-8');
            res.end(data);
            return;
          }
        }

        try {
          // gp.php よりも確実に TLE を取得できる stations.txt 等のフォーマットを使用
          // GROUP パラメータに合わせた URL 構築
          const groupName = query.match(/GROUP=([^&]+)/)?.[1] || 'active';
          const targetUrl = `https://celestrak.org/NORAD/elements/gp.php?GROUP=${groupName}&FORMAT=tle`;
          
          console.log(`Fetching from CelesTrak: ${targetUrl}`);
          const response = await axios.get(targetUrl);
          
          fs.writeFileSync(cachePath, response.data);
          res.setHeader('Content-Type', 'text/plain; charset=utf-8');
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

export default defineConfig({
  plugins: [react(), celestrakCachePlugin()],
});
