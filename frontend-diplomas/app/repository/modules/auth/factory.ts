import type {
    LoginCredentials,
    RegistrationRequest,
    LoginResponse,
    UserResponse,
    BackendTokenResponse
  } from "@/repository/modules/auth/types";
  
  const login = async (credentials: LoginCredentials): Promise<LoginResponse> => {
    try {
      const response = await $api_trap.post<BackendTokenResponse>(
        "/api/auth/login",
        credentials
      );
      if(response.status === 200){
        localStorage.setItem("trap_token", response.data.accessToken);
        localStorage.setItem("trap_token_expires_in", response.data.expiresIn);
        localStorage.setItem("trap_token_type", response.data.tokenType);
        return response;
      }
  
    } catch (error) {
      return error;
    }
  };
  
  const register = async (data: RegistrationRequest): Promise<LoginResponse> => {
    try {
      const response = await $api_trap.post<BackendTokenResponse>(
        "/api/auth/register",
        data
      );
      console.log("Registration response:", response);
      return adaptTokenResponse(response.data);
    } catch (error) {
      console.error("Registration error:", error);
      throw error;
    }
  };
  
  const validateToken = async (): Promise<boolean> => {
    try {
      await $api_trap.get("/api/auth/validate");
      return true;
    } catch (error) {
      console.error("Token validation error:", error);
      return false;
    }
  };
  
  const getCurrentUser = async (): Promise<UserResponse> => {
    try {
      const response = await $api_trap.get<UserResponse>("/api/users/me");
      return response.data;
    } catch (error) {
      console.error("Error fetching current user:", error);
      throw error;
    }
  };
  
  export { login, register, validateToken, getCurrentUser };
  