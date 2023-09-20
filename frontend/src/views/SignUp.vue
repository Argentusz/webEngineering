<template>
  <div class="signIn">
    <div class="signIn__block">
      {{ $t('signUp') }}
      <b-form-input
          v-model="user.Login"
          :placeholder="$t('login')"
          class="signIn__item"
          style="width: 300px"
      />
      <b-form-input
          v-model="user.Password"
          :placeholder="$t('password')"
          class="signIn__item"
          style="width: 300px"
      />
      <b-form-input
          v-model="user.Name"
          :placeholder="$t('nameU')"
          class="signIn__item"
          style="width: 300px"
      />
      <b-button
          @click="signUpHandler"
          variant="warning"
          class="signIn__item"
      >
        {{ $t('signUp') }}
      </b-button>
      <div style="margin-top: 10px">
        {{ $t('yesAcc') }}?<br/><a href="sign-in">{{$t('signIn')}}</a>
      </div>
    </div>
  </div>

</template>

<script>
import {url} from "@/main";

export default {
  name: "SignUp",
  data() {
    return {
      user: {
        Login: "",
        Password: "",
        Name: "",
      }
    }
  },
  created() {
    if (this.checkLoggedIn()) {
      this.$router.push("/").catch(() => {})
    }
  },
  methods: {
    signUpHandler() {
      this.$http.post(url + "/auth", this.user)
          .then(r => {
            console.log(r)
            localStorage.setItem("uid", r.data)
            localStorage.setItem("login", this.user.Login)
            localStorage.setItem("password", this.user.Password)
            localStorage.setItem("name", this.user.Name)
            this.$router.push("/").catch(() => {})
          })
          .catch(err => this.$error(err.response.data))
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