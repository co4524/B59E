import Vue from 'vue'
import Vuex from 'vuex'
import web3 from './web3'
import account from './account'


Vue.use(Vuex)
export const store = new Vuex.Store({
  strict: true,
  modules: {
    web3,
    account,
  },
})
