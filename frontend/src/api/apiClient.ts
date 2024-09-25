// import axios from "axios";

// const API_BASE_URL = "http://localhost:8080";

// const apiClient = axios.create({
//   baseURL: API_BASE_URL,
//   headers: {
//     "Content-Type": "application/json",
//   },
// });

// // Attach request interceptor only once
// apiClient.interceptors.request.use(
//   (config) => {
//     console.log("Outgoing request config (before attaching token):", config);

//     // Get the token from localStorage
//     const token = localStorage.getItem("token");
//     if (token) {
//       config.headers["Authorization"] = `Bearer ${token}`;
//       console.log("Token attached to request:", `Bearer ${token}`);
//     } else {
//       console.log("No token found in localStorage");
//     }

//     // Log the final request config
//     console.log("Final request config:", config);
//     return config;
//   },
//   (error) => {
//     console.error("Error in request interceptor:", error);
//     return Promise.reject(error);
//   }
// );

// export default apiClient;

import axios, { AxiosRequestConfig } from "axios";

const API_BASE_URL = "http://localhost:8080";

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Function to get the authorization token
const getAuthToken = () => {
  const token = localStorage.getItem("token");
  return token ? `Bearer ${token}` : null;
};

// Custom function for making API requests with token
const apiRequest = async (config: AxiosRequestConfig) => {
  // Ensure the config is an object
  const requestConfig: AxiosRequestConfig = { ...config }; // Clone config to avoid mutation
  const token = getAuthToken();

  // Attach token if it exists
  if (token) {
    requestConfig.headers = {
      ...requestConfig.headers, // Merge existing headers if any
      Authorization: token,
    };
  }

  try {
    return await apiClient(requestConfig);
  } catch (error) {
    console.error("API request error:", error);
    throw error; // Rethrow the error for further handling
  }
};

export { apiClient, apiRequest };
