import React, { useState } from "react";
import apiClient from "../../api/apiClient";
import { useNavigate } from "react-router-dom";

function Register() {
  const [name, setName] = useState("");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const navigate = useNavigate();

  const register = async (
    name: string,
    username: string,
    email: string,
    password: string
  ) => {
    const data = { name, username, email, password };
    try {
      const response = await apiClient.post("/signup", data);

      console.log("Registration response:", response);
      return response;
    } catch (error) {
      console.error("Registration error:", error);
      throw error;
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    try {
      const response = await register(name, username, email, password);
      if (response.status === 201) {
        setSuccess("Registration successful!");
        setTimeout(() => {
          navigate("/login");
        }, 1450);
        setName("");
        setUsername("");
        setEmail("");
        setPassword("");
      }
    } catch (err: any) {
      setError("Registration failed. Please try again.");
      console.error("Registration error:", err);
    }
  };

  return (
    <div className="w-full max-w-md">
      <form
        onSubmit={handleSubmit}
        className="bg-gray-800 shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4"
      >
        <h2 className="text-3xl mb-6 text-center font-bold text-blue-400">
          Register
        </h2>

        {error && <p className="text-red-500 text-center mb-4">{error}</p>}
        {success && (
          <p className="text-green-500 text-center mb-4">{success}</p>
        )}

        <div className="mb-4">
          <input
            className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
            type="text"
            placeholder="Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="mb-4">
          <input
            className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div className="mb-4">
          <input
            className="shadow appearance-none border border-gray-700 rounded w-full py-3 px-4 text-gray-300 leading-tight focus:outline-none focus:border-blue-500 bg-gray-700 transition duration-300"
            type="email"
            placeholder="Email"
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
        <div className="flex items-center justify-center">
          <button
            className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-6 rounded-full focus:outline-none focus:shadow-outline transition duration-300"
            type="submit"
          >
            Register
          </button>
        </div>
      </form>
    </div>
  );
}

export default Register;
