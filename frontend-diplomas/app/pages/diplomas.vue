<script setup lang="ts">
import { doc, getDoc, getDocs, collection } from "firebase/firestore";
import { getAuth, onAuthStateChanged } from "firebase/auth";
import AddDiploma from "~/components/AddDiploma.vue";
import VerifyDiploma from "~/components/VerifyDiploma.vue";
import ViewDiploma from "~/components/ViewDiploma.vue";

definePageMeta({
  layout: "default",
  name: "diplomas",
  middleware: ["auth"],
});

const diplomasStore = useDiplomasStore();
const usersStore = useUsersStore();
const diplomas = ref([]);
const nuxtApp = useNuxtApp();
const auth = nuxtApp.$auth;
const db = nuxtApp.$firestore;
const COLLECTION_NAME = "university"; // The name of your collection

onAuthStateChanged(auth, (user) => {
  if (user) {
    // User is signed in
    user
      .getIdToken()
      .then(async (idToken) => {
        const querySnapshot = await getDocs(
          collection(db, COLLECTION_NAME, user.uid, "diplomas")
        );
        querySnapshot.forEach((doc) => {
          diplomas.value.push(doc.data());
        });
      })
      .catch((error) => {
        console.error("Error getting ID token:", error);
      });
  }
});
</script>

<template>
  <q-page padding>
    <q-toolbar>
      <q-toolbar-title>{{ $t("diplomas.title") }}</q-toolbar-title>
      <q-space />
      <q-btn
        :label="$t('addDiploma.add')"
        unelevated
        no-caps
        color="primary"
        icon="add"
        @click="diplomasStore.setAddDiplomaDialog(true, {})"
        class="q-mr-sm"
      />
      <q-btn
        :label="$t('diplomas.verify')"
        unelevated
        no-caps
        color="primary"
        icon="check"
        @click="diplomasStore.setVerifyDiplomaDialog(true)"
      />
    </q-toolbar>

    <div class="q-mt-md">
      <div class="row q-col-gutter-md">
        <div class="col-12">
          <q-table
            flat
            :rows="diplomas"
            :columns="diplomasStore.getDiplomasColumns"
            row-key="hash"
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
                <q-td key="type" :props="props" auto-width>
                  {{ props.row.type }}
                </q-td>

                <q-td key="degreeName" :props="props">
                  {{ props.row.degree_name }}
                </q-td>

                <q-td key="gradePoint" :props="props">
                  {{ props.row.grade_point }}
                </q-td>

                <q-td key="holder" :props="props">
                  {{ props.row.holder.name }}
                </q-td>

                <q-td key="hash" :props="props">
                  {{ props.row.hash }}
                </q-td>

                <q-td key="issuedDate" :props="props">
                  {{ props.row.issued_date }}
                </q-td>

                <q-td key="actions" :props="props" auto-width>
                  <q-btn
                    unelevated
                    no-caps
                    color="primary"
                    :label="$t('diplomas.view')"
                    @click="diplomasStore.setViewDiplomaDialog(true, props.row)"
                  />
                </q-td>
              </q-tr>
            </template>
          </q-table>
        </div>
      </div>
    </div>

    <VerifyDiploma />
    <ViewDiploma />
    <AddDiploma />
  </q-page>
</template>

<style scoped></style>
