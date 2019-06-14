<template>
  <v-content>
    <v-container fluid class="pa-0 ma-0">
      <Title></Title>
      <About></About>
  </v-container>
  </v-content>
</template>

<script>
import Title from "@/components/HomePage/Title"
import About from "@/views/About"
export default {
  data() {
    return {
      bigtitle: "DECENTRALIZED COUPON CHARITY PLATFORM",
      subtitle:
        "IT'S TIME TO DONATE YOUR NEEDLESS POINTS TO EXCHANGE WHAT YOU NEED AND HELP OTHERS"
    };
  },
  components: {
    Title,
    About,
  },
};
</script>

<style lang="scss">

@import url(https://fonts.googleapis.com/css?family=Roboto+Condensed);

@mixin flex-box($jc: center, $ai: center) {
  display: flex;
  justify-content: $jc;
  align-items: $ai;
}

#intro {
  height: 100%;
  width: 100%;
}

</style>
