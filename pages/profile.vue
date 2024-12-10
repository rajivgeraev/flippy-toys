<template>
    <div class="profile">
      <div v-if="isAvailable" class="profile-container">
        <!-- User Info Section -->
        <div class="user-info">
          <img 
            v-if="userPhotoUrl" 
            :src="userPhotoUrl" 
            alt="Profile Photo"
            class="profile-photo"
          />
          <div v-else class="profile-photo-placeholder">
            <User :size="40" />
          </div>
          
          <h1 class="username">{{ userName }}</h1>
          
          <div class="info-grid">
            <!-- User ID -->
            <div class="info-item">
              <div class="info-label">
                <IdCard :size="16" />
                <span>User ID</span>
              </div>
              <div class="info-value">{{ user?.id || 'N/A' }}</div>
            </div>
  
            <!-- Username -->
            <div class="info-item">
              <div class="info-label">
                <AtSign :size="16" />
                <span>Username</span>
              </div>
              <div class="info-value">@{{ user?.username || 'N/A' }}</div>
            </div>
  
            <!-- Language -->
            <div class="info-item">
              <div class="info-label">
                <Globe :size="16" />
                <span>Language</span>
              </div>
              <div class="info-value">{{ user?.language_code?.toUpperCase() || 'N/A' }}</div>
            </div>
  
            <!-- Platform -->
            <div class="info-item">
              <div class="info-label">
                <Monitor :size="16" />
                <span>Platform</span>
              </div>
              <div class="info-value">{{ webApp?.platform || 'N/A' }}</div>
            </div>
          </div>
        </div>
  
        <!-- Color Theme -->
        <div class="theme-info">
          <div class="info-label">
            <Palette :size="16" />
            <span>Theme</span>
          </div>
          <div class="color-display">
            <div class="color-item">
              <div class="color-preview" :style="{ backgroundColor: webApp?.themeParams?.bg_color }"></div>
              <span>Background</span>
            </div>
            <div class="color-item">
              <div class="color-preview" :style="{ backgroundColor: webApp?.themeParams?.text_color }"></div>
              <span>Text</span>
            </div>
            <div class="color-item">
              <div class="color-preview" :style="{ backgroundColor: webApp?.themeParams?.hint_color }"></div>
              <span>Hint</span>
            </div>
            <div class="color-item">
              <div class="color-preview" :style="{ backgroundColor: webApp?.themeParams?.link_color }"></div>
              <span>Link</span>
            </div>
          </div>
        </div>
      </div>
  
      <div v-else class="not-available">
        <AlertTriangle :size="32" class="warning-icon" />
        <h2>Telegram WebApp is not available</h2>
        <p>This app must be opened from Telegram to access user data.</p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { computed, onMounted } from 'vue';
  import { User, IdCard, AtSign, Globe, Monitor, Palette, AlertTriangle } from 'lucide-vue-next';
  
  const { user, webApp, isAvailable } = useTelegram();
  
  const userName = computed(() => {
    if (!user) return 'Unknown User';
    return [user.first_name, user.last_name].filter(Boolean).join(' ');
  });
  
  const userPhotoUrl = computed(() => {
    return user?.photo_url || null;
  });
  </script>
  
  <style scoped>
  .profile {
    padding: 20px;
    max-width: 600px;
    margin: 0 auto;
  }
  
  .profile-container {
    background: white;
    border-radius: 12px;
    padding: 20px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }
  
  .user-info {
    text-align: center;
    margin-bottom: 24px;
  }
  
  .profile-photo {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    margin-bottom: 16px;
    object-fit: cover;
  }
  
  .profile-photo-placeholder {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    background: #f0f0f0;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 0 auto 16px;
    color: #666;
  }
  
  .username {
    font-size: 1.5rem;
    margin-bottom: 20px;
    font-weight: 600;
  }
  
  .info-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 16px;
    text-align: left;
  }
  
  .info-item {
    padding: 12px;
    background: #f8f9fa;
    border-radius: 8px;
  }
  
  .info-label {
    display: flex;
    align-items: center;
    gap: 8px;
    color: #666;
    margin-bottom: 4px;
    font-size: 0.9rem;
  }
  
  .info-value {
    font-weight: 500;
    font-size: 1rem;
  }
  
  .theme-info {
    margin-top: 24px;
    padding-top: 24px;
    border-top: 1px solid #eee;
  }
  
  .color-display {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(80px, 1fr));
    gap: 12px;
    margin-top: 12px;
  }
  
  .color-item {
    text-align: center;
  }
  
  .color-preview {
    width: 40px;
    height: 40px;
    border-radius: 8px;
    margin: 0 auto 8px;
    border: 1px solid #eee;
  }
  
  .not-available {
    text-align: center;
    padding: 40px 20px;
    background: white;
    border-radius: 12px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }
  
  .warning-icon {
    color: #ff9800;
    margin-bottom: 16px;
  }
  
  .not-available h2 {
    margin-bottom: 8px;
    font-size: 1.2rem;
  }
  
  .not-available p {
    color: #666;
  }
  </style>