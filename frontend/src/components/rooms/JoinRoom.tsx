import React, { useState, useEffect } from "react";
import apiClient from "../../api/apiClient";
import { Swiper, SwiperSlide } from "swiper/react";
import { Navigation, Pagination } from "swiper/modules"; // Import modules from swiper/modules
import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";
import "./styles.css";

function JoinRoom() {
  const [rooms, setRooms] = useState<any[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    async function fetchRooms() {
      try {
        const response = await apiClient.get("/rooms");
        console.log("Fetched rooms:", response.data);

        // Sort rooms by CreatedAt timestamp (newest first)
        const sortedRooms = response.data.sort((a: any, b: any) => {
          return (
            new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime()
          );
        });

        // Limit to the 7 most recent rooms
        const recentRooms = sortedRooms.slice(0, 7);

        setRooms(recentRooms);
        console.log("Recent rooms:", recentRooms);
      } catch (err: any) {
        console.error(err.response?.data || err.message);
        setError("Failed to fetch rooms. Please try again.");
      }
    }
    fetchRooms();
  }, []);

  return (
    <div className="w-full max-w-lg mx-auto">
      <div className="bg-gray-900 rounded-lg p-6 shadow-lg">
        <h2 className="text-3xl mb-4 text-center font-bold text-blue-400">
          Join a Room
        </h2>
        {error && <p className="text-red-500 text-center mb-4">{error}</p>}

        {rooms.length > 0 ? (
          <Swiper
            spaceBetween={30}
            slidesPerView={1}
            navigation={true} // Enable navigation
            pagination={{ clickable: true }} // Enable pagination
            modules={[Navigation, Pagination]} // Register Swiper modules
          >
            {rooms.map((room: any) => (
              <SwiperSlide key={room.id}>
                <div className="bg-gray-800 p-6 rounded-lg shadow-md">
                  <h3 className="text-2xl font-bold text-white">{room.name}</h3>
                  <p className="text-gray-300">{room.description}</p>
                  <p className="text-gray-400">
                    Type: {room.room_type === "private" ? "Private" : "Public"}
                  </p>
                  <p className="text-gray-500">
                    Created At: {new Date(room.CreatedAt).toLocaleString()}
                  </p>
                  <button className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-4">
                    Join Room
                  </button>
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
