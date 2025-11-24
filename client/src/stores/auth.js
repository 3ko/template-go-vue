import { defineStore } from "pinia";
import { createOIDCHandler } from "../utils/oidc";
import { useConfigStore } from "./config";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: null,
    accessToken: null,
    handler: null
  }),

  getters: {
    roles(state) {
      if (!state.user) return {};
      return (
        state.user["urn:zitadel:iam:roles"] ||
        state.user.roles ||
        {}
      );
    }
  },

  actions: {
    async init() {
      if (this.handler) return;

      const configStore = useConfigStore();
      const config = configStore.config || (await configStore.load());
      const authConfig = config?.auth || {};

      this.handler = createOIDCHandler({
        issuer: authConfig.issuer || import.meta.env.VITE_ZITADEL_ISSUER,
        clientId: authConfig.clientId || import.meta.env.VITE_ZITADEL_CLIENT_ID,
        redirectUri:
          authConfig.redirectUrl || window.location.origin + "/callback",
        scope: "openid email profile"
      });
    },

    async login() {
      const url = await this.handler.createAuthUrl();
      window.location.replace(url);
    },

    async handleCallback() {
      const result = await this.handler.processCallback();
      this.user = result.user;
      this.accessToken = result.accessToken;

      const redirect = sessionStorage.getItem("post_login_redirect") || "/app/profile";
      sessionStorage.removeItem("post_login_redirect");
      return redirect;
    },

    async logout() {
      await this.handler.logout();
      this.user = null;
      this.accessToken = null;
    },

    hasRole(role) {
      const roles = this.roles;
      if (Array.isArray(roles)) {
        return roles.includes(role);
      }
      if (roles && typeof roles === "object") {
        return !!roles[role];
      }
      return false;
    }
  }
});
