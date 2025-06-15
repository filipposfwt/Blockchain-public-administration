<script setup lang="ts">
import { getAuth, onAuthStateChanged } from "firebase/auth";
import {
  collection,
  doc,
  setDoc,
  addDoc,
  getDoc,
  updateDoc,
} from "firebase/firestore";
import { createDiploma } from "~/repository/modules/diplomas/factory";

const diplomasStore = useDiplomasStore();
const isAddDiploma = ref(false);
const t = useI18n();
const q = useQuasar();
const router = useRouter();
const COLLECTION_NAME = "university"; // The name of your collection
const uid = ref(null);
const nuxtApp = useNuxtApp();
const auth = nuxtApp.$auth;
const db = nuxtApp.$firestore;
const university = ref(null);

const diploma = ref({
  type: "",
  holder: "",
  issued_date: "",
  degree_name: "",
  grade_point: 0,
});

// Add watcher to debug holder value
watch(() => diploma.value.holder, (newValue) => {
  console.log('Holder value changed:', newValue);
});

onAuthStateChanged(auth, (user) => {
  if (user) {
    user
      .getIdToken()
      .then(async (idToken) => {
        const querySnapshot = await getDoc(doc(db, COLLECTION_NAME, user.uid));
        uid.value = user.uid;
        university.value = querySnapshot.data().issuer;
      })
      .catch((error) => {
        console.error("Error getting ID token:", error);
      });
  }
});

function reset() {
  diploma.value = {
    type: "",
    holder: "",
    issued_date: "",
    degree_name: "",
    grade_point: 0,
  };
  diplomasStore.setAddDiplomaDialog(false, {});
}

const addDiploma = async () => {
  try {
    console.log('Current holder value:', diploma.value.holder);
    // Fixed values for issuerDID and privateKey
    const ISSUER_DID = "did:example:123456789abcdefghi";
    const PRIVATE_KEY = "bc1394ed5bb2bcd70e2a54daeb22440abacef3ff84c977521656ea26df312f27";

    // Generate a unique ID (you might want to adjust this based on your needs)
    const diplomaId = Math.random().toString(36).substring(2, 10);

    // Prepare the diploma data according to the required structure
    const credential = {
      id: diplomaId,
      type: diploma.value.type,
      issuer: university.value,
      holder: diploma.value.holder.trim(), // Ensure no whitespace
      issuedDate: diploma.value.issued_date,
      expiryDate: "2030-05-27", // Fixed expiry date as per example
      degreeName: diploma.value.degree_name,
      university: university.value,
      gradePoint: parseFloat(diploma.value.grade_point)
    };

    console.log('Storing diploma on blockchain:', credential);
    const response = await createDiploma({
      diplomaData: credential,
      issuerDID: ISSUER_DID,
      privateKey: PRIVATE_KEY
    });
    console.log('Blockchain response:', response);

    // Then store in Firestore with the hash from blockchain
    const firestoreData = {
      id: diplomaId,
      type: diploma.value.type,
      holder: diploma.value.holder,
      issued_date: diploma.value.issued_date,
      expiry_date: "2030-05-27",
      degree_name: diploma.value.degree_name,
      grade_point: parseFloat(diploma.value.grade_point),
      hash: response.hash, // Use the hash from blockchain response
      university: university.value,
      issuer_did: ISSUER_DID,
      created_at: new Date().toISOString()
    };

    console.log('Storing in Firestore:', firestoreData);
    const docRef = await addDoc(
      collection(db, COLLECTION_NAME, uid.value, "diplomas"),
      firestoreData
    );
    console.log('Firestore document created with ID:', docRef.id);

    // Reset form and close dialog
    const currentPath = router.currentRoute.value.fullPath;
    diplomasStore.setAddDiplomaDialog(false);
    diploma.value = {
      type: "",
      holder: "",
      issued_date: "",
      degree_name: "",
      grade_point: 0,
    };

    // Refresh the page to show the new diploma
    router.replace({ path: '/_empty', query: { redirect: currentPath } })
      .then(() => {
        router.replace(currentPath);
      });

    q.notify({
      type: "positive",
      message: "Diploma added successfully!",
    });
  } catch (error) {
    console.error('Error adding diploma:', error);
    q.notify({
      type: "negative",
      message: error.response?.data?.error || "Failed to add diploma",
    });
  }
};
</script>

<template>
  <q-dialog
    persistent
    v-model="diplomasStore.getAddDiplomaDialog"
    @hide="diplomasStore.setAddDiplomaDialog(false)"
  >
    <q-card style="width: 600px; max-width: 80vw">
      <q-card-actions class="bg-primary">
        <div class="text-body1 text-white">{{ $t("addDiploma.add") }}</div>
        <q-space />
        <q-btn flat round dense icon="close" @click="reset()" color="white" />
      </q-card-actions>
      <q-card-section v-if="!isAddDiploma">
        <p class="text-body1 q-pb-none q-mt-md q-mb-xs text-weight-bold">
          Diploma info
        </p>
        <q-form class="q-mt-md" @submit.prevent="addDiploma">
          <q-input
            class="q-mt-sm"
            v-model="diploma.holder"
            label="Full Name"
            filled
            square
            dense
            bg-color="grey-2"
            :rules="[(val) => !!val || 'Full Name is required']"
            @update:model-value="(val) => console.log('Holder input updated:', val)"
          />
          <q-input
            class="q-mt-sm"
            v-model="diploma.type"
            :label="$t('diplomas.type')"
            filled
            square
            dense
            bg-color="grey-2"
            :rules="[(val) => !!val || 'Type is required']"
          />
          <q-input
            class="q-mt-sm"
            v-model="diploma.issued_date"
            :label="$t('diplomas.issuedDate')"
            filled
            square
            dense
            bg-color="grey-2"
            :rules="[(val) => !!val || 'Issued Date is required']"
          />
          <q-input
            class="q-mt-sm"
            v-model="diploma.degree_name"
            :label="$t('diplomas.degreeName')"
            filled
            square
            dense
            bg-color="grey-2"
            :rules="[(val) => !!val || 'Degree Name is required']"
          />
          <q-input
            class="q-mt-sm"
            v-model="diploma.grade_point"
            type="number"
            :label="$t('diplomas.gradePoint')"
            filled
            square
            dense
            bg-color="grey-2"
            :rules="[(val) => !!val || 'Grade Point is required']"
          />
        </q-form>
      </q-card-section>
      <q-card-section v-else>
        <div class="text-h6 text-center">{{ $t("diplomas.added") }}</div>
        <div class="flex justify-center q-mt-md">
          <q-avatar size="100px" color="green">
            <Icon name="carbon:checkmark" class="text-white" size="90" />
          </q-avatar>
        </div>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn
          flat
          :label="$t('addDiploma.cancel')"
          color="grey-8"
          unelevated
          no-caps
          @click="reset()"
        />
        <q-btn
          unelevated
          no-caps
          :label="$t('addDiploma.add')"
          color="primary"
          @click="addDiploma()"
          v-if="!isAddDiploma"
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<style scoped></style>
