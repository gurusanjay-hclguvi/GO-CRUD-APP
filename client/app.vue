<template>
  <div class="min-h-screen font-sans selection:bg-gray-200">
    <NavBar />
    <main class="pt-20 px-4 sm:px-6 lg:px-8 max-w-7xl mx-auto">
      <NuxtPage />
    </main>
  </div>
</template>

<script setup>
// Load user on app start if token exists
const { fetchUser } = useAuth()
await fetchUser()
</script>
