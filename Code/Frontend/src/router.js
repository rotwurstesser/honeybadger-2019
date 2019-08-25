import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";
import Detail from "./views/Detail.vue";
import Onboarding from "./views/Onboarding.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/home",
      name: "home",
      component: Home
    },
    {
      path: "/detail",
      name: "detail",
      component: Detail
    },
    {
      path: "/",
      name: "onboarding",
      component: Onboarding
    },
    {
      path: "/guide",
      name: "guide",
      component: () =>
        import(/* webpackChunkName: "guide" */ "./views/Guide.vue")
    }
  ]
});
