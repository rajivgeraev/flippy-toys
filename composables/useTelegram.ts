export const useTelegram = () => {
    const isInitialized = ref(false);
    const webApp = ref<any>(null);
    const user = ref<any>(null);
  
    const init = () => {
      // Only run on client side
      if (process.client) {
        try {
          if (window.Telegram?.WebApp) {
            webApp.value = window.Telegram.WebApp;
            isInitialized.value = true;
            
            if (webApp.value.initDataUnsafe?.user) {
              user.value = webApp.value.initDataUnsafe.user;
            }
            
            webApp.value.enableClosingConfirmation();
            webApp.value.expand();
          }
        } catch (error) {
          console.error('Error initializing Telegram WebApp:', error);
        }
      }
    };
  
    onMounted(() => {
      init();
    });
  
    return {
      isInitialized,
      webApp,
      user,
      init
    };
  };