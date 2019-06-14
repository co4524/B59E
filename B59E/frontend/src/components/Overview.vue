<template>
  <div id="overview">
    <div id="email-passw">
      <h2>EMAILS</h2>
      <v-form>
        <v-text-field
          v-validate="'required'"
          data-vv-name="email"
          :error-messages="errors.collect('email')"
          v-model="user.email"
          prepend-icon="email"
          color="cyan"
          class="input_text"
          readonly
          required
        ></v-text-field>
        <v-text-field
          v-validate="'required'"
          data-vv-name="new_email"
          :error-messages="errors.collect('new_email')"
          v-model="user.new_email"
          placeholder="Enter your another email"
          prepend-icon="email"
          color="cyan"
          class="input_text"
          required
        ></v-text-field>
        <div class="divider"></div>
      </v-form>
      <h2>PASSWORD</h2>
      <v-form>
        <v-text-field
          v-model="user.password"
          v-validate="'required|min:6'"
          data-vv-name="password"
          :error-messages="errors.collect('password')"
          color="cyan"
          prepend-icon="lock"
          placeholder="Enter your current password"
          type="password"
          class="input_text"
        ></v-text-field>
        <v-text-field
          v-model="user.new_password"
          v-validate="'required|min:6'"
          data-vv-name="new_password"
          :error-messages="errors.collect('new_password')"
          color="cyan"
          prepend-icon="lock"
          placeholder="Enter your new password"
          type="password"
          class="input_text"
          ref="passwd"
        ></v-text-field>
        <v-text-field
          v-model="user.match_password"
          v-validate="'required|confirmed:passwd'"
          data-vv-name="match_password"
          :error-messages="errors.collect('match_password')"
          color="cyan"
          prepend-icon="lock"
          placeholder="Retype your new password"
          type="password"
          class="input_text"
        ></v-text-field>
      </v-form>
    </div>
    <div id="information">
      <h2>ACCOUNT INFORMATION</h2>
      <v-form>
        <v-text-field
          v-model="user.name"
          v-validate="'required|max:20|alpha_spaces'"
          :counter="20"
          data-vv-name="name"
          :error-messages="errors.collect('name')"
          color="cyan"
          prepend-icon="account_circle"
          placeholder="Full Name"
          class="input_text"
        ></v-text-field>
        <v-text-field
          v-model="user.company"
          v-validate="'required|max:20|alpha_spaces'"
          :counter="20"
          data-vv-name="name"
          :error-messages="errors.collect('name')"
          color="cyan"
          prepend-icon="account_circle"
          placeholder="Type"
          class="input_text"
        ></v-text-field>
        <v-text-field
          v-model="user.phone"
          v-validate="'required|numeric'"
          data-vv-name="phone"
          :error-messages="errors.collect('phone')"
          color="cyan"
          prepend-icon="phone_in_talk"
          placeholder="Phone"
          class="input_text"
        ></v-text-field>
        <v-text-field
          v-model="user.profile_url"
          v-validate="'required'"
          data-vv-name="profile_url"
          :error-messages="errors.collect('profile_url')"
          color="cyan"
          prepend-icon="whatshot"
          placeholder="Profile Url"
          class="input_text"
        ></v-text-field>
        <div class="divider"></div>
        <v-btn style color="warning">Submit</v-btn>
      </v-form>
    </div>
  </div>
</template>


<script>
export default {
  name: "Overview",
  data() {
    return {
      user: {
        email: "",
        new_email: "",
        password: "",
        new_password: "",
        match_password: "",
        name: "",
        company: "",
        phone: "",
        profile_url: ""
      }
    };
  },
  computed: {
  },
  mounted() {
    const user = this.$store.getters["account/user"]
    this.user.email = user.email
    this.user.name = user.name
    this.user.company = user.identification
  },
};
</script>


<style lang="scss">
#overview {
  display: grid;
  grid-template-columns: 1fr 1fr;
  h2 {
    color: #cfcfcf;
    font-weight: 500;
  }
  #email-passw {
    padding: 30px;
  }
  #information {
    padding: 30px;
  }
}
</style>
