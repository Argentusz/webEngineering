<template>
  <div class="btn-icon-text text-sm">
    <img v-if="icon" :src="require(`@/assets/${icon}.svg`)" alt="" style="height: 16px"/>
    <div v-if="text"
       class="sign"
       :class="{
        'wide': breakPoint === 'wide',
        'def': !breakPoint,
        'short': breakPoint === 'short'
       }"
    >
      {{ text }}
    </div>
  </div>
</template>

<script>
export default {
  name: "ButtonIconText",
  props: {
    icon: {
      type: [String, null],
      default: null
    },
    text: {
      type: [String, null],
      default: null
    },
    // wide, null, short
    breakPoint: {
      type: [String, null],
      default: null
    }
  }
}
</script>

<style lang="scss" scoped>
.btn-icon-text {
  display: flex;
  justify-content: center;
  align-items: center;
  width: fit-content;
}

.sign {
  margin-left: 5px;
}

.wide {
  @media screen and (max-width: 1400px) {
    display: none;
  }
}

.def {
  @media screen and (max-width: 1200px) {
    display: none;
  }
}

.short {
  @media screen and (max-width: 900px) {
    display: none;
  }
}
</style>