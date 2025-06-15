<script setup lang="ts">
import { doc, getDoc,getDocs,collection,onSnapshot } from "firebase/firestore";
import { getAuth, onAuthStateChanged } from "firebase/auth";

definePageMeta({
  layout: "default",
  name: "users",
  middleware: ["auth"],
});
const usersStore = useUsersStore();
const diplomasStore = useDiplomasStore();
const students = ref([]);
const nuxtApp = useNuxtApp();
const auth = nuxtApp.$auth;
const t = useI18n();
const db = nuxtApp.$firestore;
const COLLECTION_NAME = "university"; // The name of your collection

onAuthStateChanged(auth, (user) => {
  if (user) {
    // User is signed in
    user
      .getIdToken()
      .then(async (idToken) => {


   
        const querySnapshot = await getDocs(
          collection(db, COLLECTION_NAME, user.uid, "students")
        );
        querySnapshot.forEach((doc) => {
          doc.data().uid = doc.id;
          students.value.push({id:doc.id,data:doc.data()});
        });
        usersStore.setUsers(students.value);
      })
      .catch((error) => {
        console.error("Error getting ID token:", error);
      });
  }
});

function addUser() {
  usersStore.setAddUserDialog(true);
}
</script>

<template>
  <q-page padding>

    <q-toolbar>
      
      <q-toolbar-title>{{ $t("users.title") }}</q-toolbar-title>
      <q-space />
      <q-btn
        :label="$t('users.add')"
        unelevated
        no-caps
        class="q-mr-md"
        color="primary"
        icon="add"
        @click="addUser"
      />
    </q-toolbar>

    <div class="q-mt-md">

      <div class="row q-col-gutter-md">
        <div class="col-12">
          <q-table
            flat
            :rows="usersStore.getUsers"
            :columns="usersStore.getUsersColumns"
            row-key="id"
            hide-pagination
          >
            <template v-slot:header="props">
              <q-tr :props="props" class="bg-grey-2">
                <q-th
                  v-for="col in props.cols"
                  :key="col.name"
                  class="text-weight-bold text-black"
                >
                  {{ col.label }}
                </q-th>
              </q-tr>
            </template>

            <template v-slot:body="props">
              <q-tr :props="props">
                <q-td key="firstName" :props="props" auto-width>
                  {{ props.row.data.firstName }}
                </q-td>

                <q-td key="lastName" :props="props">
                  {{ props.row.data.lastName }}
                </q-td>

                <q-td key="email" :props="props">
                  {{ props.row.data.email }}
                </q-td>

                <q-td key="phone" :props="props">
                  {{ props.row.data.phone }}
                </q-td>

                <q-td key="university" :props="props">
                  {{ props.row.data.university }}
                </q-td>

                <q-td key="actions" :props="props" auto-width>
                  <q-btn
                    unelevated
                    no-caps
                    class="q-mr-sm"
                    color="primary"
                    :disable=" props.row.data.has_diploma"

                    :label="$t('users.addDiploma')"
                    @click="diplomasStore.setAddDiplomaDialog(true, props.row)"
                  />
                  <q-btn
                    unelevated
                    no-caps
                    class="q-mr-sm"
                    color="primary"
                    :label="$t('users.view')"
                    @click="usersStore.setViewUserDialog(true, props.row.data)"
                  />
                </q-td>
              </q-tr>
            </template>
          </q-table>
        </div>
      </div>
    </div>
    <AddDiploma/> 
    <AddUser />
    <ViewUser />
  </q-page>
</template>

<style scoped></style>
