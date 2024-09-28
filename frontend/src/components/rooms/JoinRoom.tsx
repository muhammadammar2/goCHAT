import React, { useState, useEffect } from "react";
import apiClient from "../../api/apiClient";

function JoinRoom() {
  const [rooms, setRooms] = useState<any[]>([]);
  const [error, setError] = useState("");
  const [roomCode, setRoomCode] = useState<{ [key: string]: string }>({});

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

  const handleJoinRoom = async (roomId: string, roomType: string) => {
    const code = roomType === "private" ? roomCode[roomId] : undefined;
    try {
      const response = await apiClient.post("/join-room", {
        room_id: roomId,
        code,
      });
      alert(response.data.message);
    } catch (err: any) {
      console.error(err.response?.data || err.message);
      setError(
        err.response?.data?.error || "Failed to join room. Please try again."
      );
    }
  };

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
              <li key={room.ID} className="mb-4">
                <div className="bg-gray-700 p-4 rounded">
                  <h3 className="text-xl font-bold text-white">{room.name}</h3>
                  <p className="text-gray-300">{room.description}</p>
                  <p className="text-gray-500">
                    Type: {room.room_type === "private" ? "Private" : "Public"}
                  </p>
                  {room.room_type === "private" && (
                    <input
                      type="text"
                      placeholder="Enter room code"
                      value={roomCode[room.ID] || ""}
                      onChange={(e) =>
                        setRoomCode({ ...roomCode, [room.ID]: e.target.value })
                      }
                      className="mt-2 mb-2 p-2 rounded border border-gray-600"
                    />
                  )}
                  <button
                    className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-2"
                    onClick={() => handleJoinRoom(room.ID, room.room_type)}
                  >
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
