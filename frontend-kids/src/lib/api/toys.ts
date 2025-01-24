// src/lib/api/toys.ts
import { fetchApi } from './client';

interface Toy {
    id: string;
    title: string;
    description: string;
    photos: Array<{
        url: string;
        is_main: boolean;
    }>;
}

export async function getToys(): Promise<Toy[]> {
    return fetchApi<Toy[]>('/api/v1/toys');
}