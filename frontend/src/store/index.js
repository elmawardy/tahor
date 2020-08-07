import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    backendURL: 'http://localhost:8000',
    loggedIn: false,
    user_name: ""
  },
  mutations: {
    login(state) {
      state.loggedIn = true
    },
    setUsername(state, user_name) {
      state.user_name = user_name
    },
    logout(state) {
      state.loggedIn = false
    }
  },
  actions: {
  },
  modules: {
  }
})
