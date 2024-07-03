"use client";

import React, { useEffect, useState } from "react";
import validateToken from "@/utils/tools";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

const UserButton = () => {
  const [isValid, setIsValid] = useState(false);
  const router = useRouter();

  useEffect(() => {
    const checkToken = async () => {
      const valid = await validateToken();
      setIsValid(valid);
    };

    checkToken();
  }, []);

  if (!isValid) {
    return (
      <Button
        onClick={() => {
          router.push("/sign-in");
        }}
      >
        Sign in
      </Button>
    );
  }

  return <Button onClick={() => router.push("/order")}>Your order</Button>;
};

export default UserButton;
