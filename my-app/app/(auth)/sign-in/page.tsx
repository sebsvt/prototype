"use client";

import React, { useState } from "react";
import Cookies from "js-cookie";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

const Page = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const router = useRouter();

  const handleSignIn = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8000/api/users/sign-in", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: email,
          password: password,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to sign in");
      }

      const data = await response.json();
      Cookies.set("access_token", data.access_token, { expires: 7 }); // Store the token for 7 days
      setMessage(`Signed in successfully!`);
      router.push("/");
    } catch (error) {
      setMessage("Error: " + error.message);
    }
  };

  return (
    <main className="container mx-auto flex justify-center items-center h-screen">
      <section className="max-w-xl w-full px-6 py-10">
        <h1 className="text-4xl font-semibold mb-6">Sign In</h1>
        <form onSubmit={handleSignIn} className="flex flex-col gap-4">
          <label htmlFor="email" className="flex flex-col">
            Email:
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="p-2 border border-gray-300 rounded"
            />
          </label>
          <label htmlFor="password" className="flex flex-col">
            Password:
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              className="p-2 border border-gray-300 rounded"
            />
          </label>
          <Button type="submit">Sign In</Button>
        </form>
        {<p className="mt-4 text-center text-red-500">{message}</p> && (
          <p className="mt-4 text-center text-green-500">{message}</p>
        )}
      </section>
    </main>
  );
};

export default Page;
