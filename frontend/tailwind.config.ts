import type { Config } from 'tailwindcss'
import daisyui from 'daisyui'

import { addDynamicIconSelectors } from '@iconify/tailwind'

export default {
  darkMode: ['selector', '[data-theme="dark"]'],
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: ['./index.html', './src/**/*.{js,ts,vue,css}'],
  theme: {
    extend: {
      gridTemplateColumns: {
        // Simple 16 column grid
        '16': 'repeat(16, minmax(0, 1fr))',
        one: ' minmax(550px, 50%)',
        card: 'repeat(auto-fill, minmax(275px, 1fr))'
      }
    }
  },
  plugins: [daisyui, addDynamicIconSelectors()],
  safelist: ['hover:text-primary', 'hover:text-secondary', 'hover:text-accent'],

  daisyui: {
    themes: [
      {
        light: {
          primary: '#3056d3',
          secondary: '#00aeec',
          accent: '#ff6a9b',
          neutral: '#22212c',
          'neutral-content': '#f8fafc',

          'base-100': '#ebecf0',
          'base-200': '#f6f6f7',
          'base-300': '#ffffff',
          'base-content': '#0f172a',

          info: '#3056d3',
          success: '#00bd8d',
          warning: '#fd621d',
          error: '#f14767'
        }
      },
      {
        dark: {
          primary: '#3056d3',
          secondary: '#4ac7ff',
          accent: '#ff6699',
          neutral: '#a3a7af',
          'neutral-content': '#f8fafc',

          'base-100': '#16181d',
          'base-200': '#2a2e37',
          'base-300': '#2f323c',
          'base-content': '#f8fafc',

          info: '#3056d3',
          success: '#00bd8d',
          warning: '#ffa200',
          error: '#dc4244'
        }
      }
    ]
  }
} satisfies Config
