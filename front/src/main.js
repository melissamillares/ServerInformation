import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import {store} from './store/store'

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'

Vue.use(BootstrapVue)
Vue.prototype.$http = axios

Vue.config.productionTip = false

new Vue({
  router,
  // provide the store using the "store" option.
  // this will inject the store instance to all child components.
  store: store,
  components: { store },
  render: h => h(App),
}).$mount('#app')
