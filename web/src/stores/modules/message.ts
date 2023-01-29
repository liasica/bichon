import { defineStore } from "pinia";
import { useMessageList } from "@/api/message";
import { useGroupStore, useUserStore } from "@/stores";
import { useKeyDB } from "@/db";

export default defineStore("message", {
  state: () => ({
    chatMap: {} as any,
    quoteMessage: null as Model.ChatMessage | null,
    loading: {} as any,
    finished: {} as any,
  }),
  actions: {
    async initChat(groupId: string) {
      if (!this.chatMap[groupId]) {
        this.loading[groupId] = true;
        this.finished[groupId] = false;
        try {
          const res = await useMessageList({
            groupId,
          });

          const groupKey = await useKeyDB()?.getKey(groupId);
          const user = useUserStore();
          const msgs = (res || []).reverse().map((item) => {
            const content = window.ecdhDecrypt(
              groupKey!.sharedKey,
              item.content
            );
            if (item.quote) {
              item.quote.content = window.ecdhDecrypt(
                groupKey!.sharedKey,
                item.quote.content
              );
            }
            return {
              ...item,
              content,
              isSelf: user.id === item.member?.id,
            };
          });
          this.chatMap[groupId] = this.handleMsgs(msgs);
          if (msgs.length < 20) {
            this.finished[groupId] = true;
          }
        } catch (err) {
        } finally {
          this.loading[groupId] = false;
        }
      }
    },
    async loadMessage() {
      // TODO 切换到没有消息组的时候回报错
      const group = useGroupStore();
      const groupId = group.currentId;
      if (this.loading[groupId]) return;
      // state.chatMap[group.currentId]
      const params: Model.ChatMessageListReq = {
        groupId,
      };
      if (this.chatMap[groupId].length) {
        params.time = this.chatMap[groupId][0].createdAt;
      }
      const res = await useMessageList(params);
      const groupKey = await useKeyDB()?.getKey(groupId);
      const user = useUserStore();
      const msgs = (res || []).reverse().map((item) => {
        const content = window.ecdhDecrypt(groupKey!.sharedKey, item.content);
        return {
          ...item,
          content,
          isSelf: user.id === item.member?.id,
        };
      });
      const list = this.chatMap[groupId] || [];
      this.chatMap[groupId] = this.handleMsgs([...msgs, ...list]);
      if (msgs.length < 20) {
        this.finished[groupId] = true;
      }
    },
    handleMsgs(data: Model.ChatMessage[]) {
      let msgs: Model.ChatMessage[] = [];
      for (let i = 0; i < data.length; i++) {
        let item = data[i];
        const prev = i > 0 ? msgs[i - 1] : null;
        const next = i < data.length - 1 ? data[i + 1] : null;
        if (item.quote) {
          item.index = 0;
          item.isFirst = true;
          item.isLast = false;
          if (prev) {
            prev.isLast = true;
          }
        } else {
          if (!prev) {
            item.index = 0;
            item.isFirst = true;
          }
          if (!next) {
            item.isLast = true;
          }
          if (prev) {
            if (item.member?.id === prev.member?.id && !prev.quote) {
              if (prev.index === 8) {
                item.index = 9;
                item.isLast = true;
              } else if (prev.index === 9) {
                item.index = 0;
                item.isFirst = true;
              } else {
                item.index = (prev.index || 0) + 1;
              }
            } else {
              prev.isLast = true;
              item.index = 0;
              item.isFirst = true;
            }
          }
        }
        msgs.push(item);
      }
      return msgs;
    },
    async pushMsg(msg: Model.ChatMessage) {
      const groupKey = await useKeyDB()?.getKey(msg.groupId);
      const user = useUserStore();
      const content = window.ecdhDecrypt(groupKey!.sharedKey, msg.content);
      msg.content = content;
      msg.isSelf = user.id === msg.member?.id;
      if (msg.quote) {
        msg.quote.content = window.ecdhDecrypt(
          groupKey!.sharedKey,
          msg.quote.content
        );
      }
      const list = this.chatMap[msg.groupId];
      if (list?.length > 0) {
        let index = 0;
        for (let i = list.length - 1; i >= 0; i--) {
          if (
            new Date(msg.createdAt || 0).valueOf() >
            new Date(list[i].createdAt || 0).valueOf()
          ) {
            index = i;
            break;
          }
        }
        list.splice(index + 1, 0, msg);
        this.chatMap[msg.groupId] = this.handleMsgs(list);
      } else {
        this.chatMap[msg.groupId] = this.handleMsgs([msg]);
      }
      const group = useGroupStore();
      if (msg.groupId !== group.currentId) {
        group.setUnreadCount(msg.groupId);
      }
      // this.items = this.items.map((item) => {
      //   if (item.id === value) return { ...item, unreadCount: 0 };
      //   return item;
      // });
    },
    setQuoteMessage(msg: Model.ChatMessage | null) {
      this.quoteMessage = msg;
    },
  },
  getters: {
    messageList(state) {
      const group = useGroupStore();
      return state.chatMap[group.currentId];
    },
  },
});
