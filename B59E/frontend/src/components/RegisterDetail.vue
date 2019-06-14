<template>
  <div style="display:flex; flex-direction:column; align-item:center; justify-conter:center">
    <v-form style="margin-bottom:30px;width:320px;">
      <v-text-field
        v-model="register.name"
        v-validate="'required|alpha_spaces'"
        data-vv-name="name"
        :error-messages="errors.collect('name')"
        prepend-icon="account_circle"
        placeholder="Full name"
        type="text"
        color="cyan"
        class="input"
        key="reigster.name"
        required
      ></v-text-field>
      <v-text-field
        v-model="register.identification"
        v-validate="'required'"
        data-vv-name="identification"
        :error-messages="errors.collect('identification')"
        prepend-icon="credit_card"
        placeholder="身分證字號"
        color="cyan"
        class="input"
        key="reigster.identification"
        required
      ></v-text-field>
      <v-text-field
        v-model="register.birth"
        v-validate="'required|date_format:dd/MM/yyyy'"
        data-vv-name="birth"
        :error-messages="errors.collect('birth')"
        prepend-icon="cake"
        color="cyan"
        placeholder="Birth: dd/MM/yyyy"
        class="input"
        key="register.birth"
        required
      ></v-text-field>
    </v-form>
    <v-btn color="info" depressed @click="submit">continue</v-btn>
  </div>
</template>

<script>

export default {
  name: "RegisterDetail",
  data() {
    return {
      register: {
        name: "",
        identification: "",
        birth: ""
      }
    };
  },
  methods: {
    submit() {
      this.$validator.validateAll().then(valid => {
        if(valid) {
          let obj = this.register;
          console.log(obj);
          this.$emit('triggerSuccess', obj)
        }
      });
    }
  }
};
</script>