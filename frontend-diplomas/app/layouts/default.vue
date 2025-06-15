<template>
  <q-layout view="hHh LpR fFf">
    <q-header class="bg-primary">
      <q-toolbar>
        <q-toolbar-title class="text-white"> {{ $t("title") }} </q-toolbar-title>

      </q-toolbar>
    </q-header>
    <app-sidebar  />
    <q-page-container class="bg-grey-3">
      <slot></slot>
    </q-page-container>

   
  </q-layout>
</template>
