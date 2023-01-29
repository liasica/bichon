import { defineStore } from "pinia";
import { useJoinedGroupList, useGroupInfo, activeGroup } from "@/api/group";
import { useMessageStore } from "@/stores";

export default defineStore("group", {
  state: () => ({
    items: [] as Model.GroupDetail[],
    currentId: "" as string,
    groupMap: {} as any,
  }),
  actions: {
    async getGroupList() {
      const res = await useJoinedGroupList();
      this.items = res || [];
      // this.groupMap = this.items.reduce(
      //   (prev: any, next: Model.GroupDetail) => {
      //     prev[next.id] = next;
      //     return prev;
      //   },
      //   {}
      // );
    },
    async setCurrentId(value: string) {
      if (this.currentId === value) return;
      this.currentId = value;
      this.getGroupInfo(value);
      activeGroup(value);
      this.items = this.items.map((item) => {
        if (item.id === value) return { ...item, unreadCount: 0 };
        return item;
      });

      // const message = useMessageStore();
      // message.setQuoteMessage(null);
      // await message.initChat(value);
    },
    setUnreadCount(groupId: string) {
      this.items = this.items.map((item) => {
        if (item.id === groupId) {
          const unreadCount = Math.min((item.unreadCount || 0) + 1, 99);
          return { ...item, unreadCount };
        }
        return item;
      });
    },
    async getGroupInfo(groupId: string) {
      const res = await useGroupInfo(groupId);
      this.groupMap[groupId] = res;
    },
  },
  getters: {
    currentGroup(state) {
      return (
        state.groupMap[state.currentId] ||
        state.items.find((item) => item.id === state.currentId)
      );
    },
  },
});
