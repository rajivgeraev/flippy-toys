// src/lib/api/client.ts
export async function fetchApi<T>(
    endpoint: string,
    options: RequestInit = {}
): Promise<T> {
    if (typeof window === 'undefined') {
        throw new Error('API calls should be made on client side only');
    }

    const telegramInitData = window.Telegram?.WebApp?.initData || '';

    const response = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}${endpoint}`, {
        ...options,
        headers: {
            'Content-Type': 'application/json',
            'X-Telegram-Init-Data': telegramInitData,
            ...options.headers,
        },
    });

    if (!response.ok) {
        throw new Error(`API error: ${response.statusText}`);
    }

    return response.json();
}