"use client";

import { useAuth } from "@/providers/auth_provider";
import { redirect } from "next/navigation";
import React from "react";

const Page = () => {
  const { user, loading, error } = useAuth();
  if (loading) {
    return <p>loading</p>;
  }
  if (!user) {
    redirect("/sign-in");
  }
  return <div>Protected {user?.firstname}</div>;
};

export default Page;
