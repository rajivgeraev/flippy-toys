<template>
  <div class="p-4">
    <div class="grid grid-cols-2 gap-4">
      <ToyCard v-for="toy in toys" :key="toy.id" :toy="toy" />
    </div>
  </div>
</template>

<script setup>
import ToyCard from '~/components/toys/ToyCard.vue';
const { toys } = useToys()
</script>
