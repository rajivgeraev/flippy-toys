<template>
    <div class="p-4">
        <div v-if="loading" class="flex justify-center py-8">
            <Icon name="lucide:loader" size="32" class="animate-spin text-purple-600" />
        </div>

        <div v-else>
            <!-- Галерея фотографий -->
            <div class="mb-6">
                <div class="grid grid-cols-2 gap-2">
                    <div v-for="(photo, index) in toy.photos" :key="index"
                        class="aspect-square rounded-lg overflow-hidden">
                        <img :src="photo.url" class="w-full h-full object-cover" :alt="toy.title" />
                    </div>
                </div>
            </div>

            <!-- Информация об игрушке -->
            <div class="mb-6">
                <h1 class="text-xl font-bold mb-2">{{ toy.title }}</h1>
                <span :class="[
                    'inline-block px-2 py-1 rounded-full text-sm font-medium mb-4',
                    getConditionClass(toy.condition)
                ]">
                    {{ CONDITIONS.find(c => c.value === toy.condition)?.label }}
                </span>
                <p class="text-gray-600">{{ toy.description }}</p>
            </div>

            <!-- Информация о владельце -->
            <div class="bg-white rounded-lg p-4 shadow-sm">
                <div class="flex items-center gap-4">
                    <div v-if="owner?.telegram_profile.photo_url" class="w-12 h-12 rounded-full overflow-hidden">
                        <img :src="owner.telegram_profile.photo_url" class="w-full h-full object-cover"
                            :alt="owner.telegram_profile.first_name" />
                    </div>
                    <div>
                        <h2 class="font-medium">
                            {{ owner?.telegram_profile.first_name }}
                            {{ owner?.telegram_profile.last_name }}
                        </h2>
                        <p class="text-sm text-gray-500">
                            @{{ owner?.telegram_profile.username }}
                        </p>
                    </div>
                </div>
            </div>

            <!-- Кнопка связаться -->
            <div class="fixed bottom-0 left-0 right-0 p-4 bg-white border-t">
                <button class="w-full bg-purple-600 text-white py-3 rounded-lg font-medium" @click="contactOwner">
                    Связаться с владельцем
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { api } from '~/services/api';

const route = useRoute();
const loading = ref(true);
const toy = ref(null);
const owner = ref(null);

const CONDITIONS = [
    { value: 'new', label: 'Новое' },
    { value: 'like_new', label: 'Как новое' },
    { value: 'good', label: 'Хорошее' },
    { value: 'acceptable', label: 'Удовлетворительное' }
];

const getConditionClass = (condition) => {
    switch (condition) {
        case 'new': return 'bg-green-100 text-green-800';
        case 'like_new': return 'bg-blue-100 text-blue-800';
        case 'good': return 'bg-yellow-100 text-yellow-800';
        case 'acceptable': return 'bg-orange-100 text-orange-800';
        default: return 'bg-gray-100 text-gray-800';
    }
};

const fetchToyDetails = async () => {
    try {
        const toyData = await api.getToy(route.params.id);
        toy.value = toyData;
        // Получаем информацию о владельце
        const ownerData = await api.getUser(toyData.user_id);
        owner.value = ownerData;
    } catch (error) {
        console.error('Error fetching toy details:', error);
    } finally {
        loading.value = false;
    }
};

const contactOwner = () => {
    // В будущем здесь будет логика для связи с владельцем
    console.log('Contact owner:', owner.value);
};

onMounted(() => {
    fetchToyDetails();
});
</script>