import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

const config = {
  kit: {
    adapter: adapter({
      pages: 'build',   // html/css/js folder
      assets: 'build',  // assets folder
      fallback: 'index.html', // SPA fallback
    }),
    alias: {
      $components: './src/components',
      $lib: './src/lib'
    }
  },
  preprocess: vitePreprocess()
};

export default config;
