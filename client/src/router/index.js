import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Profile from "../pages/Profile.vue";
import Callback from "../pages/Callback.vue";
import Admin from "../pages/Admin.vue";
import SecureLayout from "../layouts/SecureLayout.vue";
import { useAuthStore } from "../stores/auth";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: Home },
    { path: "/callback", component: Callback },

    {
      path: "/app",
      component: SecureLayout,
      meta: { requiresAuth: true },
      children: [
        { path: "profile", name: "profile", component: Profile },
        {
          path: "admin",
          name: "admin",
          component: Admin,
          meta: { requiresAuth: true, requiresRole: "admin" }
        }
      ]
    },

    { path: "/:pathMatch(.*)*", redirect: "/" }
  ]
});

router.beforeEach(async (to, from, next) => {
  const auth = useAuthStore();

  if (!auth.handler) {
    auth.init();
  }

  if (!to.meta.requiresAuth) {
    return next();
  }

  if (!auth.accessToken) {
    const redirect = to.fullPath;
    sessionStorage.setItem("post_login_redirect", redirect);
    return next({ path: "/" });
  }

  const requiredRole = to.meta.requiresRole;
  if (requiredRole && !auth.hasRole(requiredRole)) {
    return next({ name: "profile" });
  }

  return next();
});

export default router;
