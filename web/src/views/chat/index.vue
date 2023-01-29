<template>
  <div class="relative flex flex-1 h-screen">
    <ChatList @group-select="onGroupSelect" />
    <div class="flex-1 flex">
      <div
        class="flex-1 mt-13px flex flex-col rounded-tl-10px rounded-tr-10px overflow-hidden"
      >
        <div
          class="bg-[#EDEFF3] h-20 font-medium px-6 flex items-center justify-between"
        >
          <span>{{ currentGroup?.name }}</span>
          <!-- <SvgIcon name="more-vertical" class="w-8 h-8 text-[#4E5969]" /> -->
        </div>
        <div
          class="flex-1 pb-6 flex flex-col min-w-400px max-h-[calc(100vh-91px)] bg-[#F7F8FA]"
        >
          <div ref="chatEl" class="flex-1 px-6 pt-6px pb-6 overflow-auto">
            <div v-if="group.currentId && !isFinish" class="text-center py-2">
              <SvgIcon
                name="loading"
                class="w-20px h-20px animate-spin text-primary"
              />
            </div>
            <div v-for="(item, index) in msgs" :id="item.id">
              <message-item :data="item" :index="index" />
            </div>

            <div
              v-if="!group.currentId"
              class="h-full flex items-center justify-center"
            >
              <SvgIcon name="empty" class="w-80 h-80 text-[#4E5969]" />
            </div>
          </div>
          <div v-if="group.currentId" class="px-6">
            <div v-if="quoteMessage" class="quote-message-box">
              <div class="quote-message">{{ quoteMessage?.content }}</div>
              <SvgIcon
                name="clear"
                class="w-4 h-4 text-[#E5E6EB] ml-2 cursor-pointer"
                @click="clearQuoteMessage"
              />
            </div>
            <div class="bg-white flex px-4 py-7px rounded-30px items-end">
              <textarea
                ref="inputEl"
                v-model="content"
                type="text"
                rows="1"
                class="flex-1 outline-transparent resize-none py-10px"
                placeholder="Enter your message here"
                @keydown.enter.exact="onSubmit"
              />
              <SvgIcon
                @click="onSubmit"
                name="send"
                class="w-46px h-46px text-white cursor-pointer"
              />
            </div>
          </div>
        </div>
      </div>
      <chat-details />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useWebSocket } from "@vueuse/core";
import {
  useGroupStore,
  useUserStore,
  useMessageStore,
  useAppStore,
} from "@/stores";
import { useKeyDB } from "@/db";
import { useGroupKeyShare, useGroupKeyUsed } from "@/api/group";
import { useMessageCreate } from "@/api/message";
import { useScroll } from "@vueuse/core";
import ChatDetails from "./chat-details.vue";
import ChatList from "./chat-list.vue";
import MessageItem from "./message-item.vue";
import { setTimeout } from "timers/promises";

enum Operate {
  Register = 1,
  Chat = 2,
}

enum scrollBehavior {
  Auto = "auto",
  Smooth = "smooth",
}

class Message<T> {
  operate: Operate;

  data: T;

  constructor(operate: Operate, data: T) {
    this.operate = operate;
    this.data = data;
  }

  toArrayBuffer(): ArrayBuffer {
    const str = JSON.stringify(this);
    const encoder = new TextEncoder();
    return encoder.encode(str);
  }

  static async fromBlob<T>(blob: Blob): Promise<Message<T>> {
    const buffer = await blob.arrayBuffer();
    const decoder = new TextDecoder();
    const str = decoder.decode(buffer);
    return JSON.parse(str);
  }
}

const group = useGroupStore();
const user = useUserStore();
const message = useMessageStore();
const app = useAppStore();

const content = ref("");

const chatEl = ref<HTMLElement | null>(null);
const inputEl = ref<HTMLElement | null>(null);

const { arrivedState } = useScroll(chatEl, {
  offset: { top: 30, bottom: 50 },
});
const { top: toTop, bottom: toBottom } = toRefs(arrivedState);

const currentGroup = computed(() => group.currentGroup);
const msgs = computed(() => message.messageList);
const quoteMessage = computed(() => message.quoteMessage);

const isFinish = computed(() => message.finished[group.currentId]);

const { send, status, data, open } = useWebSocket<Blob>(
  `${import.meta.env.VITE_WEBSOCKET_URL}/${user.address}`,
  {
    heartbeat: {
      message: "",
      interval: 30000,
    },
    autoReconnect: {
      retries: 10,
      delay: 1000,
      onFailed: () => {
        // TODO: 错误提示
      },
    },
  }
);

onMounted(async () => {
  group.getGroupList();
  // user.getMemberInfo();
});

const onGroupSelect = async (item: Model.GroupDetail) => {
  scrollChatBottom();
};

const onSubmit = async () => {
  const msgContent = content.value.trim();
  if (msgContent === "") return;
  const groupKey = await useKeyDB()?.getKey(group.currentId);
  if (groupKey) {
    const b64 = window.ecdhEncrypt(groupKey.sharedKey, msgContent);
    console.info(`encrypted data: ${b64}`);
    const data = window.ecdhDecrypt(groupKey.sharedKey, b64);
    console.info(`decrypted data: ${data}`);
    // send message
    const msg: Model.ChatMessage = {
      groupId: group.currentId,
      memberId: user.id,
      content: b64,
    };
    if (quoteMessage.value) {
      msg.parentId = quoteMessage.value?.id;
    }
    const res = await useMessageCreate(msg);
    if (res) {
      parseChatMessage(res);
      message.setQuoteMessage(null);
    }
    content.value = "";
  }
};

const clientRegister = async () => {
  // 提交群组keys
  try {
    if (!group.items.length) {
      await group.getGroupList();
    }
    const keys = (
      await Promise.all(
        group.items.map(async (item) => {
          let groupKey;
          try {
            groupKey = await useKeyDB()?.getKey(item.id);
          } catch (err) {}
          if (!groupKey) {
            await useGroupKeyShare(item.id);
            groupKey = await useKeyDB()?.getKey(item.id);
          }
          return {
            groupId: item.id,
            keyId: groupKey?.id,
          };
        })
      )
    ).filter((item) => item.keyId);
    if (keys.length) {
      await useGroupKeyUsed({
        keys,
      });
      await Promise.all(group.items.map((item) => message.initChat(item.id)));
      if (group.currentId) {
        group.getGroupInfo(group.currentId);
        scrollChatBottom();
      }
      //  else {
      //   await group.setCurrentId(group.items[0].id);
      // }
    }
  } catch (error) {
    console.log(error);
  }
  // group.items.forEach(async (groupItem) => {
  //   let groupKey = await useKeyDB()?.getKey(groupItem.id);
  //   if (!groupKey) {
  //     await useGroupKeyShare(groupItem.id);
  //     groupKey = await useKeyDB()?.getKey(groupItem.id);
  //   }
  //   if (groupKey) {
  //     // 设置当前使用的key
  //     await useGroupKeyUsed({
  //       keys: [{ groupId: groupItem.id, keyId: groupKey.id }],
  //     });
  //   }
  // });
};

const parseChatMessage = async (msg: Model.ChatMessage) => {
  await message.pushMsg(msg);
  // if (toBottom.value) {
  scrollChatBottom(scrollBehavior.Smooth);
  // }
  // 滚动到底部
  // 获取群组key
  // const groupKey = await useKeyDB()?.getKey(message.groupId);
  // if (groupKey) {
  //   // 解密消息
  //   const data = window.ecdhDecrypt(groupKey.sharedKey, message.content);
  //   console.info(`decrypted received message content: ${data}`);
  // }
  // console.log(message);
};

const scrollChatBottom = (behavior: scrollBehavior = scrollBehavior.Auto) => {
  nextTick(() => {
    const height = chatEl.value?.scrollHeight;
    chatEl.value?.scroll({
      behavior,
      top: height,
    });
  });
};

const clearQuoteMessage = () => {
  message.setQuoteMessage(null);
};

const changeTextAreaHeight = () => {
  nextTick(() => {
    const el = inputEl.value;
    const scrollHeight = el?.scrollHeight || 0;
    const height = el?.offsetHeight || 0;
    if (el && height <= scrollHeight) {
      el!.style.height = "auto";
      el!.style.height = Math.min(164, el?.scrollHeight) + "px";
    }
  });
};

watch(
  () => data.value,
  async (v) => {
    if (v && v.size > 0) {
      // TODO: 处理消息
      const data = await Message.fromBlob(v);
      console.info("WebSocket received message:", data);
      switch (data.operate) {
        case Operate.Register:
          await clientRegister();
          break;
        case Operate.Chat:
          await parseChatMessage(data.data as Model.ChatMessage);
          break;
        default:
          break;
      }
    }
  }
);

watch(
  () => status.value,
  (v) => {
    console.info(`WebSocket: ${v}`);
    switch (v) {
      case "OPEN":
        // 已连接，开始注册
        send(new Message(Operate.Register, user.token).toArrayBuffer());
        break;
      case "CLOSED":
        // TODO: 关闭逻辑
        break;
      default:
        break;
    }
  }
);

watch(
  () => toTop.value,
  async (value) => {
    if (value) {
      // get more message
      const lastOffset = chatEl.value?.scrollHeight || 0;
      await message.loadMessage();
      nextTick(() => {
        const height = chatEl.value?.scrollHeight || 0;
        chatEl.value?.scroll({
          top: height - lastOffset,
        });
      });
    }
  }
);

watch(
  () => app.baseUrl,
  () => {
    open();
    clientRegister();
  }
);

watch(content, () => {
  changeTextAreaHeight();
});
</script>

<style scoped lang="scss">
.quote-message-box {
  @apply mb-3 ml-4 flex items-center;
  .quote-message {
    @apply inline-block bg-[#E5E6EB] text-xs px-15px py-6px rounded-tr-10px rounded-br-10px rounded-tl-10px max-w-[70%] break-all;
  }
}
</style>
