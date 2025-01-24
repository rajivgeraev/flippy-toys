// src/app/[childId]/layout.tsx
import { Suspense } from "react";

interface LayoutProps {
  children: React.ReactNode;
  params: Promise<{
    childId: string;
  }>;
}

export default async function Layout({ children, params }: LayoutProps) {
  await params; // Дожидаемся разрешения params

  return (
    <Suspense fallback={<div>Loading...</div>}>
      <main className="min-h-screen">{children}</main>
    </Suspense>
  );
}
