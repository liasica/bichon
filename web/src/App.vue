<template>
  <router-view />
</template>
<script setup lang="ts">
import { useUserStore, useAppStore } from "@/stores";
import router from "@/router";

const app = useAppStore();

window.ethereum.on("accountsChanged", async (accounts) => {
  // Time to reload your interface with accounts[0]!
  const account = (accounts as string[])[0];
  // const user = useUserStore();
  if (account) {
    router.replace({ path: "/" });
    window.location.reload();
    // await user.login();
  }
});

onBeforeMount(() => {
  app.initDarkMode();
});
</script>
<style scoped lang="scss"></style>
