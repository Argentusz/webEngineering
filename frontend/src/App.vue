<template>
  <div id="app">
    <mt-header/>
    <router-view class="routerView"/>
  </div>
</template>

<script>

import MtHeader from "@/components/mtHeader.vue";

export default {
  name: "App",
  components: {MtHeader},
}
</script>

<style lang="scss">
@import '@/helpers/variables.scss';
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: white;
}
.routerView {
  height: calc(100vh - $mtHeaderHeight);
}
</style>
