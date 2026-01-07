<script setup lang="ts">
const email = ref('')
const password = ref('')
const success = ref(false)

const { register } = useAuth()

const submit = async () => {
  await register(email.value, password.value)
  success.value = true
}
</script>

<template>
  <div class="max-w-sm mx-auto mt-20">
    <h1 class="text-2xl font-bold mb-4">Register</h1>

    <input v-model="email" placeholder="Email" class="input" />
    <input v-model="password" type="password" placeholder="Password" class="input" />

    <button @click="submit" class="btn">Register</button>

    <p v-if="success" class="text-green-600 mt-2">
      Account created. Go to login.
    </p>
  </div>
</template>
