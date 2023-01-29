<template>
  <div :class="['sider', { collapse }]">
    <div :class="['px-4', { 'text-center !px-0': collapse }]">
      <SvgIcon
        name="back"
        :class="[
          'w-8 h-8 text-[#4E5969] cursor-pointer',
          { 'transform rotate-180': collapse },
        ]"
        @click="collapse = !collapse"
      />
      <span v-show="!collapse" class="ml-14px font-medium">Chat details</span>
    </div>
    <div v-if="currentGroup" class="px-6 flex items-center mt-56px">
      <Blockies
        :class="[
          'rounded-full overflow-hidden',
          collapse ? 'w-54px h-54px ' : 'w-64px h-64px ',
        ]"
        :seed="currentGroup.address"
        :size="collapse ? 13 : 16"
      />
      <div v-show="!collapse" class="ml-3 flex-1">
        <div class="text-xl font-medium flex items-center">
          <span
            class="max-w-90px overflow-hidden overflow-ellipsis whitespace-nowrap"
            >{{ currentGroup.name }}</span
          >
          <SvgIcon
            name="icon-edit"
            class="w-4 h-4 ml-3 text-[#4E5969] cursor-pointer"
            @click="visible = true"
          />
        </div>
        <div class="text-xs leading-5 text-[#86909C] mt-1">
          {{ currentGroup.membersCount }}个人
        </div>
      </div>
    </div>
    <div
      v-if="currentGroup?.members"
      :class="['mt-12 px-6 font-medium', { 'text-center': collapse }]"
    >
      成员
    </div>
    <div v-if="currentGroup?.members" class="mt-4 px-6 flex-1 overflow-auto">
      <div
        v-for="item in currentGroup.members"
        class="flex px-2 py-3 flex items-center mb-4 overflow-hidden overflow-ellipsis whitespace-nowrap"
      >
        <Blockies
          :class="[
            'w-9 h-9 rounded-full overflow-hidden',
            collapse ? '' : 'mr-2',
          ]"
          :seed="item.address"
          :size="10"
        />
        <span
          v-show="!collapse"
          class="flex-1 overflow-hidden overflow-ellipsis whitespace-nowrap"
          >{{ item.nickname || shortAddress(item.address) }}</span
        >
      </div>
    </div>
  </div>
  <div
    v-if="visible && currentGroup"
    class="group-info-edit"
    v-click-away="onHide"
  >
    <div class="avatar">
      <Blockies
        class="w-22 h-22 rounded-full overflow-hidden"
        :seed="currentGroup.address"
        :size="22"
      />
    </div>
    <div class="mt-3 font-medium text-xl text-center mb-1 text-[#1D2129]">
      {{ currentGroup.name }}
    </div>
    <div class="text-[#86909C] text-xs text-center mb-40px">
      <span> {{ currentGroup.membersCount }}个人</span>
    </div>
    <div class="form-item">
      <div class="label">昵称</div>
      <div class="input-box">
        <input v-if="isOwner" type="text" placeholder="请输入" v-model="name" />
        <span v-else>{{ currentGroup.name }}</span>
      </div>
    </div>
    <div class="form-item">
      <div class="label">群备注</div>
      <div class="input-box">
        <input
          v-if="isOwner"
          type="text"
          placeholder="请输入"
          v-model="intro"
        />
        <span v-else>{{ currentGroup.intro }}</span>
      </div>
    </div>
    <div class="form-item">
      <div class="label">分类</div>
      <div class="flex-1">
        <Dropdown v-if="isOwner" popupClassName="w-full border">
          <div
            class="!w-full h-9 rounded bg-[#F7F8FA] flex items-center px-5px !border-transparent"
          >
            {{ category }}
          </div>
          <template v-slot:dropdown>
            <div
              v-for="item in categories"
              :class="[
                'dropdown-item-default h-9',
                { 'text-black bg-[#E9E9E9]': category === item },
              ]"
              @click="category = item"
            >
              <span>{{ item }}</span>
            </div>
          </template>
        </Dropdown>
        <span v-else>{{ currentGroup.category }}</span>
      </div>
    </div>
    <div v-if="isOwner" class="form-item">
      <div class="label">最大成员</div>
      <div class="input-box">
        <input
          v-if="isOwner"
          type="text"
          placeholder="请输入"
          v-model="maxMembers"
        />
        <span v-else>{{ currentGroup.membersMax }}人</span>
      </div>
    </div>
    <div v-if="isOwner" class="form-item !mt-8">
      <div class="label">创建时间</div>
      <div class="input-box">{{ createTime }}</div>
    </div>
    <div v-if="isOwner" class="form-item">
      <div class="label">邀请链接</div>
      <div class="input-box">
        <span class="text-sm text-primary cursor-pointer" @click="onCopy"
          >复制</span
        ><span
          class="text-sm text-primary ml-14px cursor-pointer"
          @click="reGen"
          >重新生成</span
        >
      </div>
    </div>
    <div v-if="isOwner" class="form-item !mt-8">
      <div class="label !text-black">是否公开</div>
      <div class="ml-2">
        <Switch v-model="isPublic" />
      </div>
    </div>
    <button v-if="!isOwner" class="leave-btn" @click="onLeave">
      退出该群聊
    </button>
    <span
      v-if="isOwner"
      class="absolute top-6 right-6 text-sm text-primary cursor-pointer"
      @click="onSubmit"
      >完成</span
    >
    <LeaveGroupConfirm v-model:visible="leaveGroupVisible" />
  </div>
</template>

<script setup lang="ts">
import { useGroupStore } from "@/stores";
import { shortAddress } from "@/utils";
import { format } from "date-fns";

import {
  useGroupCategories,
  useRegenerateInviteCode,
  updateGroupInfo,
} from "@/api/group";
import { useClipboard } from "@vueuse/core";
import { useToast } from "vue-toastification";
import LeaveGroupConfirm from "./leave-group-confirm.vue";

const toast = useToast();
const group = useGroupStore();

const collapse = ref(false);

const leaveGroupVisible = ref(false);

const categories = ref<string[]>([]);
const visible = ref(false);
const name = ref("");
const intro = ref("");
const category = ref("");
const maxMembers = ref(0);
const isPublic = ref(false);

onMounted(() => {
  getCategories();
});

const getCategories = async () => {
  const res = await useGroupCategories();
  categories.value = res || [];
};

const currentGroup = computed(() => group.currentGroup as Model.GroupDetail);

const isOwner = computed(() => currentGroup.value.owner);

const createTime = computed(() =>
  currentGroup.value.createdAt
    ? format(new Date(currentGroup.value.createdAt), "yyyy.MM.dd")
    : ""
);

const source = computed(
  () => `${location.origin}?inviteCode=${currentGroup.value.inviteCode || ""}`
);

const { text, copy, copied, isSupported } = useClipboard({ source });

const onCopy = async () => {
  await copy();
  toast("Copy Success!");
};

const reGen = async () => {
  await useRegenerateInviteCode(currentGroup.value.id);
  await group.getGroupInfo(currentGroup.value.id);
  await copy();
  toast("Regenerate Success!");
};

const onHide = () => {
  visible.value = false;
};

const onSubmit = async () => {
  await updateGroupInfo({
    groupId: currentGroup.value.id,
    name: name.value,
    intro: intro.value,
    maxMembers: +maxMembers.value,
    category: category.value,
    public: isPublic.value,
  });
  toast("Successfully!");
  group.getGroupInfo(currentGroup.value.id);
  visible.value = false;
};

const onLeave = () => {
  leaveGroupVisible.value = true;
};

watch(visible, (value: boolean) => {
  if (value) {
    console.log(currentGroup.value);
    name.value = currentGroup.value.name;
    intro.value = currentGroup.value.intro;
    maxMembers.value = currentGroup.value.membersMax;
    category.value = currentGroup.value.category;
    isPublic.value = currentGroup.value.public;
  }
});
</script>

<style scoped lang="scss">
.sider {
  @apply flex flex-col w-249px pt-37px transition-all;
  &.collapse {
    @apply w-102px;
  }
}
.group-info-edit {
  @apply absolute w-368px bg-white top-216px right-48px rounded-20px px-6 pb-6;
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
    @apply w-80px text-14px text-[#86909C];
  }
  .input-box {
    @apply flex-1;
    input {
      @apply w-full rounded bg-[#F7F8FA] h-9 p-3 text-14px;
    }
  }
  & + .form-item {
    @apply mt-14px;
  }
}

.leave-btn {
  @apply mx-auto w-168px h-38px flex items-center justify-center rounded-19px text-white text-sm mt-12;
  background: linear-gradient(101.56deg, #fa7b7b 1.32%, #f15a5a 95.14%);
}
</style>
