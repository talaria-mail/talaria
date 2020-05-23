import Vue from 'vue'
import VueRouter from 'vue-router'

import router from './router/router'
import { auth } from './router/gates'
import store from './store/'
import App from './App.vue'

Vue.config.productionTip = false

Vue.use(VueRouter)

router.beforeEach(auth(store, window))

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
