



import { initializeApp } from 'firebase/app'
import { getAuth } from "firebase/auth"
import { getFirestore } from 'firebase/firestore'
import { getAnalytics } from "firebase/analytics"
const firebaseConfig = {
  apiKey: "AIzaSyCyZagqvoTm4zUgv_93n9Kc8ULW6ajTXEI",
  authDomain: "dms-university.firebaseapp.com",
  projectId: "dms-university",
  storageBucket: "dms-university.firebasestorage.app",
  messagingSenderId: "649306189814",
  appId: "1:649306189814:web:1ecc19efdb8542d6c7e1d7"
};

export default defineNuxtPlugin(nuxtApp => {
 
  const app = initializeApp(firebaseConfig)
  const analytics = getAnalytics(app)
  const auth = getAuth(app)
  const firestore = getFirestore(app)
  nuxtApp.vueApp.provide('auth', auth)
  nuxtApp.provide('auth', auth)
  nuxtApp.vueApp.provide('firestore', firestore)
  nuxtApp.provide('firestore', firestore)
})