<template>
<div class="header">
  <b-dropdown variant="warning" right>
    <template #button-content>
      <img :src="require('@/assets/language.svg')" style="height: 16px; margin-bottom: 6px" :alt="$t('language')"/>
    </template>
    <b-dropdown-item>Русский 🇷🇺</b-dropdown-item>
    <b-dropdown-item>English 🇬🇧</b-dropdown-item>
  </b-dropdown>
</div>
</template>

<script>
export default {
  name: "mtHeader"
}
</script>

<style lang="scss" scoped>
@import "@/helpers/variables.scss";
.header {
  display: flex;
  flex-direction: row-reverse;
  height: $mtHeaderHeight;
  background-color: $mtYellow;
}
</style>