<script setup lang="ts">
const email = ref('')
const password = ref('')
const error = ref('')

const { login } = useAuth()

const submit = async () => {
  try {
    await login(email.value, password.value)
    navigateTo('/')
  } catch {
    error.value = 'Invalid credentials'
  }
}
</script>

<template>
  <div class="max-w-sm mx-auto mt-20">
    <h1 class="text-2xl font-bold mb-4">Login</h1>

    <input v-model="email" placeholder="Email" class="input" />
    <input v-model="password" type="password" placeholder="Password" class="input" />

    <button @click="submit" class="btn">Login</button>

    <p v-if="error" class="text-red-500 mt-2">{{ error }}</p>
  </div>
</template>
