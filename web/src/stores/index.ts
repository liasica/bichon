import { createPinia } from "pinia";

import useAppStore from "./modules/app";
import useUserStore from "./modules/user";
import useGroupStore from "./modules/group";
import useMessageStore from "./modules/message";

const store = createPinia();

export default store;

export { store, useAppStore, useUserStore, useGroupStore, useMessageStore };
