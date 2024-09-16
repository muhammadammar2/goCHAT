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
    try {
      await apiClient.post("/create-room", {
        name,
        description,
        private: isPrivate,
        code: isPrivate ? code : undefined,
      });
      navigate("/room-options");
    } catch (err: any) {
      console.error(err.response?.data || err.message);
      setError("Failed to create room. Please try again.");
    }
  };

  return (
    <div className="w-full max-w-md">
      <form
        onSubmit={handleSubmit}
        className="bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4"
      >
        <h2 className="text-3xl mb-6 text-center font-bold text-blue-400">
          Create a Room
        </h2>
        {error && <p className="text-red-500 text-center mb-4">{error}</p>}
        <div className="mb-4">
          <input
            className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
            type="text"
            placeholder="Room Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="mb-4">
          <textarea
            className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
            placeholder="Description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            required
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-300">
            <input
              type="radio"
              name="roomType"
              value="public"
              checked={!isPrivate}
              onChange={() => setIsPrivate(false)}
              className="mr-2"
            />
            Public
          </label>
          <label className="block text-gray-300">
            <input
              type="radio"
              name="roomType"
              value="private"
              checked={isPrivate}
              onChange={() => setIsPrivate(true)}
              className="mr-2"
            />
            Private
          </label>
        </div>
        {isPrivate && (
          <div className="mb-4">
            <input
              className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
              type="text"
              placeholder="Set the Code"
              value={code}
              onChange={(e) => setCode(e.target.value)}
              required
            />
          </div>
        )}
        <div className="flex items-center justify-center mb-4">
          <button
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300"
            type="submit"
          >
            Create Room
          </button>
        </div>
      </form>
    </div>
  );
}

export default CreateRoomForm;
