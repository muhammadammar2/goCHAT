import React from "react";
import { useNavigate } from "react-router-dom";

function RoomOptions() {
  const navigate = useNavigate();

  const handleCreateRoom = () => {
    navigate("/create-room");
  };

  const handleJoinRoom = () => {
    navigate("/join-room");
  };

  return (
    <div className="w-full max-w-md">
      <div className="bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4">
        <h2 className="text-3xl mb-6 text-center font-bold text-blue-400">
          Room Options
        </h2>
        <div className="flex flex-col items-center">
          <button
            onClick={handleCreateRoom}
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300 mb-4"
          >
            Create a Room
          </button>
          <button
            onClick={handleJoinRoom}
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300"
          >
            Join a Room
          </button>
        </div>
      </div>
    </div>
  );
}

export default RoomOptions;
