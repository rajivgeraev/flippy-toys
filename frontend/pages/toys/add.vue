<template>
    <div class="p-4">
        <div class="flex items-center mb-6">
            <button @click="router.back()" class="mr-4">
                <Icon name="lucide:arrow-left" size="24" />
            </button>
            <h1 class="text-xl font-bold">Добавить игрушку</h1>
        </div>

        <Form @submit="onSubmit" v-slot="{ errors }">
            <!-- Photo Upload -->
            <div class="mb-6">
                <label class="block text-sm font-medium text-gray-700 mb-2">
                    Фотографии
                </label>
                <div class="grid grid-cols-3 gap-2 mb-2">
                    <div v-for="(preview, index) in previews" :key="index" class="relative aspect-square">
                        <img :src="preview" class="w-full h-full object-cover rounded-lg" />
                        <button @click="removePhoto(index)" type="button"
                            class="absolute top-1 right-1 bg-red-500 text-white rounded-full p-1">
                            <Icon name="lucide:x" size="16" />
                        </button>
                    </div>
                    <label v-if="previews.length < 5"
                        class="aspect-square border-2 border-dashed border-gray-300 rounded-lg flex items-center justify-center cursor-pointer hover:border-purple-500">
                        <input type="file" accept="image/*" class="hidden" @change="onPhotoSelected" multiple />
                        <Icon name="lucide:plus" size="24" class="text-gray-400" />
                    </label>
                </div>
                <p v-if="errors.photos" class="mt-1 text-sm text-red-500">
                    {{ errors.photos }}
                </p>
            </div>

            <!-- Title -->
            <div class="mb-6">
                <Field name="title" v-slot="{ field, errors }">
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                        Название
                    </label>
                    <input type="text" v-bind="field" class="w-full px-3 py-2 border rounded-lg"
                        :class="{ 'border-red-500': errors }" />
                    <p v-if="errors" class="mt-1 text-sm text-red-500">
                        {{ errors }}
                    </p>
                </Field>
            </div>

            <!-- Description -->
            <div class="mb-6">
                <Field name="description" v-slot="{ field, errors }">
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                        Описание
                    </label>
                    <textarea v-bind="field" class="w-full px-3 py-2 border rounded-lg"
                        :class="{ 'border-red-500': errors }" rows="4"></textarea>
                    <p v-if="errors" class="mt-1 text-sm text-red-500">
                        {{ errors }}
                    </p>
                </Field>
            </div>

            <!-- Condition -->
            <div class="mb-6">
                <Field name="condition" v-slot="{ field, errors }">
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                        Состояние
                    </label>
                    <select v-bind="field" class="w-full px-3 py-2 border rounded-lg"
                        :class="{ 'border-red-500': errors }">
                        <option value="">Выберите состояние</option>
                        <option value="new">Новое</option>
                        <option value="like_new">Как новое</option>
                        <option value="good">Хорошее</option>
                        <option value="acceptable">Приемлемое</option>
                    </select>
                    <p v-if="errors" class="mt-1 text-sm text-red-500">
                        {{ errors }}
                    </p>
                </Field>
            </div>

            <!-- Category -->
            <div class="mb-6">
                <Field name="category" v-slot="{ field, errors }">
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                        Категория
                    </label>
                    <select v-bind="field" class="w-full px-3 py-2 border rounded-lg"
                        :class="{ 'border-red-500': errors }">
                        <option value="">Выберите категорию</option>
                        <option value="construction_toys">Конструктор</option>
                        <option value="dolls">Куклы</option>
                        <option value="vehicles">Машинки</option>
                        <option value="educational">Развивающие</option>
                        <option value="outdoor">Для улицы</option>
                        <option value="board_games">Настольные игры</option>
                        <option value="electronic">Электронные</option>
                        <option value="stuffed_animals">Мягкие игрушки</option>
                        <option value="action_figures">Фигурки</option>
                        <option value="arts_crafts">Творчество</option>
                        <option value="musical">Музыкальные</option>
                        <option value="other">Другое</option>
                    </select>
                    <p v-if="errors" class="mt-1 text-sm text-red-500">
                        {{ errors }}
                    </p>
                </Field>
            </div>

            <button type="submit" :disabled="loading"
                class="w-full bg-purple-600 text-white py-3 rounded-lg font-medium disabled:bg-purple-400">
                {{ loading ? 'Загрузка...' : 'Добавить игрушку' }}
            </button>
        </Form>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { Form, Field } from 'vee-validate';
import { useRouter } from 'vue-router';
import * as yup from 'yup';
import { api } from '~/services/api';
import { uploadService } from '~/services/upload';

const router = useRouter();
const loading = ref(false);
const previews = ref([]);
const files = ref([]);

const schema = yup.object({
    title: yup.string().required('Введите название'),
    description: yup.string().required('Добавьте описание'),
    condition: yup.string().required('Выберите состояние'),
    category: yup.string().required('Выберите категорию')
});

const onPhotoSelected = (event) => {
    const selectedFiles = Array.from(event.target.files);
    selectedFiles.forEach(file => {
        if (files.value.length >= 5) return;

        const reader = new FileReader();
        reader.onload = (e) => {
            previews.value.push(e.target.result);
            files.value.push(file);
        };
        reader.readAsDataURL(file);
    });
};

const removePhoto = (index) => {
    previews.value.splice(index, 1);
    files.value.splice(index, 1);
};

const onSubmit = async (values) => {
    try {
        loading.value = true;

        if (files.value.length === 0) {
            throw new Error('Добавьте хотя бы одну фотографию');
        }

        // Загружаем фотографии
        const uploadedPhotos = await uploadService.uploadMultiple(files.value);

        // Создаем игрушку
        const toyData = {
            title: values.title,
            description: values.description,
            condition: values.condition,
            category: values.category,
            photos: uploadedPhotos
        };

        await api.createToy(toyData);
        router.push('/my-toys');
    } catch (error) {
        console.error('Error creating toy:', error);
        // Здесь можно добавить уведомление об ошибке
    } finally {
        loading.value = false;
    }
};
</script>