<script setup lang="ts">
const usersStore = useUsersStore();
const t = useI18n();


  
</script>

<template>
  <q-dialog
    v-model="usersStore.getViewUserDialog"
    @hide="usersStore.setViewUserDialog(false)"
  >
    <q-card style="width: 140mm; max-width: 160mm">
      <q-card-actions class="bg-primary">
        <div class="text-body1 text-white">{{ $t("users.view") }}</div>
        <q-space />
        <q-btn
          flat
          round
          dense
          icon="close"
          @click="usersStore.setViewUserDialog(false)"
          color="white"
        />
      </q-card-actions>

      <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <div class="text-body1 text-bold text-black">{{ $t("users.firstName") }}</div>
              <div>{{ usersStore.getUser.firstName }}</div>
            </div>  
            <div class="col-12">
              <div class="text-body1 text-bold text-black">{{ $t("users.lastName") }}</div>
              <div>{{ usersStore.getUser.lastName }}</div>
            </div>  
            <div class="col-12">
              <div class="text-body1 text-bold text-black">{{ $t("users.email") }}</div>
              <div>{{ usersStore.getUser.email }}</div>
            </div>  
            <div class="col-12">
              <div class="text-body1 text-bold text-black">{{ $t("users.phone") }}</div>
              <div>{{ usersStore.getUser.phone }}</div>
            </div>  
            <div class="col-12">
              <div class="text-body1 text-bold text-black">{{ $t("users.university") }}</div>
              <div>{{ usersStore.getUser.university }}</div>
            </div>  
          </div>  
        </q-card-section>  
      </q-card>  
    </q-dialog>  

</template>

