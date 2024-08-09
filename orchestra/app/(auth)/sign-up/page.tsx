"use client";

import SignUpForm from "@/components/forms/sign-up";
import { useAuth } from "@/providers/auth_provider";
import Link from "next/link";
import { redirect } from "next/navigation";
import React from "react";

const Page = () => {
  const { user, loading, error } = useAuth();
  if (loading) {
    return <p>loading</p>;
  }
  if (user) {
    return redirect("/");
  }
  return (
    <div className="flex flex-col w-90">
      <h1 className="font-semibold text-3xl my-6">Sign up</h1>
      <SignUpForm />
      <Link href={"/sign-in"} className="text-blue-500 text-sm my-5">
        Already have an account
      </Link>
    </div>
  );
};

export default Page;
