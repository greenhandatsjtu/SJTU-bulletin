import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'
import './registerServiceWorker'

Vue.config.productionTip = false

new Vue({
    vuetify,
    render: h => h(App)
}).$mount('#app')
