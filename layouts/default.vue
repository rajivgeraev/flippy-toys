<template>
    <div class="app-layout">
      <main class="main">
        <slot />
      </main>
  
      <nav class="bottom-nav">
        <NuxtLink to="/" class="nav-item">
          <div class="nav-icon">
            <Home :size="24" />
          </div>
          <span class="nav-label">Home</span>
        </NuxtLink>
        <NuxtLink to="/browse" class="nav-item">
          <div class="nav-icon">
            <Search :size="24" />
          </div>
          <span class="nav-label">Browse</span>
        </NuxtLink>
        <NuxtLink to="/post" class="nav-item">
          <div class="nav-icon">
            <PlusSquare :size="24" />
          </div>
          <span class="nav-label">Post</span>
        </NuxtLink>
        <NuxtLink to="/profile" class="nav-item">
          <div class="nav-icon">
            <User :size="24" />
          </div>
          <span class="nav-label">Profile</span>
        </NuxtLink>
      </nav>
    </div>
  </template>
  
  <script setup>
  import { Home, Search, PlusSquare, User } from 'lucide-vue-next';
  </script>
  
  <style scoped>
  .app-layout {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    padding-bottom: 60px;
  }
  
  .main {
    flex: 1;
    padding: 1rem;
  }
  
  .bottom-nav {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    height: 60px;
    background-color: white;
    display: flex;
    justify-content: space-around;
    align-items: center;
    box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
  }
  
  .nav-item {
    text-decoration: none;
    color: #666;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 8px;
    width: 25%;
  }
  
  .nav-icon {
    margin-bottom: 4px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .nav-label {
    font-size: 0.75rem;
  }
  
  /* Active link styling */
  .nav-item.router-link-active {
    color: #007bff;
  }
  
  .router-link-active .nav-icon {
    color: #007bff;
  }
  
  /* Optional: Add a safe area padding for iOS devices */
  @supports (padding: max(0px)) {
    .bottom-nav {
      padding-bottom: max(0px, env(safe-area-inset-bottom));
      height: calc(60px + env(safe-area-inset-bottom, 0px));
    }
  }
  </style>