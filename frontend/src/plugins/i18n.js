import Vue from 'vue'
import VueI18n from "vue-i18n";

Vue.use(VueI18n)

const msg = {
    ru: {
        signIn: "Войти",
        login: "Логин",
        password: "Пароль",
        name: "Имя",
        nameU: "Название",
        language: "Язык"
    },
    en: {
        signIn: "Sign In",
        login: "Login",
        password: "Password",
        name: "Name",
        nameU: "Name",
        language: "Language"
    }
}

let lcl = localStorage.getItem('lang')
if (lcl === null) {
    lcl = 'ru'
}
export const i18n = new VueI18n({
    locale: lcl,
    fallbackLocale: 'ru',
    messages: msg
})