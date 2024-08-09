"use client";

import React from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useAuth } from "@/providers/auth_provider";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import Cookies from "js-cookie";
import { redirect, useRouter } from "next/navigation";
import { ComponentIcon, Settings2Icon, UserCircle } from "lucide-react";

const UserDropDown = () => {
  const router = useRouter();
  const { user, setUser } = useAuth();
  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <Avatar className="h-7 w-7">
          <AvatarImage
            className="object-cover"
            src="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/bd8acd46-3677-468e-a68b-bec2f9f9170d/dd2qa1j-4b05e8a0-e454-4ab2-8bda-2d9fa11ea0a8.jpg/v1/fill/w_1095,h_730,q_70,strp/chapter_142__ice_queen_kaguya_by_saiplayssupport_dd2qa1j-pre.jpg?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7ImhlaWdodCI6Ijw9MjMzNyIsInBhdGgiOiJcL2ZcL2JkOGFjZDQ2LTM2NzctNDY4ZS1hNjhiLWJlYzJmOWY5MTcwZFwvZGQycWExai00YjA1ZThhMC1lNDU0LTRhYjItOGJkYS0yZDlmYTExZWEwYTguanBnIiwid2lkdGgiOiI8PTM1MDgifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6aW1hZ2Uub3BlcmF0aW9ucyJdfQ.UZe9g_n4HNZ9rivZu1XSJCNemgTKGbxAVX2oao7tat0"
            alt="@shadcn"
          />
          <AvatarFallback>{user?.firstname}</AvatarFallback>
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-48 md:w-52 mx-2">
        <DropdownMenuLabel>
          {user?.firstname} {user?.surname}
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem className="flex gap-2 items-center">
          <Settings2Icon size={"18"} />
          Settings
        </DropdownMenuItem>
        <DropdownMenuItem
          className="flex gap-2 items-center"
          onClick={() => router.push("/studio")}
        >
          <ComponentIcon size={"18"} />
          My studio
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          className="text-red-600 hover:cursor-pointer hover:!text-red-500"
          onClick={() => {
            Cookies.remove("access_token");
            setUser(null);
            redirect("/");
          }}
        >
          Sign out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default UserDropDown;
