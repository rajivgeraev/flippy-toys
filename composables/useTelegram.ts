export const useTelegram = () => {
  const isInitialized = ref(false);
  const webApp = window.Telegram?.WebApp;
  const user = ref<any>(null);
  
  const init = () => {
    if (webApp) {
      webApp.value = window.Telegram.WebApp;
      isInitialized.value = true;
      
      if (webApp.value.initDataUnsafe?.user) {
        user.value = webApp.value.initDataUnsafe.user;
      }
      
      // Enable closing confirmation if needed
      webApp.value.enableClosingConfirmation();
      
      // Expand webapp to full height
      webApp.value.expand();
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