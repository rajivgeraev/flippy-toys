// src/lib/api/client.ts
export async function fetchApi<T>(
    endpoint: string,
    options: RequestInit = {}
): Promise<T> {
    if (typeof window === 'undefined') {
        throw new Error('API calls should be made on client side only');
    }

    const telegramInitData = window.Telegram?.WebApp?.initData || 'query_id=AAEw6icIAAAAADDqJwj_1ZO6&user=%7B%22id%22%3A136833584%2C%22first_name%22%3A%22Rajiv%22%2C%22last_name%22%3A%22Geraev%22%2C%22username%22%3A%22RajivGeraev%22%2C%22language_code%22%3A%22en%22%2C%22is_premium%22%3Atrue%2C%22allows_write_to_pm%22%3Atrue%2C%22photo_url%22%3A%22https%3A%5C%2F%5C%2Ft.me%5C%2Fi%5C%2Fuserpic%5C%2F320%5C%2F0iH3m1Z1r2RmgHqCRUXvQ3K55sTZamIwyhEYq7yGq6w.svg%22%7D&auth_date=1737578479&signature=ULAKqmvcIstT7x_WOhK5gkpkRwIgCa7wrCxqRvX_S1xI2AhVmzR0RpRwLgs4Jm8AA8q2-p_ebQokqGWRIYriCw&hash=d326a5d0fd184046d7ac5435d7758bc90b222fa9f7d925152a4121ee75238d4a';
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