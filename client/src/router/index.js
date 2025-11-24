import { createRouter, createWebHistory } from "vue-router";
import Config from "../pages/Config.vue";
import Profile from "../pages/Profile.vue";
import Callback from "../pages/Callback.vue";
import Admin from "../pages/Admin.vue";
import SecureLayout from "../layouts/SecureLayout.vue";
import { useAuthStore } from "../stores/auth";
import { useConfigStore } from "../stores/config";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/config", name: "config", component: Config },
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
    { path: "/", redirect: "/app/profile" },
    { path: "/:pathMatch(.*)*", redirect: "/app/profile" }
  ]
});

router.beforeEach(async (to, from, next) => {
  const auth = useAuthStore();
  const configStore = useConfigStore();
  const isConfigRoute = to.name === "config" || to.path === "/config";

  try {
    if (!configStore.config) {
      await configStore.load();
    }
    if (configStore.config?.configured && !auth.handler) {
      await auth.init();
    }
  } catch (err) {
    console.error("Configuration error", err);
    return next();
  }

  if (!configStore.config?.configured && !isConfigRoute) {
    return next({ name: "config" });
  }

  if (!to.meta.requiresAuth) {
    return next();
  }

  if (!auth.accessToken) {
    const redirect = to.fullPath;
    sessionStorage.setItem("post_login_redirect", redirect);
    if (configStore.config?.configured) {
      return auth.login();
    }
    return next({ name: "config" });
  }

  const requiredRole = to.meta.requiresRole;
  if (requiredRole && !auth.hasRole(requiredRole)) {
    return next({ name: "profile" });
  }

  return next();
});

export default router;
