import { createApp } from 'vue/dist/vue.esm-bundler.js'
import {createRouter, createWebHistory} from 'vue-router'
import Pin from "./components/Pin.vue";
import Home from "./components/Home.vue";
import Tx from "./components/Tx.vue";


const routes = [
    { path: '/', component: Home },
    { path: '/pin', component: Pin },
    { path: '/tx', component: Tx },
]


const router = createRouter({
    history: createWebHistory(),
    routes,
})

// 5. Create and mount the root instance.
const app = createApp({
}).
   use(router).
   mount('#app')
