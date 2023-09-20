import Vue from "vue"
import VueRouter from "vue-router"
import SignIn from "@/views/SignIn.vue";
import FacultiesGridView from "@/views/facultiesGridView.vue";
import SignUp from "@/views/SignUp.vue";
Vue.use(VueRouter)
const routes = [
    {
        path: "/sign-in",
        component: SignIn
    },
    {
        path: "/sign-up",
        component: SignUp
    },
    {
        path: "/",
        component: FacultiesGridView
    }
]

const router = new VueRouter({
    mode: "history",
    routes
})

export default router