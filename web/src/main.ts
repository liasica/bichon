import "virtual:windi-base.css";
import "virtual:windi-components.css";
import "virtual:windi-utilities.css";
import "virtual:svg-icons-register";

import "@/styles/index.scss";

import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { store, useUserStore } from "@/stores";
import Blockies from "blockies-vue";
import Toast, { PluginOptions, POSITION } from "vue-toastification";
import "vue-toastification/dist/index.css";

import VueClickAway from "vue3-click-away";

const options: PluginOptions = {
  position: POSITION.TOP_CENTER,
  hideProgressBar: true,
  toastClassName: "custom-toast-class",
  timeout: 3000,
};

const app = createApp(App);

router.beforeEach(async (to, from) => {
  if (!to.meta || !to.meta.requireLogin) {
    return true;
  }

  const user = useUserStore();
  if (!user.isLogin) {
    try {
      const token = await user.login();
      await user.getMemberInfo();
      return !!token;
    } catch (err) {
      return false;
    }
  } else {
    await user.getMemberInfo();
  }
  return true;
});

app
  .use(router)
  .use(store)
  .use(Blockies)
  .use(VueClickAway)
  .use(Toast, options)
  .mount("#app");
