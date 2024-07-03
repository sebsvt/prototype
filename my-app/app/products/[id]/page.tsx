"use client";

import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie";

interface Product {
  product_id: number;
  sku: string;
  name: string;
  description: string;
  price: number;
  is_available: boolean;
}

const Page = ({ params }: { params: { id: number } }) => {
  const router = useRouter();
  const [product, setProduct] = useState<Product | null>(null);
  const [duration, setDuration] = useState<number>(1);
  const [totalPrice, setTotalPrice] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchProduct = async () => {
      setLoading(true);
      try {
        const res = await fetch(
          `http://localhost:8000/api/products/${params.id}`
        );
        if (!res.ok) {
          throw new Error("Failed to fetch");
        }
        const product: Product = await res.json();
        setProduct(product);
        setTotalPrice(product.price);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    fetchProduct();
  }, [params.id]);

  const handleOrderSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const token = Cookies.get("access_token");
    if (!token) {
      console.error("No access token found");
      return;
    }

    const orderData = {
      customer_id: 1, // Mocked customer ID
      product_sku: product?.sku,
      product_cost: product?.price,
      duration: duration,
    };
    try {
      const res = await fetch("http://localhost:8000/api/order/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${token}`,
        },
        body: JSON.stringify(orderData),
      });
      if (!res.ok) {
        throw new Error("Failed to create order");
      }
      const order = await res.json();
      router.push(`/order/${order.order_id}`);
    } catch (error) {
      console.error(error);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!product) {
    return <div>Product not found</div>;
  }

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
                <p className="mb-4 text-muted-foreground">
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
                    onChange={(e) => {
                      setDuration(Number(e.target.value));
                      setTotalPrice(Number(e.target.value) * product.price);
                    }}
                  />
                  <p className="text-xl font-semibold mt-4">
                    ${totalPrice.toFixed(2)}
                  </p>
                  <span className="text-sm text-muted-foreground">
                    Monthly payment
                  </span>
                </div>
              </section>
            </div>
            <section className="footer mt-4">
              <form onSubmit={handleOrderSubmit}>
                <Button type="submit" className="w-full">
                  Buy
                </Button>
              </form>
            </section>
          </div>
        </div>
      </section>
    </main>
  );
};

export default Page;
