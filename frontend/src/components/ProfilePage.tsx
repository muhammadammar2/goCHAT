import apiClient from "../api/apiClient";
import { useAuth } from "../auth/AuthContext";
import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const ProfilePage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!isAuthenticated) {
      navigate("/login");
      return;
    }

    apiClient
      .get("/profile")
      .then((response) => {
        setUsername(response.data.username);
        setName(response.data.name);
        setEmail(response.data.email);
        setLoading(false);
      })
      .catch((error) => {
        console.error("Error fetching profile data:", error);
        setError("Failed to load profile data.");
        setLoading(false);
      });
  }, [isAuthenticated, navigate]);

  if (loading) return <p className="text-center text-gray-400">Loading...</p>;

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-900 text-white p-8">
      <div className="flex flex-col items-center bg-gray-800 rounded-lg p-8 shadow-lg max-w-md w-full">
        <h1 className="text-3xl font-bold text-center mb-6">Your Profile</h1>
        {error && <p className="text-red-500 mb-4 text-center">{error}</p>}

        <div className="text-lg mb-4">
          <strong>Username:</strong>{" "}
          <span className="text-gray-300">{username}</span>
        </div>
        <div className="text-lg mb-4">
          <strong>Name:</strong> <span className="text-gray-300">{name}</span>
        </div>
        <div className="text-lg mb-4">
          <strong>Email:</strong> <span className="text-gray-300">{email}</span>
        </div>

        <button
          onClick={() => navigate("/update-profile")}
          className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded transition duration-300 mt-4"
        >
          Edit Profile
        </button>
      </div>
    </div>
  );
};

export default ProfilePage;
