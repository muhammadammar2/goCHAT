import React, { useState } from "react";
import apiClient from "../../api/apiClient";
import { useNavigate } from "react-router-dom";

function CreateRoomForm() {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [isPrivate, setIsPrivate] = useState(false);
  const [code, setCode] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Creating room with token:", localStorage.getItem("token")); // Log the token

    try {
      console.log("Sending request to create room with data:", {
        name,
        description,
        type: isPrivate ? "private" : "public",
        code: isPrivate ? code : undefined,
      });

      await apiClient.post("/create-room", {
        name,
        description,
        type: isPrivate ? "private" : "public",
        code: isPrivate ? code : undefined,
      });

      console.log("Room successfully created!");
      navigate("/profile");
    } catch (err: any) {
      console.error(
        "Error occurred while creating room:",
        err.response?.data || err.message
      );
      setError("Failed to create room. Please try again.");
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-900 px-4">
      <div className="w-full max-w-lg">
        <form
          onSubmit={handleSubmit}
          className="bg-gray-800 shadow-xl rounded-lg p-6 sm:p-8"
        >
          <h2 className="text-2xl sm:text-3xl mb-6 text-center font-semibold text-blue-500">
            Create a Room
          </h2>
          {error && (
            <p className="text-red-500 text-center mb-4 animate-pulse">
              {error}
            </p>
          )}
          <div className="mb-4">
            <input
              className="w-full p-3 sm:p-4 rounded-lg bg-gray-700 text-gray-300 placeholder-gray-500 border-2 border-gray-700 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
              type="text"
              placeholder="Room Name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
          <div className="mb-4">
            <textarea
              className="w-full p-3 sm:p-4 rounded-lg bg-gray-700 text-gray-300 placeholder-gray-500 border-2 border-gray-700 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200 resize-none"
              placeholder="Description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              required
              rows={4}
            />
          </div>
          <div className="mb-4 flex items-center justify-between">
            <span className="text-gray-300">Room Type:</span>
            <label className="relative inline-flex items-center cursor-pointer">
              <input
                type="checkbox"
                checked={isPrivate}
                onChange={() => setIsPrivate(!isPrivate)}
                className="sr-only peer"
              />
              <div className="w-11 h-6 bg-gray-700 rounded-full peer peer-focus:ring-2 peer-focus:ring-blue-500 dark:bg-gray-700 peer-checked:bg-blue-600 after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:after:translate-x-full peer-checked:after:border-white"></div>
              <span className="ml-3 text-gray-300">
                {isPrivate ? "Private" : "Public"}
              </span>
            </label>
          </div>
          {isPrivate && (
            <div className="mb-4 transition-all duration-300 ease-in-out">
              <input
                className="w-full p-3 sm:p-4 rounded-lg bg-gray-700 text-gray-300 placeholder-gray-500 border-2 border-gray-700 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
                type="text"
                placeholder="Set the Code"
                value={code}
                onChange={(e) => setCode(e.target.value)}
                required
              />
            </div>
          )}
          <div className="flex justify-center mb-4">
            <button
              className="w-full sm:w-auto bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-300"
              type="submit"
            >
              Create Room
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}

export default CreateRoomForm;
