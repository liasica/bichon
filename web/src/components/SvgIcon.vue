<template>
  <svg aria-hidden="true" class="icon">
    <use :xlink:href="symbolId" />
  </svg>
</template>

<script setup lang="ts">
interface Props {
  prefix?: string;
  name: string;
}

const props = withDefaults(defineProps<Props>(), {
  prefix: "icon",
});

const symbolId = computed(() => `#${props.prefix}-${props.name}`);
</script>
<style scoped lang="scss">
.icon {
  @apply inline-block fill-current;
}
</style>
