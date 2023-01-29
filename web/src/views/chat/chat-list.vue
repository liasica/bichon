<template>
  <div class="flex flex-col w-374px px-6 py-8 overflow-auto">
    <div class="flex items-center mb-34px">
      <Blockies
        class="w-64px h-64px rounded-full overflow-hidden"
        :seed="user.address"
        :size="16"
      />
      <div class="ml-3">
        <div class="text-xl font-semibold flex items-center">
          {{ user.nickname || shortAddress(user.address) }}
          <SvgIcon
            name="icon-edit"
            class="w-4 h-4 ml-3 text-[#4E5969] cursor-pointer"
            @click="visible = true"
          />
        </div>
        <div class="text-xs leading-5 text-[#86909C] mt-1">
          {{ user.intro }}
        </div>
      </div>
    </div>
    <div
      class="rounded-19px bg-[#F7F8FA] min-h-38px px-16px flex items-center mr-2"
    >
      <SvgIcon name="search" class="w-16px h-16px text-[#4E5969] mr-2" />
      <input
        v-model="keyword"
        type="text"
        placeholder="Search Group"
        class="bg-transparent text-sm outline-none w-full h-full"
      />
    </div>
    <div class="flex-1 mt-4">
      <div
        v-for="item in groupItems"
        :class="['group-item', { active: item.id === group.currentId }]"
        @click="onGroupSelect(item)"
      >
        <Blockies
          class="w-50px h-50px rounded-full overflow-hidden mr-3"
          :seed="item.address"
          :size="13"
        />
        <div class="flex-1 overflow-hidden overflow-ellipsis whitespace-nowrap">
          <div class="flex items-center justify-between">
            <span class="font-medium text-base">{{ item.name }}</span>
            <span class="text-[#86909C]">{{ lastUnreadTime(item) }}</span>
          </div>
          <div class="flex items-center justify-between">
            <div
              class="text-[#86909C] overflow-hidden overflow-ellipsis whitespace-nowrap"
            >
              {{ lastMessage(item) }}
            </div>
            <span
              v-if="item.unreadCount"
              class="min-w-5 min-h-5 bg-[#F15A5A] text-xs leading-5 text-white text-center rounded-full"
              >{{ item.unreadCount }}</span
            >
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-if="visible" class="user-info-edit" v-click-away="onHide">
    <div class="avatar">
      <Blockies
        class="w-22 h-22 rounded-full overflow-hidden"
        :seed="user.address"
        :size="22"
      />
    </div>
    <div class="mt-3 font-medium text-xl text-center mb-1 text-[#1D2129]">
      {{ shortAddress(user.address) }}
    </div>
    <div class="text-[#86909C] text-xs text-center mb-40px">
      <span class="max-w-[80%] overflow-hidden">{{ user.intro }}</span>
    </div>
    <div class="form-item">
      <div class="label">昵称</div>
      <div class="input-box">
        <input type="text" placeholder="请输入" v-model="nickname" />
      </div>
    </div>
    <div class="form-item">
      <div class="label">说明</div>
      <div class="input-box">
        <input type="text" placeholder="请输入" v-model="intro" />
      </div>
    </div>
    <span
      class="absolute top-6 right-6 text-sm text-primary cursor-pointer"
      @click="onSubmit"
      >完成</span
    >
  </div>
</template>

<script setup lang="ts">
import { format } from "date-fns";
import { updateMemberInfo } from "@/api";
import { useGroupStore, useUserStore, useMessageStore } from "@/stores";
import { shortAddress } from "@/utils";
import { useToast } from "vue-toastification";

const emit = defineEmits(["group-select"]);

const toast = useToast();
const user = useUserStore();
const group = useGroupStore();
const message = useMessageStore();

const keyword = ref("");
const visible = ref(false);

const nickname = ref("");
const intro = ref("");

const groupItems = computed(() => {
  if (keyword.value) {
    return group.items.filter((item) => item.name.indexOf(keyword.value) > -1);
  }
  return group.items;
});

const lastMessage = computed(() => {
  return (group: Model.GroupDetail) => {
    const msgList = message.chatMap[group.id];
    return msgList?.length ? msgList[msgList.length - 1].content : "";
  };
});

const lastUnreadTime = computed(() => {
  return (group: Model.GroupDetail) => {
    const msgList = message.chatMap[group.id];
    const unreadTime = msgList?.length
      ? msgList[msgList.length - 1].createdAt
      : group.unreadTime;
    return unreadTime ? format(new Date(unreadTime), "yy.MM.dd") : "";
  };
});

const onGroupSelect = async (item: Model.GroupDetail) => {
  await group.setCurrentId(item.id);
  emit("group-select", item);
};

const onSubmit = async () => {
  await updateMemberInfo({
    nickname: nickname.value,
    intro: intro.value,
  });
  toast("Successfully!");
  user.getMemberInfo();
  visible.value = false;
};

const onHide = () => {
  visible.value = false;
};

watch(visible, (value: boolean) => {
  if (value) {
    nickname.value = user.nickname;
    intro.value = user.intro;
  }
});
</script>

<style scoped lang="scss">
.group-item {
  @apply relative flex items-center px-3 py-11px mb-4 cursor-pointer border border-transparent rounded-[10px] bg-white;
  background-clip: padding-box;
  &:hover,
  &.active {
    &::after {
      @apply absolute -top-1px -bottom-1px -left-1px -right-1px content-[''] -z-1 rounded-[10px];
      background: linear-gradient(101.56deg, #a373f9 1.32%, #864deb 95.14%);
    }
  }
}
.user-info-edit {
  @apply absolute top-128px left-6 w-368px px-6 bg-white rounded-20px border pb-7;
  box-shadow: 0px -10px 12px rgba(45, 13, 106, 0.04),
    0px 12px 14px rgba(45, 13, 106, 0.06);
  .avatar {
    @apply inline-block border-2 ml-[50%] -mt-20px transform -translate-x-1/2 rounded-full border-white;
    filter: drop-shadow(0px 12px 14px rgba(45, 13, 106, 0.08));
  }
}

.form-item {
  @apply flex items-center;
  .label {
    @apply w-44px text-14px text-[#86909C];
  }
  .input-box {
    @apply flex-1;
    input {
      @apply w-full rounded bg-[#F7F8FA] h-9 p-3 text-14px;
    }
  }
  & + .form-item {
    @apply mt-5;
  }
}
</style>
