import { createRouter, createWebHistory } from 'vue-router';
import Home from '../components/HomeView.vue';
import About from '../components/AboutView.vue';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'Home',
            component: Home
        },
        {
            path: '/about',
            name: 'About',
            component: About
        }
    ],
});

export default router;
