<template>
<b-modal
    :title="!changeFaculty ? $t('newFaculty') : `${$t('changeFaculty')}: ${changeFaculty.Name}`"
    id="faculty-new-edit-modal"
    header-bg-variant="warning"
    header-text-variant="dark"
    header-border-variant="warning"
    body-bg-variant="dark"
    body-text-variant="white"
    footer-bg-variant="dark"
    footer-text-variant="white"
    footer-border-variant="warning"
    cancel-disabled
    ok-disabled
>
    <template #modal-footer>
      <b-button
          class="text-sm"
          @click="$bvModal.hide('faculty-new-edit-modal'); $emit('hide')"
      >
        {{ $t('cancel') }}
      </b-button>
      <b-button
        variant="success"
        class="text-sm"
        @click="saveHandler"
      >
        {{ $t('save') }}
      </b-button>
    </template>
  <b-input v-model="faculty.Name" :placeholder="$t('nameU')" class="text-sm input"/>
  <b-input v-model="faculty.ExamDate" :placeholder="$t('examDate')" class="text-sm input" />
  <b-input v-model="faculty.ExamAud" :placeholder="$t('examAud')" class="text-sm input" />

</b-modal>
</template>

<script>
import {url} from "@/main";

export default {
  name: "FacultyNewEditModal",
  data() {
    return {
      faculty: {
        Name: "",
        ExamDate: "",
        ExamAud: ""
      },
      method: "POST"
    }
  },
  props: {
    changeFaculty: {
      type: [Object, null],
      default: null
    }
  },
  watch: {
    changeFaculty() {
      if (this.changeFaculty) {
        this.faculty = this.changeFaculty
        this.method = "PATCH"
        return
      }
      this.faculty = {
        Name: "",
        ExamDate: "",
        ExamAud: ""
      }
      this.method = "POST"
    }
  },
  methods: {
    allowedSymbols(str, spacesAllowed = false) {
      if (spacesAllowed) {
        return /^[a-zA-Z0-9а-яА-Я!@#^*()., ]+$/.test(str);
      } else {
        return /^[a-zA-Z0-9!@#^*().,]+$/.test(str);
      }
    },
    saveHandler() {
      // TODO: Оповещение об ошибке
      if (this.faculty.Name === "" || !this.allowedSymbols(this.faculty.Name, true)) {
        console.log("*")
        return
      }
      if (this.faculty.ExamDate === "" || !this.allowedSymbols(this.faculty.ExamDate, true)) {
        console.log("**")
        return
      }
      if (this.faculty.ExamAud === "" || !this.allowedSymbols(this.faculty.ExamAud, true)) {
        console.log("***")
        return
      }
      this.faculty.UniversityID = +localStorage.getItem("uid")
      if (this.method === "POST") {
        this.post()
      } else {
        this.patch()
      }

    },
    post() {
      this.$http.post(url + "/api/faculties", this.faculty, this.addBasicAuth({}))
          .then(() => {
            this.$bvModal.hide("faculty-new-edit-modal")
            this.$emit("hide")
            this.$emit("refresh")
            this.faculty = {
              Name: "",
              ExamDate: "",
              ExamAud: ""
            }
          })
          .catch(err => this.$error(err.response.data))
    },
    patch() {
      this.$http.patch(url + "/api/faculties", this.faculty, this.addBasicAuth({}))
          .then(() => {
            this.$bvModal.hide("faculty-new-edit-modal")
            this.$emit("hide")
            this.$emit("refresh")
            this.faculty = {
              Name: "",
              ExamDate: "",
              ExamAud: ""
            }
          })
          .catch(err => this.$error(err.response.data))
    }
  }
}
</script>

<style lang="scss" scoped>
.input:not(:last-child) {
  margin-bottom: 15px;
}
</style>