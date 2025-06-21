/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    colors: {
      'background-primary': '#ffffff',
      'background-secondary': '#f0f0f0',
      'background-tertiary': '#f5f5f5',
      'background-inv-primary': '#252525',
      'border-primary': '#d4d4d4',
      'border-secondary': '#eaeaea',
      'brand-primary': '#38a3ee',
      'brand-secondary': '#2167a3',
      'brand-light-primary': '#fff3f3',
      'text-primary': '#2a2a2a',
      'text-secondary': '#4a4a4a',
      'text-tertiary': '#9a9a9a',
      'text-inv-primary': '#f0f0f0',
      'status-error': '#e02a2a',
      red: '#ff0000',
      white: '#ffffff',
      transparent: 'transparent'
    },
    fontFamily: {
      primary: ['Open Sans', 'Noto Sans JP', 'sans-serif']
    },
    extend: {
      spacing: {
        75: '300px',
      },
      borderRadius: {
        15: '15px'
      },
      height: {
        'header-offset': 'calc(100vh - 56px)'
      },
      maxWidth: {
        'form-max': '600px',
        'profile-max': '500px'
      },
      padding: {
        'container-x': '120px'
      }
    }
  },
  plugins: []
}
