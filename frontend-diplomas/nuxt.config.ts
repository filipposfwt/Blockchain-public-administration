// https://nuxt.com/docs/api/configuration/nuxt-config
import pkg from "./package.json";
import pallete from "./app/theme/pallete";

export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  ssr: false,
  future: {
    compatibilityVersion: 4,
  },
  devtools: { enabled: false },
  app: {
    pageTransition: false,
    layoutTransition: false,
    head: {
      charset: "utf-16",
      viewport: "width=500, initial-scale=1",
      title: "dms-university",
      meta: [
        { hid: "title", name: "title", content: "dms-university" },
        {
          hid: "description",
          name: "description",
          content: "dms-university",
        },
        {
          hid: "keywords",
          name: "keywords",
          content: ["dms-university"].join(","),
        },
      ],
    },
  },
  modules: [
    "@nuxt/fonts",
    "@nuxt/icon",
    "@nuxt/eslint",
    "@nuxt/image",
    "nuxt-quasar-ui",
    "@pinia/nuxt",
    "nuxt-lodash",
    "@nuxtjs/i18n",
  ],
  quasar: {
    plugins: ["Loading", "LoadingBar", "Notify"],
    extras: { fontIcons: ["line-awesome"] },
    config: {
      loading: {},
      notify: {},
      loadingBar: { color: "green", size: "5px", position: "top" },
      brand: pallete,
    },
  },
  i18n: {
    strategy: "no_prefix",
    defaultLocale: "el",
    lazy: true,
    langDir: "../app/i18n/locales/",

    compilation: { strictMessage: false },

    locales: [
      { code: "en", file: "en.json" },
      { code: "el", file: "el.json" },
    ],
  },

  lodash: {
    prefix: "_",
  },
});
