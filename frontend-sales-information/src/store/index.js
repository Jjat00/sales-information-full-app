import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        loading: {
            title: '',
            state: false
        }
    },
    mutations: {
        showDialog(state, payload) {
            state.loading.title = payload.title
            state.loading.state = true
        },
        quitDialog(state) {
            state.loading.state = false
        }
    },
    actions: {

    },
    modules: {}
})