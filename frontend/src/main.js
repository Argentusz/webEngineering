import Vue from 'vue'
import App from './App.vue'
import http from "@/plugins/http"
import router from './router'
import {i18n} from './plugins/i18n'
import { BootstrapVue } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.config.productionTip = false
Vue.use(BootstrapVue);
Vue.use(http, {
  baseUrl: "https://localhost:8081"
})
Vue.prototype.$baseUrl = "https://localhost:8081"
Vue.mixin({
  methods: {
    $log(...opts) {
      console.log(...opts)
    }
  }
})
export const url = Vue.prototype.$baseUrl
new Vue({
  i18n,
  router,
  render: h => h(App),
}).$mount('#app')
