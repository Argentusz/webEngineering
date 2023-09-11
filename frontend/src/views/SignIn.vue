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
import axios from "axios";

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
  methods: {
    signInHandler() {
      const base64auth = btoa(`${this.user.Login}:${this.user.Password}`)
      const auth = axios.create({
        headers: {
          Authorization: `Basic ${base64auth}`
        }
      })
      auth.get(url + "/auth")
          .then(data => {
            console.log("data", data)
          })
          .catch(err => {
            console.log("err", err)
          })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "@/helpers/variables.scss";

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