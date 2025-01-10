<template>
  <div class="p-4">
    <!-- Фильтры -->
    <div class="mb-4">
      <div class="flex gap-2 overflow-x-auto">
        <button v-for="category in CATEGORIES" :key="category.value" @click="toggleCategory(category.value)" :class="[
          'px-4 py-2 rounded-full whitespace-nowrap',
          selectedCategories.includes(category.value)
            ? 'bg-purple-600 text-white'
            : 'bg-gray-100 text-gray-700'
        ]">
          {{ category.label }}
        </button>
      </div>
    </div>

    <!-- Сетка игрушек -->
    <div v-if="loading" class="flex justify-center py-8">
      <Icon name="lucide:loader" size="32" class="animate-spin text-purple-600" />
    </div>

    <div v-else-if="toys.length === 0" class="text-center py-8 text-gray-500">
      Игрушки не найдены
    </div>

    <div v-else class="grid grid-cols-2 gap-4">
      <div v-for="toy in toys" :key="toy.id" @click="router.push(`/toys/${toy.id}`)"
        class="bg-white rounded-lg shadow-sm overflow-hidden cursor-pointer">
        <div class="aspect-square">
          <img :src="toy.photos[0]?.url || '/placeholder.jpg'" class="w-full h-full object-cover" :alt="toy.title" />
        </div>
        <div class="p-3">
          <h3 class="font-medium text-sm mb-1">{{ toy.title }}</h3>
          <p class="text-gray-500 text-xs">{{ CONDITIONS.find(c => c.value === toy.condition)?.label }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { api } from '~/services/api';

const router = useRouter();
const loading = ref(true);
const toys = ref([]);
const selectedCategories = ref([]);

const CATEGORIES = [
  { value: 'construction_toys', label: 'Конструкторы' },
  { value: 'dolls', label: 'Куклы' },
  { value: 'vehicles', label: 'Транспорт' },
  { value: 'educational', label: 'Обучающие' },
  { value: 'outdoor', label: 'Уличные' }
];

const CONDITIONS = [
  { value: 'new', label: 'Новое' },
  { value: 'like_new', label: 'Как новое' },
  { value: 'good', label: 'Хорошее' },
  { value: 'acceptable', label: 'Удовлетворительное' }
];

const toggleCategory = (category) => {
  const index = selectedCategories.value.indexOf(category);
  if (index === -1) {
    selectedCategories.value.push(category);
  } else {
    selectedCategories.value.splice(index, 1);
  }
  fetchToys();
};

const fetchToys = async () => {
  loading.value = true;
  try {
    const response = await api.listToys({
      categories: selectedCategories.value
    });
    toys.value = response;
  } catch (error) {
    console.error('Error fetching toys:', error);
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchToys();
});
</script>