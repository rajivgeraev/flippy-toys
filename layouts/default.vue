<template>
    <div class="app-layout">
      <header class="header">
        <nav class="nav">
          <NuxtLink to="/" class="nav-link">Home</NuxtLink>
          <NuxtLink to="/browse" class="nav-link">Browse</NuxtLink>
          <NuxtLink to="/post" class="nav-link">Post Item</NuxtLink>
          <NuxtLink to="/profile" class="nav-link">Profile</NuxtLink>
        </nav>
      </header>
      
      <main class="main">
        <slot />
      </main>
  
      <footer class="footer">
        <p>© 2024 Flippy - Toy Exchange Platform</p>
      </footer>
    </div>
  </template>
  
  <style scoped>
  .app-layout {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }
  
  .header {
    padding: 1rem;
    background-color: #f8f9fa;
    border-bottom: 1px solid #eee;
  }
  
  .nav {
    display: flex;
    gap: 1rem;
    justify-content: center;
  }
  
  .nav-link {
    text-decoration: none;
    color: #333;
    padding: 0.5rem 1rem;
  }
  
  .nav-link:hover {
    color: #007bff;
  }
  
  .main {
    flex: 1;
    padding: 2rem;
  }
  
  .footer {
    padding: 1rem;
    background-color: #f8f9fa;
    text-align: center;
    border-top: 1px solid #eee;
  }
  </style>