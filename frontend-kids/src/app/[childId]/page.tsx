// src/app/[childId]/page.tsx
import { ClientPage } from "@/components/pages/ClientPage";

interface PageProps {
  params: Promise<{
    childId: string;
  }>;
}

export default async function Page({ params }: PageProps) {
  const resolvedParams = await params;
  return <ClientPage childId={resolvedParams.childId} />;
}
