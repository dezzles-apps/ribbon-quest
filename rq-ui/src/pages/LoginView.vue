<script setup lang="ts">

import { ref } from 'vue';
import { useApi } from '@/composables/useApi';
import { useAuthStore } from '@/stores/authStore';
import { useRouter } from 'vue-router';

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
    const response = await apiFetch('/api/auth/v1/login', {
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
  <div class="about">
    <div class="box mb-5">
    <h1 class="title is-3">Login</h1>
      <div class="field">
        <p class="control has-icons-left has-icons-right">
          <input class="input" type="username" placeholder="Username" v-model="username" @keyup.enter="handleLogin">
          <span class="icon is-small is-left">
            <i class="mdi mdi-account"></i>
          </span>
          <span class="icon is-small is-right">
            <i class="mdi mdi-check"></i>
          </span>
        </p>
      </div>
      <div class="field">
        <p class="control has-icons-left">
          <input class="input" type="password" placeholder="Password" v-model="password" @keyup.enter="handleLogin" >
          <span class="icon is-small is-left">
            <i class="mdi mdi-lock"></i>
          </span>
        </p>
      </div>
      {{ viewError }}
      <div class="field">
        <p class="control">
          <button class="button is-success" @click="handleLogin" :class="{ 'is-loading': loading }">
            Login
          </button>
        </p>
      </div>
    </div>
  </div>
</template>