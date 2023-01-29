<template>
  <teleport to="body">
    <transition
      enter-active-class="transition ease-out duration-200 transform"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition ease-in duration-200 transform"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-show="visible"
        ref="modal-wrapper"
        class="fixed z-1000 inset-0 overflow-y-auto bg-black bg-opacity-50"
      >
        <div
          class="flex items-center justify-center min-h-screen text-center"
          @click="onClickAway"
        >
          <transition
            enter-active-class="transition ease-out duration-300 transform "
            enter-from-class="opacity-0 translate-y-10 scale-10"
            enter-to-class="opacity-100 translate-y-0 scale-100"
            leave-active-class="ease-in duration-200"
            leave-from-class="opacity-100 translate-y-0 scale-100"
            leave-to-class="opacity-0 translate-y-10 translate-y-0 scale-10"
            @before-enter="beforeEnter"
            @after-leave="afterLeave"
          >
            <div
              v-show="visible"
              class="relative w-450px bg-white py-5 px-10 rounded"
              role="dialog"
              ref="modal"
              aria-modal="true"
              aria-labelledby="modal-headline"
              @click.stop
            >
              <div
                v-if="closable"
                class="absolute right-0 top-0 cursor-pointer p-8px leading-0"
                @click="onClose"
              >
                <SvgIcon
                  name="close"
                  class="w-14px h-14px"
                  dark-color="#000000"
                />
              </div>
              <div class="font-semibold">{{ title }}</div>
              <slot></slot>
            </div>
          </transition>
        </div>
      </div>
    </transition>
  </teleport>
</template>

<script setup lang="ts">
interface Props {
  visible?: boolean;
  title?: string;
  clickaway?: boolean;
  closable?: boolean;
}

const emit = defineEmits(["update:visible", "close"]);

const props = withDefaults(defineProps<Props>(), {
  visible: true,
  title: "",
  clickaway: true,
  closable: true,
});

const onClickAway = () => {
  if (props.clickaway) {
    emit("update:visible", false);
    emit("close");
  }
};
const onClose = () => {
  emit("update:visible", false);
  emit("close");
};

const beforeEnter = () => {
  if (hasScrollBar()) {
    document.body.classList.add("modal-block-scroll");
  }
};
const afterLeave = () => {
  document.body.classList.remove("modal-block-scroll");
};

const hasScrollBar = () => {
  return (
    document.body.scrollHeight >
    (window.innerHeight || document.documentElement.clientHeight)
  );
};
</script>

<style scoped lang="scss"></style>
