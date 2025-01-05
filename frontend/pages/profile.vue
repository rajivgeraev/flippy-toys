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
                        <div v-if="user.telegram_profile.photo_url"
                            class="w-20 h-20 rounded-full overflow-hidden bg-gray-100">
                            <img :src="user.telegram_profile.photo_url" :alt="user.telegram_profile.username"
                                class="w-full h-full object-cover" />
                        </div>
                        <div v-else
                            class="w-20 h-20 rounded-full bg-gray-200 flex items-center justify-center text-gray-600">
                            <User :size="40" />
                        </div>
                        <div>
                            <h1 class="text-2xl font-semibold m-0">
                                {{ user.telegram_profile.first_name }} {{ user.telegram_profile.last_name }}
                            </h1>
                            <p v-if="user.telegram_profile.username" class="text-gray-600 mt-1">
                                @{{ user.telegram_profile.username }}
                            </p>
                        </div>
                    </div>

                    <!-- Details -->
                    <div class="bg-white rounded-xl p-4 shadow-sm space-y-3">
                        <div class="flex items-center gap-3 py-3 border-b border-gray-100">
                            <MessageCircle :size="20" />
                            <span>Language: {{ user.telegram_profile.language_code }}</span>
                        </div>
                        <div class="flex items-center gap-3 py-3 border-b border-gray-100">
                            <ShieldCheck :size="20" />
                            <span>Access Level: {{ user.access_level }}</span>
                        </div>
                        <div v-if="user.telegram_profile.is_premium"
                            class="flex items-center gap-3 py-3 text-yellow-500">
                            <Star :size="20" />
                            <span>Premium User</span>
                        </div>
                    </div>
                </div>

                <!-- Остальной код без изменений -->
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