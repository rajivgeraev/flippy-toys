// src/app/[childId]/page.tsx
import { getChildren } from "@/lib/api/children";
import { notFound } from "next/navigation";
import { KidsHomePage } from "@/components/pages/KidsHomePage";

interface PageProps {
  params: Promise<{
    childId: string;
  }>;
}

export default async function Page({ params }: PageProps) {
  try {
    const resolvedParams = await params;
    const children = await getChildren();
    const child = children.find((c) => c.id === resolvedParams.childId);

    if (!child) {
      notFound();
    }

    return <KidsHomePage child={child} />;
  } catch (error) {
    console.error("Failed to fetch child data:", error);
    notFound();
  }
}
