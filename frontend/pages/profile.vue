<template>
    <div class="max-w-2xl mx-auto p-5">
        <ClientOnly>
            <div v-if="initError" class="flex items-center gap-2 p-4 bg-red-100 text-red-600 rounded-lg mx-4">
                <AlertCircle :size="24" />
                <p>{{ initError }}</p>
            </div>

            <template v-else>
                <div v-if="isInitialized && user" class="space-y-6">
                    <!-- Profile Header -->
                    <div class="flex items-center gap-5">
                        <div v-if="user.photo_url" class="w-20 h-20 rounded-full overflow-hidden bg-gray-100">
                            <img :src="user.photo_url" :alt="user.username" class="w-full h-full object-cover" />
                        </div>
                        <div v-else
                            class="w-20 h-20 rounded-full bg-gray-200 flex items-center justify-center text-gray-600">
                            <User :size="40" />
                        </div>
                        <div>
                            <h1 class="text-2xl font-semibold m-0">{{ user.first_name }} {{ user.last_name }}</h1>
                            <p v-if="user.username" class="text-gray-600 mt-1">@{{ user.username }}</p>
                        </div>
                    </div>

                    <!-- Details -->
                    <div class="bg-white rounded-xl p-4 shadow-sm space-y-3">
                        <div class="flex items-center gap-3 py-3 border-b border-gray-100">
                            <MessageCircle :size="20" />
                            <span>Language Code: {{ user.language_code }}</span>
                        </div>
                        <div class="flex items-center gap-3 py-3 border-b border-gray-100">
                            <ShieldCheck :size="20" />
                            <span>Telegram ID: {{ user.id }}</span>
                        </div>
                        <div v-if="user.is_premium" class="flex items-center gap-3 py-3 text-yellow-500">
                            <Star :size="20" />
                            <span>Premium User</span>
                        </div>
                    </div>

                    <!-- Stats -->
                    <div class="grid grid-cols-3 gap-4">
                        <div class="bg-white p-4 rounded-xl text-center shadow-sm">
                            <h3 class="text-2xl text-blue-500 m-0">0</h3>
                            <p class="text-gray-600 text-sm mt-1">Items Posted</p>
                        </div>
                        <div class="bg-white p-4 rounded-xl text-center shadow-sm">
                            <h3 class="text-2xl text-blue-500 m-0">0</h3>
                            <p class="text-gray-600 text-sm mt-1">Exchanges</p>
                        </div>
                        <div class="bg-white p-4 rounded-xl text-center shadow-sm">
                            <h3 class="text-2xl text-blue-500 m-0">0</h3>
                            <p class="text-gray-600 text-sm mt-1">Reviews</p>
                        </div>
                    </div>
                </div>

                <!-- Login Prompt -->
                <div v-else-if="isInitialized && !user" class="text-center py-10 px-5">
                    <User :size="64" class="text-gray-600 mb-5 mx-auto" />
                    <h2 class="text-xl text-gray-800 mb-3">Login Required</h2>
                    <p class="text-gray-600 mb-6">
                        Please open this app through the Telegram bot to access your profile
                        and start exchanging toys.
                    </p>
                    <div class="bg-white p-5 rounded-xl text-left">
                        <p class="text-gray-800 my-2">1. Open Telegram</p>
                        <p class="text-gray-800 my-2">2. Find @FlippyBot</p>
                        <p class="text-gray-800 my-2">3. Start the bot and open the web app</p>
                    </div>
                </div>

                <!-- Loading State -->
                <div v-else class="flex flex-col items-center justify-center min-h-[300px] gap-4 text-gray-600">
                    <Loader2 :size="32" class="animate-spin" />
                    <p>Loading...</p>
                    <small class="text-xs text-gray-600 mt-2">
                        Init: {{ isInitialized ? "Yes" : "No" }}, User:
                        {{ user ? "Yes" : "No" }}
                    </small>
                </div>
            </template>
        </ClientOnly>
    </div>
</template>

<script setup>
import {
    User,
    MessageCircle,
    ShieldCheck,
    Star,
    Loader2,
    AlertCircle,
} from "lucide-vue-next";

import { telegram } from '~/composables/useTelegram';
const { isInitialized, user, initError } = telegram;

watch([isInitialized, user], ([newInit, newUser]) => {
    console.log("State changed:", {
        isInitialized: newInit,
        hasUser: !!newUser,
        userData: newUser,
    });
});
</script>