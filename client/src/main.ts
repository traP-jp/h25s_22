import './assets/main.css'
import { GoogleMap } from 'vue3-google-map'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(router)
app.use(GoogleMap.default, {
  load: {
    key: 'YOUR_Maps_API_KEY',
  },
})
app.mount('#app')
