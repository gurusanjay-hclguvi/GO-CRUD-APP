<template>
  <div class="flex items-center justify-center min-h-[calc(100vh-5rem)]">
    <div class="glass-panel p-8 rounded-2xl w-full max-w-md relative overflow-hidden">
      <h2 class="text-3xl font-bold text-center mb-8 text-black">
        Create Account
      </h2>

      <form @submit.prevent="handleRegister" class="space-y-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Full Name</label>
          <input 
            v-model="name" 
            type="text" 
            required
            class="input-field"
            placeholder="John Doe"
          />
        </div>

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
          <span v-if="loading">Creating account...</span>
          <span v-else>Register</span>
        </button>
      </form>

      <p class="mt-6 text-center text-gray-600 text-sm">
        Already have an account? 
        <NuxtLink to="/login" class="text-black hover:underline font-medium">Sign in</NuxtLink>
      </p>
    </div>
  </div>
</template>

<script setup>
const name = ref('')
const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const { register } = useAuth()

const handleRegister = async () => {
  loading.value = true
  error.value = ''
  try {
    await register(name.value, email.value, password.value)
  } catch (e) {
    error.value = e.statusMessage || 'An error occurred during registration'
  } finally {
    loading.value = false
  }
}
</script>
