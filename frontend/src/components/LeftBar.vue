<template>
  <b-sidebar
      id="faculty-sidebar"
      bg-variant="dark"
      text-variant="white"
      width="50vw"
      @change="$emit('change')"
  >
    <div class="sidebar">
      <div class="grid-panel">
        <controls-bar
          v-model="searchStr"
          breakPoint="wide"
          :disableEdit="!selectedStudent"
          :disableDelete="!selectedStudent"
          :disableKick="!selectedStudent"
          showDelete
          showKick
          @refresh="$emit('refresh')"
          @edit="$emit('edit', selectedStudent)"
          @delete="$emit('delete', selectedStudent); selectedStudent = null"
          @kick="$emit('kick', selectedStudent); selectedStudent = null"
        >
          <b-button
              class="actions-btn text-sm"
              variant="warning"
              @click="$emit('clickMore')"
          >
            <button-icon-text :text="`${$t('more')}...`" icon="group-add"/>
          </b-button>
        </controls-bar>
        <ag-grid-vue
            style="width: calc(50vw - 60px); height: calc(100vh - 100px); text-align: left; margin-top: 10px"
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
  name: "LeftBar",
  components: {ButtonIconText, ControlsBar},
  data() {
    return {
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