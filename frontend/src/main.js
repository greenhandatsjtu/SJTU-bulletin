import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'
import axios from "axios";

Vue.config.productionTip = false

Vue.prototype.$api = axios.create()

new Vue({
    vuetify,
    render: h => h(App)
}).$mount('#app')
