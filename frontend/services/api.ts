interface ValidateResponse {
  user: {
    id: number;
    first_name: string;
    last_name: string;
    username: string;
    language_code: string;
    is_premium: boolean;
    allows_write_to_pm: boolean;
    photo_url: string;
  };
  auth_date: number;
  hash: string;
  // другие поля если нужны
}

export const api = {
  async validateUser(initData: string): Promise<ValidateResponse> {
    try {
      const response = await fetch('https://api.flippy.toys/api/v1/auth/validate', {
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
  }
};