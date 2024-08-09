"use client";

import SignInForm from "@/components/forms/sign-in";
import { redirect, useRouter } from "next/navigation";
import { useAuth } from "@/providers/auth_provider";
import Link from "next/link";

const Page = () => {
  const { user, loading, error } = useAuth();
  if (loading) {
    return <p>loading</p>;
  }
  if (user) {
    return redirect("/");
  }

  return (
    <div className="flex flex-col w-72">
      <h1 className="font-semibold text-3xl my-6">Sign in</h1>
      <SignInForm />
      <Link href={"/sign-up"} className="text-blue-500 text-sm my-5">
        Don{"'"}t have an account?
      </Link>
    </div>
  );
};

export default Page;
