<template>
  <div class="flex items-center justify-center min-h-[calc(100vh-5rem)]">
    <div class="glass-panel p-8 rounded-2xl w-full max-w-md relative overflow-hidden">
      <h2 class="text-3xl font-bold text-center mb-8 text-black">
        Welcome Back
      </h2>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Email Address</label>
          <input 
            v-model="email" 
            type="email" 
            required
            class="input-field"
            placeholder="you@example.com"
          />
        </div>

        <div>
           <label class="block text-sm font-medium text-gray-700 mb-2">Password</label>
           <input 
            v-model="password" 
            type="password" 
            required
            class="input-field"
            placeholder="••••••••"
          />
        </div>

        <div v-if="error" class="bg-red-50 text-red-600 px-4 py-3 rounded-lg text-sm border border-red-100">
          {{ error }}
        </div>

        <button 
          type="submit" 
          class="w-full btn-primary"
          :disabled="loading"
        >
          <span v-if="loading">Signing in...</span>
          <span v-else>Sign In</span>
        </button>
      </form>

      <p class="mt-6 text-center text-gray-600 text-sm">
        Don't have an account? 
        <NuxtLink to="/register" class="text-black hover:underline font-medium">Create one</NuxtLink>
      </p>
    </div>
  </div>
</template>

<script setup>
const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const { login } = useAuth()

const handleLogin = async () => {
  loading.value = true
  error.value = ''
  try {
    await login(email.value, password.value)
  } catch (e) {
    error.value = e.statusMessage || 'An error occurred during login'
  } finally {
    loading.value = false
  }
}
</script>
