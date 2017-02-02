import Vue from 'vue'
import vueFilter from 'vue-filter'
import App from './components/App.vue'

new Vue({
  el: '#app',
  render: h => h(App)
})

Vue.use(vueFilter);
