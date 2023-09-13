<template>
  <b-modal
      :title="!changeStudent ? $t('newStudent') : `${$t('changeStudent')}: ${changeStudent.Name}`"
      id="student-new-edit-modal"
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
    <b-input v-model="student.Name" :placeholder="$t('name')" class="text-sm input"/>
    <b-input v-model="student.Exam" :placeholder="$t('exam')" class="text-sm input" />
  </b-modal>
</template>

<script>
import {url} from "@/main";

export default {
  name: "StudentNewEditModal",
  data() {
    return {
      student: {
        Name: "",
        Exam: ""
      },
      method: "POST"
    }
  },
  props: {
    changeStudent: {
      type: [Object, null],
      default: null
    }
  },
  watch: {
    changeStudent() {
      if (this.changeStudent) {
        this.student = this.changeStudent
        this.method = "PATCH"
        return
      }
      this.student = {
        Name: "",
        Exam: ""
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
      if (this.student.Name === "" || !this.allowedSymbols(this.student.Name, true)) {
        return
      }
      if (this.student.Exam === "" || !this.allowedSymbols(this.student.Exam, true)) {
        return
      }
      this.student.Exam = +this.student.Exam
      if (this.method === "POST") {
        this.post()
      } else {
        this.patch()
      }
    },
    post() {
      this.$http.post(url + "/api/students", this.student, this.addBasicAuth({}))
          .then(() => {
            this.$bvModal.hide("student-new-edit-modal")
            this.$emit("hide")
            this.$emit("refresh")
            this.student = {
              Name: "",
              Exam: ""
            }
          })
          .catch(err => console.log(err))
    },
    patch() {
      this.$http.patch(url + "/api/students", this.student, this.addBasicAuth({}))
          .then(() => {
            this.$bvModal.hide("student-new-edit-modal")
            this.$emit("hide")
            this.$emit("refresh")
            this.student = {
              Name: "",
              Exam: ""
            }
          })
          .catch(err => console.log(err))
    }
  }
}
</script>

<style scoped>
.input:not(:last-child) {
  margin-bottom: 15px;
}
</style>