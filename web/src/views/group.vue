<template>
  <div class="relative flex-1 h-screen overflow-auto flex flex-col">
    <div class="w-808px mx-auto h-[calc(100vh-58px)] flex flex-col">
      <div class="pt-8 font-medium text-xl mb-6 text-default">Explore</div>
      <div class="flex h-38px mb-5">
        <div
          class="rounded-19px bg-[#F7F8FA] px-16px flex items-center w-351px mr-2 dark:(bg-[rgba(255,255,255,0.16)])"
        >
          <SvgIcon name="search" class="w-16px h-16px text-[#4E5969] mr-2" />
          <input
            type="text"
            v-model="keyword"
            placeholder="Search"
            class="bg-transparent h-full w-full outline-none text-default"
            @keydown.enter="onKeywordChange"
          />
        </div>
        <Dropdown popupClassName="w-full border">
          <div
            class="rounded-19px bg-[#F7F8FA] px-27px py-5px text-14px flex items-center text-[#86909C] dark:(bg-[rgba(255,255,255,0.16)] text-[#ffffff])"
          >
            {{ category }}
            <SvgIcon
              name="icon-down"
              class="w-16px h-16px text-[#4E5969] ml-2 dark:(text-[#ffffff])"
            />
          </div>
          <template v-slot:dropdown>
            <div
              v-for="item in categories"
              :class="[
                'dropdown-item-default uppercase h-9',
                { 'text-black bg-[#E9E9E9]': category === item },
              ]"
              @click="onCategoryChange(item)"
            >
              <span>{{ item }}</span>
            </div>
          </template>
        </Dropdown>
      </div>
      <div ref="chatEl" class="flex-1 overflow-auto">
        <div class="grid grid-cols-4 gap-5 mb-10px">
          <div
            v-for="item in items"
            class="py-6 border border-[#F7F8FA] rounded-10px flex flex-col items-center bg-[#F7F8FA] hover:(!border-[#864DEB] bg-transparent) dark:(bg-transparent border-[#333333])"
          >
            <Blockies
              class="w-68px h-68px rounded-full overflow-hidden mb-3"
              :seed="item.address"
              :size="17"
            />
            <span
              class="text-[#1D2129] font-semibold leading-6 font-nunito text-default"
            >
              {{ item.name }}
            </span>
            <span class="text-14px leading-22px text-[#86909C] mt-1px">
              {{ item.membersCount }}个成员（{{ item.membersMax }}满）
            </span>
            <button
              :class="['btn', { joined: item.joined }]"
              @click="joinGroup(item)"
            >
              <template v-if="item.joined">
                <span class="hidden">进入</span>
                <span>已加入</span>
              </template>
              <span v-else>加入</span>
            </button>
          </div>
        </div>
        <div v-if="!isFinish" class="text-center py-2">
          <SvgIcon
            name="loading"
            class="w-20px h-20px animate-spin text-primary"
          />
        </div>
      </div>
    </div>
    <div class="sticky bg-white bottom-0 w-full h-58px dark:(bg-[#111111])">
      <div class="w-808px h-full mx-auto flex items-center justify-between">
        <div class="text-[#C9CDD4]">© 2022 Puppy chat.</div>
        <div>
          <a href="https://github.com/chatpuppy" target="_blank"
            ><SvgIcon name="github" class="icon"
          /></a>
          <a href="https://twitter.com/chatpuppynft" target="_blank"
            ><SvgIcon name="twitter" class="icon"
          /></a>
          <a href="https://discord.com/invite/QN658sJWkk" target="_blank"
            ><SvgIcon name="discord" class="icon"
          /></a>
          <SvgIcon
            :name="app.darkMode ? 'sun' : 'moon'"
            class="icon"
            @click="app.toggleDarkMode"
          />
          <SvgIcon name="copyright" class="icon" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useGroupJoin, useGroupList, useGroupCategories } from "@/api/group";
import { useUserStore, useAppStore, useGroupStore } from "@/stores";
import { useScroll } from "@vueuse/core";
import router from "@/router";

const user = useUserStore();
const app = useAppStore();
const group = useGroupStore();

const AllCate = "全部";
const route = useRoute();

const pagination = ref({
  current: 0,
  pages: 1,
  total: 1,
});
const pageSize = 16;
const isFinish = ref(false);

const keyword = ref("");
const categories = ref<string[]>([]);
const category = ref(AllCate);

const items = ref<Model.GroupListRes[]>([]);

const inviteCode = computed(() => route.query.inviteCode || "");

const chatEl = ref<HTMLElement | null>(null);

const { arrivedState } = useScroll(chatEl, {
  offset: { top: 30, bottom: 30 },
});

const { top: toTop, bottom: toBottom } = toRefs(arrivedState);

onMounted(() => {
  getList();
  getCategories();

  if (inviteCode.value) {
    joinInviteGroup();
  }
});

const getCategories = async () => {
  const res = await useGroupCategories();
  categories.value = [AllCate].concat(res || []);
};

const joinInviteGroup = async () => {
  await useGroupJoin({ inviteCode: inviteCode.value });
  getList();
};

const getList = async (reset: boolean = true) => {
  if (reset) {
    isFinish.value = false;
  }
  const res = await useGroupList({
    pageSize,
    name: keyword.value,
    category: AllCate === category.value ? "" : category.value,
    current: reset ? 1 : pagination.value.current + 1,
  });
  const list = res?.items || [];
  pagination.value = res!.pagination;
  items.value = reset ? list : [...items.value, ...list];
  if (list.length < pageSize) {
    isFinish.value = true;
  }
};

const joinGroup = async (item: any) => {
  if (!user.isLogin) {
    await user.login();
    await getList();
    item = items.value.find((group) => group.id === item.id);
  }
  if (item.joined) {
    group.setCurrentId(item.id);
    router.push({ path: "/chat" });
  } else {
    await useGroupJoin({ groupId: item.id });
    getList();
  }
};

const onCategoryChange = (item: string) => {
  category.value = item;
  getList();
};

const onKeywordChange = () => {
  getList();
};

watch(
  () => toBottom.value,
  async (value) => {
    if (value && !isFinish.value) {
      // get more message
      getList(false);
    }
  }
);

watch(
  () => user.token,
  () => {
    getList();
  }
);

watch(
  () => app.baseUrl,
  () => {
    getList();
  }
);
</script>

<style scoped lang="scss">
.icon {
  @apply w-7 h-7 text-[#C9CDD4] ml-4  cursor-pointer;
}
.btn {
  @apply rounded-20px border border-[#864DEB] text-[#864DEB] py-7px w-100px mt-3 disabled:(bg-[#E5E6EB] text-white border-[#E5E6EB] cursor-not-allowed) dark:(bg-transparent);
  &.joined:hover {
    span {
      &:first-child {
        @apply block;
      }
      &:last-child {
        @apply hidden;
      }
    }
  }
}
</style>
