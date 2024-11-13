import type { Config } from 'tailwindcss'
import daisyui from 'daisyui'

import { addDynamicIconSelectors } from '@iconify/tailwind'

export default {
  darkMode: ['selector', '[data-theme="dark"]'],
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: ['./index.html', './src/**/*.{js,ts,vue,css}'],
  theme: {
    extend: {}
  },
  plugins: [daisyui, addDynamicIconSelectors()],
  safelist: ['hover:text-primary', 'hover:text-secondary', 'hover:text-accent'],

  daisyui: {
    themes: [
      {
        light: {
          primary: '#3056d3',
          secondary: '#00aeec',
          accent: '#ff6699',
          neutral: '#a3a7af',
          'neutral-content': '#f8fafc',

          'base-100': '#f3f3f3',
          'base-200': '#f7f8f9',
          'base-300': '#ffffff',
          'base-content': '#0f172a',

          info: '#3056d3',
          success: '#00bd8d',
          warning: '#ffa200',
          error: '#ff6a3d'
        }
      },
      {
        dark: {
          primary: '#3056d3',
          secondary: '#4ac7ff',
          accent: '#ff6699',
          neutral: '#22212c',
          'neutral-content': '#f8fafc',

          'base-100': '#16181d',
          'base-200': '#2a2e37',
          'base-300': '#2f323c',
          'base-content': '#f8fafc',

          info: '#3056d3',
          success: '#00bd8d',
          warning: '#ffa200',
          error: '#ff6a3d'
        }
      }
    ]
  }
} satisfies Config
