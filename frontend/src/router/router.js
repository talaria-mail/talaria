import Router from "vue-router";

import Login from "../views/Login.vue";

export default new Router({
  mode: "history",
  base: "/",
  routes: [
    {
      path: "/login",
      name: "login",
      component: Login
    }
  ]
});
