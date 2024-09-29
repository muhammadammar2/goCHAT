import React, { useState, useEffect } from "react";
import apiClient from "../../api/apiClient";
import { useNavigate } from "react-router-dom";
import { Swiper, SwiperSlide } from "swiper/react";
import { Navigation, Pagination } from "swiper/modules";
import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";
import "./styles.css";

function JoinRoom() {
  const [rooms, setRooms] = useState<any[]>([]);
  const [error, setError] = useState("");
  const [code, setCode] = useState("");
  const [selectedRoom, setSelectedRoom] = useState<any>(null);
  const navigate = useNavigate();

  useEffect(() => {
    async function fetchRooms() {
      try {
        const response = await apiClient.get("/rooms");

        const sortedRooms = response.data.sort((a: any, b: any) => {
          return (
            new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime()
          );
        });

        const recentRooms = sortedRooms.slice(0, 7);
        setRooms(recentRooms);
      } catch (err: any) {
        console.error(err.response?.data || err.message);
        setError("Failed to fetch rooms. Please try again.");
      }
    }
    fetchRooms();
  }, []);

  const handleJoinRoom = (room: any) => {
    setSelectedRoom(room);
    if (room.room_type === "public") {
      navigate("/profile");
    } else {
      setCode("");
    }
  };

  const handleCodeSubmit = async () => {
    try {
      if (!selectedRoom) {
        setError("No room selected.");
        return;
      }

      if (selectedRoom.room_type === "private" && !code) {
        setError("Code is required for private room.");
        return;
      }

      const response = await apiClient.post("/join-room", {
        room_id: selectedRoom.ID,
        code: code,
      });

      if (
        response.status === 200 &&
        response.data.message === "Successfully joined the room"
      ) {
        navigate("/profile");
      } else {
        setError("Unable to join the room. Please try again.");
      }
    } catch (error: any) {
      console.error("Error from API:", error);
      if (error.response && error.response.status === 403) {
        setError("Wrong code. Please try again.");
      } else {
        setError("An unexpected error occurred. Please try again.");
      }
    }
  };

  return (
    <div className="w-full max-w-xl mx-auto">
      <div className="bg-gray-900 rounded-lg p-8 shadow-xl">
        <h2 className="text-3xl mb-6 text-center font-bold text-blue-400">
          Join a Room
        </h2>
        {error && <p className="text-red-500 text-center mb-4">{error}</p>}
        {rooms.length > 0 ? (
          <Swiper
            spaceBetween={40}
            slidesPerView={1}
            navigation={true}
            pagination={{ clickable: true }}
            modules={[Navigation, Pagination]}
          >
            {rooms.map((room: any) => (
              <SwiperSlide key={room.id}>
                <div className="bg-gray-800 p-8 rounded-lg shadow-lg">
                  <h3 className="text-2xl font-bold text-white">{room.name}</h3>
                  <p className="text-gray-300 mt-2">{room.description}</p>
                  <p className="text-gray-400 mt-1">
                    Type: {room.room_type === "private" ? "Private" : "Public"}
                  </p>
                  <p className="text-gray-500 mt-1">
                    Created At: {new Date(room.CreatedAt).toLocaleString()}
                  </p>
                  <button
                    className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded mt-6"
                    onClick={() => handleJoinRoom(room)}
                  >
                    Select Room
                  </button>
                  {selectedRoom?.id === room.id &&
                    room.room_type === "private" && (
                      <div className="mt-4">
                        <input
                          type="text"
                          placeholder="Enter room code"
                          value={code}
                          onChange={(e) => setCode(e.target.value)}
                          className="border border-gray-400 p-2 rounded w-full bg-gray-800 text-white"
                        />
                        <button
                          onClick={handleCodeSubmit}
                          className="bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mt-2"
                        >
                          Submit Code
                        </button>
                      </div>
                    )}
                </div>
              </SwiperSlide>
            ))}
          </Swiper>
        ) : (
          <p className="text-gray-300 text-center">No rooms available.</p>
        )}
      </div>
    </div>
  );
}

export default JoinRoom;
