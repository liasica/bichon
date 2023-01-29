<template>
  <div v-if="visible" v-click-away="onClickAway">
    <div
      class="absolute top-0 left-0 right-0 bottom-0 bg-[rgba(0,0,0,0.1)] rounded-20px"
    ></div>
    <div
      class="absolute w-268px left-50px bottom-6 rounded-20px bg-white flex flex-col justify-center items-center p-6"
    >
      <SvgIcon name="logo" class="w-58px h-58px" />
      <span class="mt-2 leading-6 font-medium mb-30px">Puppy Chat</span>
      <span class="text-sm">{{
        isSuccess
          ? `您已成功退出群聊 ${currentGroup.name}`
          : "您真的要退出该群聊吗"
      }}</span>
      <button
        v-if="isSuccess"
        class="w-full btn btn-primary mt-5"
        @click="onOk"
      >
        我知道了
      </button>
      <div v-else class="btn-box">
        <button class="btn" @click="onCancel">取消</button>
        <button class="btn btn-primary" @click="onConfirm">确定</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useLeaveGroup } from "@/api/group";
import { useGroupStore } from "@/stores";

interface Props {
  visible?: boolean;
}

const emit = defineEmits(["update:visible"]);

const props = withDefaults(defineProps<Props>(), {
  visible: true,
});

const group = useGroupStore();

const currentGroup = computed(() => group.currentGroup as Model.GroupDetail);

const isSuccess = ref(false);

const onCancel = () => {
  emit("update:visible", false);
};

const onConfirm = async () => {
  await useLeaveGroup(currentGroup.value.id);
  isSuccess.value = true;
};

const onOk = () => {
  window.location.reload();
};

const onClickAway = () => {
  isSuccess.value && onOk();
};
</script>

<style scoped lang="scss">
.btn-box {
  @apply flex w-full mt-5;
  .btn {
    @apply flex-1;
  }
}
.btn {
  @apply h-38px rounded-19px bg-[#F2F3F5] text-sm text-[#4E5969];
  &.btn-primary {
    @apply text-white;
    background: linear-gradient(101.56deg, #a373f9 1.32%, #864deb 95.14%);
  }
  & + .btn {
    @apply ml-4;
  }
}
</style>
