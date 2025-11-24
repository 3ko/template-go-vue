<template>
  <div class="min-h-screen bg-gray-50 text-gray-800">
    <header class="bg-white shadow-sm">
      <div class="max-w-5xl mx-auto px-6 py-4 flex items-center justify-between">
        <div>
          <p class="text-sm text-gray-500">Page de configuration</p>
          <h1 class="text-2xl font-bold">Paramétrage de l'application</h1>
        </div>
        <div class="flex items-center gap-3">
          <button
            class="px-4 py-2 rounded-lg bg-indigo-600 text-white hover:bg-indigo-700"
            :class="{ 'opacity-60 cursor-not-allowed': !isConfigured }"
            :disabled="!isConfigured"
            @click="login"
          >
            Se connecter
          </button>
          <RouterLink
            to="/app/profile"
            class="px-4 py-2 rounded-lg border border-gray-300 hover:bg-gray-100"
          >
            Aller à l'espace sécurisé
          </RouterLink>
        </div>
      </div>
    </header>

    <main class="max-w-5xl mx-auto px-6 py-8">
      <section class="bg-white shadow rounded-xl p-6 mb-6">
        <div class="flex items-start justify-between gap-4 flex-wrap">
          <div>
            <p class="text-sm text-gray-500">État</p>
            <h2 class="text-xl font-semibold">Configuration actuelle</h2>
          </div>
          <div class="flex items-center gap-3 text-sm">
            <span class="px-3 py-1 bg-blue-50 text-blue-700 rounded-full" v-if="loading">
              Chargement...
            </span>
            <span
              class="px-3 py-1 bg-yellow-50 text-yellow-700 rounded-full"
              v-else-if="!isConfigured"
            >
              À configurer
            </span>
            <span class="px-3 py-1 bg-green-50 text-green-700 rounded-full" v-else>
              Configurée
            </span>
          </div>
        </div>

        <div class="mt-4 grid gap-3 md:grid-cols-3">
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="text-sm text-gray-500">Base de données</p>
            <p class="font-semibold">{{ form.database.name }}@{{ form.database.host }}</p>
            <p class="text-xs text-gray-500">Port {{ form.database.port }} — Utilisateur {{ form.database.user }}</p>
          </div>
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="text-sm text-gray-500">Authentification</p>
            <p class="font-semibold">Zitadel</p>
            <p class="text-xs text-gray-500 truncate">{{ form.auth.issuer }}</p>
          </div>
          <div class="p-4 bg-gray-50 rounded-lg">
            <p class="text-sm text-gray-500">Utilisateurs actifs</p>
            <p class="font-semibold">{{ config?.activeUsers?.length || 0 }}</p>
            <p class="text-xs text-gray-500">Mis à jour dynamiquement</p>
          </div>
        </div>

        <p class="text-red-600 mt-3" v-if="error">{{ error }}</p>
      </section>

      <section class="bg-white shadow rounded-xl p-6 space-y-6">
        <div class="flex items-start justify-between gap-4 flex-wrap">
          <div>
            <p class="text-sm text-gray-500">Édition</p>
            <h2 class="text-xl font-semibold">Mettre à jour les paramètres</h2>
          </div>
          <button
            class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 disabled:opacity-50"
            :disabled="saving"
            @click="save"
          >
            {{ saving ? "Sauvegarde..." : "Sauvegarder" }}
          </button>
        </div>

        <div class="grid gap-6 md:grid-cols-2">
          <div class="space-y-4">
            <h3 class="text-lg font-semibold">Base de données</h3>
            <div class="space-y-2">
              <label class="block text-sm text-gray-600">Hôte</label>
              <input v-model="form.database.host" class="w-full px-3 py-2 border rounded-lg" />
            </div>
            <div class="space-y-2">
              <label class="block text-sm text-gray-600">Nom</label>
              <input v-model="form.database.name" class="w-full px-3 py-2 border rounded-lg" />
            </div>
            <div class="space-y-2">
              <label class="block text-sm text-gray-600">Utilisateur</label>
              <input v-model="form.database.user" class="w-full px-3 py-2 border rounded-lg" />
            </div>
            <div class="space-y-2">
              <label class="block text-sm text-gray-600">Port</label>
              <input v-model="form.database.port" class="w-full px-3 py-2 border rounded-lg" />
            </div>
          </div>

          <div class="space-y-4">
            <h3 class="text-lg font-semibold">Connexion Zitadel</h3>
            <div class="space-y-2">
              <label class="block text-sm text-gray-600">Issuer</label>
              <input v-model="form.auth.issuer" class="w-full px-3 py-2 border rounded-lg" />
            </div>
            <div class="space-y-2">
              <label class="block text-sm text-gray-600">Client ID</label>
              <input v-model="form.auth.clientId" class="w-full px-3 py-2 border rounded-lg" />
            </div>
            <div class="space-y-2">
              <label class="block text-sm text-gray-600">Redirect URL</label>
              <input v-model="form.auth.redirectUrl" class="w-full px-3 py-2 border rounded-lg" />
            </div>
          </div>
        </div>

        <div class="space-y-2">
          <h3 class="text-lg font-semibold">Métadonnées (JSON)</h3>
          <textarea
            v-model="metadataText"
            rows="4"
            class="w-full px-3 py-2 border rounded-lg font-mono text-sm"
            placeholder="{\n  \"env\": \"dev\"\n}"
          />
          <p class="text-sm text-gray-500">
            Les métadonnées permettent d'ajouter des paires clé/valeur additionnelles.
          </p>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref, watch } from "vue";
import { RouterLink } from "vue-router";
import { useAuthStore } from "../stores/auth";
import { useConfigStore } from "../stores/config";

const auth = useAuthStore();
const configStore = useConfigStore();
const isConfigured = computed(() => configStore.config?.configured);

const loading = ref(false);
const saving = ref(false);
const error = ref(null);
const metadataText = ref("{}");
const config = ref(null);

const form = reactive({
  database: { host: "", name: "", user: "", port: "" },
  auth: { issuer: "", clientId: "", redirectUrl: "" }
});

const populateForm = (cfg) => {
  if (!cfg) return;
  form.database.host = cfg.database?.host || "";
  form.database.name = cfg.database?.name || "";
  form.database.user = cfg.database?.user || "";
  form.database.port = cfg.database?.port || "";
  form.auth.issuer = cfg.auth?.issuer || "";
  form.auth.clientId = cfg.auth?.clientId || "";
  form.auth.redirectUrl = cfg.auth?.redirectUrl || "";
  metadataText.value = JSON.stringify(cfg.metadata || {}, null, 2);
};

onMounted(async () => {
  loading.value = true;
  error.value = null;
  try {
    const cfg = await configStore.load();
    config.value = cfg;
    populateForm(cfg);
    if (cfg?.configured) {
      await auth.init();
    }
  } catch (err) {
    error.value = err.message;
  } finally {
    loading.value = false;
  }
});

watch(
  () => configStore.config,
  (cfg) => {
    if (cfg) {
      config.value = cfg;
      populateForm(cfg);
    }
  }
);

const save = async () => {
  saving.value = true;
  error.value = null;
  try {
    let metadata = {};
    try {
      metadata = JSON.parse(metadataText.value || "{}") || {};
    } catch (err) {
      throw new Error("Métadonnées invalides (JSON)");
    }

    const updated = await configStore.updateConfig({
      database: form.database,
      auth: form.auth,
      metadata
    });
    config.value = updated;
    if (updated?.configured && !auth.handler) {
      await auth.init();
    }
  } catch (err) {
    error.value = err.message;
  } finally {
    saving.value = false;
  }
};

const login = () => {
  auth.login();
};
</script>
