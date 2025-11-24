import { defineStore } from "pinia";

export const useConfigStore = defineStore("config", {
  state: () => ({
    config: null,
    loading: false,
    error: null
  }),

  actions: {
    async load() {
      if (this.config || this.loading) return this.config;
      this.loading = true;
      this.error = null;
      try {
        const res = await fetch("/api/config");
        if (!res.ok) {
          throw new Error("Impossible de charger la configuration");
        }
        this.config = await res.json();
        return this.config;
      } catch (err) {
        this.error = err.message;
        throw err;
      } finally {
        this.loading = false;
      }
    },

    async updateConfig(payload) {
      this.loading = true;
      this.error = null;
      try {
        const res = await fetch("/api/config", {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload)
        });
        if (!res.ok) {
          const msg = await res.text();
          throw new Error(msg || "Échec de la mise à jour");
        }
        this.config = await res.json();
        return this.config;
      } catch (err) {
        this.error = err.message;
        throw err;
      } finally {
        this.loading = false;
      }
    }
  }
});
