<template>
  <div
    class="flex flex-col items-center min-w-98px bg-[#F7F8FA] h-screen px-20px py-32px dark:(bg-[#090909])"
  >
    <!-- <SvgIcon name="logo" class="w-58px h-58px" /> -->
    <img src="/logo.jpg" class="w-58px h-58px rounded-8px" alt="chatpuppy" />
    <router-link
      to="/"
      class="mt-100px"
      v-slot="{ href, route, navigate, isActive, isExactActive }"
    >
      <SvgIcon
        :name="isExactActive ? 'menu-active' : 'menu'"
        class="w-32px h-32px"
      />
    </router-link>
    <router-link
      to="/chat"
      class="mt-40px"
      v-slot="{ href, route, navigate, isActive, isExactActive }"
    >
      <SvgIcon
        :name="isActive ? 'users-active' : 'users'"
        class="w-32px h-32px"
      />
    </router-link>
    <router-link
      to="/create-group"
      class="mt-40px"
      v-slot="{ href, route, navigate, isActive, isExactActive }"
    >
      <SvgIcon
        :name="isActive ? 'add-circle-active' : 'add-circle'"
        class="w-32px h-32px"
      />
    </router-link>
    <div class="flex-1 flex flex-col items-center justify-end">
      <SvgIcon
        name="setting"
        class="w-32px h-32px mt-40px cursor-pointer"
        @click="visible = true"
      />
      <!-- <SvgIcon name="account" class="w-32px h-32px mt-40px cursor-pointer" /> -->
      <!-- <SvgIcon name="turnoff" class="w-32px h-32px mt-40px cursor-pointer" /> -->
    </div>
  </div>
  <Modal v-model:visible="visible" title="设置">
    <div class="form-item mt-5">
      <div class="label">节点设置</div>
      <div class="flex-1">
        <Dropdown popupClassName="w-full border">
          <div
            class="!w-full h-9 rounded bg-[#F7F8FA] flex items-center px-5px !border-transparent"
          >
            {{ baseUrl }}
          </div>
          <template v-slot:dropdown>
            <div
              v-for="item in app.apiUrlList"
              :class="[
                'dropdown-item-default h-9',
                { 'text-black bg-[#E9E9E9]': baseUrl === item },
              ]"
              @click="baseUrl = item"
            >
              <span>{{ item }}</span>
            </div>
          </template>
        </Dropdown>
      </div>
    </div>
    <button class="btn mt-10" @click="onConfirm">Confirm</button>
  </Modal>
</template>

<script setup lang="ts">
import { useAppStore } from "@/stores";
const app = useAppStore();

const visible = ref(false);

const baseUrl = ref(app.baseUrl);

onBeforeMount(() => {
  app.getApiUrlList();
});

const onConfirm = () => {
  app.setBaseUrl(baseUrl.value);
  visible.value = false;
};

watch(visible, (value: boolean) => {
  if (value) {
    app.getApiUrlList();
    baseUrl.value = app.baseUrl;
  }
});
</script>

<style scoped lang="scss">
.btn {
  @apply mx-auto  block w-220px h-38px rounded-19px text-white;
  background: linear-gradient(101.56deg, #a373f9 1.32%, #864deb 95.14%);
  &:disabled {
    opacity: 0.5;
  }
}
.form-item {
  @apply flex items-center;
  .label {
    @apply w-72px text-14px text-[#86909C];
  }
  .input-box {
    @apply flex-1;
    input {
      @apply w-full rounded bg-[#F7F8FA] h-9 p-3 text-14px;
    }
  }
}
</style>
