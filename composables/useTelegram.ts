export const useTelegram = () => {
    const webApp = (window as any).Telegram?.WebApp;
    
    if (!webApp) {
      console.warn('Telegram WebApp is not available');
      return {
        user: null,
        webApp: null,
        isAvailable: false
      };
    }
  
    const user = webApp.initDataUnsafe?.user || null;
    
    return {
      user,
      webApp,
      isAvailable: true
    };
  };