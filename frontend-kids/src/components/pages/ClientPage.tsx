// src/components/pages/ClientPage.tsx
"use client";
import { getChildren } from "@/lib/api/children";
import { notFound } from "next/navigation";
import { KidsHomePage } from "@/components/pages/KidsHomePage";
import { useEffect, useState } from "react";
import { Child } from "@/types/child";

export function ClientPage({ childId }: { childId: string }) {
  const [child, setChild] = useState<Child | null>(null);
  const [error, setError] = useState<boolean>(false);

  useEffect(() => {
    const fetchData = async () => {
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
  }, [childId]);

  if (error) return notFound();
  if (!child) return <div>Loading...</div>;

  return <KidsHomePage child={child} />;
}
