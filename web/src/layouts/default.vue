<template>
  <div class="flex">
    <Sidebar></Sidebar>
    <router-view :key="route.name || ''" />
  </div>
</template>

<script setup lang="ts">
const route = useRoute();
</script>

<style scoped lang="scss"></style>
