<script setup lang="ts">

import { ref } from 'vue';
import { useApi } from '@/composables/useApi';
import { useAuthStore } from '@/stores/authStore';
import { useRouter } from 'vue-router';
import { Auth } from '@/composables/endpoints';

const { apiFetch } = useApi();
const router = useRouter();
const authStore = useAuthStore();

const username = ref<string>('');
const password = ref<string>('');
const loading = ref<boolean>(false);
const viewError = ref<string | null>(null);

async function handleLogin() {
  if (loading.value) {
    return;
  }
  if (!username.value || !password.value) {
    viewError.value = 'Username and password are required.';
    return;
  }
  viewError.value = null;
  loading.value = true;
  try {
    const response = await apiFetch(Auth.Login, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: username.value,
        password: password.value
      })
    });

    const data = await response.json();
    if (!response.ok) {
      throw new Error(`${data.error || 'Unknown error'}`);
    }

    authStore.setToken(data.token);
    router.push('/');
  } catch (error) {
    console.error('Error during login:', error);
    viewError.value = (error as Error).message;
  } finally {
    loading.value = false;
  }
}

</script>

<template>
    <v-sheet
    class="align-center flex-wrap text-center mx-auto px-4"
    elevation="2"
    maxWidth="800"
    width="100%"
    rounded
  >
    <div class="pb-4 pt-4">
      <h2 class="text-headline-large font-weight-black my-0">Login</h2>

      <v-text-field
        label="Username"
        type="text"
        v-model="username"
        @keyup.enter="handleLogin"
      ></v-text-field>

      <v-text-field
        label="Password"
        type="password"
        v-model="password"
        @keyup.enter="handleLogin"
      ></v-text-field>
      <v-alert
        v-if="viewError"
        type="error"
        class="my-2"
      >
        {{ viewError }}
      </v-alert>
      <v-btn
        color="orange"
        variant="flat"
        @click="handleLogin"
      >
        Login
    </v-btn>
    </div>
  </v-sheet>
</template>