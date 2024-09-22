import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

apiClient.interceptors.request.use(
  (config) => {
    console.log("Outgoing request config:", config);

    const token = localStorage.getItem("token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
      console.log("Token attached to request:", `Bearer {token}`); //k
    } else {
      console.log("No token found in localStorage");
    }
    console.log("Request confid", config);
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default apiClient;
