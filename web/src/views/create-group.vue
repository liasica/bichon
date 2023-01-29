<template>
  <div class="w-full flex items-center justify-center">
    <div
      class="w-808px h-304px rounded-20px border border-[#F2F3F5] shadow p-6 dark:(bg-[#373739] border-[#48484] shadow-dark)"
    >
      <div class="font-medium text-xl mb-10 text-default">Create Group</div>

      <div class="grid grid-cols-2 gap-x-120px gap-y-5">
        <div class="form-item">
          <div class="label">群组昵称</div>
          <div class="input-box">
            <input type="text" placeholder="请输入" v-model="groupName" />
          </div>
        </div>
        <div class="form-item">
          <div class="label">最大成员</div>
          <div class="input-box">
            <input
              type="text"
              placeholder="请输入"
              v-model.number="maxMembers"
            />
          </div>
        </div>
      </div>
      <div class="grid grid-cols-2 gap-x-120px mt-7 mb-10">
        <div class="form-item">
          <div class="label">分类</div>
          <div class="flex-1">
            <Dropdown popupClassName="w-full border dark:(border-[#aaa])">
              <div
                class="!w-full h-9 rounded bg-[#F7F8FA] flex items-center px-5px !border-transparent text-default dark:(bg-[rgba(255,255,255,0.16)])"
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
          </div>
        </div>
        <div class="form-item">
          <div class="label">是否公开</div>
          <div class="ml-2">
            <Switch v-model="isPublic" />
          </div>
        </div>
      </div>
      <button class="btn" :disabled="false" @click="onCreate">完成</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useGroupCreate, useGroupCategories } from "@/api/group";
import router from "@/router";
import { useToast } from "vue-toastification";

const toast = useToast();

const isPublic = ref(true);

const groupName = ref("");
const maxMembers = ref();
const loading = ref(false);
const categories = ref<string[]>([]);
const category = ref("");

onMounted(() => {
  getCategories();
});

const getCategories = async () => {
  const res = await useGroupCategories();
  categories.value = res || [];
};

const onCreate = async () => {
  try {
    loading.value = true;
    await useGroupCreate({
      name: groupName.value,
      category: category.value,
      maxMembers: maxMembers.value,
    } as Model.GroupCreateReq);
    router.replace({ path: "/" });
  } catch (err: any) {
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped lang="scss">
.image-add {
  @apply ml-2 w-88px h-88px flex items-center justify-center border-2px border-white bg-[#F7F8FA] rounded-full;
  box-shadow: 0px 12px 14px rgba(45, 13, 106, 0.04);
}
.form-item {
  @apply flex items-center;
  .label {
    @apply w-72px text-14px text-[#86909C];
  }
  .input-box {
    @apply flex-1;
    input {
      @apply w-full rounded bg-[#F7F8FA] h-9 p-3 text-14px text-default dark:(bg-[rgba(255,255,255,0.16)]);
    }
  }
}
.btn {
  @apply mx-auto  block w-220px h-38px rounded-19px text-white;
  background: linear-gradient(101.56deg, #a373f9 1.32%, #864deb 95.14%);
  &:disabled {
    opacity: 0.5;
  }
}
</style>
