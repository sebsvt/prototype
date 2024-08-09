import type { Metadata } from "next";
import { Inter as FontSans } from "next/font/google";

import "./globals.css";

import { cn } from "@/lib/utils";
import { AuthProvider } from "@/providers/auth_provider";
import Navigation from "@/components/shared/navigation";
import { ThemeProvider } from "@/providers/theme_provider";

const fontSans = FontSans({
  subsets: ["latin"],
  variable: "--font-sans",
});

export const metadata: Metadata = {
  title: "Orchestra",
  description: "Connect your business with your customer",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={cn(
          "min-h-screen bg-background font-sans antialiased",
          fontSans.variable
        )}
      >
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <AuthProvider>
            <Navigation />
            {children}
          </AuthProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
