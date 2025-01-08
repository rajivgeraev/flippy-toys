interface TelegramProfile {
  id: string;
  user_id: string;
  telegram_id: number;
  username: string;
  first_name: string;
  last_name: string;
  photo_url: string;
  language_code: string;
  is_premium: boolean;
  phone?: string;
}

interface User {
  id: string;
  display_name: string;
  email?: string;
  phone?: string;
  real_first_name?: string;
  real_last_name?: string;
  access_level: string;
  created_at: string;
  updated_at: string;
  telegram_profile: TelegramProfile;
}

interface ToyCreate {
  title: string;
  description?: string;
  age_min?: number;
  age_max?: number;
  condition?: string;
  category?: string;
  photos?: File[];
}


interface PhotoData {
  secure_url: string;
  public_id: string;
  asset_id: string;
}

interface CreateToyRequest {
  title: string;
  description?: string;
  age_min?: number;
  age_max?: number;
  condition?: string;
  category?: string;
  photos: PhotoData[];
}

const BASE_URL = 'https://api.flippy.toys/api/v1';

export const api = {
  async validateUser(initData: string): Promise<User> {
    try {
      const response = await fetch(`${BASE_URL}/auth/validate`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ init_data: initData }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      console.error('Validation error:', error);
      throw error;
    }
  },

  async createToy(data: CreateToyRequest): Promise<any> {
    const response = await fetch(`${BASE_URL}/toys`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'X-Telegram-Init-Data': window.Telegram.WebApp.initData
      },
      body: JSON.stringify(data)
    });

    if (!response.ok) {
      throw new Error('Failed to create toy');
    }

    return response.json();
  },

  async getUserToys(data: CreateToyRequest): Promise<any> {
    const response = await fetch(`${BASE_URL}/toys/my`, {
      method: 'GET',
      headers: {
        'X-Telegram-Init-Data': window.Telegram.WebApp.initData
      },
      body: JSON.stringify(data)
    });

    if (!response.ok) {
      throw new Error('Failed to create toy');
    }

    return response.json();
  },

  async getToy(id: string): Promise<any> {
    const response = await fetch(`${BASE_URL}/toys/id/${id}`, {
      method: 'GET',
      headers: {
        'X-Telegram-Init-Data': window.Telegram.WebApp.initData
      }
    });

    if (!response.ok) {
      throw new Error('Failed to fetch toy');
    }

    return response.json();
  },

  async updateToy(id: string, data: any): Promise<any> {
    const response = await fetch(`${BASE_URL}/toys/id/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'X-Telegram-Init-Data': window.Telegram.WebApp.initData
      },
      body: JSON.stringify(data),
      credentials: 'include' // ВАЖНО!
    });

    if (!response.ok) {
      throw new Error('Failed to update toy');
    }

    return response.json();
  }

};
