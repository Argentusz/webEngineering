<template>
<div class="facultiesGridView">
<!--  Модальные окна  -->
  <faculty-new-edit-modal
    :changeFaculty="facultyForChange"
    @hide="facultyForChange = null"
    @refresh="getData"
  />
  <student-new-edit-modal
    :changeStudent="studentForChange"
    @hide="studentForChange = null"
    @refresh="sidebarRefreshHandler"
  />
  <div class="grid-panel">
    <controls-bar
      v-model="searchStr"
      :disableDelete="!selectedFID"
      :disableEdit="!selectedFID"
      breakPoint="short"
      showAdd
      showDelete
      @refresh="getData"
      @add="addHandler"
      @edit="editHandler"
      @delete="deleteHandler"
    >
      <b-button
        :disabled="!selectedFID"
        class="actions-btn text-sm"
        variant="warning"
        v-b-toggle.faculty-sidebar
        @click="getStudentsByFID"
      >
        <button-icon-text :text="$t('students')" icon="group" breakPoint="short"/>
      </b-button>
    </controls-bar>
    <ag-grid-vue
      :rowData="faculties"
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
  <left-bar
    v-model="showLeftBar"
    :rowData="facultyStudents"
    :columnDefs="facultyStudentsColumnDefs"
    @refresh="sidebarRefreshHandler"
    @edit="editStudentHandler"
    @delete="deleteStudentHandler"
    @kick="kickHandler"
    @clickMore="clickMoreHandler"
    @change="() => {facultyStudents = null; otherStudents = null; showRightBar = false}"
  />
  <right-bar
    v-model="showRightBar"
    :rowData="otherStudents"
    :columnDefs="facultyStudentsColumnDefs"
    @refresh="sidebarRefreshHandler"
    @add="addStudentHandler"
    @edit="editStudentHandler"
    @delete="deleteStudentHandler"
    @invite="inviteHandler"
    @change="() => {otherStudents = null}"
  />
</div>
</template>

<script>
import {url} from "@/main";
import ControlsBar from "@/components/ControlsBar.vue";
import LeftBar from "@/components/LeftBar.vue";
import RightBar from "@/components/RightBar.vue";
import ButtonIconText from "@/components/ButtonIconText.vue";
import FacultyNewEditModal from "@/components/Modals/FacultyNewEditModal.vue";
import StudentNewEditModal from "@/components/Modals/StudentNewEditModal.vue";
export default {
  name: "FacultiesGridView",
  components: {StudentNewEditModal, FacultyNewEditModal, ButtonIconText, RightBar, LeftBar, ControlsBar},
  data() {
    return {
      faculties: null,
      uid: null,
      selectedFID: null,
      selectedFaculty: null,
      facultyForChange: null,
      studentForChange: null,
      showLeftBar: false,
      showRightBar: false,
      searchStr: "",
      gridApi: null,
      vw: 0,
      facultyStudents: null,
      otherStudents: null,
    }
  },
  created() {
    if (!this.checkLoggedIn()) {
      this.$router.push("/sign-in").catch(() => {})
    }
    this.uid = localStorage.getItem("uid")
    this.vw = window.innerWidth;
    window.addEventListener("resize", this.resizeHandler);
    this.getData()
  },
  unmounted() {
    window.removeEventListener("resize", this.resizeHandler);
  },
  computed: {
    columnDefs() {
      return [
        {field: "Name", headerName: this.$t("nameU"), width: (this.vw - 65) * 0.4, sortable: true},
        {field: "ExamDate", headerName: this.$t("examDate"), width: (this.vw - 65) * 0.3, sortable: true},
        {field: "ExamAud", headerName: this.$t("examAud"), width: (this.vw - 65) * 0.3, sortable: true}
      ]
    },
    facultyStudentsColumnDefs() {
      return [
        {field: "Name", headerName: this.$t("name"), width: (this.vw - 130) * 0.7 / 2, sortable: true},
        {field: "Exam", headerName: this.$t("exam"), width: (this.vw - 130) * 0.3 / 2, sortable: true},
      ]
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
      this.selectedFID = selectedRows[0].ID;
      this.selectedFaculty = selectedRows[0]
    },
    addHandler() {
      this.facultyForChange = null
      this.$bvModal.show("faculty-new-edit-modal")
    },
    editHandler() {
      if (!this.selectedFaculty) {
        return
      }
      this.facultyForChange = this.selectedFaculty
      this.$bvModal.show("faculty-new-edit-modal")
    },
    deleteHandler() {
      this.$http.delete(url + "/api/faculties", this.addBasicAuth({
        data: this.selectedFaculty
      }))
          .then(() => {
            this.selectedFaculty = null
            this.selectedFID = null
            this.getData()
          })
          .catch(err => this.$error(err.response.data))

    },
    clickMoreHandler() {
      if (!this.showRightBar) {
        this.getOtherStudents()
        this.showRightBar = true
      } else {
        this.showRightBar = false
        this.otherStudents = null
      }
    },
    addStudentHandler() {
      this.studentForChange = null
      this.$bvModal.show("student-new-edit-modal")
    },
    editStudentHandler(student) {
      if (!student) {
        return
      }
      this.studentForChange = student
      this.$bvModal.show("student-new-edit-modal")
    },
    deleteStudentHandler(student) {
      if (!student) {
        return
      }
      this.$http.delete(url + "/api/students", this.addBasicAuth({data: student}))
          .then(() => {
            this.sidebarRefreshHandler()
          })
          .catch(err => this.$error(err.response.data))
    },
    kickHandler(student) {
      if (!student?.ID || !this.selectedFID) {
        return
      }
      this.$http.delete(url + "/api/faculties_to_students", this.addBasicAuth({data: {Fid: this.selectedFID, Sid: student.ID}}))
          .then(() => {
            this.sidebarRefreshHandler()
          })
          .catch(err => this.$error(err.response.data))
    },
    inviteHandler(student) {
      if (!student?.ID || !this.selectedFID) {
        return
      }
      this.$http.post(url + "/api/faculties_to_students",
          {Fid: this.selectedFID, Sid: student.ID}, this.addBasicAuth({}))
          .then(() => {
            this.sidebarRefreshHandler()
          })
    },
    resizeHandler() {
      this.vw = window.innerWidth;
    },
    onGridReady(params) {
      this.gridApi = params.api
    },
    doesExternalFilterPass(el) {
      return `${el.data.Name} ${el.data.ExamDate} ${el.data.ExamDate}`.toLowerCase().includes(this.searchStr.toLowerCase())
    },
    getData() {
      this.$http.get(url + "/api/faculties", this.addBasicAuth({params: {uid: this.uid}}))
          .then(r => {
            this.faculties = r.data || []
          })
          .catch(err => this.$error(err.response.data))
    },
    getStudentsByFID() {
      if (!this.selectedFID) {
        this.facultyStudents = []
        return
      }
      this.facultyStudents = null
      this.$http.get(url + "/api/faculties_to_students", this.addBasicAuth({params: {id: this.selectedFID}}))
          .then(r => {
            this.facultyStudents = r.data || []
          })
          .catch(err => this.$error(err.response.data))
    },
    getOtherStudents() {
      this.otherStudents = null
      this.$http.get(url + "/api/students", this.addBasicAuth({}))
          .then(r => {
            const facultyStudentsIDs = (this.facultyStudents || []).map(el => el.ID) || []
            this.otherStudents = r.data.filter(el => !facultyStudentsIDs.includes(el.ID))
          })
          .catch(err => this.$error(err.response.data))
    },
    async sidebarRefreshHandler() {
      await this.getStudentsByFID()
      if (this.showRightBar) {
        this.getOtherStudents()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/helpers/variables.scss";
.facultiesGridView {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: $mtBlack;
  padding-left: 30px;
  padding-right: 30px;
}
</style>