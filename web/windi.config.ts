import { defineConfig } from "vite-plugin-windicss";

export default defineConfig({
  darkMode: "class",
  shortcuts: {
    "text-default": "text-[#000000] dark:text-[#FFFFFF]",
    "text-primary": "text-[#864DEB]",
    "dropdown-item-default":
      "flex items-center justify-between p-5px text-[#999999] leading-14px font-semibold hover:(text-black cursor-pointer bg-[#E9E9E9]) dark:hover:(text-white bg-[#2F2F2F])",
  },
  theme: {
    boxShadow: {
      DEFAULT:
        "0px -10px 12px rgba(45, 13, 106, 0.04), 0px 12px 14px rgba(45, 13, 106, 0.06)",
      dark: "0px -10px 12px rgba(0, 0, 0, 0.18), 0px 12px 14px rgba(0, 0, 0, 0.24)",
    },
    extend: {
      colors: {},
      fontFamily: {
        nunito: "Nunito Sans",
      },
    },
  },
  plugins: [require("windicss/plugin/line-clamp")],
});
