// middleware/auth.js
export default defineNuxtRouteMiddleware(async (to, from) => {
  const accessToken = localStorage.getItem("dms_accessToken");
 
  if (to.path === '/login' && accessToken) {
    return navigateTo('/diplomas')
  }

  // Handle the /login page itself: if already logged in, redirect away from it
  if (to.path === '/login') {
    if (accessToken) {
      console.log(`[AuthMiddleware] Already logged in (user: ${user.value.uid}) on /login. Redirecting to /diplomas.`);
      return navigateTo('/diplomas'); // Or '/diplomas' if that's the primary authenticated page
    }
    // If not logged in, or auth not ready yet, allow access to /login page
    return;
  }


  // No token at all? Redirect to login
  if (!accessToken) {
      return navigateTo('/login')
  }

});