import { defineConfig } from "vite"
import { resolve } from "path"
import vue from "@vitejs/plugin-vue"
import AutoImport from "unplugin-auto-import/vite"
import Components from "unplugin-vue-components/vite"
import { ElementPlusResolver } from "unplugin-vue-components/resolvers"

// https://vitejs.dev/config/
export default defineConfig({
        plugins: [
                vue(),
                AutoImport({
                        resolvers: [ElementPlusResolver()],
                }),
                Components({
                        resolvers: [ElementPlusResolver()],
                }),
        ],
        resolve: {
                alias: {
                        "@": resolve(__dirname, "src"), // 路径别名
                },
                extensions: [".js", ".json", ".ts"], // 使用路径别名时要省略的后缀名
        },
})
