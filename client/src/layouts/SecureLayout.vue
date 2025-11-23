<template>
  <div class="min-h-screen bg-gray-100">
    <header class="bg-white shadow">
      <div class="max-w-6xl mx-auto px-4 py-4 flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <span class="text-xl font-bold text-blue-600">MonApp</span>
        </div>

        <div class="flex items-center space-x-4">
          <div v-if="user" class="flex items-center space-x-3">
            <div
              class="w-10 h-10 rounded-full bg-blue-600 text-white flex items-center justify-center font-bold"
            >
              {{ initials }}
            </div>
            <div class="text-sm">
              <p class="font-semibold text-gray-800">
                {{ user.name || user.preferred_username || "Utilisateur" }}
              </p>
              <p class="text-gray-500">
                {{ user.email }}
              </p>
            </div>
          </div>

          <button
            @click="logout"
            class="px-4 py-2 bg-red-600 text-white rounded-xl text-sm hover:bg-red-700 transition"
          >
            Logout
          </button>
        </div>
      </div>
    </header>

    <main class="max-w-6xl mx-auto p-6">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { useAuthStore } from "../stores/auth";

const auth = useAuthStore();

const user = computed(() => auth.user);

const initials = computed(() => {
  if (!user.value) return "U";
  const n = user.value.name || user.value.preferred_username || "U";
  return n.toString().charAt(0).toUpperCase();
});

function logout() {
  auth.logout();
}
</script>
