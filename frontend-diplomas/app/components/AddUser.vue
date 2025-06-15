<script setup lang="ts">
import { getAuth, onAuthStateChanged } from "firebase/auth";
import { collection, doc, setDoc, addDoc, getDoc } from "firebase/firestore";

const usersStore = useUsersStore();
const t = useI18n();
const nuxtApp = useNuxtApp();
const q = useQuasar();
const db = nuxtApp.$firestore;
const auth = nuxtApp.$auth;
const uid = ref(null);
const COLLECTION_NAME = "university"; // The name of your collection
const university = ref(null);
const addUser = ref({
  firstName: "",
  lastName: "",
  email: "",
  has_diploma: false,
  phone: "",
  university: null,
});

onAuthStateChanged(auth, (user) => {
  if (user) {
    user
      .getIdToken()
      .then(async (idToken) => {
        const querySnapshot = await getDoc(doc(db, COLLECTION_NAME, user.uid));
        university.value = querySnapshot.data().issuer;
        uid.value = user.uid;
        addUser.value.university = university.value;
      })
      .catch((error) => {
        console.error("Error getting ID token:", error);
      });
  }
});

const createUser = async () => {
  // Add a new document with a generated id.
  const docRef = await addDoc(
    collection(db, COLLECTION_NAME, uid.value, "students"),
    addUser.value
  );
  usersStore.setAddUserDialog(false);
  addUser.value = {
    firstName: "",
    lastName: "",
    email: "",  
    has_diploma: false,
    phone: "",    
    university: null,
  };
  q.notify({
    type: "positive",
    message: "User added successfully!",
  });
};
</script>

<template>
  <q-dialog
    v-model="usersStore.getAddUserDialog"
    persistent
    @hide="usersStore.setAddUserDialog(false)"
  >
    <q-card style="width: 150mm; max-width: 160mm">
      <q-form @submit="createUser">
        <q-card-actions class="bg-primary">
          <div class="text-body1 text-white">{{ $t("users.add") }}</div>
          <q-space />
          <q-btn
            flat
            round
            dense
            icon="close"
            @click="usersStore.setAddUserDialog(false)"
            color="white"
          />
        </q-card-actions>

        <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <div class="text-body2 text-bold text-black">
                {{ $t("users.firstName") }}
              </div>
              <q-input
                v-model="addUser.firstName"
                filled
                class="q-mt-xs"
                square
                :rules="[ val => val && val.length > 0 ||   $t('users.isRequired')]"
                dense
                bg-color="grey-2"
              />
            </div>
            <div class="col-12">
              <div class="text-body2 text-bold text-black q-mt-sm">
                {{ $t("users.lastName") }}
              </div>
              <q-input
                v-model="addUser.lastName"
                filled
                square
                class="q-mt-sm"
                dense
                bg-color="grey-2"
                :rules="[ val => val && val.length > 0 || $t('users.isRequired')]"
              />
            </div>
            <div class="col-12">
              <div class="text-body2 text-bold text-black q-mt-sm">
                {{ $t("users.email") }}
              </div>
              <q-input
                class="q-mt-sm"
                v-model="addUser.email"
                type="email"
                filled
                square
                dense
                bg-color="grey-2"
                :rules="[ val => !!val || $t('users.isRequired')]"
              />
            </div>
            <div class="col-12">
              <div class="text-body2 text-bold text-black q-mt-sm">
                {{ $t("users.phone") }}
              </div>
              <q-input
                v-model="addUser.phone"   
                type="tel"
                filled
                class="q-mt-xs"
                square
                dense
                bg-color="grey-2"
                :rules="[  val => !!val || $t('users.isRequired')]"
              />
            </div>
          </div>
        </q-card-section>
        <q-card-actions>
          <q-space />
          <q-btn
            :label="$t('users.cancel')"
            @click="usersStore.setAddUserDialog(false)"
            color="grey-8"
            flat
            unelevated
            no-caps
          />

          <q-btn
            :label="$t('users.add')"
            type="submit"
            color="primary"
            unelevated
            no-caps
          />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>
