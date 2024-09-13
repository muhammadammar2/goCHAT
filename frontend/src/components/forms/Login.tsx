import React, { useState } from "react";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Handle login logic here
    console.log("Login submitted:", { username, password });
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
