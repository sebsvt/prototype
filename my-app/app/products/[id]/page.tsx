import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Image from "next/image";
import React from "react";

interface Product {
  product_id: number;
  sku: string;
  name: string;
  description: string;
  price: number;
  is_available: boolean;
}

async function getProduct(product_id: number) {
  const res = await fetch(`http://localhost:8000/api/products/${product_id}`);
  if (!res.ok) {
    throw new Error("Failed to fetch");
  }
  const product: Product = await res.json();
  return product;
}

const Page = async ({ params }: { params: { id: number } }) => {
  const product = await getProduct(params.id);
  return (
    <main className="container mx-auto w-full py-6">
      <section className="mt-6">
        <h1 className="scroll-m-20 text-3xl md:text-4xl font-bold tracking-tight">
          {product.name}
        </h1>
        <p className="text-red-400">{product.sku}</p>
      </section>
      <section className="mt-8">
        <div className="flex flex-col md:flex-row gap-6">
          <div className="w-full md:w-1/2">
            <img
              src="https://images.autofun.co.th/file1/48c6adaea6a4401991b2ab646821cf86_1125x630.jpg"
              alt="product image"
              className="w-full h-auto rounded-md"
            />
          </div>
          <div className="w-full md:w-1/2 flex flex-col justify-between">
            <div>
              <section className="header">
                <h2 className="text-2xl font-bold mb-4">{product.name}</h2>
              </section>
              <section className="body">
                <p className="mb-4 text-muted-foreground ">
                  {product.description} Lorem ipsum dolor, sit amet consectetur
                  adipisicing elit. Aliquam ut error excepturi natus velit autem
                  molestiae minima laboriosam impedit dolorum sunt sequi
                  distinctio, odio voluptatibus praesentium quo ipsa, molestias
                  accusamus!
                </p>
                <div>
                  <h3 className="font-semibold mb-2">Duration time</h3>
                  <Input
                    type="number"
                    placeholder="duration times"
                    defaultValue={1}
                  />
                  <p className="text-xl font-semibold mt-4">
                    ${product.price.toFixed(2)}
                  </p>
                  <span className="text-sm text-muted-foreground">
                    Monthly payment
                  </span>
                </div>
              </section>
            </div>
            <section className="footer mt-4">
              <Button className="w-full">Buy</Button>
            </section>
          </div>
        </div>
      </section>
    </main>
  );
};

export default Page;
