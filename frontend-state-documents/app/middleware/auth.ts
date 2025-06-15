// middleware/auth.js
import { getAuth } from 'firebase/auth';

export default defineNuxtRouteMiddleware((to) => {
  const auth = getAuth();
  const user = auth.currentUser;

  // If user is not logged in and trying to access a protected route
  if (!user && to.path !== '/login' && to.path !== '/signup') {
    return navigateTo('/login');
  }

  // If user is logged in and trying to access auth pages
  if (user && (to.path === '/login' || to.path === '/signup')) {
    return navigateTo('/documents');
  }
});