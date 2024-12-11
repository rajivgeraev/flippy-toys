// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      script: [
        {
          src: "https://telegram.org/js/telegram-web-app.js",
          defer: true,
        },
      ],
    },
  },

  css: ["~/assets/css/main.css"],

  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
});
