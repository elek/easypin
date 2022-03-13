import {createApp} from 'vue/dist/vue.esm-bundler.js'
import {createRouter, createWebHistory} from 'vue-router'
import Pin from "./components/Pin.vue";
import Home from "./components/Home.vue";
import Tx from "./components/Tx.vue";
import Block from "./components/Block.vue";
import Blocks from "./components/Blocks.vue";


const routes = [
    {path: '/', component: Home},
    {path: '/pin/:hash', component: Pin, props: true},
    {path: '/pin', component: Pin},
    {path: '/tx', component: Tx},
    {path: '/block/:hash', component: Block, props: true},
    {path: '/blocks', component: Blocks},
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
