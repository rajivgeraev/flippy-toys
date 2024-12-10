<template>
    <div class="profile">
      <ClientOnly>
        <template v-if="isInitialized">
          <div v-if="user" class="profile-content">
            <div class="profile-header">
            <div class="avatar" v-if="user.photo_url">
              <img :src="user.photo_url" :alt="user.username" />
            </div>
            <div class="avatar placeholder" v-else>
              <User :size="40" />
            </div>
            <div class="user-info">
              <h1>{{ user.first_name }} {{ user.last_name }}</h1>
              <p class="username" v-if="user.username">@{{ user.username }}</p>
            </div>
          </div>
  
          <div class="details">
            <div class="detail-item">
              <MessageCircle :size="20" />
              <span>Language Code: {{ user.language_code }}</span>
            </div>
            <div class="detail-item">
              <ShieldCheck :size="20" />
              <span>Telegram ID: {{ user.id }}</span>
            </div>
            <div class="detail-item premium" v-if="user.is_premium">
              <Star :size="20" />
              <span>Premium User</span>
            </div>
          </div>
  
          <div class="stats">
            <div class="stat-item">
              <h3>0</h3>
              <p>Items Posted</p>
            </div>
            <div class="stat-item">
              <h3>0</h3>
              <p>Exchanges</p>
            </div>
            <div class="stat-item">
              <h3>0</h3>
              <p>Reviews</p>
            </div>
          </div>
        </div>
  
        <div v-else class="login-prompt">
          <User :size="64" class="login-icon" />
          <h2>Login Required</h2>
          <p>Please open this app through the Telegram bot to access your profile and start exchanging toys.</p>
          <div class="steps">
            <p>1. Open Telegram</p>
            <p>2. Find @FlippyBot</p>
            <p>3. Start the bot and open the web app</p>
          </div>
        </div>
      </template>
  
      <div v-else class="loading">
        <Loader2 :size="32" class="spin" />
        <p>Loading...</p>
      </div>
     </ClientOnly>
    </div>
  </template>
  
  <script setup>
  import { User, MessageCircle, ShieldCheck, Star, Loader2 } from 'lucide-vue-next';
  const { isInitialized, user } = process.client ? useTelegram() : { isInitialized: ref(false), user: ref(null) };
  </script>
  
  <style scoped>
  .profile {
    padding: 20px;
    max-width: 600px;
    margin: 0 auto;
  }
  
  .profile-header {
    display: flex;
    align-items: center;
    gap: 20px;
    margin-bottom: 24px;
  }
  
  .avatar {
    width: 80px;
    height: 80px;
    border-radius: 50%;
    overflow: hidden;
    background: #f0f0f0;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .avatar img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  
  .avatar.placeholder {
    background: #e0e0e0;
    color: #666;
  }
  
  .user-info h1 {
    font-size: 1.5rem;
    margin: 0;
  }
  
  .username {
    color: #666;
    margin-top: 4px;
  }
  
  .details {
    background: white;
    border-radius: 12px;
    padding: 16px;
    margin-bottom: 24px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }
  
  .detail-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 0;
    border-bottom: 1px solid #eee;
    color: #333;
  }
  
  .detail-item:last-child {
    border-bottom: none;
  }
  
  .detail-item.premium {
    color: #FFD700;
  }
  
  .stats {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 16px;
    margin-top: 24px;
  }
  
  .stat-item {
    background: white;
    padding: 16px;
    border-radius: 12px;
    text-align: center;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  }
  
  .stat-item h3 {
    font-size: 1.5rem;
    margin: 0;
    color: #007bff;
  }
  
  .stat-item p {
    margin: 4px 0 0;
    color: #666;
    font-size: 0.9rem;
  }
  
  .login-prompt {
    text-align: center;
    padding: 40px 20px;
  }
  
  .login-icon {
    color: #666;
    margin-bottom: 20px;
  }
  
  .login-prompt h2 {
    margin-bottom: 12px;
    color: #333;
  }
  
  .login-prompt p {
    color: #666;
    margin-bottom: 24px;
  }
  
  .steps {
    background: white;
    padding: 20px;
    border-radius: 12px;
    text-align: left;
  }
  
  .steps p {
    margin: 8px 0;
    color: #333;
  }
  
  .loading {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 300px;
    gap: 16px;
    color: #666;
  }
  
  .spin {
    animation: spin 1s linear infinite;
  }
  
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
  </style>