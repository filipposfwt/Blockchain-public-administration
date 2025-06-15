<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'auth',
  name: 'signup'
})

const router = useRouter()
const { register, loading, error } = useAuth()

const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const name = ref('')
const organization = ref('')

const handleRegister = async () => {
  if (password.value !== confirmPassword.value) {
    error.value = 'Passwords do not match'
    return
  }

  try {
    await register(email.value, password.value, {
      name: name.value,
      organization: organization.value
    })
    router.push('/documents')
  } catch (e) {
    console.error('Registration failed:', e)
  }
}
</script>

<template>
  <q-page class="flex flex-center">
    <q-card class="signup-card">
      <q-card-section class="bg-deep-red text-white">
        <div class="text-h6">Register</div>
        <div class="text-subtitle2">Create your account</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit="handleRegister" class="q-gutter-md">
          <q-input
            v-model="name"
            label="Full Name"
            :rules="[val => !!val || 'Name is required']"
          />

          <q-input
            v-model="email"
            type="email"
            label="Email"
            :rules="[val => !!val || 'Email is required']"
          />

          <q-input
            v-model="organization"
            label="Organization"
            :rules="[val => !!val || 'Organization is required']"
          />

          <q-input
            v-model="password"
            type="password"
            label="Password"
            :rules="[val => !!val || 'Password is required']"
          />

          <q-input
            v-model="confirmPassword"
            type="password"
            label="Confirm Password"
            :rules="[val => !!val || 'Please confirm your password']"
          />

          <div class="text-negative" v-if="error">{{ error }}</div>

          <div>
            <q-btn
              type="submit"
              label="Register"
              class="full-width"
              :loading="loading"
              style="background-color: #8B0000 !important; color: white"
            />
          </div>

          <div class="text-center q-mt-sm">
            <router-link to="/login" class="text-deep-red">
              Already have an account? Login
            </router-link>
          </div>
        </q-form>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<style scoped>
.signup-card {
  width: 100%;
  max-width: 400px;
}

.bg-deep-red {
  background-color: #8B0000 !important;
}

.text-deep-red {
  color: #8B0000;
}
</style> 