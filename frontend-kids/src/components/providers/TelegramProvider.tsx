// src/components/providers/TelegramProvider.tsx
"use client";

import { useEffect } from "react";
import Script from "next/script";

export default function TelegramProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  useEffect(() => {
    // Инициализируем приложение когда Telegram Web App готов
    if (window.Telegram?.WebApp) {
      window.Telegram.WebApp.ready();
      window.Telegram.WebApp.expand();
    }
  }, []);

  return (
    <>
      <Script src="https://telegram.org/js/telegram-web-app.js" />
      {children}
    </>
  );
}
