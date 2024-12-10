export const useTelegram = () => {
    const isInitialized = ref(false);
    const webApp = ref<any>(null);
    const user = ref<any>(null);
    const initError = ref<string | null>(null);
  
    const init = () => {
      if (process.client) {
        try {
          console.log('Initializing Telegram WebApp...');
          
          // Log if Telegram object exists
          console.log('Telegram object exists:', !!window.Telegram);
          console.log('WebApp object exists:', !!window.Telegram?.WebApp);
          
          if (window.Telegram?.WebApp) {
            webApp.value = window.Telegram.WebApp;
            isInitialized.value = true;
            
            // Log the raw initData
            console.log('WebApp initData:', webApp.value.initData);
            console.log('WebApp initDataUnsafe:', webApp.value.initDataUnsafe);
            
            if (webApp.value.initDataUnsafe?.user) {
              user.value = webApp.value.initDataUnsafe.user;
              console.log('User data received:', user.value);
            } else {
              console.log('No user data in initDataUnsafe');
            }
            
            webApp.value.enableClosingConfirmation();
            webApp.value.expand();
            
            // Log the final state
            console.log('Initialization completed. State:', {
              isInitialized: isInitialized.value,
              hasUser: !!user.value,
              webAppReady: !!webApp.value
            });
          } else {
            initError.value = 'Telegram WebApp is not available';
            console.warn('Telegram WebApp is not available');
          }
        } catch (error) {
          initError.value = error instanceof Error ? error.message : 'Unknown error';
          console.error('Error initializing Telegram WebApp:', error);
        }
      } else {
        console.log('Running on server side, skipping initialization');
      }
    };
  
    onMounted(() => {
      console.log('Component mounted, starting initialization...');
      init();
    });
  
    return {
      isInitialized,
      webApp,
      user,
      initError,
      init
    };
  };