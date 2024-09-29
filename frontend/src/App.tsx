import React from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  useNavigate,
  useLocation,
} from "react-router-dom";
import Login from "./components/forms/Login";
import Register from "./components/forms/Register";
import RoomOptions from "./components/rooms/RoomOptions";
import CreateRoomForm from "./components/rooms/createRoomForm";
import JoinRoom from "./components/rooms/JoinRoom";
import ProtectedRoute from "./components/ProtectedRoute";
import { AuthProvider, useAuth } from "./auth/AuthContext";
import ProfilePage from "./components/ProfilePage";
import Chat from "./components/Messages/Chat";

function App() {
  return (
    <AuthProvider>
      <Router>
        <AppContent />
      </Router>
    </AuthProvider>
  );
}

function AppContent() {
  const location = useLocation();
  const showLeftContent = location.pathname !== "/profile";

  return (
    <div className="h-screen bg-gray-900 text-white flex flex-col lg:flex-row overflow-hidden">
      {showLeftContent && (
        <div className="w-full lg:w-1/2 flex flex-col items-center justify-center p-8 bg-gradient-to-br from-gray-800 to-gray-700">
          <div className="text-center">
            <h1 className="text-6xl lg:text-7xl font-extrabold tracking-tight text-gray-500 mb-6 p-2 relative">
              <span className="glowing-text" data-text="go">
                go
              </span>
              <span className="glowing-text" data-text="CHAT">
                CHAT
              </span>
            </h1>
            <p className="text-xl font-semibold leading-relaxed text-gray-300 mt-6 max-w-lg mx-auto">
              You're free to sell drugs and smuggle weapons here
            </p>
          </div>
        </div>
      )}

      <div
        className={`w-full ${
          showLeftContent ? "lg:w-1/2" : ""
        } flex flex-col items-center justify-center p-8 bg-gray-900 relative`}
      >
        <div className="absolute top-4 right-4 flex space-x-4">
          <ProfileButton />
          <LogoutButton />
        </div>
        <div className="w-full max-w-md">
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/" element={<HomePage />} />

            <Route element={<ProtectedRoute />}>
              <Route path="/room-options" element={<RoomOptions />} />
              <Route path="/create-room" element={<CreateRoomForm />} />
              <Route path="/join-room" element={<JoinRoom />} />
              <Route path="/profile" element={<ProfilePage />} />
              <Route path="/chat" element={<Chat />} /> {/* Added chat route */}
            </Route>
          </Routes>
        </div>
      </div>
    </div>
  );
}

function LogoutButton() {
  const { isAuthenticated, logout } = useAuth();
  const navigate = useNavigate();

  if (!isAuthenticated) return null;

  return (
    <button
      onClick={async () => {
        await logout();
        navigate("/login");
      }}
      className="bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-full focus:outline-none focus:shadow-outline transition duration-300"
    >
      Logout
    </button>
  );
}

function ProfileButton() {
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();

  if (!isAuthenticated) return null;

  return (
    <button
      onClick={() => navigate("/profile")}
      className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full focus:outline-none focus:shadow-outline transition duration-300"
    >
      Profile
    </button>
  );
}

function HomePage() {
  const navigate = useNavigate();

  return (
    <div className="flex flex-col items-center space-y-4">
      <h2 className="text-3xl mb-6 text-center font-bold text-white">Home</h2>
      <button
        onClick={() => navigate("/login")}
        className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300 w-full"
      >
        Login
      </button>
      <button
        onClick={() => navigate("/register")}
        className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300 w-full"
      >
        Register
      </button>
    </div>
  );
}

export default App;
