<template>
  <div class="relative flex items-center" @click="onChange">
    <SvgIcon
      :name="value ? 'checkbox-checked' : 'checkbox'"
      :class="['w-5 h-5', value ? 'text-white' : 'text-[#C9CDD4]']"
    />
    <input
      type="radio"
      :name="name"
      :value="value"
      class="opacity-0 absolute w-full h-full cursor-pointer"
    />
    <span class="px-5px"><slot></slot></span>
  </div>
</template>

<script setup lang="ts">
interface Props {
  modelValue?: string;
  name: string;
  label?: string;
  disabled?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  label: "",
  modelValue: "",
  disabled: false,
});

const emit = defineEmits(["update:modelValue", "change"]);

const value = computed(() => props.modelValue === props.label);

const onChange = () => {
  emit("update:modelValue", props.label);
};
</script>

<style scoped lang="scss"></style>
