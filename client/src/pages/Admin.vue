<template>
  <div class="min-h-screen bg-gray-100 p-6 flex justify-center items-center">
    <div class="bg-white p-10 rounded-2xl shadow-xl w-full max-w-3xl">
      <h1 class="text-3xl font-bold mb-4 text-gray-800">Admin</h1>
      <p class="text-gray-600 mb-6">
        Espace réservé aux utilisateurs ayant le rôle
        <span class="font-semibold">admin</span> dans Zitadel.
      </p>

      <div class="grid md:grid-cols-2 gap-6">
        <div class="bg-gray-50 rounded-xl p-4 border">
          <h2 class="text-lg font-semibold mb-2">Infos utilisateur</h2>
          <pre class="text-xs text-gray-700 whitespace-pre-wrap">
{{ user }}
          </pre>
        </div>
        <div class="bg-gray-50 rounded-xl p-4 border">
          <h2 class="text-lg font-semibold mb-2">Rôles</h2>
          <pre class="text-xs text-gray-700 whitespace-pre-wrap">
{{ roles }}
          </pre>
          <p class="mt-2 text-sm" v-if="isAdmin">✅ Vous êtes admin.</p>
          <p class="mt-2 text-sm" v-else>⛔ Vous n'êtes pas admin.</p>
        </div>
      </div>

      <div class="mt-8 text-sm text-gray-500">
        <p>
          La vérification principale du rôle se fait côté front (guard router) et doit aussi être faite côté backend.
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted } from "vue";
import axios from "axios";
import { useAuthStore } from "../stores/auth";

const auth = useAuthStore();

const user = computed(() => auth.user);
const roles = computed(() => auth.roles);
const isAdmin = computed(() => auth.hasRole("admin"));

onMounted(async () => {
  if (auth.accessToken) {
    try {
      await axios.get("/api/secure/admin/stats", {
        headers: {
          Authorization: `Bearer ${auth.accessToken}`
        }
      });
    } catch (e) {
      console.error("Backend admin check failed", e);
    }
  }
});
</script>
