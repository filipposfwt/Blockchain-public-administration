<script setup lang="ts">
import { verifyDiploma } from "~/repository/modules/diplomas/factory";
const diplomasStore = useDiplomasStore();
const t = useI18n();
const q = useQuasar();
const hash = ref('');
const isVerifyDiploma = ref(false);
const error = ref('');
const issuerDID = ref('');

const verify = async () => {
  try {
    error.value = '';
    // Validate hash format
    if (!hash.value || hash.value.length !== 64) {
      error.value = 'Please enter a valid 64-character hash';
      q.notify({
        type: 'negative',
        message: 'Please enter a valid 64-character hash',
      });
      return;
    }

    console.log('Starting verification process...');
    console.log('Hash to verify:', hash.value);
    console.log('Making API call...');
    
    const response = await verifyDiploma(hash.value);
    console.log('Raw verification response:', response);

    if (response && response.verified) {
      console.log('Verification successful');
      hash.value = '';
      isVerifyDiploma.value = true;
      issuerDID.value = response.issuerDID;
      q.notify({
        type: 'positive',
        message: 'Diploma verified successfully!',
      });
    } else {
      console.log('Verification failed - diploma not found');
      error.value = 'Diploma not found';
      q.notify({
        type: 'negative',
        message: 'Diploma not found',
      });
    }
  } catch (err) {
    console.error('Verification error caught:');
    console.error('Error object:', err);
    console.error('Error response:', err.response);
    console.error('Error data:', err.response?.data);
    error.value = err.response?.data?.error || 'An error occurred while verifying the diploma';
    q.notify({
      type: 'negative',
      message: error.value,
    });
  }
};

const reset = () => {
  hash.value = '';
  error.value = '';
  issuerDID.value = '';
  isVerifyDiploma.value = false;
  diplomasStore.setVerifyDiplomaDialog(false);
};
</script>

<template>
<q-dialog v-model="diplomasStore.getVerifyDiplomaDialog"> 
  <q-card style="width: 400px; max-width: 80vw">
    <q-card-actions class="bg-primary">
      <div class="text-body1 text-white">{{ $t("diplomas.verify") }}</div>
      <q-space />
      <q-btn flat round dense icon="close" @click="reset()" color="white" />
    </q-card-actions>
    <q-card-section v-if="!isVerifyDiploma">
      <q-form @submit.prevent="verify">
        <p class="q-mb-xs text-weight-bold">{{ $t("diplomas.hash") }}</p>
        <q-input
          v-model="hash"
          filled
          square
          dense
          bg-color="grey-2"
          placeholder="Enter Diploma Hash"
          type="text"
          :error="!!error"
          :error-message="error"
          :rules="[val => val.length === 64 || 'Hash must be 64 characters long']"
        />
      </q-form>
    </q-card-section>
    <q-card-section v-else>
      <div class="text-h6 text-center">{{ $t("diplomas.verified") }}</div>
      <div class="flex justify-center q-mt-md">
        <q-avatar size="100px" color="green">
          <Icon name="carbon:checkmark" class="text-white" size="90" />
        </q-avatar>
      </div>
      <div class="text-center q-mt-md">
        <p class="text-weight-bold">Issuer DID:</p>
        <p>{{ issuerDID }}</p>
      </div>
    </q-card-section>
    <q-card-actions align="right">  
      <q-btn flat :label="$t('diplomas.cancel')" color="grey-8" unelevated no-caps @click="reset()" />
      <q-btn unelevated no-caps :label="$t('diplomas.verify')" color="primary" @click="verify()" v-if="!isVerifyDiploma" />
    </q-card-actions>
  </q-card>
</q-dialog>
</template>

<style scoped></style>
