<template>
  <div
    v-if="data"
    :id="data.id"
    :class="[
      ' flex',
      { 'justify-end': data.isSelf, 'mt-8': data.isFirst && index !== 0 },
    ]"
  >
    <Blockies
      v-if="!data.isSelf"
      :class="[
        'w-40px h-40px rounded-full overflow-hidden mr-5',
        { 'opacity-0': !data.isFirst },
      ]"
      :seed="address"
      :size="10"
    />
    <div
      :class="[
        'msg-box',
        {
          self: data.isSelf,
          other: !data.isSelf,
          first: data.isFirst,
          last: data.isLast,
        },
      ]"
    >
      <div v-if="data.isFirst" class="name">
        {{ name(data.member) }}
      </div>
      <div
        class="msg"
        @mouseover="showMenu = true"
        @mouseout="showMenu = false"
      >
        {{ data.content }}
        <div
          :class="[
            'menu',
            showMenu ? 'flex' : 'hidden',
            { self: data.isSelf, other: !data.isSelf },
          ]"
        >
          <span @click="onCopy">复制</span>
          <span @click="onQuote">引用</span>
        </div>
      </div>
      <div v-if="data.isLast" class="text-[#86909C] text-xs">
        {{ date }}
      </div>
    </div>
    <Blockies
      v-if="data.isSelf"
      :class="[
        'w-40px h-40px rounded-full overflow-hidden ml-5',
        { 'opacity-0': !data.isFirst },
      ]"
      :seed="address"
      :size="10"
    />
  </div>
  <div v-if="data?.quote">
    <div
      :class="[
        'msg-box px-15 last',
        {
          self: data.isSelf,
          other: !data.isSelf,
        },
      ]"
    >
      <div class="msg quote-msg">
        <span>{{ name(data.quote?.member) }}：</span>
        {{ data.quote.content }}
      </div>
      <div class="text-[#86909C] text-xs">
        {{ date }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { shortAddress } from "@/utils";
import { useMessageStore } from "@/stores";
import { format } from "date-fns";
import { useClipboard } from "@vueuse/core";
import { useToast } from "vue-toastification";
const toast = useToast();
const message = useMessageStore();

interface Props {
  data?: Model.ChatMessage | null;
  index: number;
}

const showMenu = ref(false);

const props = withDefaults(defineProps<Props>(), {
  data: null,
});

const address = computed(() => props.data?.member?.address || "");

const date = computed(() =>
  props.data?.createdAt
    ? format(new Date(props.data?.createdAt), "yy.MM.dd HH:mm:ss")
    : ""
);

const name = computed(() => (member: Model.Member | undefined) => {
  return member ? member?.nickname || shortAddress(member?.address) : "";
});

const content = computed(() => props.data?.content || "");

const { text, copy, copied, isSupported } = useClipboard({ source: content });

const onCopy = async () => {
  await copy();
  toast("Copy Success!");
  showMenu.value = false;
  console.log(props.data);
};
const onQuote = () => {
  message.setQuoteMessage(props.data);
};
</script>

<style scoped lang="scss">
.msg-box {
  @apply flex flex-1 flex-col mb-6px;
  .name {
    @apply text-[#86909C] text-xs mb-2px;
  }
  .msg {
    @apply relative  min-h-12 max-w-[80%] bg-[#ABB5C1] text-white px-18px py-13px text-sm leading-22px break-all whitespace-pre-wrap;
    &.quote-msg {
      @apply max-w-[70%] bg-[#E5E6EB] text-[#86909C]; // TODO max-h-60px line-clamp-2
    }
  }
  &.self {
    @apply items-end;
    .msg {
      // text-right
      @apply rounded-tl-10px rounded-bl-10px;
    }
  }
  &.other {
    .msg {
      @apply rounded-tr-10px rounded-br-10px;
    }
  }
  &.first {
    &.self {
      .msg {
        @apply rounded-tr-10px;
      }
    }
    &.other {
      .msg {
        @apply rounded-tl-10px;
      }
    }
  }
  &.last {
    &.self {
      .msg {
        @apply rounded-br-10px;
      }
    }
    &.other {
      .msg {
        @apply rounded-bl-10px;
      }
    }
  }
  .menu {
    @apply absolute  flex-col items-center w-12 rounded bg-[rgba(247,248,250,0.4)] border border-[rgba(134,144,156,0.4)] text-sm text-[#4E5969] py-10px -bottom-40px z-1;
    box-shadow: 0px -3px 12px rgba(0, 0, 0, 0.1),
      0px 6px 14px rgba(0, 0, 0, 0.1);
    backdrop-filter: blur(25px);
    span {
      @apply w-full text-center cursor-pointer;
      & + span {
        @apply mt-1;
      }
      &:hover {
        @apply text-primary;
      }
    }
    &.self {
      @apply -left-30px;
    }
    &.other {
      @apply -right-30px;
    }
  }
}
</style>
