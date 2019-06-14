<template>
  <v-content id="proj">
    <v-container fluid fill-height grid-list-xl style="maxWidth: 80%;">
      <v-layout justify-space-around row wrap v-if="item.length">
        <div v-for="(val, idx) in item" :key="idx" class="row">
          <v-card>
            <v-card-title>
              <div>
                <div class="headline">{{ val.title }}</div>
              </div>
            </v-card-title>
            <v-img :src="val.url" height="300px"></v-img>
            <v-card-text>{{ val.description }}</v-card-text>
            <v-progress-linear color="cyan" height="5" :value="progress(val.currentFund, val.targetFund)"></v-progress-linear>
            <v-list-tile>
              <v-list-tile-content>
                <v-list-tile-sub-title ><v-icon>pregnant_woman</v-icon> Launcher : {{val.name}}</v-list-tile-sub-title>
                <v-list-tile-sub-title ><v-icon>attach_money</v-icon> Price: {{val.targetFund}}</v-list-tile-sub-title>
              </v-list-tile-content>
              <v-list-tile-content>
                <v-list-tile-sub-title ><v-icon>update</v-icon> Deadline: {{val.enddate}}</v-list-tile-sub-title>
                <v-list-tile-sub-title ><v-icon>attach_money</v-icon> Price: {{val.currentFund}}</v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-card-actions>
              <!--<v-btn flat color="cyan" v-on:click="donate($event, val)">Donate</v-btn>-->
              <v-dialog v-model="dialog" persistent max-width="400">
                <template v-slot:activator="{ on }">
                  <!-- <v-btn color="primary" dark v-on="on">Open Dialog</v-btn>-->
                  <v-btn flat color="cyan" v-on="on">Buy</v-btn>
                </template>
                <v-card>
                  <v-layout justify-space-around row wrap>
                    <v-card-title class="headline">BUY TREASURE</v-card-title>
                    <v-flex xs10>
                    <v-text-field
                      v-model="did"
                      label="directoryID"
                      prepend-icon="folder"
                      color="cyan"
                    ></v-text-field>
                    <v-text-field
                      v-model="id"
                      label="User Id"
                      prepend-icon="person"
                      color="cyan"
                    ></v-text-field>
                    <v-text-field
                      v-model="pwd"
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
                    </v-flex>
                    <v-flex xs10>
                      <v-card-text>Are Your Sure to Buy {{val.title}} ?</v-card-text>
                    </v-flex>
                    <v-card-actions>
                      <v-spacer></v-spacer>
                      <v-btn color="red" flat @click="dialog = false">Disagree</v-btn>
                      <v-btn color="green darken-1" flat @click="dialog = false" v-on:click="donate">Agree</v-btn>
                    </v-card-actions>
                  </v-layout>
                </v-card>
              </v-dialog>
              <v-btn color="orange" flat @click="init">DETAIL</v-btn>
            </v-card-actions>
          </v-card>
        </div>
        <div class="row"></div>
        <div class="row"></div>
      </v-layout>
      <v-layout justify-space-around align-center wrap v-else>
        <h1>No Current Project Found</h1>
      </v-layout>
    </v-container>
  </v-content>
</template>

<script>
import {address} from "../web3/contract/BCPContract"
import URL from "../parameter/ip"
import { mapState } from "vuex";
import { EAS } from "../test.js"
import { retrieveEAS } from "../test.js"
export default {
  data() {
    return {
      item: [],
      dialog: false,
      did: "",
      id: "",
      pwd: "",

    };
  },
  async beforeCreate() {
    const url = URL.launch.propose
    this.axios.get(url).then(response => {
      console.log(response.data);
      this.item = response.data;
    });

    await this.$store.dispatch("web3/getPlatformContractInstance");
  },

  computed: {
    ...mapState('web3/', {
      PlatformContractInstance: state => state.PlatformContractInstance,
      coinbase: state => state.web3.coinbase,
    }),
  },
  methods: {
    progress(currentFund, targetFund) {
      return parseInt(currentFund) / parseInt(targetFund)
    },
    async donate() {
      let result = await EAS(
        this.did,
        this.id,
        'dataCertificate',
        2020,
        'providerAgreement',
        this.pwd
      );
      alert(result);
    },
    async init() {
      console.log("debug");
    }
  }
};
</script>

<style lang="scss">
#proj {
  .row {
    width: 430px;
    margin-top: 30px;
  }
}
</style>
