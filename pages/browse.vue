<template>
  <div class="home">
    <h1>Popular Toys</h1>
    <div class="toys-grid">
      <ToyCard v-for="toy in toys" :key="toy.id" :toy="toy" />
    </div>
  </div>
</template>

<script setup>
import ToyCard from '~/components/toys/ToyCard.vue';
const { toys } = useToys();
</script>

<style scoped>
.home {
  padding: 16px;
}

.toys-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  margin-top: 20px;
}
</style>