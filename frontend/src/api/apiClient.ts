import axios from "axios";

const API_BASE_URL = "http://localhost:6969";

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

apiClient.interceptors.request.use(
  (config) => {
    console.log("Outgoing request config (before attaching token):", config);

    const token = localStorage.getItem("token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
      console.log("Token attached to request:", `Bearer ${token}`);
    } else {
      console.log("No token found in localStorage");
    }

    console.log("Final request config:", config);
    return config;
  },
  (error) => {
    console.error("Error in request interceptor:", error);
    return Promise.reject(error);
  }
);

export default apiClient;
