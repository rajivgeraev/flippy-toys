import type { RefSymbol } from "@vue/reactivity";
import { api } from '~/services/api';

const useGlobalTelegram = () => {
  const isInitialized = ref(false);
  const webApp = ref<any>(null);
  const user = ref<any>(null);
  const initError = ref<string | null>(null);

  const waitForTelegramScript = () => {
    return new Promise((resolve) => {
      if (window.Telegram?.WebApp) {
        resolve(true);
        return;
      }

      const checkInterval = setInterval(() => {
        if (window.Telegram?.WebApp) {
          clearInterval(checkInterval);
          resolve(true);
        }
      }, 100);

      setTimeout(() => {
        clearInterval(checkInterval);
        resolve(false);
      }, 5000);
    });
  };

  const init = async () => {
    if (import.meta.client) {
      console.log('Init started:', {
        timestamp: new Date().toISOString(),
        hasWindow: typeof window !== 'undefined',
        hasTelegram: !!window.Telegram
      });

      try {
        const scriptLoaded = await waitForTelegramScript();

        if (!scriptLoaded) {
          throw new Error('Telegram WebApp script failed to load');
        }

        if (window.Telegram?.WebApp) {
          webApp.value = window.Telegram.WebApp;

          const initData = new URLSearchParams(webApp.value.initData);
          const authDate = parseInt(initData.get('auth_date') || '0');
          const currentTime = Math.floor(Date.now() / 1000);

          console.log('Auth validation:', {
            authDate,
            currentTime,
            timeDifference: currentTime - authDate,
            initDataUnsafe: webApp.value.initDataUnsafe
          });

          if (webApp.value.initDataUnsafe?.user) {
            const validationResult = await api.validateUser(webApp.value.initData);

            user.value = validationResult.user;
            // user.value = webApp.value.initDataUnsafe.user;
            isInitialized.value = true;

            webApp.value.expand();

            const platform = webApp.value.platform;
            if (platform === 'android' || platform === 'ios') {
              // webApp.value.requestFullscreen();
            }

          } else {
            throw new Error('No user data in initDataUnsafe');
          }
        }
      } catch (error) {
        console.error('Initialization error:', {
          error,
          timestamp: new Date().toISOString(),
          windowState: Object.keys(window)
        });
        initError.value = error instanceof Error ? error.message : 'Unknown error';
      }
    }
  };

  return {
    isInitialized,
    webApp,
    user,
    initError,
    init
  };
};

export const telegram = useGlobalTelegram();