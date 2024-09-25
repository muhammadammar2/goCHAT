import React, { useEffect } from "react";
// import apiClient from "../api/apiClient";
import { apiClient } from "../api/apiClient"; // Import the apiRequest function

const FetchDataComponent: React.FC = () => {
  useEffect(() => {
    apiClient
      .get("/api/resource", {
        headers: {
          Authorization: "12hg3v1h23vh12v3h1v3gh12",
        },
      })
      .then((response) => console.log(response.data))
      .catch((error) => console.error("Error:", error));
  }, []);

  return (
    <div>
      <h1>Data Fetch Example</h1>
      <p>Check the console for the fetched data.</p>
    </div>
  );
};

export default FetchDataComponent;
