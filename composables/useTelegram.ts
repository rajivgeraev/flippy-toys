export const useTelegram = () => {
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

      // Таймаут через 5 секунд
      setTimeout(() => {
        clearInterval(checkInterval);
        resolve(false);
      }, 5000);
    });
  };

  const init = async () => {
    if (import.meta.client) {
      console.log("Init started:", {
        timestamp: new Date().toISOString(),
        hasWindow: typeof window !== "undefined",
        hasTelegram: !!window.Telegram,
      });

      try {
        const scriptLoaded = await waitForTelegramScript();

        if (!scriptLoaded) {
          throw new Error("Telegram WebApp script failed to load");
        }

        if (window.Telegram?.WebApp) {
          webApp.value = window.Telegram.WebApp;

          // Проверяем валидность данных
          const initData = new URLSearchParams(webApp.value.initData);
          const authDate = parseInt(initData.get("auth_date") || "0");
          const currentTime = Math.floor(Date.now() / 1000);

          console.log("Auth validation:", {
            authDate,
            currentTime,
            timeDifference: currentTime - authDate,
            initDataUnsafe: webApp.value.initDataUnsafe,
          });

          if (webApp.value.initDataUnsafe?.user) {
            user.value = webApp.value.initDataUnsafe.user;
            isInitialized.value = true;
            webApp.value.expand();
          } else {
            throw new Error("No user data in initDataUnsafe");
          }
        }
      } catch (error) {
        console.error("Initialization error:", {
          error,
          timestamp: new Date().toISOString(),
          windowState: Object.keys(window),
        });
        initError.value =
          error instanceof Error ? error.message : "Unknown error";
      }
    }
  };

  // Добавляем автоматическую переинициализацию при потере данных
  if (import.meta.client) {
    const checkInterval = setInterval(() => {
      if (isInitialized.value && !user.value) {
        console.log("Detected data loss, reinitializing...");
        init();
      }
    }, 5000);

    onUnmounted(() => {
      clearInterval(checkInterval);
    });
  }

  onMounted(() => {
    init();
  });

  return {
    isInitialized,
    webApp,
    user,
    initError,
    init,
  };
};
