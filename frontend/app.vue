<script setup>
import { onMounted } from 'vue';
import { telegram } from '~/composables/useTelegram';

onMounted(() => {
  telegram.init();
});
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>