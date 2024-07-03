"use client";

import { Button } from "@/components/ui/button";
import React, { useState, useEffect } from "react";
import Cookies from "js-cookie";

interface Order {
  order_id: number;
  customer_id: number;
  product_sku: string;
  product_cost: number;
  duration: number;
  total_cost: number;
  is_paid: boolean;
  created_at: string; // Use string to avoid serialization issues with Date
}

const Page = ({ params }: { params: { ref: number } }) => {
  const [order, setOrder] = useState<Order | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchOrder = async () => {
      setLoading(true);
      const token = Cookies.get("access_token");

      if (!token) {
        setError("No access token found. Please sign in first.");
        return;
      }

      try {
        const response = await fetch("http://localhost:8000/api/order", {
          headers: {
            "Content-Type": "application/json",
            Authorization: `${token}`,
          },
        });

        if (!response.ok) {
          throw new Error("Failed to fetch orders");
        }
        const res = await fetch(
          `http://localhost:8000/api/order/${params.ref}`,
          {
            headers: {
              "Content-Type": "application/json",
              Authorization: `${token}`,
            },
          }
        );
        if (!res.ok) {
          throw new Error(await res.json());
        }
        const data: Order = await res.json();
        setOrder(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchOrder();
  }, [params.ref]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;
  if (!order) return <div>No order found</div>;

  return (
    <div className="container mx-auto w-full py-6">
      <h1 className="text-3xl md:text-4xl font-bold tracking-tight">
        Order Details
      </h1>
      <div className="mt-4">
        <p>Order ID: {order.order_id}</p>
        <p>Customer ID: {order.customer_id}</p>
        <p>Product SKU: {order.product_sku}</p>
        <p>Product Cost: ${order.product_cost.toFixed(2)}</p>
        <p>Duration: {order.duration} months</p>
        <p>Total Cost: ${order.total_cost.toFixed(2)}</p>
        <p>Status: {order.is_paid ? "Paid" : "Unpaid"}</p>
        <p>Created At: {new Date(order.created_at).toLocaleString()}</p>
      </div>
      <Button className="mt-4">Check Out</Button>
    </div>
  );
};

export default Page;
