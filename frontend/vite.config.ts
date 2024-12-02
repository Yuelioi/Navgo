import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'

import IconsResolver from 'unplugin-icons/resolver'

import { addDynamicIconSelectors } from '@iconify/tailwind'

// https://vitejs.dev/config/
export default defineConfig({
  base: '/',
  plugins: [
    vue(),
    vueJsx(),
    vueDevTools(),
    addDynamicIconSelectors(),
    AutoImport({
      dts: 'auto-imports.d.ts',
      include: [
        /\.[t]sx?$/, // .ts, .tsx
        /\.vue$/,
        /\.vue\?vue/, // .vue
        /\.md$/ // .md
      ],
      eslintrc: {
        enabled: true
      },
      imports: [
        'vue',
        'vue-router',
        {
          pinia: ['defineStore', 'storeToRefs']
        },
        {
          from: './src/models/',
          imports: ['Nav'],
          type: true
        }
      ],

      resolvers: [IconsResolver({})],
      dirs: ['./src/hooks/**', './src/stores/', './src/plugins/*']
    }),
    Components({
      dts: 'components.d.ts',
      dirs: ['src/components/**', 'src/views'],
      extensions: ['vue'],
      include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
      exclude: [/[\\/]node_modules[\\/]/, /[\\/]\.git[\\/]/],
      resolvers: [IconsResolver({})]
    })
  ],

  resolve: {
    alias: {
      vue: 'vue/dist/vue.esm-bundler.js',
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:9200',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  }
})
