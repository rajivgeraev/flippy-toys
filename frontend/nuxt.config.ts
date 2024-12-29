// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      script: [
        {
          src: "https://telegram.org/js/telegram-web-app.js",
          async: true,
          onload: "window.telegramScriptLoaded = true;",
        },
      ],
    },
  },

  css: ["~/assets/css/main.css"],

  // Добавляем новые модули
  modules: [
    '@nuxtjs/tailwindcss',
    'nuxt-icon'
  ],

  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
});