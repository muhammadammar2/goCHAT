// import React, {
//   createContext,
//   useContext,
//   useState,
//   ReactNode,
//   useEffect,
// } from "react";
// import apiClient from "../api/apiClient";

// interface AuthContextType {
//   isAuthenticated: boolean;
//   login: () => void;
//   logout: () => Promise<void>;
// }

// const AuthContext = createContext<AuthContextType | undefined>(undefined);

// export const AuthProvider: React.FC<{ children: ReactNode }> = ({
//   children,
// }) => {
//   const [isAuthenticated, setIsAuthenticated] = useState<boolean>(
//     !!localStorage.getItem("token")
//   );

//   useEffect(() => {
//     const validateToken = async () => {
//       const token = localStorage.getItem("token");
//       if (token) {
//         try {
//           await apiClient.post("/validate-token", { token });
//         } catch (error) {
//           console.error("Token validation failed:", error);
//           logout();
//         }
//       }
//     };

//     validateToken();
//   }, []);

//   const login = async () => {
//     setIsAuthenticated(true);
//   };

//   const logout = async () => {
//     try {
//       await apiClient.post("/logout");
//       localStorage.removeItem("token");
//       setIsAuthenticated(false);
//     } catch (error) {
//       console.error("Failed to log out:", error);
//     }
//   };

//   return (
//     <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
//       {children}
//     </AuthContext.Provider>
//   );
// };

// export const useAuth = () => {
//   const context = useContext(AuthContext);
//   if (!context) {
//     throw new Error("useAuth must be used within an AuthProvider");
//   }
//   return context;
// };

import React, { createContext, useContext, useState, ReactNode } from "react";
import apiClient from "../api/apiClient";

interface AuthContextType {
  isAuthenticated: boolean;
  login: () => Promise<void>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(
    !!localStorage.getItem("token")
  );

  const login = async () => {
    setIsAuthenticated(true);
  };

  const logout = async () => {
    try {
      await apiClient.post("/logout");
      localStorage.removeItem("token");
      setIsAuthenticated(false);
    } catch (error) {
      console.error("Failed to log out:", error);
    }
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
