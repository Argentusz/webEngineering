<template>
  <div class="signIn">
      <div class="signIn__block">
        {{ $t('signIn') }}
        <b-form-input
            v-model="user.Login"
            :placeholder="$t('login')"
            class="signIn__item"
        />
        <b-form-input
            v-model="user.Password"
            :placeholder="$t('password')"
            class="signIn__item"
        />
        <b-button
            @click="signInHandler"
            variant="warning"
            class="signIn__item"
        >
          {{ $t('signIn') }}
        </b-button>
      </div>
  </div>

</template>

<script>
import {url} from "@/main";

export default {
  name: "SignIn",
  data() {
    return {
      user: {
        Login: "",
        Password: ""
      }
    }
  },
  created() {
    if (this.checkLoggedIn()) {
      this.$router.push("/").catch(() => {})
    }
  },
  methods: {
    signInHandler() {
      const base64auth = btoa(`${this.user.Login}:${this.user.Password}`)
      this.$http.get(url + "/auth", {headers: {Authorization: `Basic ${base64auth}`}})
          .then(r => {
            localStorage.setItem("uid", r.data.ID)
            localStorage.setItem("login", r.data.Login)
            localStorage.setItem("password", this.user.Password)
            localStorage.setItem("name", r.data.Name)
            this.$router.push("/").catch(() => {})
          })
          .catch(err => {
            console.log(err)
          })
    }
  }
}
</script>

<style lang="scss" scoped>


.signIn {
  display: flex;
  justify-content: center;
  align-items: center;
  background-image: url("@/assets/signInBG.png");
  width: 100%;
  height: 100vh;
}

.signIn__block {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.signIn__item {
  margin-top: 20px;
}
</style>