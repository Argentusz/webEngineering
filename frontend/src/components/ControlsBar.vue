<template>
  <div class="controls-bar">
    <b-input v-model="searchStr" :placeholder="$t('find')" class="finder text-sm" />
    <div class="controls-bar__all-btns">
      <b-button-group class="controls-bar__btn-group">
        <b-button v-if="showRefresh" class="controls-bar__btn text-sm" variant="warning" @click="$emit('refresh')">
          <button-icon-text :breakPoint="breakPoint" :text="$t('refresh')" icon="refresh"/>
        </b-button>
        <b-button v-if="showAdd" class="controls-bar__btn text-sm" variant="warning" @click="$emit('add')">
          <button-icon-text :breakPoint="breakPoint" :text="$t('add')" icon="add"/>
        </b-button>
        <b-button v-if="showEdit" :disabled="disableEdit" variant="warning" class="controls-bar__btn text-sm" @click="$emit('edit')">
          <button-icon-text :breakPoint="breakPoint" :text="$t('edit')" icon="edit"/>
        </b-button>
        <b-button v-if="showDelete" :disabled="disableDelete" variant="danger" class="controls-bar__btn text-sm" @click="$emit('delete')">
          <button-icon-text :breakPoint="breakPoint" :text="$t('delete')" icon="delete"/>
        </b-button>
        <b-button v-if="showKick" :disabled="disableKick" variant="danger" class="controls-bar__btn text-sm" @click="$emit('kick')">
          <button-icon-text :breakPoint="breakPoint" :text="$t('kick')" icon="kick"/>
        </b-button>
      </b-button-group>
      <slot/>
    </div>
  </div>
</template>

<script>
import ButtonIconText from "@/components/ButtonIconText.vue";

export default {
  name: "ControlsBar",
  components: {ButtonIconText},
  data() {
    return {
      searchStr: ""
    }
  },
  props: {
    value: {
      type: String,
      default: ""
    },
    showRefresh: {
      type: Boolean,
      default: true
    },
    showAdd: {
      type: Boolean,
      default: false
    },
    showDelete: {
      type: Boolean,
      default: false
    },
    disableDelete: {
      type: Boolean,
      default: false,
    },
    showEdit: {
      type: Boolean,
      default: true
    },
    disableEdit: {
      type: Boolean,
      default: false
    },
    showKick: {
      type: Boolean,
      default: false
    },
    disableKick: {
      type: Boolean,
      default: false
    },
    // wide, null, short
    breakPoint: {
      type: [String, null],
      default: null
    }
  },
  created() {
    if (this.value) {
      this.searchStr = this.value
    }
  },
  watch: {
    searchStr() {
      this.$emit("input", this.searchStr)
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/helpers/variables.scss";
.finder {
  width: 100%;
  @media screen and (min-width: 901px) {
    max-width: 400px;
  }
}
.controls-bar {
  display: flex;
  align-items: center;
  flex-wrap: nowrap;
  @media screen and (max-width: 900px) {
    flex-wrap: wrap;
  }
}

.controls-bar:first-line {
  margin-bottom: 15px;
}

.controls-bar__btn {
  height: 32px;
  width: fit-content;
  border: 1px solid $mtBlack;
  @media screen and (min-width: 901px) {
    margin-left: 10px;
  }
}

.controls-bar__all-btns {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  @media screen and (max-width: 900px) {
    margin-top: 15px;
  }
}

</style>