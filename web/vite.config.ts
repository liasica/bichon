import path from "path";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import WindiCSS from "vite-plugin-windicss";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    // https://github.com/antfu/unplugin-auto-import
    AutoImport({
      imports: ["vue", "vue-router"],
      dts: "src/auto-imports.d.ts",
    }),
    // https://github.com/antfu/unplugin-vue-components
    Components({
      resolvers: [],
      dts: "src/components.d.ts",
    }),

    // https://cn.windicss.org/integrations/vite.html
    WindiCSS(),

    // https://github.com/vbenjs/vite-plugin-svg-icons
    createSvgIconsPlugin({
      iconDirs: [path.resolve(process.cwd(), "src/icons")],
      symbolId: "icon-[dir]-[name]",
    }),
  ],
  resolve: {
    alias: [
      {
        find: "@/",
        replacement: `${path.resolve(__dirname, "src")}/`,
      },
    ],
  },
  server: {
    host: "0.0.0.0",
    proxy: {
      "/api": {
        target: "https://chatpuppy.rarely.work",
        changeOrigin: true,
        // secure: false,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});
