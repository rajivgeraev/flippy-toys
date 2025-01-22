// src/app/[childId]/page.tsx
import { ClientPage } from "@/components/pages/ClientPage";

export const dynamic = "force-dynamic"; // Отключаем статическую генерацию
export const runtime = "edge"; // Используем edge runtime для быстрого старта

interface PageProps {
  params: Promise<{
    childId: string;
  }>;
}

export default async function Page({ params }: PageProps) {
  const resolvedParams = await params;
  return <ClientPage childId={resolvedParams.childId} />;
}
