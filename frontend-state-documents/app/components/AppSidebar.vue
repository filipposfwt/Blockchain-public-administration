<script setup lang="ts">
import { ref } from 'vue'

const leftDrawerOpen = ref(false)

const router = useRouter();
const route = useRoute();
const logout = () => {
  localStorage.removeItem("dms_accessToken");
  router.push({ name: "login" });
};  
</script>

<template>
  <q-drawer
    v-model="leftDrawerOpen"
    show-if-above
    bordered
    class="bg-grey-1"
  >
    <q-list>
      <q-item-label header>Navigation</q-item-label>

      <q-item
        to="/"
        clickable
        v-ripple
      >
        <q-item-section avatar>
          <q-icon name="home" />
        </q-item-section>
        <q-item-section>
          {{ $t('sidebar.home') }}
        </q-item-section>
      </q-item>

      <q-item
        to="/documents"
        clickable
        v-ripple
      >
        <q-item-section avatar>
          <q-icon name="description" />
        </q-item-section>
        <q-item-section>
          {{ $t('documents.title') }}
        </q-item-section>
      </q-item>
    </q-list>

    <div class="absolute-bottom">
      <q-item v-ripple clickable @click="logout()">
        <q-item-section avatar>
          <Icon name="carbon:logout" size="22px" />
        </q-item-section>
        <q-item-section> {{ $t("auth.logout") }} </q-item-section>
      </q-item>
    </div>
  </q-drawer>
</template>

<style lang="scss">
.sidebar {
  background-color: #fff !important;
}

.logo {
  filter: brightness(0) saturate(100%) invert(99%) sepia(9%) saturate(124%)
    hue-rotate(194deg) brightness(114%) contrast(100%);
}
</style>
