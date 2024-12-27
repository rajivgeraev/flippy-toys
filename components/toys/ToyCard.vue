<template>
    <div class="toy-card">
        <img :src="`/images/toys/${toy.image}`" :alt="toy.title" class="toy-image">
        <div class="toy-content">
            <h3>{{ toy.title }}</h3>
            <p>{{ toy.description }}</p>
            <div class="toy-info">
                <span>{{ toy.age }}</span>
                <span>{{ toy.location }}</span>
            </div>
        </div>
    </div>
</template>

<script setup>
defineProps(['toy']);
</script>

<style scoped>
.toy-card {
    background: white;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.toy-image {
    width: 100%;
    height: 200px;
    object-fit: cover;
}

.toy-content {
    padding: 16px;
}

.toy-info {
    display: flex;
    justify-content: space-between;
    margin-top: 8px;
    color: #666;
}
</style>