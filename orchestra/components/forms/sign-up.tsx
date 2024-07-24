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
} from "../ui/form";
import { useForm } from "react-hook-form";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { useToast } from "../ui/use-toast";
import { redirect, useRouter } from "next/navigation";
import Cookies from "js-cookie";
import { useAuth } from "@/providers/auth_provider";

const SignUpFormSchema = z.object({
  firstname: z.string().min(1, { message: "First name is required" }),
  surname: z.string().min(1, { message: "Surname is required" }),
  gender: z.enum(["Male", "Female", "Other"], {
    message: "Gender is required",
  }),
  phone: z
    .string()
    .min(10, { message: "Phone number must be at least 10 digits" }),
  date_of_birth: z.string().refine((val) => !isNaN(Date.parse(val)), {
    message: "Invalid date format",
  }),
  email: z.string().email({ message: "Invalid email format" }),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters" }),
});

const SignUpForm = () => {
  // const router = useRouter();
  const form = useForm<z.infer<typeof SignUpFormSchema>>({
    resolver: zodResolver(SignUpFormSchema),
  });

  const { toast } = useToast();

  const { user, loading } = useAuth();
  if (loading) {
    return <></>;
  }
  if (user) {
    return redirect("/");
  }
  const onSubmit = async (values: z.infer<typeof SignUpFormSchema>) => {
    try {
      const response = await fetch("http://localhost:8000/api/auth/sign-up", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(values),
      });
      const data = await response.json();
      if (!response.ok) {
        throw new Error(data.error); // Throw the error message received from the server
      }
      toast({
        title: "Sign up successfully",
      });
      window.location.reload();
      Cookies.set("access_token", data.access_token, { expires: 15 }); // Expires in 1 day
      // console.log("Success:", data);
    } catch (error: any) {
      // Display error message using toast
      toast({
        title: "Sign up failed",
        description: error.message, // Use error.message to display the error
      });
      console.error("Error:", error);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <div className="flex space-x-4">
          <FormField
            control={form.control}
            name="firstname"
            render={({ field }) => (
              <FormItem className="w-1/2">
                <FormLabel>First Name</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="First Name"
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
            name="surname"
            render={({ field }) => (
              <FormItem className="w-1/2">
                <FormLabel>Surname</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="Surname"
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
            name="phone"
            render={({ field }) => (
              <FormItem className="w-full">
                <FormLabel>Phone</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="Phone"
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
            name="gender"
            render={({ field }) => (
              <FormItem className="w-1/3">
                <FormLabel>Gender</FormLabel>
                <Select
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder="gender" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectItem value="Male">Male</SelectItem>
                    <SelectItem value="Female">Female</SelectItem>
                    <SelectItem value="Other">Other</SelectItem>
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className="flex space-x-4">
          <FormField
            control={form.control}
            name="date_of_birth"
            render={({ field }) => (
              <FormItem className="w-1/3">
                <FormLabel>Date of Birth</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="YYYY-MM-DD"
                    {...field}
                    type="date"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem className="w-full">
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input
                    className="w-full"
                    placeholder="string@example.com"
                    {...field}
                    type="email"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input
                  placeholder="your password"
                  {...field}
                  type="password"
                  className="w-full"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="w-full">
          Sign Up
        </Button>
      </form>
    </Form>
  );
};

export default SignUpForm;
