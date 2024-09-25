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

import axios from "axios";

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
const apiRequest = async (config: any) => {
  const token = getAuthToken();

  if (token) {
    config.headers["Authorization"] = token;
  }

  return await apiClient(config);
};

export { apiClient, apiRequest };
