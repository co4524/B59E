import Vue from 'vue'
import './plugins/vuetify'
import VeeValidate from 'vee-validate'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueCookies from 'vue-cookies'
import App from './App.vue'
import router from './router'
import {store} from './store/store'
import URL from './parameter/ip'

Vue.use(VueAxios, axios);
Vue.use(VeeValidate);
Vue.use(VueCookies)

Vue.config.productionTip = false

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    // If has cookie in JWT token , validate it!
    const url = URL.auth
    console.log(url)
    const token = VueCookies.get("token")
    const form = new FormData();
    form.append("token", token)
    axios.post(
        url, 
        form,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
      .then((response) => {
        if(response.status == 200) {
          const user = store.getters["account/user"]
          if (user === "") {
            store.commit("account/setUserInfo", response.data)
          }
          store.commit("account/toggleLoginToken", token)
          next()
        } else {
          next({
            path: "/login",
            query: {
              redirect: to.fullPath
            }
          });
        }
      })
      .catch((error) => {
        console.log(error)
        next({
          path: "/login",
          query: {
            redirect: to.fullPath
          }
        });
      })
  } else {
    next(); // make sure to always call next()!
  }
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
