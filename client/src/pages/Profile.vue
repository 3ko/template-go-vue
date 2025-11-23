<template>
  <div class="min-h-screen bg-gray-100 p-6 flex justify-center items-center">
    <div class="bg-white p-10 rounded-2xl shadow-xl w-full max-w-2xl">
      <h1 class="text-3xl font-bold mb-6 text-gray-800">Mon Profil</h1>

      <div v-if="user" class="space-y-4">
        <div class="flex items-center space-x-4">
          <div
            class="w-16 h-16 rounded-full bg-blue-600 text-white flex items-center justify-center text-2xl font-bold"
          >
            {{ initial }}
          </div>
          <div>
            <p class="text-xl font-semibold">
              {{ user.name || user.preferred_username }}
            </p>
            <p class="text-gray-500">
              {{ user.email }}
            </p>
          </div>
        </div>

        <div class="bg-gray-50 p-4 rounded-xl border mt-6">
          <h2 class="text-lg font-semibold mb-2">Détails du compte</h2>
          <pre class="text-sm text-gray-700 whitespace-pre-wrap">
{{ user }}
          </pre>
        </div>
      </div>

      <div v-else class="text-center text-gray-500">Aucun utilisateur chargé.</div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { useAuthStore } from "../stores/auth";

const auth = useAuthStore();

const user = computed(() => auth.user);
const initial = computed(() => {
  if (!user.value) return "U";
  const n = user.value.name || user.value.preferred_username || "U";
  return n.toString().charAt(0).toUpperCase();
});
</script>
