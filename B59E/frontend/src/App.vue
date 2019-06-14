<template>
  <v-app dark>
    <Navigation/>
    <transition name="fade" mode="out-in">
      <router-view/>
    </transition>
  </v-app>
</template>

<script>
import Navigation from "./components/Navigation";
export default {
  name: "App",
  components: {
    Navigation
  },
  data() {
    return {
      //
    };
  },
  async beforeCreate() {
    try {
      await this.$store.dispatch("web3/registerWeb3");
    } catch (error) {
      console.log("error")
    }
  }
};
</script>


<style lang="scss">
@import url("https://fonts.googleapis.com/css?family=Roboto+Condensed");
@import url("https://fonts.googleapis.com/css?family=Lato:300");
.fade-enter-active,
.fade-leave-active {
  transition-duration: 0.3s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
  opacity: 0;
}
</style>
