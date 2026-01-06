<script setup>
import { useTodos } from '../composables/useTodos'
import { ref } from 'vue'
const { createTodo } = useTodos()
const emit = defineEmits(['created'])

const title = ref('')

const submit = async () => {
  if (!title.value.trim()) return

  await createTodo({ title: title.value })
  title.value = ''
  emit('created')
}
</script>

<template>
  <div class="flex gap-2 mb-4">
    <input
      v-model="title"
      placeholder="New todo"
      class="flex-1 border px-3 py-2 rounded"
    />
    <button
      @click="submit"
      class="bg-black text-white px-4 py-2 rounded"
    >
      Add
    </button>
  </div>
</template>
