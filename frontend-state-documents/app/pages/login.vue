<template>
  <q-page class="flex flex-center">
    <q-card class="login-card">
      <q-card-section class="bg-primary text-white">
        <div class="text-h6">Login</div>
        <div class="text-subtitle2">Welcome to Public Document Service</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit="handleLogin" class="q-gutter-md">
          <q-input
            v-model="email"
            type="email"
            label="Email"
            :rules="[val => !!val || 'Email is required']"
          />

          <q-input
            v-model="password"
            type="password"
            label="Password"
            :rules="[val => !!val || 'Password is required']"
          />

          <div class="text-negative" v-if="error">{{ error }}</div>

          <div>
            <q-btn
              type="submit"
              label="Login"
              class="full-width"
              :loading="loading"
            />
          </div>

          <div class="text-center q-mt-sm">
            <router-link to="/signup">
              Don't have an account? Register
            </router-link>
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'auth',
  name: 'login'
})

const router = useRouter()
const { login, loading, error } = useAuth()

const email = ref('')
const password = ref('')

const handleLogin = async () => {
  try {
    await login(email.value, password.value)
    router.push('/documents')
  } catch (e) {
    console.error('Login failed:', e)
  }
}
</script>

<style scoped>
.login-card {
  width: 100%;
  max-width: 400px;
}
</style>
