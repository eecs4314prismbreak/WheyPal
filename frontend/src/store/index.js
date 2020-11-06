// entrypoint to vuex... where we combine our modules and add them to our vuex store
import Vuex from "vuex";
import Vue from "vue";
import wheypal from "./modules/wheypal";
import createPersistedState from "vuex-persistedstate";

// Load vuex
Vue.use(Vuex);

// create store
export default new Vuex.Store({
  modules: {
    wheypal
  },
  plugins: [createPersistedState()]
});
