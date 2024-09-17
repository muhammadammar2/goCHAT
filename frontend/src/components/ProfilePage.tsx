// components/ProfilePage.tsx
import apiClient from "../api/apiClient";
import { useAuth } from "../auth/AuthContext";

// components/ProfilePage.tsx
import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const ProfilePage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
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
        setLoading(false);
      })
      .catch((error) => {
        console.error("Error fetching profile data:", error);
        setError("Failed to load profile data.");
        setLoading(false);
      });
  }, [isAuthenticated, navigate]);

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    apiClient
      .put("/update-profile", { username, name, password })
      .then((response) => {
        alert("Profile updated successfully!");
      })
      .catch((error) => {
        console.error("Error updating profile:", error);
        setError("Failed to update profile.");
      });
  };

  if (loading) return <p>Loading...</p>;

  return (
    <div className="profile-page min-h-screen bg-gray-900 text-white p-8">
      <div className="w-full max-w-md mx-auto bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4">
        <h1 className="text-3xl font-bold mb-6">Edit Profile</h1>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <form onSubmit={handleSubmit} className="space-y-4">
          <label className="block">
            Username:
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="mt-1 p-2 border border-gray-600 rounded w-full"
            />
          </label>
          <label className="block">
            Name:
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="mt-1 p-2 border border-gray-600 rounded w-full"
            />
          </label>
          <label className="block">
            Password:
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="mt-1 p-2 border border-gray-600 rounded w-full"
            />
          </label>
          <button
            type="submit"
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >
            Update Profile
          </button>
        </form>
      </div>
    </div>
  );
};

export default ProfilePage;
