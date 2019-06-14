<template>
  <div style="display:flex; flex-direction:column; align-item:center; justify-conter:center">
    <v-form style="margin-bottom:5px;width:400px;">
      <v-text-field
        v-model="register.did"
        v-validate="'required|alpha_num|'"
        data-vv-name="id"
        prepend-icon="folder"
        name="login"
        placeholder="DirectoryID"
        type="text"
        color="cyan"
        class="input"
        required
      ></v-text-field>
      <v-text-field
        v-model="register.id"
        v-validate="'required|alpha_num|'"
        data-vv-name="id"
        :error-messages="errors.collect('id')"
        prepend-icon="person"
        name="login"
        placeholder="User ID"
        type="text"
        color="cyan"
        class="input"
        required
      ></v-text-field>
      <v-text-field
        v-model="register.password"
        v-validate="'required|min:6'"
        data-vv-name="password"
        :error-messages="errors.collect('password')"
        prepend-icon="lock"
        name="password"
        placeholder="Password"
        type="password"
        color="cyan"
        class="input"
        required
      ></v-text-field>
      <v-text-field
        v-model="register.userType"
        v-validate="'required|alpha_num|'"
        data-vv-name="id"
        prepend-icon="face"
        name="login"
        placeholder="Provider/Consumer"
        type="text"
        color="cyan"
        class="input"
        required
      ></v-text-field>
      <v-text-field
        v-model="register.email"
        v-validate="'required|email'"
        data-vv-name="email"
        :error-messages="errors.collect('email')"
        prepend-icon="email"
        color="cyan"
        placeholder="Email"
        class="input"
        required
      ></v-text-field>
    </v-form>
    <v-btn color="info" depressed @click="submit">continue</v-btn>
    <v-btn color="info" depressed @click="init">Init</v-btn>
  </div>
</template>


<script>
import URL from "../parameter/ip"
import { createDir } from "../test.js"
import { regProvider , regConsumer } from "../test.js"
export default {
  name: "RegisterName",
  data() {
    return {
      register: {
        did: "",
        id: "",
        password: "",
        userType: "",
        email: ""
      }
    };
  },
  methods: {
    async submit() {
      let result = await regProvider( this.register.did , this.register.userType , this.register.id , this.register.password ) ;
      alert(result);
      this.$validator.validateAll().then(valid => {
        if (valid) {
          console.log("post queryName");
          const url = URL.register.queryName
          let obj = this.register;
          console.log(obj);
          this.axios
            .post(url, obj)
            .then(response => {
              if (response.data.status === "OK") {
                console.log("You can continue");
                this.$emit('triggerSuccess', obj)
              }
            })
            .catch(error => {
              console.log(error);
            });
        }
      });
    },
    async init() {
      let DEBUG2 = await createDir();
      console.log(DEBUG2);
      alert(DEBUG2);
    }
  }
};
</script>