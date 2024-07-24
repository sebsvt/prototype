"use client";

import React from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import Cookies from "js-cookie";
import { useForm } from "react-hook-form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { useRouter } from "next/navigation";
import { Textarea } from "@/components/ui/textarea";
import { useAuth } from "@/providers/auth_provider";

const StudioFormSchema = z.object({
  subdomain: z.string().min(1, { message: "Subdomain is required" }),
  name: z.string().min(1, { message: "Name is required" }),
  description: z
    .string()
    .min(10, {
      message: "Address must be at least 10 characters.",
    })
    .max(160, {
      message: "Address must not be longer than 30 characters.",
    }),
  address: z.string().min(1, { message: "Address is required" }),
  city: z.string().min(1, { message: "City is required" }),
  zipcode: z.string().min(1, { message: "Zipcode is required" }),
  state: z.string().min(1, { message: "State is required" }),
  country: z.string().min(1, { message: "Country is required" }),
});

const OpenStudioForm = () => {
  const form = useForm<z.infer<typeof StudioFormSchema>>({
    resolver: zodResolver(StudioFormSchema),
  });

  const { toast } = useToast();
  const { user, loading, error } = useAuth();
  const router = useRouter();

  if (loading) {
    return <div>loading</div>;
  }

  if (!user) {
    return router.push("/sign-in");
  }

  const onSubmit = async (values: z.infer<typeof StudioFormSchema>) => {
    const access_token = Cookies.get("access_token");
    if (!access_token) {
      return router.push("/sign-in");
    }
    try {
      const response = await fetch("http://localhost:8000/api/studio/opening", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${access_token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify(values),
      });
      const data = await response.json();
      if (!response.ok) {
        throw new Error(data.error);
      }
      toast({
        title: "Studio created successfully",
      });
      // router.push("/success");
    } catch (error: any) {
      toast({
        title: "Creation failed",
        description: error.message,
      });
      console.error("Error:", error);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <div className="flex flex-col space-y-8 md:flex-row md:space-x-4 md:space-y-0">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem className="w-full md:w-1/2">
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="Name"
                    {...field}
                    type="text"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="subdomain"
            render={({ field }) => (
              <FormItem className="w-full md:w-1/2">
                <FormLabel>Subdomain</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="Subdomain"
                    {...field}
                    type="text"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className="">
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem className="w-full">
                <FormLabel>Description</FormLabel>
                <FormControl>
                  <Textarea
                    className="w-full"
                    placeholder="Description"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div>
          <FormField
            control={form.control}
            name="address"
            render={({ field }) => (
              <FormItem className="w-full">
                <FormLabel>Address</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="Address"
                    type="text"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className="flex space-x-4">
          <FormField
            control={form.control}
            name="city"
            render={({ field }) => (
              <FormItem className="w-1/2">
                <FormLabel>City</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="City"
                    {...field}
                    type="text"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="zipcode"
            render={({ field }) => (
              <FormItem className="w-1/2">
                <FormLabel>Zipcode</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="Zipcode"
                    {...field}
                    type="text"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className="flex space-x-4">
          <FormField
            control={form.control}
            name="state"
            render={({ field }) => (
              <FormItem className="w-1/2">
                <FormLabel>State</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="State"
                    {...field}
                    type="text"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="country"
            render={({ field }) => (
              <FormItem className="w-1/2">
                <FormLabel>Country</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="Country"
                    {...field}
                    type="text"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <Button type="submit" className="w-full">
          Create Studio
        </Button>
      </form>
    </Form>
  );
};

export default OpenStudioForm;
