import Vue from "vue"
import App from "./App.vue"
import http from "@/plugins/http"
import router from "./router"
import {i18n} from "./plugins/i18n"
import { BootstrapVue } from "bootstrap-vue"
import VueCookies from "vue-cookies"
import "bootstrap/dist/css/bootstrap.css"
import "bootstrap-vue/dist/bootstrap-vue.css"
import { AgGridVue } from "ag-grid-vue"
import "ag-grid-community/styles/ag-grid.css";
import "ag-grid-community/styles/ag-theme-balham.css";
import "@/helpers/styles.scss"
import "@/helpers/variables.scss"

Vue.config.productionTip = true
Vue.use(VueCookies);
Vue.component("ag-grid-vue", AgGridVue)
Vue.use(BootstrapVue);

Vue.use(http, {
  baseUrl: "https://localhost:8081"
})
Vue.prototype.$baseUrl = "https://localhost:8081"
Vue.mixin({
  data() {
    return {
      username: null,
    }
  },
  created() {
    this.username = localStorage.getItem("name")
  },
  methods: {
    setUsername(username) {
      this.username = username
    },
    checkLoggedIn() {
      return localStorage.getItem("login") && localStorage.getItem("password") &&
          localStorage.getItem("uid") && localStorage.getItem("name")
    },
    addBasicAuth(obj) {
      const login = localStorage.getItem("login")
      const password = localStorage.getItem("password")
      if (login == null || password == null) {
        return null
      }
      const base64auth = btoa(`${login}:${password}`)
      return Object.assign(obj, {
        headers: {Authorization: `Basic ${base64auth}`}
      })
    },
    $log(...opts) {
      console.log(...opts)
    },
    $error(error, title = this.$t("error")) {
      this.$bvToast.toast(error, {
        title: title,
        variant: "danger",
        solid: true,
        toaster: "b-toaster-bottom-right"
      })
    }
  }
})
export const url = Vue.prototype.$baseUrl
new Vue({
  i18n,
  router,
  render: h => h(App),
}).$mount("#app")
