import { Button } from "@/components/ui/button";
import { getStudioByDomain } from "@/lib/queries";
import Link from "next/link";
import { notFound } from "next/navigation";
import React from "react";

type PageProps = {
  params: {
    studio_domain: string;
  };
};

const Page = async ({ params }: PageProps) => {
  try {
    const studioData = await getStudioByDomain(params.studio_domain);

    return (
      <div>
        <nav className="flex items-center justify-center gap-8 md:px-4">
          <Link
            href={"/"}
            className="font-medium text-xs hover:underline underline-offset-4"
          >
            Overview
          </Link>
          <Link
            href={"/"}
            className="font-medium text-xs hover:underline underline-offset-4"
          >
            iDental
          </Link>
          <Link
            href={"/"}
            className="font-medium text-xs hover:underline underline-offset-4"
          >
            iStore
          </Link>
        </nav>
        <main className="container mx-auto">
          <section className="mt-5 px-4">
            <div className="space-y-1 pb-4">
              <h2 className="text-xl md:text-2xl font-semibold">
                {studioData.name}
              </h2>
              <p className="text-muted-foreground text-xs">
                Software · Aerospace engineer
              </p>
              <p className="text-muted-foreground text-xs">
                {studioData.address} {studioData.state} {studioData.city}{" "}
                {studioData.zipcode} {studioData.country}
              </p>
            </div>
            <div>
              <p className="text-sm">{studioData.description} 😹</p>
            </div>
          </section>
          <section className="mt-4 py-4"></section>
        </main>
      </div>
    );
  } catch (error) {
    return notFound();
  }
};

export default Page;
