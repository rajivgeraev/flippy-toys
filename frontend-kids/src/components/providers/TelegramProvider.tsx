// src/components/providers/TelegramProvider.tsx
"use client";

import { createContext, useContext, useEffect, useState } from "react";
import Script from "next/script";

const TelegramAuthContext = createContext<string>("");
export const useTelegramAuth = () => useContext(TelegramAuthContext);

export default function TelegramProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [initData, setInitData] = useState("");

  useEffect(() => {
    if (window.Telegram?.WebApp) {
      window.Telegram.WebApp.ready();
      window.Telegram.WebApp.expand();
      setInitData(window.Telegram.WebApp.initData);
    }
  }, []);

  return (
    <TelegramAuthContext.Provider value={initData}>
      <Script src="https://telegram.org/js/telegram-web-app.js" />
      {children}
    </TelegramAuthContext.Provider>
  );
}
