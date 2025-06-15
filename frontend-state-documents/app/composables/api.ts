import axios from "axios";

const config = useRuntimeConfig();
const $api_dms_diplomas = axios.create({
  baseURL: "http://localhost:8080",
});

$api_dms_diplomas.interceptors.request.use(
  (config) => {
    console.log('Making request to:', config.url);
    console.log('Request config:', config);
    const token = localStorage.getItem("trap_token");
    if (token !== null) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    console.error('Request error:', error);
    return Promise.reject(error);
  }
);

$api_dms_diplomas.interceptors.response.use(
  (response) => {
    console.log('Response received:', response);
    return response;
  },
  (error) => {
    console.error('Response error:', error);
    console.error('Error config:', error.config);
    console.error('Error response:', error.response);
    return Promise.reject(error);
  }
);
export { $api_dms_diplomas };
