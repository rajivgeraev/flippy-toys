// src/components/pages/ClientPage.tsx
// src/components/pages/ClientPage.tsx
"use client";

import { getChildren } from "@/lib/api/children";
import { notFound } from "next/navigation";
import { KidsHomePage } from "@/components/pages/KidsHomePage";
import { useEffect, useState } from "react";
import { Child } from "@/types/child";
import { useTelegramAuth } from "@/components/providers/TelegramProvider";

export function ClientPage({ childId }: { childId: string }) {
  const [child, setChild] = useState<Child | null>(null);
  const [error, setError] = useState<boolean>(false);
  const { isReady } = useTelegramAuth();

  useEffect(() => {
    const fetchData = async () => {
      if (!isReady) return; // Ждем готовности Telegram WebApp

      try {
        const children = await getChildren();
        const foundChild = children.find((c) => c.id === childId);

        if (!foundChild) {
          setError(true);
          return;
        }

        setChild(foundChild);
      } catch (error) {
        console.error("Failed to fetch child data:", error);
        setError(true);
      }
    };

    fetchData();
  }, [childId, isReady]); // Добавляем isReady в зависимости

  if (!isReady) return <div>Loading Telegram WebApp...</div>;
  if (error) return notFound();
  if (!child) return <div>Loading child data...</div>;

  return <KidsHomePage child={child} />;
}
