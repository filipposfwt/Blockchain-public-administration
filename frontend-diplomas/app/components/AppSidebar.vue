<script setup lang="ts">
const router = useRouter();
const route = useRoute();
const logout = () => {
  localStorage.removeItem("dms_accessToken");
  router.push({ name: "login" });
};  
</script>

<template>
  <q-drawer
    class="sidebar"
    :width="200"
    side="left"
    bordered
    show-if-above
    persistent
  >

    <q-scroll-area style="height: calc(100% - 0px); margin-top: 0px">
      <q-list >
    
        <q-item
          v-ripple
          clickable
          :to="{ name: 'diplomas' }"
          :active-class="route.name === 'diplomas' ? 'bg-grey-3' : 'bg-white'"
        >
          <q-item-section avatar>
            <Icon name="carbon:document" size="22px" />
          </q-item-section>
          <q-item-section> {{ $t("sidebar.diplomas") }} </q-item-section>
        </q-item>

        
        <q-item
          v-ripple
          clickable
          :to="{ name: 'users' }"
          :active-class="route.name === 'users' ? 'bg-grey-3' : 'bg-white'"
        >
          <q-item-section avatar>
            <Icon name="carbon:events-alt" size="22px" />
          </q-item-section>

          <q-item-section> {{ $t("sidebar.users") }} </q-item-section>
        </q-item>
      </q-list>
    </q-scroll-area>

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
