import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import { collection, getDocs, query, orderBy, addDoc, updateDoc } from 'firebase/firestore'
import { useNuxtApp } from '#app'

export interface BaseCertificate {
  fullName: string
  registrationNumber: string
  municipality: string
  issuedDate: string
  expiryDate: string
}

export interface BirthCertificate extends BaseCertificate {
  fatherName: string
  motherName: string
  birthDate: string
  birthPlace: string
}

export interface DeathCertificate extends BaseCertificate {
  deathDate: string
  deathPlace: string
}

export interface MarriageCertificate extends BaseCertificate {
  spouseName: string
  marriageDate: string
  marriagePlace: string
}

export interface StateInfo {
  type: 'birthCertificate' | 'deathCertificate' | 'marriageCertificate'
  birthCertificate?: BirthCertificate
  deathCertificate?: DeathCertificate
  marriageCertificate?: MarriageCertificate
}

export interface State {
  stateHash: string
  stateInfo: StateInfo
  recordId?: string
}

export const useDocumentsStore = defineStore('documents', () => {
  const $q = useQuasar()
  const { $firebase } = useNuxtApp()
  const documents = ref<State[]>([])
  const showAddDocumentDialog = ref(false)
  const showViewDocumentDialog = ref(false)
  const showVerifyDocumentDialog = ref(false)
  const currentDocument = ref<State | null>(null)
  const selectedDocument = ref<State | null>(null)

  const setAddDocumentDialog = (value: boolean) => {
    showAddDocumentDialog.value = value
  }

  const setViewDocumentDialog = (value: boolean) => {
    showViewDocumentDialog.value = value
  }

  const setVerifyDocumentDialog = (value: boolean) => {
    showVerifyDocumentDialog.value = value
  }

  const setCurrentDocument = (document: State) => {
    currentDocument.value = document
  }

  const setSelectedDocument = (document: State) => {
    selectedDocument.value = document
    currentDocument.value = document
  }

  const addDocument = async (document: State) => {
    try {
      documents.value.push(document)
    } catch (error) {
      console.error('Error adding document:', error)
      throw error
    }
  }

  const verifyDocument = async (recordId: string, hash: string) => {
    try {
      const response = await fetch(`http://localhost:8080/documents/${recordId}?hash=${hash}`)
      if (!response.ok) {
        throw new Error('Verification failed')
      }
      return await response.json()
    } catch (error) {
      console.error('Error verifying document:', error)
      throw error
    }
  }

  return {
    documents,
    showAddDocumentDialog,
    showViewDocumentDialog,
    showVerifyDocumentDialog,
    currentDocument,
    selectedDocument,
    setViewDocumentDialog,
    setVerifyDocumentDialog,
    setAddDocumentDialog,
    setCurrentDocument,
    setSelectedDocument,
    addDocument,
    verifyDocument
  }
}) 