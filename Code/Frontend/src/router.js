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
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/detail",
      name: "detail",
      component: Detail
    },
    {
      path: "/onboarding",
      name: "onboarding",
      component: Onboarding
    },
    {
      path: "/onboarding",
      name: "onboarding",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/Onboarding.vue")
    },
    {
      path: "/guide",
      name: "guide",
      component: () =>
        import(/* webpackChunkName: "guide" */ "./views/Guide.vue")
    }
  ]
});
