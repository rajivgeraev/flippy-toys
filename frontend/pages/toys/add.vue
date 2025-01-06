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
                        <button @click="removePhoto(index)"
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

            <!-- Age Range -->
            <div class="mb-6">
                <label class="block text-sm font-medium text-gray-700 mb-2">
                    Возраст
                </label>
                <div class="flex gap-4">
                    <Field name="ageMin" v-slot="{ field, errors }">
                        <div class="flex-1">
                            <input type="number" v-bind="field" placeholder="От"
                                class="w-full px-3 py-2 border rounded-lg" :class="{ 'border-red-500': errors }" />
                            <p v-if="errors" class="mt-1 text-sm text-red-500">
                                {{ errors }}
                            </p>
                        </div>
                    </Field>
                    <Field name="ageMax" v-slot="{ field, errors }">
                        <div class="flex-1">
                            <input type="number" v-bind="field" placeholder="До"
                                class="w-full px-3 py-2 border rounded-lg" :class="{ 'border-red-500': errors }" />
                            <p v-if="errors" class="mt-1 text-sm text-red-500">
                                {{ errors }}
                            </p>
                        </div>
                    </Field>
                </div>
            </div>

            <!-- Condition -->
            <div class="mb-6">
                <Field name="condition" v-slot="{ field, errors }">
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                        Состояние
                    </label>
                    <select v-bind="field" class="w-full px-3 py-2 border rounded-lg"
                        :class="{ 'border-red-500': errors }">
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

            <button type="submit" class="w-full bg-purple-600 text-white py-3 rounded-lg font-medium">
                Добавить игрушку
            </button>
        </Form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { Form, Field } from 'vee-validate'
import { useRouter } from 'vue-router'
import { api } from '~/services/api'
import * as yup from 'yup'

const router = useRouter()
const previews = ref([])
const photos = ref([])

const schema = yup.object({
    title: yup.string().required('Введите название'),
    description: yup.string().required('Добавьте описание'),
    ageMin: yup.number()
        .required('Укажите минимальный возраст')
        .min(0, 'Минимальный возраст не может быть отрицательным'),
    ageMax: yup.number()
        .required('Укажите максимальный возраст')
        .min(yup.ref('ageMin'), 'Максимальный возраст должен быть больше минимального'),
    condition: yup.string().required('Выберите состояние'),
    category: yup.string().required('Выберите категорию'),
    photos: yup.array()
        .min(1, 'Добавьте хотя бы одно фото')
        .max(5, 'Максимум 5 фотографий')
})

const onPhotoSelected = (event) => {
    const files = Array.from(event.target.files)
    files.forEach(file => {
        if (previews.value.length >= 5) return

        const reader = new FileReader()
        reader.onload = (e) => {
            previews.value.push(e.target.result)
            photos.value.push(file)
        }
        reader.readAsDataURL(file)
    })
}

const removePhoto = (index) => {
    previews.value.splice(index, 1)
    photos.value.splice(index, 1)
}

const onSubmit = async (values) => {
    try {
        await api.createToy({
            title: values.title,
            description: values.description,
            age_min: values.ageMin,
            age_max: values.ageMax,
            condition: values.condition,
            category: values.category,
            photos: photos.value
        });
        router.push('/my-toys');
    } catch (error) {
        console.error('Error creating toy:', error);
    }
}

</script>