<template>
  <v-content>
    <v-container fluid fill-height id="login-wrapper">
      <div id="login">
        <img src="../assets/owl.svg">
        <h2>Sign in with your</h2>
        <h1>BCP ID</h1>
        <v-form style="margin-bottom:30px;width:300px;">
          <v-text-field
            v-model="user.id"
            v-validate="'required|alpha_num'"
            data-vv-name="id"
            :error-messages="errors.collect('id')"
            prepend-icon="person"
            name="id"
            label="id"
            type="text"
            color="cyan"
            required
            ></v-text-field>
          <v-text-field
            v-model="user.password"
            v-validate="'required|min:6'"
            data-vv-name="password"
            :error-messages="errors.collect('password')"
            prepend-icon="lock"
            v-on:keyup.enter="submit"
            name="password"
            label="Password"
            id="password"
            type="password"
            color="cyan"
            required
          ></v-text-field>
        </v-form>
        <v-btn color="info" v-on:click="submit"  depressed>Sigin</v-btn>
        <div id="createPassword">
          <a>Forgot Password?</a>
          <span>|</span>
          <a>Create Account</a>
        </div>
      </div>
    </v-container>
  </v-content>
</template>


<script>
import URL from "../parameter/ip"
export default {
  data() {
    return {
      user: {
        id: "",
        password: "",
      }
    }
  },
  methods: {
    submit() {
      this.$validator.validateAll().then(valid => {
         const url = URL.login
         let obj = this.user
         this.axios.post(url, obj)
         .then((response) => {
           if (response.status === 200) {
              this.$store.commit('account/toggleLoginToken', response.data.token)
              this.$store.commit('account/setUserInfo', response.data)
              this.$cookies.set("token", response.data.token, "4m")
              this.$router.push("/account")
           }
         })
         .catch((error) => {
           console.log(error)
         }) 
      })
    }
  },

};
</script>


<style lang="scss">
#login-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
}
#login {
  padding: 30px;
  width: 500px;
  height: 600px;
  background-color: #424242;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  img {
      width: 120px;
  }

  h2 {
    font-size: 2em;
  }
  h1 {
    font-size: 2.5em;
    font-weight: 800;
  }
  #createPassword {
    margin-top: 20px;
    width: 60%;
    display: flex;
    justify-content: space-around;
    /* unvisited link */
    a {
        width: 120px;
    }
    a:link {
      color: green;
    }
  }
}
</style>
