import { defineConfig, presetWind, presetTypography, presetIcons } from 'unocss'

export default defineConfig({
  presets: [
    presetWind({
      dark: 'class',
    }),
    presetTypography(),
    presetIcons(),
  ],

  theme: {
    colors: {
      brand: {
        50: '#eef4ff',
        100: '#d9e3ff',
        200: '#b7c8ff',
        300: '#94adff',
        400: '#6a8bff',
        500: '#3b6aff',
        600: '#2854db',
        700: '#1e40af',
        800: '#1a358a',
        900: '#172c6e',
      },
    },
  },

  shortcuts: {
    'btn-primary':
      'px-4 py-2 rounded-lg bg-blue-600 text-white hover:bg-blue-700 transition shadow-sm hover:shadow-md font-medium dark:(bg-blue-500 hover:bg-blue-600)',
    'card':
      'bg-white dark:bg-gray-800 rounded-2xl shadow-md p-6 border border-gray-100 dark:border-gray-700 hover:shadow-lg transition',
  },
})