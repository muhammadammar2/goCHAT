import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../auth/AuthContext";
import apiClient from "../../api/apiClient";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { login } = useAuth();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const response = await apiClient.post("/login", { email, password });
      const { token } = response.data;

      if (token) {
        localStorage.setItem("token", token);
        login();
        navigate("/room-options");
      } else {
        throw new Error("Login failed. Token not received.");
      }
    } catch (err: any) {
      console.error(err.response?.data || err.message);
      setError(
        err.response?.data?.message ||
          "Login failed. Please check your credentials."
      );
    }
  };

  return (
    <div className="w-full max-w-md">
      <form
        onSubmit={handleSubmit}
        className="bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4"
      >
        <h2 className="text-3xl mb-6 text-center font-bold text-blue-400">
          Login
        </h2>
        {error && <p className="text-red-500 text-center mb-4">{error}</p>}
        <div className="mb-4">
          <input
            className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
            type="text"
            placeholder="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="mb-6">
          <input
            className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 mb-3 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <div className="flex items-center justify-center mb-4">
          <button
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300"
            type="submit"
          >
            Login
          </button>
        </div>
      </form>
    </div>
  );
}

export default Login;
