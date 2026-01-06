<script setup lang="ts">
import type { Todo } from '~/types/todo'

const { getTodos } = useTodos()

const todos = ref<Todo[]>([])
const loading = ref(false)

const fetchTodos = async () => {
  loading.value = true
  todos.value = await getTodos()
  loading.value = false
}

onMounted(fetchTodos)
</script>


<template>
  <div class="max-w-2xl mx-auto p-6">
    <h1 class="text-2xl font-bold mb-4">Todos</h1>

    <TodoCreate @created="fetchTodos" />

    <p v-if="loading" class="text-gray-500">Loading...</p>

    <ul v-else class="space-y-2">
      <li
        v-for="todo in todos"
        :key="todo.id"
        class="flex justify-between items-center bg-white p-3 rounded shadow"
      >
        <span>{{ todo.title }}</span>

        <TodoActions
          :todo="todo"
          @updated="fetchTodos"
        />
      </li>
    </ul>

  </div>
</template>
