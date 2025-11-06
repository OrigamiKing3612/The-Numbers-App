export default defineNuxtConfig({
    compatibilityDate: '2025-07-15',
    experimental: {
        viteEnvironmentApi: true
    },
    devtools: { enabled: true },
    css: ['@/assets/styles/main.css'],

    vite: {
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: `
                      @use "@/assets/styles/lib/_variables.scss" as *;
                      @use "@/assets/styles/lib/_colors.scss" as *;
                      @use "@/assets/styles/lib/_housecolors.scss" as *;
                  `
                }
            }
        }
    },

    components: [
        { path: "~/components/", pathPrefix: false }
    ],

    modules: ["nuxt-lucide-icons"],
    lucide: {
        namePrefix: 'Icon'
    }
})
