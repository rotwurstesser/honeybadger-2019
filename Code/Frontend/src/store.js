import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    weather: {}
  },
  mutations: {
    setWeather(state, weather) {
      state.weather = weather;
    }
  }
});
