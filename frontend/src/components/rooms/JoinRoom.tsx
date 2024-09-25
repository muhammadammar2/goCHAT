import React, { useState, useEffect } from "react";
import { apiClient } from "../../api/apiClient";
// import apiClient from "../../api/apiClient";

function JoinRoom() {
  const [rooms, setRooms] = useState<any[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    async function fetchRooms() {
      try {
        const response = await apiClient.get("/rooms");
        setRooms(response.data);
      } catch (err: any) {
        console.error(err.response?.data || err.message);
        setError("Failed to fetch rooms. Please try again.");
      }
    }
    fetchRooms();
  }, []);

  return (
    <div className="w-full max-w-md">
      <div className="bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4">
        <h2 className="text-3xl mb-6 text-center font-bold text-blue-400">
          Join a Room
        </h2>
        {error && <p className="text-red-500 text-center mb-4">{error}</p>}
        <ul>
          {rooms.length > 0 ? (
            rooms.map((room: any) => (
              <li key={room.id} className="mb-4">
                <div className="bg-gray-700 p-4 rounded">
                  <h3 className="text-xl font-bold text-white">{room.name}</h3>
                  <p className="text-gray-300">{room.description}</p>
                  <p className="text-gray-500">
                    Type: {room.private ? "Private" : "Public"}
                  </p>
                  <button className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-2">
                    Join Room
                  </button>
                </div>
              </li>
            ))
          ) : (
            <p className="text-gray-300 text-center">No rooms available.</p>
          )}
        </ul>
      </div>
    </div>
  );
}

export default JoinRoom;
