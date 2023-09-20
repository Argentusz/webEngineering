import Vue from "vue"
import VueI18n from "vue-i18n";

Vue.use(VueI18n)

const msg = {
    ru: {
        signIn: "Войти",
        signUp: "Зарегистрироваться",
        login: "Логин",
        password: "Пароль",
        name: "Имя",
        nameU: "Название",
        language: "Язык",
        exit: "Выйти",
        examDate: "Дата экзамена",
        examAud: "Аудитория экзамена",
        controlsTitle: "Панель управления",
        find: "Найти",
        refresh: "Обновить",
        add: "Добавить",
        edit: "Изменить",
        delete: "Удалить",
        kick: "Отклонить",
        exam: "Экзамен",
        students: "Абитуриенты",
        more: "Ещё",
        invite: "Пригласить",
        save: "Сохранить",
        cancel: "Отменить",
        newStudent: "Новый абитуриент",
        changeStudent: "Изменить студента",
        newFaculty: "Новый факультет",
        changeFaculty: "Изменить факультет",
        settings: "Настройки",
        noAcc: "Нет аккаунта",
        yesAcc: "Есть аккаунт",
        error: "Ошибка",
        errWrongLoginPassword: "Неверный логин или пароль"
    },
    en: {
        signIn: "Sign In",
        signUp: "Sign Up",
        login: "Login",
        password: "Password",
        name: "Name",
        nameU: "Name",
        language: "Language",
        exit: "Exit",
        examDate: "Exam date",
        examAud: "Exam auditory",
        controlsTitle: "Controls panel",
        find: "Find",
        refresh: "Refresh",
        add: "Add",
        edit: "Edit",
        delete: "Delete",
        kick: "Decline",
        exam: "Exam",
        students: "Students",
        more: "More",
        invite: "Invite",
        save: "Save",
        cancel: "Cancel",
        newStudent: "New student",
        changeStudent: "Change student",
        newFaculty: "New faculty",
        changeFaculty: "Change faculty",
        settings: "Settings",
        noAcc: "Don`t have account",
        yesAcc: "Already have account",
        error: "Error",
        errWrongLoginPassword: "Wrong login or password"
    }
}

let lcl = localStorage.getItem("lang")
if (lcl === null) {
    lcl = "ru"
}
export const i18n = new VueI18n({
    locale: lcl,
    fallbackLocale: "ru",
    messages: msg
})