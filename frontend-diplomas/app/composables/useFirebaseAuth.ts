// composables/useFirebaseAuth.js
import {
  onAuthStateChanged,
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
  signOut,
} from 'firebase/auth';

export const useFirebaseAuth = () => {
  const nuxtApp = useNuxtApp();
  const auth = nuxtApp.$auth; // Get the auth instance provided by the plugin

  const user = ref(null); // Reactive reference for the Firebase User object (null if not logged in)
  const loading = ref(true); // True while Firebase auth state is being determined
  const authReady = ref(false); // True once the initial auth state (logged in/out) is known





  // Method to create a new user with email and password
  const signUp = async (email, password) => {
    try {
      const userCredential = await createUserWithEmailAndPassword(auth, email, password);
      return userCredential.user;
    } catch (error) {
      throw error;
    }
  };

  // Method to sign in an existing user with email and password
  const signIn = async (email, password) => {
    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password);
      return userCredential.user;
    } catch (error) {
      throw error;
    }
  };

  const signOutUser = async () => {
    try {
      await signOut(auth);
    } catch (error) {
      throw error;
    }
  };

  // Return reactive state and methods to be used in Vue components
  return {
    user,      // Current Firebase User object (reactive)
    loading,   // True while initial auth state is being determined (reactive)
    authReady, // True once initial auth state is known (reactive)
    signUp,
    signIn,
    signOut: signOutUser, // Renamed to avoid potential global conflicts if any
  };
};