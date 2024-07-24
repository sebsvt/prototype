"use client";

import SkeletonCard from "@/components/shared/skeleton-card";
import { getUserStudios } from "@/lib/queries";
import { useAuth } from "@/providers/auth_provider";
import Link from "next/link";
import { redirect } from "next/navigation";
import React, { useEffect, useState } from "react";

const Page = () => {
  const [studios, setStudios] = useState<StudioResponse[] | null>(null);

  useEffect(() => {
    const fetching_user_studios = async () => {
      try {
        const rest = await getUserStudios();
        setStudios(rest);
      } catch (error) {
        console.log(studios);
      }
    };
    fetching_user_studios();
  }, []);

  const { user, loading } = useAuth();
  if (loading) {
    return (
      <main className="container mx-auto flex items-center justify-center min-h-screen">
        <SkeletonCard />;
      </main>
    );
  }
  if (!user) {
    return redirect("/sign-in");
  }
  return (
    <main className="container mx-auto">
      <section className="">
        <h3 className="mt-8 scroll-m-20 text-3xl font-semibold tracking-tight">
          My Studio
        </h3>
      </section>
      <section>
        {studios ? (
          studios.map((studio, key) => (
            <div key={key}>
              <Link href={`${studio.subdomain}`}>{studio.name}</Link>
            </div>
          ))
        ) : (
          <p className="leading-7 [&:not(:first-child)]:mt-6">
            None of Studio.
            <Link
              href={"/studio/create"}
              className="text-blue-500 hover:text-blue-600"
            >
              Feel free to create a new one
            </Link>
          </p>
        )}
      </section>
    </main>
  );
};

export default Page;
