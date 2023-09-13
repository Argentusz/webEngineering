<template>
<div class="header">
  <div class="header__btns">
    <b-button v-if="loggedIn" variant="warning" class="text-sm">
      <button-icon-text breakPoint="short" :text="name" icon="user-settings"/>
    </b-button>
    <b-dropdown variant="warning" style="height: 32px" right no-caret>
      <template #button-content>
        <button-icon-text :text="$t('language')" breakPoint="short" icon="language"/>
      </template>
      <b-dropdown-item @click="changeLangHandler('ru')">Русский 🇷🇺</b-dropdown-item>
      <b-dropdown-item @click="changeLangHandler('en')">English 🇬🇧</b-dropdown-item>
    </b-dropdown>
    <b-button
        v-if="loggedIn"
        variant="danger"
        style="height: 32px"
        class="text-sm"
        @click="exitHandler"
    >
      {{ $t('exit') }}
    </b-button>
  </div>
  <div v-if="loggedIn" style="align-self: center" class="header__btns">
    <b>{{ $t('controlsTitle') }} {{ name }}</b>
  </div>
</div>
</template>

<script>
import ButtonIconText from "@/components/ButtonIconText.vue";

export default {
  name: "mtHeader",
  components: {ButtonIconText},
  data() {
    return {
      name: ""
    }
  },
  created() {
    this.name = localStorage.getItem("name")
  },
  computed: {
    loggedIn() {
      return this.checkLoggedIn()
    }
  },
  methods: {
    exitHandler() {
      localStorage.clear()
      this.$router.push("/sign-in").catch(() => {})
    },
    changeLangHandler(lang) {
      this.$i18n.locale = lang
      localStorage.setItem("lang", lang)
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/helpers/variables.scss";
.header {
  display: flex;
  flex-direction: row-reverse;
  justify-content: space-between;
  height: $mtHeaderHeight;
  background-color: $mtYellow;
}
.header__btns {
  margin: 5px;
}
</style>