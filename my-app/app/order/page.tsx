"use client";

import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";

interface Order {
  order_id: number;
  customer_id: number;
  product_sku: string;
  product_cost: number;
  duration: number;
  total_cost: number;
  is_paid: boolean;
  created_at: string;
}

const Page = () => {
  const [orders, setOrders] = useState<Order[] | null>(null);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    const fetchOrders = async () => {
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

        const data: Order[] = await response.json();
        setOrders(data);
      } catch (error: any) {
        setError(error.message);
      }
    };

    fetchOrders();
  }, []);

  return (
    <main className="container mx-auto mt-10">
      <h1 className="text-3xl font-semibold">Checkout</h1>
      {orders === null ? (
        <p>- No orders found.</p>
      ) : orders.length === 0 ? (
        <p>- No orders found.</p>
      ) : (
        <div className="mt-4">
          <table className="min-w-full bg-white border border-gray-300">
            <thead>
              <tr>
                <th className="py-2 px-4 border-b">Order ID</th>
                <th className="py-2 px-4 border-b">Product SKU</th>
                <th className="py-2 px-4 border-b">Product Cost</th>
                <th className="py-2 px-4 border-b">Duration</th>
                <th className="py-2 px-4 border-b">Total Cost</th>
                <th className="py-2 px-4 border-b">Is Paid</th>
                <th className="py-2 px-4 border-b">Created At</th>
              </tr>
            </thead>
            <tbody>
              {orders.map((order) => (
                <tr
                  className="hover:cursor-pointer"
                  key={order.order_id}
                  onClick={() => router.push(`/order/${order.order_id}`)}
                >
                  <td className="py-2 px-4 border-b">{order.order_id}</td>
                  <td className="py-2 px-4 border-b">{order.product_sku}</td>
                  <td className="py-2 px-4 border-b">
                    ${order.product_cost.toFixed(2)}
                  </td>
                  <td className="py-2 px-4 border-b">{order.duration}</td>
                  <td className="py-2 px-4 border-b">
                    ${order.total_cost.toFixed(2)}
                  </td>
                  <td className="py-2 px-4 border-b">
                    {order.is_paid ? "Yes" : "No"}
                  </td>
                  <td className="py-2 px-4 border-b">
                    {new Date(order.created_at).toLocaleString()}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </main>
  );
};

export default Page;
