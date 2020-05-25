import Router from "vue-router"

import Login from '@/views/Login.vue'
import Main from '@/views/Main.vue'
import UserProfile from '@/components/UserProfile.vue'

export default new Router({
  mode: "history",
  base: "/",
  routes: [
    {
      path: "/login",
      name: "login",
      component: Login
    },
    {
      path: "/",
      name: "main",
      component: Main,
      children: [
        {
          path: "/profile",
          name: "profile",
          component: UserProfile,
          meta: { authenticated: true },
        }
      ]
    }
  ]
});
