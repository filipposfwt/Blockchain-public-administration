<template>
  <q-layout view="lHh Lpr lFf">
    <q-page-container>
      <slot />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'default'
})
</script>
