import { createI18n } from 'vue-i18n'
import el from '../i18n/locales/el.json'
import en from '../i18n/locales/en.json'

export default defineNuxtPlugin(({ vueApp }) => {
  const i18n = createI18n({
    legacy: false,
    locale: 'el',
    fallbackLocale: 'en',
    messages: {
      el,
      en
    }
  })

  vueApp.use(i18n)
}) 