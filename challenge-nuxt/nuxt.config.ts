// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    head: {
        title: 'Challenge nuxt',
        htmlAttrs: {
            lang: 'es'
        },
        meta: [
            {charset: 'utf-8'},
            {name: 'viewport', content: 'width=device-width,initial-scale=1,maximum-scale=5,minimum-scale=1'},
            {hid: 'description', name: 'description', content: ''},
            {name: 'robots', content: 'all'},
            {"http-equiv": 'Content-type', content: 'text/html; charset=utf-8'},
            {"http-equiv": 'X-UA-Compatible', content: 'IE=edge,chrome=1'},
        ],
    },
    devtools: {enabled: true},
    modules: ['@nuxtjs/tailwindcss']


})