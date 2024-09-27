import React, { useEffect } from "react";
import apiClient from "../api/apiClient";

const FetchDataComponent: React.FC = () => {
  useEffect(() => {
    const fetchData = async () => {
      const token = localStorage.getItem("token");
      try {
        const response = await apiClient.get("/api/resource", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        console.log(response.data);
      } catch (error) {
        console.error("Error:", error);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      <h1>Data Fetch Example</h1>
      <p>Check the console for the fetched data.</p>
    </div>
  );
};

export default FetchDataComponent;
