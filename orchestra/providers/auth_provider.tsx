// context/authContext.tsx

"use client";

import React, { createContext, useContext, useState, useEffect } from "react";
import axios from "axios";

interface User {
  user_ref: number;
  firstname: string;
  surname: string;
  gender: string;
  phone: string;
  date_of_birth: string;
}

interface AuthContextType {
  user: User | null;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
  loading: boolean;
  error: string | null;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchUserInfo = async (userId: string) => {
    try {
      const response = await axios.get(
        `http://localhost:8000/api/profile/${userId}`
      );
      setUser(response.data);
    } catch (err) {
      setError("Failed to fetch user information");
    }
  };

  const verifyToken = async (token: string) => {
    try {
      const response = await axios.get(
        "http://localhost:8000/api/auth/verify_token",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      const { iss } = response.data; // Assuming 'iss' is the user ID
      await fetchUserInfo(iss); // Fetch additional user information
      setError(null);
    } catch (err) {
      setUser(null);
      setError("Invalid credential");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const token = document.cookie
      .split("; ")
      .find((row) => row.startsWith("access_token="))
      ?.split("=")[1];
    if (token) {
      verifyToken(token);
    } else {
      setLoading(false);
    }
  }, []);

  return (
    <AuthContext.Provider value={{ user, setUser, loading, error }}>
      {children}
    </AuthContext.Provider>
  );
};
