<template>
  <div class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="bg-white shadow-lg rounded-2xl p-8 text-center">
      <div class="mb-4">
        <div class="w-10 h-10 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto"></div>
      </div>
      <p class="text-gray-700">Connexion en cours...</p>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";

const router = useRouter();
const auth = useAuthStore();

auth.init();

onMounted(async () => {
  try {
    const redirect = await auth.handleCallback();
    router.push(redirect);
  } catch (e) {
    console.error(e);
    router.push("/");
  }
});
</script>
