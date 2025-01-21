// src/app/layout.tsx
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import TelegramProvider from "@/components/providers/TelegramProvider";

const inter = Inter({ subsets: ["latin", "cyrillic"] });

export const metadata: Metadata = {
  title: "Flippy Toys Kids",
  description: "Детский режим Flippy Toys",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ru">
      <body className={inter.className}>
        <TelegramProvider>{children}</TelegramProvider>
      </body>
    </html>
  );
}
