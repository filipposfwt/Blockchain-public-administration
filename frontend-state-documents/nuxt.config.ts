// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  ssr: false,
  future: {
    compatibilityVersion: 4,
  },
  devtools: { enabled: true },
  app: {
    pageTransition: false,
    layoutTransition: false,
    head: {
      charset: "utf-16",
      viewport: "width=500, initial-scale=1",
      title: "Public Document Service",
      meta: [
        { name: "title", content: "Public Document Service" },
        {
          name: "description",
          content: "Public Document Service",
        },
        {
          name: "keywords",
          content: ["Public Document Service"].join(","),
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
    '@nuxtjs/i18n'
  ],
  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE || 'http://localhost:8080'
    }
  },
  quasar: {
    plugins: [
      'Notify',
      'Dialog'
    ],
    extras: {
      font: 'roboto-font',
      fontIcons: ['material-icons']
    },
    config: {
      loading: {},
      notify: {},
      loadingBar: { color: "#8B0000", size: "5px", position: "top" },
      brand: {
        primary: "#8B0000",
        secondary: "#26A69A",
        accent: "#9C27B0",
        dark: "#1d1d1d",
        positive: "#21BA45",
        negative: "#C10015",
        info: "#31CCEC",
        warning: "#F2C037"
      },
    },
  },
  i18n: {
    strategy: 'no_prefix',
    defaultLocale: 'el',
    lazy: true,
    langDir: './app/locales',
    compilation: { strictMessage: false },
    locales: [
      { code: 'el', file: 'el.json' },
      { code: 'en', file: 'en.json' }
    ]
  },
  lodash: {
    prefix: "_",
  },
  devServer: {
    port: 3001
  }
})
