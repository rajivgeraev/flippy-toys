<template>
    <div class="p-4">
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-xl font-bold">Мои игрушки</h1>
            <button @click="navigateToAdd"
                class="bg-purple-600 text-white px-4 py-2 rounded-full flex items-center gap-2">
                <Icon name="lucide:plus" size="20" />
                <span>Добавить</span>
            </button>
        </div>

        <div v-if="loading" class="flex justify-center py-8">
            <Icon name="lucide:loader" size="32" class="animate-spin text-purple-600" />
        </div>

        <div v-else-if="toys.length === 0" class="text-center py-8">
            <Icon name="lucide:package" size="48" class="text-gray-400 mx-auto mb-4" />
            <p class="text-gray-500">У вас пока нет добавленных игрушек</p>
            <button @click="navigateToAdd" class="mt-4 text-purple-600 font-medium">
                Добавить первую игрушку
            </button>
        </div>

        <div v-else class="grid grid-cols-2 gap-4">
            <div v-for="toy in toys" :key="toy.id" class="bg-white rounded-lg shadow overflow-hidden">
                <img :src="toy.photos && toy.photos.length > 0 ? toy.photos[0].url : '/placeholder.jpg'"
                    class="w-full h-32 object-cover" />
                <div class="p-3">
                    <h3 class="font-medium text-sm mb-1">{{ toy.title }}</h3>
                    <p class="text-gray-500 text-xs">
                        {{ toy.age_range?.min ?? 0 }}-{{ toy.age_range?.max ?? '+' }} лет
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { api } from '~/services/api';
import { telegram } from '~/composables/useTelegram';

const router = useRouter();
const toys = ref([]);
const loading = ref(true);

const fetchToys = async () => {
    try {
        // Ждем инициализации
        if (!telegram.isInitialized.value) {
            await new Promise(resolve => {
                const unwatch = watch(telegram.isInitialized, (newValue) => {
                    if (newValue) {
                        unwatch();
                        resolve(true);
                    }
                });
            });
        }

        toys.value = await api.getUserToys();
    } catch (error) {
        console.error('Failed to fetch toys:', error);
    } finally {
        loading.value = false;
    }
};

const navigateToAdd = () => {
    router.push('/toys/add');
};

onMounted(() => {
    fetchToys();
});
</script>