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
                        <option value="new">Новое</option>
                        <option value="like_new">Как новое</option>
                        <option value="good">Хорошее</option>
                        <option value="acceptable">Удовлетворительное</option>
                    </select>
                </div>

                <div class="mb-4">
                    <label class="block text-sm font-medium text-gray-700">Категория</label>
                    <select v-model="toy.category" :disabled="!isEditing" class="mt-1 block w-full p-2 border rounded">
                        <option value="construction_toys">Конструкторы</option>
                        <option value="dolls">Куклы</option>
                        <option value="vehicles">Транспорт</option>
                        <option value="educational">Обучающие</option>
                        <option value="outdoor">Уличные</option>
                    </select>
                </div>

                <!-- Кнопка сохранения -->
                <div v-if="isEditing" class="flex justify-end">
                    <button type="submit" class="bg-purple-600 text-white px-4 py-2 rounded-full">Сохранить</button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { api } from '~/services/api';

const route = useRoute();
const router = useRouter();
const toy = ref({});
const loading = ref(true);
const isEditing = ref(false);

// Загрузка данных игрушки
const fetchToy = async () => {
    try {
        const response = await api.getToy(route.params.id);
        toy.value = response;
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

// Сохранение изменений
const saveToy = async () => {
    try {
        console.log('Отправка данных для обновления:', toy.value);
        await api.updateToy(route.params.id, toy.value);
        isEditing.value = false;
    } catch (error) {
        console.error('Ошибка сохранения:', error);
    }
};

onMounted(() => {
    fetchToy();
});
</script>
