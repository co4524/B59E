<template>
  <v-content>
    <v-container fluid fill-height id="register-wrapper">
      <div id="register">
        <div id="login_header">
          <img v-if="pass==0" src="../assets/chicken_small.svg">
          <img v-if="pass==1" src="../assets/chicken.svg">
          <img v-if="pass==2" src="../assets/chicken_hot.svg">
          <div v-if="pass!=2">
            <h2>Register B59E account</h2>
            <p>
              If you already have B59E account, click
              <router-link to="/login">sigin</router-link>
            </p>
          </div>
          <div v-else>
            <h2>You succeesfully register!</h2>
          </div>
        </div>
        <RegisterName v-if="pass==0" v-on:triggerSuccess="success_and_next"></RegisterName>
        <RegisterDetail v-if="pass==1" v-on:triggerSuccess="post_and_next"></RegisterDetail>
      </div>
    </v-container>
    <v-footer style="position:absolute;bottom:0px;width:100%">
      <v-progress-linear v-model="progress_bar" color="cyan" height="8"></v-progress-linear>
    </v-footer>
  </v-content>
</template>


<script>
import RegisterName from "@/components/RegisterName"
import RegisterDetail from "@/components/RegisterDetail"
import URL from "../parameter/ip"
export default {
  data() {
    return {
      progress_bar: 33,
      pass: 0,
      register: {
        id: "",
        password: "",
        email: "",
        name: "",
        identification: "",
        birth: "",
        registertime: "",
      }
    };
  },
  components: {
    RegisterName,
    RegisterDetail,
  },

  methods: {
    success() {
      this.pass = this.pass + 1;
      this.progress_bar = this.progress_bar + 33;
    },
    success_and_next(arg){
      this.register.id = arg.id
      this.register.password = arg.password
      this.register.email = arg.email
      this.success()
    },
    post_and_next(arg) {
      this.register.name = arg.name
      this.register.identification = arg.identification
      this.register.birth = arg.birth
      this.register.registertime = this.getdate()
      console.log(this.register.registertime)
      const url = URL.register.register
      let obj = this.register;
      this.axios.post(url, obj)
        .then(response => {
          if (response.data.status === "OK") {
            // Set token 
            this.$store.commit('account/toggleLoginToken', response.data.token)
            this.$store.commit('account/setUserInfo', obj)
            // Set cookie
            this.$cookies.remove("token");
            this.$cookies.set("token", response.data.token, "4m")
            this.success()
            setTimeout(() => {
              this.$router.push({path:'/project'})
            }, 1000);
          }
        })
        .catch(error => {
          console.log(error)
        });
    },
    getdate() {
      var d = new Date();
      var months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
      return months[d.getMonth()] + " " + d.getDate().toString() + ", " + d.getFullYear().toString()
    }
  }
};
</script>

<style lang="scss">
#register-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
}

#register {
  padding: 30px;
  width: 500px;
  height: 700px;
  background-color: #424242;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  #login_header {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  img {
    margin-right: 15px;
    width: 80px;
  }

  .input {
    height: 80px;
  }
}
</style>


