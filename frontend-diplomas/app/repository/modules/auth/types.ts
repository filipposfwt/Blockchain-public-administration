export type LoginCredentials = {
    email: string;
    password: string;
  }
  
  export type RegistrationRequest = {
    email: string;
    password: string;
  }
  
  export type LoginResponse = {
    access_token: string;
    token_type: string;
    expires_in: number;
  }
  
  export type UserResponse = {
    id: number;
    email: string;
    enabled: boolean;
    firstName?: string;
    lastName?: string;
    mobilePhone?: string;
    afm?: string;
  }
  
  // Interface that matches the backend response
  export type BackendTokenResponse = {
    accessToken: string;
    tokenType: string;
    expiresIn: number;
  }
  
  
  
  