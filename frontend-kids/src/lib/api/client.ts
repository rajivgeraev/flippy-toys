// src/lib/api/client.ts
const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080';

export async function fetchApi<T>(
    endpoint: string,
    options: RequestInit = {}
): Promise<T> {
    const telegramInitData = typeof window !== 'undefined'
        ? window.Telegram?.WebApp?.initData
        : process.env.HEADER_X_TELEGRAM_INIT_DATA; // Only for test
        // : ''

    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...(telegramInitData && { 'X-Telegram-Init-Data': telegramInitData }),
            ...options.headers,
        },
    });

    if (!response.ok) {
        throw new Error(`API error: ${response.statusText}`);
    }

    return response.json();
}