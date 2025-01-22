// src/lib/api/client.ts
export async function fetchApi<T>(
    endpoint: string,
    options: RequestInit = {}
): Promise<T> {
    if (typeof window === 'undefined') {
        throw new Error('API calls should be made on client side only');
    }

    const telegramInitData = window.Telegram?.WebApp?.initData || '';
    const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}${endpoint}`;
    const headers = {
        'Content-Type': 'application/json',
        'X-Telegram-Init-Data': telegramInitData,
        ...options.headers,
    };

    console.log('API Request:', {
        url,
        method: options.method || 'GET',
        headers,
        telegramInitData,
        windowTelegram: !!window.Telegram,
        windowWebApp: !!window.Telegram?.WebApp
    });

    const response = await fetch(url, {
        ...options,
        headers
    });

    if (!response.ok) {
        console.error('API Response Error:', {
            status: response.status,
            statusText: response.statusText,
            headers: Object.fromEntries(response.headers.entries())
        });
        throw new Error(`API error: ${response.statusText}`);
    }

    return response.json();
}