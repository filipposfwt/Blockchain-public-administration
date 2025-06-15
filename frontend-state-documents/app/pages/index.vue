<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

definePageMeta({
  layout: 'auth',
  name: 'index',
  middleware: ['auth']
});

const router = useRouter();

onMounted(() => {
  // Force immediate navigation
  window.location.href = '/login';
});
</script>

<template>
  <q-page class="flex flex-center">
    <div class="text-center">
      <h1 class="text-h4 q-mb-md">Welcome to Public Document Service</h1>
      <p class="text-subtitle1">Please log in to continue</p>
    </div>
  </q-page>
</template>

<style scoped>
.text-h4 {
  font-weight: 500;
}

.page-enter-active,
.page-leave-active {
  transition: all 0.2s;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
  filter: blur(1rem);
}
</style>
