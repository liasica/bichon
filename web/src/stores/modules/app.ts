import { defineStore } from "pinia";
import Cookies from "js-cookie";
import { useNodes } from "@/api";

const KEY_THEME_MODE = "theme_mode";
const KEY_BASE_URL = "base_url";

const baseUrl = Cookies.get(KEY_BASE_URL) || import.meta.env.VITE_BASE_API_URL;

export default defineStore("app", {
  state: () => ({
    darkMode: Cookies.get(KEY_THEME_MODE) === "dark",
    baseUrl,
    apiUrlList: [] as string[],
  }),
  actions: {
    initDarkMode() {
      this.darkMode = Cookies.get(KEY_THEME_MODE) === "dark";
      if (this.darkMode) {
        document.documentElement.classList.add("dark");
      } else {
        document.documentElement.classList.remove("dark");
      }
    },
    toggleDarkMode() {
      this.darkMode = !this.darkMode;
      document.documentElement.classList.toggle("dark");
      const modeName = this.darkMode ? "dark" : "light";
      Cookies.set(KEY_THEME_MODE, modeName);
    },
    async getApiUrlList() {
      const res = await useNodes();
      this.apiUrlList = res || [];
    },
    setBaseUrl(url: string) {
      this.baseUrl = url;
      Cookies.set(KEY_BASE_URL, url);
    },
  },
  getters: {},
});
