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
import Cookies from "js-cookie";
import { useToast } from "../ui/use-toast";
import { useRouter } from "next/navigation";
import { useAuth } from "@/providers/auth_provider";

const SignInFormScheme = z.object({
  email: z.string().email({ message: "invalid email format" }),
  password: z.string().min(8, { message: "password length is less than 8" }),
});

const SignInForm = () => {
  const router = useRouter();
  const form = useForm<z.infer<typeof SignInFormScheme>>({
    resolver: zodResolver(SignInFormScheme),
  });

  const { toast } = useToast();
  const onSubmit = async (values: z.infer<typeof SignInFormScheme>) => {
    try {
      const response = await fetch("http://localhost:8000/api/auth/sign-in", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(values),
      });
      const data = await response.json();
      if (!response.ok) {
        throw new Error(data.error);
      }
      // console.log("Success:", data);
      Cookies.set("access_token", data.access_token, { expires: 15 }); // Expires in 1 day
      window.location.reload();
      router.push("/");
    } catch (error: any) {
      // Display error message using toast
      toast({
        variant: "destructive",
        title: "Incorrect email or password.",
        description: error.message, // Use error.message to display the error
      });
      console.error("Error:", error);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
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
          Sign in
        </Button>
      </form>
    </Form>
  );
};

export default SignInForm;
