import { createApp } from 'vue/dist/vue.esm-bundler.js'
import {createRouter, createWebHashHistory} from 'vue-router'
import Pin from "./components/Pin.vue";
import Home from "./components/Home.vue";


const routes = [
    { path: '/', component: Home },
    { path: '/pin', component: Pin },
]


const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

// 5. Create and mount the root instance.
const app = createApp({
}).
   use(router).
   mount('#app')
