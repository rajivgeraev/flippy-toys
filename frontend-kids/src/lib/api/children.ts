// src/lib/api/children.ts
import { Child } from '@/types/child';
import { fetchApi } from './client';

export async function getChildren(): Promise<Child[]> {
    try {
        const response = await fetchApi<Child[]>('/api/v1/children');
        console.log('API Response:', response); // для отладки
        return response;
    } catch (error) {
        console.error('API Error:', error);
        throw error;
    }
}