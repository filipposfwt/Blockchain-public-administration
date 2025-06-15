<template>
  <q-dialog :model-value="modelValue" @update:model-value="$emit('update:modelValue', $event)" persistent>
    <q-card style="width: 400px; max-width: 80vw">
      <q-card-actions class="bg-primary">
        <div class="text-body1 text-white">{{ $t('verifyDocument.title') }}</div>
        <q-space />
        <q-btn flat round dense icon="close" @click="onClose" color="white" />
      </q-card-actions>

      <q-card-section v-if="!isVerified">
        <q-form @submit="verify" class="q-gutter-y-md">
          <q-input
            v-model="hash"
            :label="$t('verifyDocument.inputLabel')"
            type="textarea"
            outlined
            :rules="[val => !!val || $t('verifyDocument.inputLabel') + ' ' + $t('login.errors.required')]"
            :hint="$t('verifyDocument.inputHint')"
            autogrow
          />
        </q-form>
      </q-card-section>

      <q-card-section v-else>
        <div class="text-h6 text-center">{{ $t('verifyDocument.verified') }}</div>
        <div class="flex justify-center q-mt-md">
          <q-avatar size="100px" color="green">
            <q-icon name="check" size="90px" color="white" />
          </q-avatar>
        </div>
        <div class="text-center q-mt-md">
          <p class="text-weight-bold">{{ $t('documents.recordId') }}:</p>
          <p class="font-mono q-pa-sm bg-grey-2 rounded-borders" style="word-break: break-all; white-space: pre-wrap;">{{ verificationResult.recordId }}</p>
          <p class="text-weight-bold q-mt-sm">{{ $t('documents.hash') }}:</p>
          <p class="font-mono q-pa-sm bg-grey-2 rounded-borders" style="word-break: break-all; white-space: pre-wrap;">{{ verificationResult.hash }}</p>
        </div>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat :label="$t('verifyDocument.cancel')" color="grey-8" unelevated no-caps @click="onClose" />
        <q-btn unelevated no-caps :label="$t('verifyDocument.verify')" color="primary" @click="verify()" v-if="!isVerified" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import { useDocumentsStore } from '../stores/documents'
import { useI18n } from 'vue-i18n'

const props = defineProps<{
  modelValue: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'close'): void
}>()

const $q = useQuasar()
const documentsStore = useDocumentsStore()
const { t } = useI18n()

const hash = ref('')
const isVerified = ref(false)
const error = ref('')
const verificationResult = ref(null)
const loading = ref(false)

const verify = async () => {
  try {
    loading.value = true
    
    // Parse the combined input
    const lines = hash.value.split('\n')
    let recordId = ''
    let hashValue = ''
    
    for (const line of lines) {
      if (line.startsWith('recordId:')) {
        recordId = line.replace('recordId:', '').trim()
      } else if (line.startsWith('hash:')) {
        hashValue = line.replace('hash:', '').trim()
      }
    }

    if (!recordId || !hashValue) {
      throw new Error(t('verifyDocument.invalidFormat'))
    }

    const response = await fetch(`http://localhost:8080/documents/${recordId}?hash=${hashValue}`)
    
    if (!response.ok) {
      throw new Error(t('verifyDocument.error'))
    }

    const data = await response.json()
    
    if (data.verified) {
      isVerified.value = true
      verificationResult.value = {
        recordId: recordId,
        hash: hashValue
      }
      $q.notify({
        type: 'positive',
        message: t('verifyDocument.success')
      })
    } else {
      $q.notify({
        type: 'negative',
        message: t('verifyDocument.error')
      })
    }
  } catch (error) {
    console.error('Verification error:', error)
    $q.notify({
      type: 'negative',
      message: error.message || t('verifyDocument.error')
    })
  } finally {
    loading.value = false
  }
}

const onClose = () => {
  emit('update:modelValue', false)
  emit('close')
  hash.value = ''
  error.value = ''
  verificationResult.value = null
  isVerified.value = false
}
</script>

<style scoped>
.font-mono {
  font-family: monospace;
}
</style>
