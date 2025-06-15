import { initializeApp, getApps, getApp } from 'firebase/app';
import { getAuth, browserLocalPersistence } from 'firebase/auth';
import { getFirestore } from 'firebase/firestore';

export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig();

  const firebaseConfig = {
    apiKey: config.public.firebaseApiKey,
    authDomain: config.public.firebaseAuthDomain,
    projectId: config.public.firebaseProjectId,
    storageBucket: config.public.firebaseStorageBucket,
    messagingSenderId: config.public.firebaseMessagingSenderId,
    appId: config.public.firebaseAppId
  };

  // Initialize Firebase only if it hasn't been initialized yet
  const app = getApps().length === 0 ? initializeApp(firebaseConfig) : getApp();
  const auth = getAuth(app);
  const db = getFirestore(app);

  // Configure auth persistence
  auth.setPersistence(browserLocalPersistence);

  nuxtApp.vueApp.config.globalProperties.$firebase = {
    app,
    auth,
    db
  };

  return {
    provide: {
      firebase: {
        app,
        auth,
        db
      }
    }
  };
}); 