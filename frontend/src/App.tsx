import React from "react";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  useNavigate,
} from "react-router-dom";
import Login from "./components/forms/Login";
import Register from "./components/forms/Register";
import RoomOptions from "./components/rooms/RoomOptions";
import CreateRoomForm from "./components/rooms/createRoomForm";
import JoinRoom from "./components/rooms/JoinRoom";

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-900 text-white flex flex-col lg:flex-row">
        <div className="w-full lg:w-1/2 flex flex-col items-center justify-center p-8 bg-gradient-to-br from-gray-800 to-gray-700">
          <div className="text-center">
            <h1 className="text-6xl lg:text-7xl font-extrabold tracking-tight text-white mb-6 p-2">
              go<span className="align-baseline">CHAT</span>
            </h1>
            <p className="text-xl font-semibold leading-relaxed text-gray-300 mt-6 max-w-lg mx-auto">
              You're free to sell drugs and smuggle weapons here
            </p>
          </div>
        </div>

        <div className="w-full lg:w-1/2 flex flex-col items-center justify-center p-8 bg-gray-900">
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/room-options" element={<RoomOptions />} />
            <Route path="/create-room" element={<CreateRoomForm />} />
            <Route path="/join-room" element={<JoinRoom />} />
            <Route path="/" element={<HomePage />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

function HomePage() {
  const navigate = useNavigate();

  return (
    <div className="w-full max-w-md">
      <div className="bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4">
        <h2 className="text-3xl mb-6 text-center font-bold text-white">Home</h2>
        <div className="flex flex-col items-center">
          <button
            onClick={() => navigate("/login")}
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300 mb-4"
          >
            Login
          </button>
          <button
            onClick={() => navigate("/register")}
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300"
          >
            Register
          </button>
        </div>
      </div>
    </div>
  );
}

export default App;
