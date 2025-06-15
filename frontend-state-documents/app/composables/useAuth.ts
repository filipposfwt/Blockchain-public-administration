import { ref } from 'vue'
import { signInWithEmailAndPassword, createUserWithEmailAndPassword, signOut } from 'firebase/auth'
import { useStateUserStore } from '~/stores/stateUser'

export const useAuth = () => {
  const { $firebase } = useNuxtApp()
  const userStore = useStateUserStore()
  const loading = ref(false)
  const error = ref<string | null>(null)

  const login = async (email: string, password: string) => {
    try {
      loading.value = true
      error.value = null
      
      const userCredential = await signInWithEmailAndPassword($firebase.auth, email, password)
      userStore.setUser(userCredential.user)
      await userStore.fetchUser(userCredential.user.uid)
    } catch (e: any) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  const register = async (email: string, password: string, userData: any) => {
    try {
      loading.value = true
      error.value = null
      
      const userCredential = await createUserWithEmailAndPassword($firebase.auth, email, password)
      const userId = await userStore.createUser({
        ...userData,
        email,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString()
      })
      
      userStore.setUser(userCredential.user)
      return userId
    } catch (e: any) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    try {
      loading.value = true
      error.value = null
      
      await signOut($firebase.auth)
      userStore.clearUser()
    } catch (e: any) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  return {
    login,
    register,
    logout,
    loading,
    error
  }
} 