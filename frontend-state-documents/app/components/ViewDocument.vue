<script setup lang="ts">
import { getAuth, onAuthStateChanged } from "firebase/auth";
import { collection, doc, getDoc } from "firebase/firestore";

const documentsStore = useDocumentsStore();
const t = useI18n();
const nuxtApp = useNuxtApp();
const auth = nuxtApp.$auth;
const db = nuxtApp.$firestore;
const COLLECTION_NAME = "state"; // The name of your collection
const issuer = ref("");

onMounted(async () => {
  const user = auth.currentUser;
  if (user) {
    try {
      const idToken = await user.getIdToken();
      const docRef = doc(db, COLLECTION_NAME, user.uid);
      const docSnap = await getDoc(docRef);
      if (docSnap.exists()) {
        issuer.value = docSnap.data().issuer || "Δήμος";
      } else {
        issuer.value = "Δήμος";
      }
    } catch (error) {
      console.error("Error getting issuer:", error);
      issuer.value = "Δήμος";
    }
  }
});
</script>

<template>
  <q-dialog
    v-model="documentsStore.showViewDocumentDialog"
    @hide="documentsStore.setViewDocumentDialog(false)"
  >
    <q-card style="width: 220mm; max-width: 240mm">
      <q-card-actions class="bg-primary">
        <div class="text-body1 text-white">{{ $t("documents.view") }}</div>
        <q-space />
        <q-btn
          flat
          round
          dense
          icon="close"
          @click="documentsStore.setViewDocumentDialog(false)"
          color="white"
        />
      </q-card-actions>

      <q-card-section v-if="documentsStore.currentDocument">
        <div class="document-content">
          <div class="header-section">
            <h1 class="issuer-name">
              {{ issuer }}
            </h1>
          </div>

          <div class="main-content">
            <div class="document-details">
              <h2 class="document-title">{{ $t(`documents.types.${documentsStore.currentDocument.stateInfo.type}`) }}</h2>
              
              <div class="document-info">
                <p class="info-item">
                  <span class="label">{{ $t('documents.fullName') }}:</span>
                  <span class="value">{{ documentsStore.currentDocument.stateInfo[documentsStore.currentDocument.stateInfo.type]?.fullName }}</span>
                </p>
                <p class="info-item">
                  <span class="label">{{ $t('documents.registrationNumber') }}:</span>
                  <span class="value">{{ documentsStore.currentDocument.stateInfo[documentsStore.currentDocument.stateInfo.type]?.registrationNumber }}</span>
                </p>
                <p class="info-item">
                  <span class="label">{{ $t('documents.municipality') }}:</span>
                  <span class="value">{{ documentsStore.currentDocument.stateInfo[documentsStore.currentDocument.stateInfo.type]?.municipality }}</span>
                </p>
                <p class="info-item">
                  <span class="label">{{ $t('documents.issuedDate') }}:</span>
                  <span class="value">{{ documentsStore.currentDocument.stateInfo[documentsStore.currentDocument.stateInfo.type]?.issuedDate }}</span>
                </p>
                <p class="info-item">
                  <span class="label">{{ $t('documents.expiryDate') }}:</span>
                  <span class="value">{{ documentsStore.currentDocument.stateInfo[documentsStore.currentDocument.stateInfo.type]?.expiryDate }}</span>
                </p>

                <!-- Birth Certificate specific fields -->
                <template v-if="documentsStore.currentDocument.stateInfo.type === 'birthCertificate'">
                  <p class="info-item">
                    <span class="label">{{ $t('documents.fatherName') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.birthCertificate?.fatherName }}</span>
                  </p>
                  <p class="info-item">
                    <span class="label">{{ $t('documents.motherName') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.birthCertificate?.motherName }}</span>
                  </p>
                  <p class="info-item">
                    <span class="label">{{ $t('documents.birthDate') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.birthCertificate?.birthDate }}</span>
                  </p>
                  <p class="info-item">
                    <span class="label">{{ $t('documents.birthPlace') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.birthCertificate?.birthPlace }}</span>
                  </p>
                </template>

                <!-- Death Certificate specific fields -->
                <template v-if="documentsStore.currentDocument.stateInfo.type === 'deathCertificate'">
                  <p class="info-item">
                    <span class="label">{{ $t('documents.deathDate') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.deathCertificate?.deathDate }}</span>
                  </p>
                  <p class="info-item">
                    <span class="label">{{ $t('documents.deathPlace') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.deathCertificate?.deathPlace }}</span>
                  </p>
                </template>

                <!-- Marriage Certificate specific fields -->
                <template v-if="documentsStore.currentDocument.stateInfo.type === 'marriageCertificate'">
                  <p class="info-item">
                    <span class="label">{{ $t('documents.spouseName') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.marriageCertificate?.spouseName }}</span>
                  </p>
                  <p class="info-item">
                    <span class="label">{{ $t('documents.marriageDate') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.marriageCertificate?.marriageDate }}</span>
                  </p>
                  <p class="info-item">
                    <span class="label">{{ $t('documents.marriagePlace') }}:</span>
                    <span class="value">{{ documentsStore.currentDocument.stateInfo.marriageCertificate?.marriagePlace }}</span>
                  </p>
                </template>
              </div>
            </div>

            <div class="document-footer">
              <div class="document-id">
                <p class="text-caption text-grey-8 q-mb-xs" style="font-size: 12px">
                  {{ $t('documents.recordId') }}: {{ documentsStore.currentDocument.recordId }}
                </p>
                <p class="text-caption text-grey-8 q-mb-xs" style="font-size: 12px">
                  {{ $t('documents.stateHash') }}: {{ documentsStore.currentDocument.stateHash }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<style scoped>
.document-content {
  padding: 2rem;
}

.header-section {
  text-align: center;
  margin-bottom: 2rem;
}

.issuer-name {
  font-size: 1.5rem;
  color: #333;
  margin-bottom: 0.5rem;
}

.main-content {
  margin: 2rem 0;
}

.document-details {
  margin-bottom: 2rem;
}

.document-title {
  font-size: 1.8rem;
  color: #333;
  text-align: center;
  margin-bottom: 2rem;
}

.document-info {
  display: grid;
  gap: 1rem;
}

.info-item {
  display: flex;
  gap: 1rem;
  margin-bottom: 0.5rem;
}

.label {
  font-weight: bold;
  min-width: 150px;
  color: #666;
}

.value {
  color: #333;
}

.document-footer {
  margin-top: 2rem;
  padding-top: 1rem;
  border-top: 1px solid #eee;
}

.document-id {
  text-align: center;
}
</style>
