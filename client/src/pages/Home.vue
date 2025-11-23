<template>
  <div class="min-h-screen bg-gray-50 flex flex-col items-center justify-center p-10">
    <div class="bg-white shadow-xl rounded-2xl p-10 w-full max-w-md text-center">
      <h1 class="text-3xl font-bold mb-6 text-gray-800">Bienvenue</h1>
      <p class="text-gray-600 mb-6">
        Connectez-vous pour accéder à votre profil et à la zone admin.
      </p>
      <button
        @click="login"
        class="px-6 py-3 bg-blue-600 text-white rounded-xl shadow hover:bg-blue-700 transition"
      >
        Login avec Zitadel (PKCE)
      </button>
    </div>
  </div>
</template>

<script setup>
import { useRoute } from "vue-router";
import { useAuthStore } from "../stores/auth";

const auth = useAuthStore();
auth.init();

const route = useRoute();

const login = () => {
  const redirect = route.query.redirect || "/app/profile";
  sessionStorage.setItem("post_login_redirect", redirect);
  auth.login();
};
</script>
