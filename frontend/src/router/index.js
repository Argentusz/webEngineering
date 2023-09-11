import Vue from 'vue'
import VueRouter from 'vue-router'
import SignIn from "@/views/SignIn.vue";
Vue.use(VueRouter)
const routes = [
    {
        path: '/sign-in',
        component: SignIn
    },
]

const router = new VueRouter({
    mode: 'history',
    routes
})

export default router