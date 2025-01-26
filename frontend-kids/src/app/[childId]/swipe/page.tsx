// src/app/[childId]/swipe/page.tsx
import { ToySwiperPage } from '@/components/pages/toys/ToySwiperPage';

interface PageProps {
    params: Promise<{
        childId: string;
    }>;
}

export default async function SwipePage({ params }: PageProps) {
    const resolvedParams = await params;
    return <ToySwiperPage childId={resolvedParams.childId} />;
}