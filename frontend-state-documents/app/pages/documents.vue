<template>
  <q-page padding>
    <div class="bg-primary text-white q-pa-md q-mb-md">
      <div class="text-h4">Public Document Service</div>
    </div>

    <q-toolbar>
      <q-toolbar-title>{{ $t('documents.myDocuments') }}</q-toolbar-title>
      <q-space />
      <q-btn
        :label="$t('documents.add')"
        unelevated
        no-caps
        color="primary"
        icon="add"
        @click="documentsStore.setAddDocumentDialog(true)"
        class="q-mr-sm"
      />
      <q-btn
        :label="$t('documents.verify')"
        unelevated
        no-caps
        color="primary"
        icon="check"
        @click="documentsStore.setVerifyDocumentDialog(true)"
      />
    </q-toolbar>

    <div class="q-mt-md">
      <q-table
        :rows="documentsStore.documents"
        :columns="columns"
        row-key="stateHash"
        :loading="loading"
      >
        <template v-slot:body-cell-actions="props">
          <q-td :props="props">
            <q-btn-group flat>
              <q-btn
                flat
                round
                color="primary"
                icon="visibility"
                @click="viewDocument(props.row)"
              >
                <q-tooltip>{{ $t('documents.view') }}</q-tooltip>
              </q-btn>
              <q-btn
                flat
                round
                color="primary"
                icon="content_copy"
                @click="copyToClipboard(props.row)"
              >
                <q-tooltip>{{ $t('documents.copyVerification') }}</q-tooltip>
              </q-btn>
            </q-btn-group>
          </q-td>
        </template>
      </q-table>
    </div>

    <div class="fixed-bottom-left q-pa-md">
      <q-btn
        :label="$t('sidebar.logout')"
        unelevated
        no-caps
        color="negative"
        icon="logout"
        @click="handleLogout"
      />
    </div>

    <!-- Add Document Dialog -->
    <q-dialog v-model="documentsStore.showAddDocumentDialog">
      <q-card style="min-width: 500px">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">{{ $t('documents.add') }}</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>

        <q-card-section class="q-pt-md">
          <q-select
            v-model="selectedType"
            :options="documentTypes"
            :label="$t('documents.type')"
            emit-value
            map-options
            outlined
            class="q-mb-md"
          />

          <q-btn
            flat
            color="primary"
            icon="auto_fix_high"
            :label="$t('documents.generateRandom')"
            class="q-mb-md"
            @click="generateRandomData"
          />

          <div class="row q-col-gutter-md">
            <template v-if="selectedType === 'birthCertificate'">
              <div class="col-12 col-md-6">
                <q-input
                  v-model="birthCertificate.fullName"
                  :label="$t('documents.fullName')"
                  outlined
                  :placeholder="$t('documents.fullName')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="birthCertificate.registrationNumber"
                  :label="$t('documents.registrationNumber')"
                  outlined
                  :placeholder="$t('documents.registrationNumber')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="birthCertificate.municipality"
                  :label="$t('documents.municipality')"
                  outlined
                  :placeholder="$t('documents.municipality')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="birthCertificate.fatherName"
                  :label="$t('documents.fatherName')"
                  outlined
                  :placeholder="$t('documents.fatherName')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="birthCertificate.motherName"
                  :label="$t('documents.motherName')"
                  outlined
                  :placeholder="$t('documents.motherName')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="birthCertificate.birthDate"
                  :label="$t('documents.birthDate')"
                  type="date"
                  outlined
                />
              </div>
              <div class="col-12">
                <q-input
                  v-model="birthCertificate.birthPlace"
                  :label="$t('documents.birthPlace')"
                  outlined
                  :placeholder="$t('documents.birthPlace')"
                />
              </div>
            </template>

            <template v-if="selectedType === 'deathCertificate'">
              <div class="col-12 col-md-6">
                <q-input
                  v-model="deathCertificate.fullName"
                  :label="$t('documents.fullName')"
                  outlined
                  :placeholder="$t('documents.fullName')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="deathCertificate.registrationNumber"
                  :label="$t('documents.registrationNumber')"
                  outlined
                  :placeholder="$t('documents.registrationNumber')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="deathCertificate.municipality"
                  :label="$t('documents.municipality')"
                  outlined
                  :placeholder="$t('documents.municipality')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="deathCertificate.deathDate"
                  :label="$t('documents.deathDate')"
                  type="date"
                  outlined
                />
              </div>
              <div class="col-12">
                <q-input
                  v-model="deathCertificate.deathPlace"
                  :label="$t('documents.deathPlace')"
                  outlined
                  :placeholder="$t('documents.deathPlace')"
                />
              </div>
            </template>

            <template v-if="selectedType === 'marriageCertificate'">
              <div class="col-12 col-md-6">
                <q-input
                  v-model="marriageCertificate.fullName"
                  :label="$t('documents.fullName')"
                  outlined
                  :placeholder="$t('documents.fullName')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="marriageCertificate.registrationNumber"
                  :label="$t('documents.registrationNumber')"
                  outlined
                  :placeholder="$t('documents.registrationNumber')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="marriageCertificate.municipality"
                  :label="$t('documents.municipality')"
                  outlined
                  :placeholder="$t('documents.municipality')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="marriageCertificate.spouseName"
                  :label="$t('documents.spouseName')"
                  outlined
                  :placeholder="$t('documents.spouseName')"
                />
              </div>
              <div class="col-12 col-md-6">
                <q-input
                  v-model="marriageCertificate.marriageDate"
                  :label="$t('documents.marriageDate')"
                  type="date"
                  outlined
                />
              </div>
              <div class="col-12">
                <q-input
                  v-model="marriageCertificate.marriagePlace"
                  :label="$t('documents.marriagePlace')"
                  outlined
                  :placeholder="$t('documents.marriagePlace')"
                />
              </div>
            </template>

            <div class="col-12 col-md-6">
              <q-input
                v-model="issuedDate"
                :label="$t('documents.issuedDate')"
                type="date"
                outlined
              />
            </div>
            <div class="col-12 col-md-6">
              <q-input
                v-model="expiryDate"
                :label="$t('documents.expiryDate')"
                type="date"
                outlined
              />
            </div>
          </div>
        </q-card-section>

        <q-card-actions align="right" class="q-pa-md">
          <q-btn
            flat
            :label="$t('documents.cancel')"
            color="primary"
            v-close-popup
          />
          <q-btn
            unelevated
            :label="$t('documents.save')"
            color="primary"
            @click="saveDocument"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- View Document Dialog -->
    <q-dialog v-model="documentsStore.showViewDocumentDialog">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">{{ $t('documents.view') }}</div>
        </q-card-section>

        <q-card-section v-if="documentsStore.selectedDocument">
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <div class="text-subtitle2">{{ $t('documents.type') }}</div>
              <div>{{ 
                documentsStore.selectedDocument.stateInfo.type === 'birthCertificate' ? $t('documents.types.birthCertificate') :
                documentsStore.selectedDocument.stateInfo.type === 'deathCertificate' ? $t('documents.types.deathCertificate') :
                documentsStore.selectedDocument.stateInfo.type === 'marriageCertificate' ? $t('documents.types.marriageCertificate') :
                documentsStore.selectedDocument.stateInfo.type
              }}</div>
            </div>
            <div class="col-12">
              <div class="text-subtitle2">{{ $t('documents.stateHash') }}</div>
              <div>{{ documentsStore.selectedDocument.stateHash }}</div>
            </div>
            <template v-if="documentsStore.selectedDocument.stateInfo.type === 'birthCertificate'">
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.fullName') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.birthCertificate?.fullName }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.fatherName') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.birthCertificate?.fatherName }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.motherName') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.birthCertificate?.motherName }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.birthDate') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.birthCertificate?.birthDate }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.birthPlace') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.birthCertificate?.birthPlace }}</div>
              </div>
            </template>
            <template v-if="documentsStore.selectedDocument.stateInfo.type === 'deathCertificate'">
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.fullName') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.deathCertificate?.fullName }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.deathDate') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.deathCertificate?.deathDate }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.deathPlace') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.deathCertificate?.deathPlace }}</div>
              </div>
            </template>
            <template v-if="documentsStore.selectedDocument.stateInfo.type === 'marriageCertificate'">
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.fullName') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.marriageCertificate?.fullName }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.spouseName') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.marriageCertificate?.spouseName }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.marriageDate') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.marriageCertificate?.marriageDate }}</div>
              </div>
              <div class="col-12">
                <div class="text-subtitle2">{{ $t('documents.marriagePlace') }}</div>
                <div>{{ documentsStore.selectedDocument.stateInfo.marriageCertificate?.marriagePlace }}</div>
              </div>
            </template>
          </div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn
            flat
            :label="$t('documents.close')"
            color="primary"
            v-close-popup
          />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- Verify Document Dialog -->
    <q-dialog v-model="documentsStore.showVerifyDocumentDialog">
      <q-card style="min-width: 500px">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">{{ $t('documents.verify') }}</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>

        <q-card-section class="q-pt-md">
          <q-input
            v-model="verificationHash"
            :label="$t('verifyDocument.inputLabel')"
            :hint="$t('verifyDocument.inputHint')"
            outlined
            class="q-mb-md"
          />
        </q-card-section>

        <q-card-actions align="right" class="q-pa-md">
          <q-btn
            flat
            :label="$t('verifyDocument.cancel')"
            color="primary"
            v-close-popup
          />
          <q-btn
            unelevated
            :label="$t('verifyDocument.verify')"
            color="primary"
            @click="submitVerification"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- Verification Result Dialog -->
    <q-dialog v-model="showVerificationResult">
      <q-card style="min-width: 350px">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">{{ $t('verifyDocument.verified') }}</div>
          <q-space />
          <q-btn icon="close" flat round dense v-close-popup />
        </q-card-section>

        <q-card-section v-if="verificationResult" class="q-pt-md">
          <div class="row justify-center q-mb-md">
            <q-icon
              name="verified"
              color="positive"
              size="4rem"
            />
          </div>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <div class="text-subtitle2">{{ $t('documents.stateHash') }}</div>
              <div class="text-body1" style="word-break: break-all; white-space: pre-wrap;">{{ verificationResult.hash }}</div>
            </div>
            <div class="col-12">
              <div class="text-subtitle2">{{ $t('documents.recordId') }}</div>
              <div class="text-body1" style="word-break: break-all; white-space: pre-wrap;">{{ verificationResult.recordId }}</div>
            </div>
            <div class="col-12">
              <div class="text-subtitle2">{{ $t('verifyDocument.status') }}</div>
              <div class="text-body1 text-positive">
                {{ verificationResult.verified ? $t('verifyDocument.verified') : $t('verifyDocument.notVerified') }}
              </div>
            </div>
          </div>
        </q-card-section>

        <q-card-actions align="right" class="q-pa-md">
          <q-btn
            flat
            :label="$t('documents.close')"
            color="primary"
            v-close-popup
          />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { collection, getDocs, query, orderBy, addDoc, updateDoc } from 'firebase/firestore'
import VerifyDocument from '~/components/VerifyDocument.vue'
import { useFirebaseAuth } from '~/composables/useFirebaseAuth'
import { useRouter } from 'vue-router'
import { useDocumentsStore } from '~/stores/documents'

definePageMeta({
  layout: 'default'
})

const $q = useQuasar()
const { t } = useI18n()
const { $firebase } = useNuxtApp()
const documentsStore = useDocumentsStore()
const { signOut } = useFirebaseAuth()
const router = useRouter()

const documentTypes = [
  { label: t('documents.types.birthCertificate'), value: 'birthCertificate' },
  { label: t('documents.types.deathCertificate'), value: 'deathCertificate' },
  { label: t('documents.types.marriageCertificate'), value: 'marriageCertificate' }
]

const loading = ref(false)
const selectedType = ref('')
const verificationHash = ref('')

const birthCertificate = ref({
  fullName: '',
  registrationNumber: '',
  municipality: '',
  issuedDate: '',
  expiryDate: '',
  fatherName: '',
  motherName: '',
  birthDate: '',
  birthPlace: ''
})

const deathCertificate = ref({
  fullName: '',
  registrationNumber: '',
  municipality: '',
  issuedDate: '',
  expiryDate: '',
  deathDate: '',
  deathPlace: ''
})

const marriageCertificate = ref({
  fullName: '',
  registrationNumber: '',
  municipality: '',
  issuedDate: '',
  expiryDate: '',
  spouseName: '',
  marriageDate: '',
  marriagePlace: ''
})

const issuedDate = ref('')
const expiryDate = ref('')

const columns = [
  {
    name: 'type',
    label: t('documents.type'),
    field: (row: any) => {
      switch (row.stateInfo.type) {
        case 'birthCertificate':
          return t('documents.types.birthCertificate')
        case 'deathCertificate':
          return t('documents.types.deathCertificate')
        case 'marriageCertificate':
          return t('documents.types.marriageCertificate')
        default:
          return ''
      }
    },
    align: 'left'
  },
  {
    name: 'fullName',
    label: t('documents.fullName'),
    field: (row: any) => {
      switch (row.stateInfo.type) {
        case 'birthCertificate':
          return row.stateInfo.birthCertificate?.fullName
        case 'deathCertificate':
          return row.stateInfo.deathCertificate?.fullName
        case 'marriageCertificate':
          return row.stateInfo.marriageCertificate?.fullName
        default:
          return ''
      }
    },
    align: 'left'
  },
  {
    name: 'recordId',
    label: t('documents.recordId'),
    field: 'recordId',
    align: 'left'
  },
  {
    name: 'stateHash',
    label: t('documents.stateHash'),
    field: 'stateHash',
    align: 'left'
  },
  {
    name: 'actions',
    label: t('documents.actions'),
    field: 'actions',
    align: 'right'
  }
]

const showVerificationResult = ref(false)
const verificationResult = ref(null)

onMounted(async () => {
  await fetchDocuments()
})

const fetchDocuments = async () => {
  try {
    loading.value = true
    const docsRef = collection($firebase.db, 'documents')
    const q = query(docsRef, orderBy('createdAt', 'desc'))
    const snapshot = await getDocs(q)
    const docs = snapshot.docs.map(doc => ({
      id: doc.id,
      ...doc.data()
    }))
    console.log('Fetched documents:', docs) // Debug log
    documentsStore.documents = docs
  } catch (error) {
    console.error('Error fetching documents:', error)
    $q.notify({
      type: 'negative',
      message: t('documents.fetchError')
    })
  } finally {
    loading.value = false
  }
}

const viewDocument = (document: any) => {
  documentsStore.setSelectedDocument(document)
  documentsStore.setViewDocumentDialog(true)
}

const verifyDocument = (document: any) => {
  documentsStore.setSelectedDocument(document)
  documentsStore.setVerifyDocumentDialog(true)
}

const submitVerification = async () => {
  if (!verificationHash.value) {
    return
  }

  try {
    loading.value = true
    console.log('Input value:', verificationHash.value) // Debug log

    // Parse the input string more robustly
    const input = verificationHash.value.trim()
    const recordIdMatch = input.match(/recordId:\s*([^\s]+)/)
    const hashMatch = input.match(/hash:\s*([^\s]+)/)

    console.log('Matches:', { recordIdMatch, hashMatch }) // Debug log

    if (!recordIdMatch || !hashMatch) {
      throw new Error('Invalid input format')
    }

    const recordId = recordIdMatch[1]
    const hash = hashMatch[1]

    console.log('Extracted values:', { recordId, hash }) // Debug log
    
    const response = await fetch(`http://localhost:8080/documents/${recordId}?hash=${hash}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      }
    })

    if (!response.ok) {
      throw new Error('Verification failed')
    }

    const result = await response.json()
    verificationResult.value = result
    showVerificationResult.value = true
    documentsStore.setVerifyDocumentDialog(false)
    verificationHash.value = ''

    $q.notify({
      type: 'positive',
      message: t('verifyDocument.success')
    })
  } catch (error) {
    console.error('Verification error:', error)
    $q.notify({
      type: 'negative',
      message: t('verifyDocument.error')
    })
  } finally {
    loading.value = false
  }
}

const saveDocument = async () => {
  try {
    loading.value = true
    const document = {
      type: selectedType.value
    }

    switch (selectedType.value) {
      case 'birthCertificate':
        Object.assign(document, {
          birthCertificate: {
            ...birthCertificate.value,
            issuedDate: issuedDate.value,
            expiryDate: expiryDate.value
          }
        })
        break
      case 'deathCertificate':
        Object.assign(document, {
          deathCertificate: {
            ...deathCertificate.value,
            issuedDate: issuedDate.value,
            expiryDate: expiryDate.value
          }
        })
        break
      case 'marriageCertificate':
        Object.assign(document, {
          marriageCertificate: {
            ...marriageCertificate.value,
            issuedDate: issuedDate.value,
            expiryDate: expiryDate.value
          }
        })
        break
    }

    // Prepare request body with stateInfo wrapper and privateKey
    const requestBody = {
      stateInfo: document,
      privateKey: "bc1394ed5bb2bcd70e2a54daeb22440abacef3ff84c977521656ea26df312f27"
    }

    console.log('Sending request body:', requestBody) // Debug log

    // Send POST request to backend
    const response = await fetch('http://localhost:8080/documents', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestBody)
    })

    // Log the raw response for debugging
    const responseText = await response.text()
    console.log('Raw response:', responseText)

    if (!response.ok) {
      throw new Error(`Failed to create state record: ${responseText}`)
    }

    let result
    try {
      result = JSON.parse(responseText)
    } catch (parseError) {
      console.error('Error parsing response:', parseError)
      throw new Error('Invalid response format from server')
    }

    console.log('Backend response:', result) // Debug log

    // Add the document to Firestore with the hash and recordId from the backend
    const docRef = await addDoc(collection($firebase.db, 'documents'), {
      stateInfo: document,
      stateHash: result.hash,
      recordId: result.recordId,
      createdAt: new Date().toISOString()
    })

    console.log('Created Firestore document:', docRef.id) // Debug log

    $q.notify({
      type: 'positive',
      message: t('documents.addSuccess')
    })

    documentsStore.setAddDocumentDialog(false)
    resetForm()
    await fetchDocuments() // Refresh the documents list
  } catch (error) {
    console.error('Error saving document:', error)
    $q.notify({
      type: 'negative',
      message: error.message || t('documents.addError')
    })
  } finally {
    loading.value = false
  }
}

const resetForm = () => {
  selectedType.value = ''
  birthCertificate.value = {
    fullName: '',
    registrationNumber: '',
    municipality: '',
    issuedDate: '',
    expiryDate: '',
    fatherName: '',
    motherName: '',
    birthDate: '',
    birthPlace: ''
  }
  deathCertificate.value = {
    fullName: '',
    registrationNumber: '',
    municipality: '',
    issuedDate: '',
    expiryDate: '',
    deathDate: '',
    deathPlace: ''
  }
  marriageCertificate.value = {
    fullName: '',
    registrationNumber: '',
    municipality: '',
    issuedDate: '',
    expiryDate: '',
    spouseName: '',
    marriageDate: '',
    marriagePlace: ''
  }
  issuedDate.value = ''
  expiryDate.value = ''
}

const handleLogout = async () => {
  try {
    await signOut()
    $q.notify({
      type: 'positive',
      message: t('documents.logoutSuccess')
    })
    router.push('/login')
  } catch (error) {
    console.error('Logout error:', error)
    $q.notify({
      type: 'negative',
      message: t('documents.logoutError')
    })
  }
}

const generateRandomData = () => {
  const today = new Date()
  const newIssuedDate = today.toISOString().split('T')[0]
  const newExpiryDate = new Date(today.setFullYear(today.getFullYear() + 5)).toISOString().split('T')[0]

  const greekNames = [
    'Γεώργιος Παπαδόπουλος',
    'Μαρία Κωνσταντίνου',
    'Νικόλαος Αλεξίου',
    'Ελένη Δημητρίου',
    'Δημήτρης Γεωργίου'
  ]

  const greekPlaces = [
    'Αθήνα',
    'Θεσσαλονίκη',
    'Πάτρα',
    'Ηράκλειο',
    'Λάρισα'
  ]

  const randomName = greekNames[Math.floor(Math.random() * greekNames.length)]
  const randomPlace = greekPlaces[Math.floor(Math.random() * greekPlaces.length)]
  const randomRegNumber = Math.floor(Math.random() * 10000).toString().padStart(4, '0')

  switch (selectedType.value) {
    case 'birthCertificate':
      birthCertificate.value = {
        fullName: randomName,
        registrationNumber: `BC${randomRegNumber}`,
        municipality: randomPlace,
        issuedDate: newIssuedDate,
        expiryDate: newExpiryDate,
        fatherName: greekNames[Math.floor(Math.random() * greekNames.length)],
        motherName: greekNames[Math.floor(Math.random() * greekNames.length)],
        birthDate: new Date(today.setFullYear(today.getFullYear() - 30)).toISOString().split('T')[0],
        birthPlace: randomPlace
      }
      break
    case 'deathCertificate':
      deathCertificate.value = {
        fullName: randomName,
        registrationNumber: `DC${randomRegNumber}`,
        municipality: randomPlace,
        issuedDate: newIssuedDate,
        expiryDate: newExpiryDate,
        deathDate: new Date(today.setFullYear(today.getFullYear() - 1)).toISOString().split('T')[0],
        deathPlace: randomPlace
      }
      break
    case 'marriageCertificate':
      marriageCertificate.value = {
        fullName: randomName,
        registrationNumber: `MC${randomRegNumber}`,
        municipality: randomPlace,
        issuedDate: newIssuedDate,
        expiryDate: newExpiryDate,
        spouseName: greekNames[Math.floor(Math.random() * greekNames.length)],
        marriageDate: new Date(today.setFullYear(today.getFullYear() - 5)).toISOString().split('T')[0],
        marriagePlace: randomPlace
      }
      break
  }

  issuedDate.value = newIssuedDate
  expiryDate.value = newExpiryDate
}

const copyToClipboard = (document: any) => {
  const text = `recordId: ${document.recordId} hash: ${document.stateHash}`
  navigator.clipboard.writeText(text).then(() => {
    $q.notify({
      type: 'positive',
      message: t('documents.copiedToClipboard'),
      position: 'top',
      timeout: 2000
    })
  }).catch(() => {
    $q.notify({
      type: 'negative',
      message: t('documents.copyError'),
      position: 'top',
      timeout: 2000
    })
  })
}
</script>