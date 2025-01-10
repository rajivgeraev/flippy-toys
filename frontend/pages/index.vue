<template>
  <div class="p-4">
    <!-- Фильтры -->
    <div class="mb-4">
      <div class="flex gap-2 overflow-x-auto pb-2">
        <button v-for="category in CATEGORIES" :key="category.value" @click="toggleCategory(category.value)" :class="[
          'px-4 py-2 rounded-full whitespace-nowrap transition-colors',
          selectedCategories.includes(category.value)
            ? 'bg-purple-600 text-white'
            : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
        ]">
          {{ category.label }}
        </button>
      </div>
    </div>

    <!-- Игрушки -->
    <div v-if="loading" class="flex justify-center py-8">
      <Icon name="lucide:loader" size="32" class="animate-spin text-purple-600" />
    </div>

    <div v-else-if="toys.length === 0" class="text-center py-12">
      <Icon name="lucide:package" size="48" class="mx-auto mb-4 text-gray-400" />
      <p class="text-gray-500">Игрушки не найдены</p>
    </div>

    <div v-else class="grid grid-cols-2 gap-4" :class="{ 'opacity-50': loading }">
      <div v-for="toy in toys" :key="toy.id" @click="router.push(`/toys/view/${toy.id}`)"
        class="bg-white rounded-lg shadow-sm overflow-hidden cursor-pointer">
        <div class="aspect-square relative">
          <img :src="toy.photos[0]?.url || '/placeholder.jpg'" class="w-full h-full object-cover" :alt="toy.title" />
          <span :class="[
            'absolute top-2 right-2 px-2 py-1 rounded-full text-xs font-medium',
            getConditionClass(toy.condition)
          ]">
            {{ CONDITIONS.find(c => c.value === toy.condition)?.label }}
          </span>
        </div>
        <div class="p-3">
          <h3 class="font-medium text-sm mb-1">{{ toy.title }}</h3>
          <p class="text-gray-500 text-xs truncate">{{ toy.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { api } from '~/services/api';

const router = useRouter();
const route = useRoute();
const loading = ref(true);
const toys = ref([]);
const selectedCategories = ref([]);

// Инициализация фильтров из URL
onMounted(() => {
  const categories = route.query.categories;
  if (categories) {
    selectedCategories.value = Array.isArray(categories) ? categories : [categories];
  }
  fetchToys();
});

// Синхронизация URL с фильтрами
watch(selectedCategories, (newCategories) => {
  router.push({
    query: {
      ...route.query,
      categories: newCategories.length ? newCategories : undefined
    }
  });
}, { deep: true });

const CATEGORIES = [
  { value: 'construction_toys', label: 'Конструктор' },
  { value: 'dolls', label: 'Куклы' },
  { value: 'vehicles', label: 'Машинки' },
  { value: 'educational', label: 'Развивающие' },
  { value: 'outdoor', label: 'Для улицы' },
  { value: 'board_games', label: 'Настольные игры' },
  { value: 'electronic', label: 'Электронные' },
  { value: 'stuffed_animals', label: 'Мягкие игрушки' },
  { value: 'action_figures', label: 'Фигурки' },
  { value: 'arts_crafts', label: 'Творчество' },
  { value: 'musical', label: 'Музыкальные' },
  { value: 'other', label: 'Другое' }
];

const CONDITIONS = [
  { value: 'new', label: 'Новое' },
  { value: 'like_new', label: 'Как новое' },
  { value: 'good', label: 'Хорошее' },
  { value: 'acceptable', label: 'Удовлетворительное' }
];

const getConditionClass = (condition) => {
  switch (condition) {
    case 'new':
      return 'bg-green-100 text-green-800';
    case 'like_new':
      return 'bg-blue-100 text-blue-800';
    case 'good':
      return 'bg-yellow-100 text-yellow-800';
    case 'acceptable':
      return 'bg-orange-100 text-orange-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
};

const toggleCategory = async (category) => {
  const index = selectedCategories.value.indexOf(category);
  if (index === -1) {
    selectedCategories.value.push(category);
  } else {
    selectedCategories.value.splice(index, 1);
  }
  // Сразу вызываем fetchToys после изменения категорий
  await fetchToys();
};


// Обновляем fetchToys при изменении фильтров
// watch(selectedCategories, async () => {
//   await fetchToys();
// });


const fetchToys = async () => {
  loading.value = true;
  console.log('Fetching toys with categories:', selectedCategories.value); // добавим лог
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
</script>