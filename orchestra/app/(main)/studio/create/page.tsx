import OpenStudioForm from "@/components/forms/studio/open-studio";
import React from "react";

const Page = () => {
  return (
    <main className="flex items-center justify-center min-h-screen">
      <div className="container mx-auto">
        <h3 className="mt-8 scroll-m-20 text-3xl font-semibold tracking-tight">
          Studio Information
        </h3>
        <section className="mt-6 mb-10">
          <OpenStudioForm />
        </section>
      </div>
    </main>
  );
};

export default Page;
