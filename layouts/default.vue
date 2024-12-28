<template>
    <div class="flex flex-col h-screen bg-gray-50">
        <!-- Header -->
        <header class="px-4 py-3 bg-white shadow">
            <div class="flex items-center justify-between">
                <h1 class="text-xl font-bold text-purple-600">Flippy Toys</h1>
                <div class="flex items-center space-x-2">
                    <div class="relative w-64">
                        <input type="text" placeholder="Поиск игрушек..."
                            class="w-full px-4 py-2 border rounded-full focus:outline-none focus:ring-2 focus:ring-purple-300" />
                        <Icon name="lucide:search" class="absolute right-3 top-2.5 text-gray-400" size="20" />
                    </div>
                </div>
            </div>
        </header>

        <!-- Main Content -->
        <main class="flex-1 overflow-y-auto">
            <slot />
        </main>

        <!-- Bottom Navigation -->
        <nav class="bg-white border-t">
            <div class="flex justify-around py-3">
                <NuxtLink to="/" class="flex flex-col items-center"
                    :class="route.path === '/' ? 'text-purple-600' : 'text-gray-500'">
                    <Icon name="lucide:home" size="24" />
                    <span class="text-xs mt-1">Главная</span>
                </NuxtLink>
                <NuxtLink to="/favorites" class="flex flex-col items-center"
                    :class="route.path === '/favorites' ? 'text-purple-600' : 'text-gray-500'">
                    <Icon name="lucide:heart" size="24" />
                    <span class="text-xs mt-1">Избранное</span>
                </NuxtLink>
                <NuxtLink to="/messages" class="flex flex-col items-center"
                    :class="route.path === '/messages' ? 'text-purple-600' : 'text-gray-500'">
                    <Icon name="lucide:message-circle" size="24" />
                    <span class="text-xs mt-1">Сообщения</span>
                </NuxtLink>
                <NuxtLink to="/profile" class="flex flex-col items-center"
                    :class="route.path === '/profile' ? 'text-purple-600' : 'text-gray-500'">
                    <Icon name="lucide:user" size="24" />
                    <span class="text-xs mt-1">Профиль</span>
                </NuxtLink>
            </div>
        </nav>
    </div>
</template>

<script setup>
const route = useRoute()
</script>