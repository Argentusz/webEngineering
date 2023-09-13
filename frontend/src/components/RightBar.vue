<template>
  <b-sidebar
      v-model="show"
      id="students-sidebar"
      bg-variant="dark"
      text-variant="white"
      width="50vw"
      right
      @change="$emit('change')"
  >
    <div class="sidebar">
      <div class="grid-panel">
        <controls-bar
            v-model="searchStr"
            breakPoint="wide"
            :disableEdit="!selectedStudent"
            :disableDelete="!selectedStudent"
            showDelete
            showAdd
            @delete="$emit('delete', selectedStudent); selectedStudent = null"
            @add="$emit('add')"
            @edit="$emit('edit', selectedStudent)"
            @refresh="$emit('refresh')"
        >
          <b-button class="actions-btn text-sm" variant="warning" @click="$emit('invite', selectedStudent); selectedStudent = null">
            <button-icon-text :text="$t('invite')" icon="invite-to-group"/>
          </b-button>
        </controls-bar>
        <ag-grid-vue
            style="width: calc(50vw - 60px); height: calc(100vh - 100px); margin-top: 10px"
            :rowData="rowData"
            :isExternalFilterPresent="() => this.searchStr !== ''"
            :doesExternalFilterPass="doesExternalFilterPass"
            :columnDefs="columnDefs"
            class="ag-theme-balham-dark grid"
            rowSelection="single"
            animateRows
            @selectionChanged="selectionChangedHandler"
            @gridReady="onGridReady"
        />
      </div>
    </div>
  </b-sidebar>

</template>

<script>
import ControlsBar from "@/components/ControlsBar.vue";
import ButtonIconText from "@/components/ButtonIconText.vue";

export default {
  name: "RightBar",
  components: {ButtonIconText, ControlsBar},
  data() {
    return {
      show: this.value,
      selectedStudent: null,
      searchStr: "",
      gridApi: null
    }
  },
  props: {
    value: {
      type: Boolean,
      required: true
    },
    rowData: {
      type: [Array, null],
      default: null
    },
    columnDefs: {
      type: Array,
      default: () => []
    }
  },
  watch: {
    show() {
      this.$emit("input", this.show)
    },
    value() {
      this.show = this.value
    },
    searchStr() {
      if (this.gridApi) {
        this.gridApi.onFilterChanged()
      }
    }
  },
  methods: {
    selectionChangedHandler() {
      const selectedRows = this.gridApi.getSelectedRows();
      this.selectedStudent = selectedRows[0];
    },
    onGridReady(params) {
      this.gridApi = params.api
    },
    doesExternalFilterPass(el) {
      return `${el.data.Name} ${el.data.Exam}`.toLowerCase().includes(this.searchStr.toLowerCase())
    }
  }

}
</script>

<style scoped>

</style>