import { defineStore } from "pinia";
import { ref, computed } from "vue";
import type { User } from "firebase/auth";
import { collection, getDocs, query, where, doc, getDoc, addDoc, updateDoc } from "firebase/firestore";

export interface StateUser {
  id: string;
  email: string;
  organization: string;
  role: string;
  createdAt: string;
  updatedAt: string;
}

export const useStateUserStore = defineStore("stateUser", () => {
  const user = ref<StateUser | null>(null);
  const authUser = ref<User | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const isAuthenticated = computed(() => !!authUser.value);

  const setUser = (user: User | null) => {
    authUser.value = user;
  };

  const fetchUser = async (userId: string) => {
    try {
      loading.value = true;
      error.value = null;
      const { $firebase } = useNuxtApp();
      
      const userDoc = await getDoc(doc($firebase.db, "state_users", userId));
      if (userDoc.exists()) {
        user.value = {
          id: userDoc.id,
          ...userDoc.data()
        } as StateUser;
      } else {
        error.value = "User not found";
      }
    } catch (err) {
      console.error("Error fetching user data:", err);
      error.value = "Failed to fetch user data";
    } finally {
      loading.value = false;
    }
  };

  const createUser = async (userData: Omit<StateUser, "id">) => {
    try {
      loading.value = true;
      error.value = null;
      const { $firebase } = useNuxtApp();
      
      const docRef = await addDoc(collection($firebase.db, "state_users"), userData);
      user.value = {
        id: docRef.id,
        ...userData
      } as StateUser;
      return docRef.id;
    } catch (err) {
      console.error("Error creating user:", err);
      error.value = "Failed to create user";
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const updateUser = async (userId: string, userData: Partial<StateUser>) => {
    try {
      loading.value = true;
      error.value = null;
      const { $firebase } = useNuxtApp();
      
      const userRef = doc($firebase.db, "state_users", userId);
      await updateDoc(userRef, userData);
      
      if (user.value) {
        user.value = {
          ...user.value,
          ...userData
        } as StateUser;
      }
    } catch (err) {
      console.error("Error updating user:", err);
      error.value = "Failed to update user";
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const clearUser = () => {
    user.value = null;
    authUser.value = null;
    error.value = null;
  };

  return {
    user,
    authUser,
    loading,
    error,
    isAuthenticated,
    setUser,
    fetchUser,
    createUser,
    updateUser,
    clearUser
  };
}); 