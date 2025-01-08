<template>
    <div class="p-4">
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-xl font-bold">{{ isEditing ? 'Редактировать игрушку' : 'Просмотр игрушки' }}</h1>
            <button v-if="!isEditing" @click="toggleEdit" class="bg-purple-600 text-white px-4 py-2 rounded-full">
                Редактировать
            </button>
        </div>

        <div v-if="loading" class="flex justify-center py-8">
            <Icon name="lucide:loader" size="32" class="animate-spin text-purple-600" />
        </div>

        <div v-else>
            <!-- Фотографии -->
            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700">Фотографии</label>
                <div class="grid grid-cols-3 gap-4 mt-2">
                    <div v-for="photo in toy.photos" :key="photo.id" class="aspect-square rounded-lg overflow-hidden">
                        <img :src="photo.url" class="w-full h-full object-cover" :alt="toy.title" />
                    </div>
                </div>
            </div>

            <!-- Форма редактирования -->
            <form @submit.prevent="saveToy">
                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700">Название</label>
                    <input v-model="toy.title" :disabled="!isEditing" class="mt-1 block w-full p-2 border rounded" />
                </div>

                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700">Описание</label>
                    <textarea v-model="toy.description" :disabled="!isEditing"
                        class="mt-1 block w-full p-2 border rounded"></textarea>
                </div>

                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700">Состояние</label>
                    <select v-model="toy.condition" :disabled="!isEditing" class="mt-1 block w-full p-2 border rounded">
                        <option v-for="option in CONDITIONS" :key="option.value" :value="option.value">
                            {{ option.label }}
                        </option>
                    </select>
                </div>

                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700">Категория</label>
                    <select v-model="toy.category" :disabled="!isEditing" class="mt-1 block w-full p-2 border rounded">
                        <option v-for="option in CATEGORIES" :key="option.value" :value="option.value">
                            {{ option.label }}
                        </option>
                    </select>
                </div>

                <!-- Кнопки управления -->
                <div v-if="isEditing" class="flex justify-end gap-2">
                    <button type="button" @click="cancelEdit" class="bg-gray-500 text-white px-4 py-2 rounded-full">
                        Отмена
                    </button>
                    <button type="submit" :disabled="isSaving"
                        class="bg-purple-600 text-white px-4 py-2 rounded-full disabled:bg-purple-400">
                        {{ isSaving ? 'Сохранение...' : 'Сохранить' }}
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { api } from '~/services/api';

const CONDITIONS = [
    { value: 'new', label: 'Новое' },
    { value: 'like_new', label: 'Как новое' },
    { value: 'good', label: 'Хорошее' },
    { value: 'acceptable', label: 'Удовлетворительное' }
];

const CATEGORIES = [
    { value: 'construction_toys', label: 'Конструктор'},
    { value: 'dolls', label: 'Куклы'},
    { value: 'vehicles', label: 'Машинки'},
    { value: 'educational', label: 'Развивающие'},
    { value: 'outdoor', label: 'Для улицы'},
    { value: 'board_games', label: 'Настольные игры'},
    { value: 'electronic', label: 'Электронные'},
    { value: 'stuffed_animals', label: 'Мягкие игрушки'},
    { value: 'action_figures', label: 'Фигурки'},
    { value: 'arts_crafts', label: 'Творчество'},
    { value: 'musical', label: 'Музыкальные'},
    { value: 'other', label: 'Другое'}
];

const route = useRoute();
const router = useRouter();
const toy = ref({});
const originalToy = ref(null);
const loading = ref(true);
const isEditing = ref(false);
const isSaving = ref(false);

// Загрузка данных игрушки
const fetchToy = async () => {
    try {
        const response = await api.getToy(route.params.id);
        toy.value = response;
        originalToy.value = JSON.parse(JSON.stringify(response));
    } catch (error) {
        console.error('Ошибка загрузки игрушки:', error);
    } finally {
        loading.value = false;
    }
};

// Переключить режим редактирования
const toggleEdit = () => {
    isEditing.value = !isEditing.value;
};

// Отмена редактирования
const cancelEdit = () => {
    toy.value = JSON.parse(JSON.stringify(originalToy.value));
    isEditing.value = false;
};

// Сохранение изменений
const saveToy = async () => {
    if (isSaving.value) return;

    isSaving.value = true;
    try {
        await api.updateToy(route.params.id, toy.value);
        originalToy.value = JSON.parse(JSON.stringify(toy.value));
        isEditing.value = false;
    } catch (error) {
        console.error('Ошибка сохранения:', error);
    } finally {
        isSaving.value = false;
    }
};

onMounted(() => {
    fetchToy();
});
</script>