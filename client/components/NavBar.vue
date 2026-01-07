<template>
  <nav class="fixed top-0 left-0 right-0 z-50 glass-panel">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <div class="flex items-center">
          <NuxtLink to="/" class="text-xl font-bold text-black tracking-tight">
            TodoApp
          </NuxtLink>
        </div>
        
        <div class="flex items-center gap-4">
          <template v-if="user">
            <span class="text-gray-600 text-sm hidden sm:block">Welcome, {{ user.name }}</span>
            <NuxtLink to="/profile" class="text-gray-600 hover:text-black transition-colors">
              Profile
            </NuxtLink>
            <button @click="logout" class="text-gray-600 hover:text-red-600 transition-colors">
              Logout
            </button>
          </template>
          <template v-else>
            <NuxtLink to="/login" class="text-gray-600 hover:text-black transition-colors">
              Login
            </NuxtLink>
            <NuxtLink to="/register" class="btn-primary py-1.5 px-3 text-sm">
              Register
            </NuxtLink>
          </template>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
const { user, logout } = useAuth()
</script>
