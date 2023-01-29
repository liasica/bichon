<template>
  <div class="relative" v-click-away="onHide">
    <div :class="triggerClassName" @click="show = !show">
      <slot :dropdownVisiable="show"></slot>
    </div>
    <transition name="zoom-in-top">
      <div
        v-if="show"
        :class="[
          'absolute right-0  shadow-lg z-10 overflow-y-auto scrollbar-x-none whitespace-nowrap bg-white dark:(bg-[#222222])',
          popupClassName,
        ]"
      >
        <slot name="prefix"></slot>
        <div ref="menu" :class="menuClassName" @click="onHide">
          <slot name="dropdown"></slot>
        </div>
        <slot name="suffix"></slot>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
interface Props {
  triggerClassName?: string;
  popupClassName?: string;
  menuClassName?: string;
}

const props = withDefaults(defineProps<Props>(), {
  popupClassName: "",
  menuClassName: "",
});

const show = ref(false);

function onHide() {
  show.value = false;
}
</script>

<style scoped lang="scss">
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

$fade-transition: transform 300ms cubic-bezier(0.23, 1, 0.32, 1),
  opacity 300ms cubic-bezier(0.23, 1, 0.32, 1) !default;

.zoom-in-top-enter-active,
.zoom-in-top-leave-active {
  opacity: 1;
  transform: scaleY(1);
  transition: $fade-transition;
  transform-origin: center top;
}
.zoom-in-top-enter,
.zoom-in-top-leave-active {
  opacity: 0;
  transform: scaleY(0);
}
</style>
