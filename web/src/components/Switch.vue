<template>
  <label class="cursor-pointer flex items-center">
    <input
      v-bind="$attrs"
      class="input"
      type="checkbox"
      :checked="modelValue"
      @change="onChange"
    />
    <span class="switch"></span>
    <span
      class="mr-3 text-[#1a202c] overflow-hidden overflow-ellipsis whitespace-nowrap"
      >{{ label }}</span
    >
  </label>
</template>

<script setup lang="ts">
interface Props {
  label?: string;
  modelValue: boolean;
}

const props = withDefaults(defineProps<Props>(), {});

const emit = defineEmits(["update:modelValue"]);

const onChange = (e: any) => {
  emit("update:modelValue", e.target.checked);
};
</script>

<style scoped lang="scss">
/* Visually hide the checkbox input */
.input {
  @apply absolute w-1px h-1px p-0 -ml-1px overflow-hidden whitespace-nowrap border-0;
  clip: rect(0, 0, 0, 0);
  &:checked + .switch {
    @apply bg-[#864deb];
    &::before {
      border-color: #864deb;
      /* Move the inner circle to the right */
      transform: translateX(
        calc(var(--switch-container-width) - var(--switch-size))
      );
    }
  }
  &:focus + .switch::before {
    border-color: var(--gray);
  }
  &:focus:checked + .switch::before {
    border-color: var(--dark-teal);
  }
  &:disabled + .switch {
    background-color: var(--gray);
    &::before {
      background-color: var(--dark-gray);
      border-color: var(--dark-gray);
    }
  }
}
.switch {
  @apply relative flex flex-shrink-0 items-center bg-[#C9CDD4];
  --switch-container-width: 40px;
  --switch-size: calc(var(--switch-container-width) / 2);
  --light-gray: #e2e8f0;
  --gray: #cbd5e0;
  --dark-gray: #a0aec0;
  --teal: #4fd1c5;
  --dark-teal: #319795;
  /* Vertically center the inner circle */
  height: var(--switch-size);
  flex-basis: var(--switch-container-width);
  /* Make the container element rounded */
  border-radius: var(--switch-size);
  /* In case the label gets really long, the toggle shouldn't shrink. */
  transition: background-color 0.25s ease-in-out;
  &::before {
    @apply absolute content-[''] bg-white rounded-full left-3px;
    transition: transform 0.375s ease-in-out;
    height: calc(var(--switch-size) - 6px);
    width: calc(var(--switch-size) - 6px);
  }
}
</style>
