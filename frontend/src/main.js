import {createApp} from 'vue'
import App from './App.vue'
import router from './router';
import './assets/base.min.css';
import './assets/style.css';

createApp(App).use(router).mount('#app');