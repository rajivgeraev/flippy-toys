<template>
    <div class="bg-white rounded-lg shadow p-3">
        <img :src="`/images/toys/${toy.image}`" :alt="toy.title" class="w-full h-40 object-cover rounded-lg mb-2" />
        <h3 class="font-semibold text-gray-800">{{ toy.title }}</h3>
        <p class="text-sm text-gray-500">{{ toy.age }}</p>
        <div class="flex items-center justify-between mt-2">
            <span class="text-xs bg-green-100 text-green-600 px-2 py-1 rounded">
                {{ toy.condition }}
            </span>
            <span class="text-xs text-gray-500">{{ toy.location }}</span>
        </div>
    </div>
</template>

<script setup>
defineProps({
    toy: {
        type: Object,
        required: true
    }
})
</script>