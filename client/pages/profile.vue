<script setup lang="ts">
definePageMeta({
  middleware: 'auth',
  ssr: false
})

const { user, fetchProfile, logout } = useAuth()

onMounted(fetchProfile)
</script>

<template>
  <div class="max-w-xl mx-auto p-6">
    <h1 class="text-2xl font-bold mb-4">Profile</h1>

    <pre>{{ user }}</pre>

    <button @click="logout" class="btn mt-4">
      Logout
    </button>
  </div>
</template>
