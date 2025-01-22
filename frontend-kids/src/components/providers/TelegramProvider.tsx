// src/components/providers/TelegramProvider.tsx
"use client";

import { createContext, useContext, useEffect, useState, useMemo } from "react";
import Script from "next/script";

const TelegramAuthContext = createContext<{
  initData: string;
  isReady: boolean;
}>({ initData: "", isReady: false });

export const useTelegramAuth = () => useContext(TelegramAuthContext);

export default function TelegramProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [initData, setInitData] = useState("");
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    const initTelegram = () => {
      if (window.Telegram?.WebApp) {
        window.Telegram.WebApp.ready();
        window.Telegram.WebApp.expand();
        setInitData(window.Telegram.WebApp.initData);
        setIsReady(true);
      }
    };

    // Пробуем инициализировать сразу
    initTelegram();

    // И добавляем слушатель для скрипта
    window.addEventListener("telegram-web-app-script-loaded", initTelegram);

    return () => {
      window.removeEventListener(
        "telegram-web-app-script-loaded",
        initTelegram
      );
    };
  }, []);

  const value = useMemo(() => ({ initData, isReady }), [initData, isReady]);

  return (
    <TelegramAuthContext.Provider value={value}>
      <Script
        src="https://telegram.org/js/telegram-web-app.js"
        strategy="lazyOnload"
        onLoad={() => {
          window.dispatchEvent(new Event("telegram-web-app-script-loaded"));
        }}
      />
      {children}
    </TelegramAuthContext.Provider>
  );
}
