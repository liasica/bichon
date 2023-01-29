import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import DefaultLayout from "@/layouts/default.vue";

export const routes = [
  {
    path: "/",
    component: DefaultLayout,
    meta: {
      title: "HomeLayout",
    },
    children: [
      {
        path: "",
        name: "Group",
        component: () => import("@/views/group.vue"),
        meta: {
          title: "group",
        },
      },
      {
        path: "chat",
        name: "Chat",
        component: () => import("@/views/chat/index.vue"),
        meta: {
          title: "chat",
          requireLogin: true,
        },
      },
      {
        path: "create-group",
        name: "CreateGroup",
        component: () => import("@/views/create-group.vue"),
        meta: {
          title: "CreateGroup",
          requireLogin: true,
        },
      },
    ],
  },
] as RouteRecordRaw[];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }
    return to.path !== from.path
      ? {
          behavior: "smooth",
          top: 0,
          left: 0,
        }
      : {};
  },
});

export default router;
