<script setup lang="ts">
import { useTodos } from '../composables/useTodos'
const props = defineProps<{ todo: any }>()
const emit = defineEmits(['updated'])

const { updateTodo, deleteTodo } = useTodos()

const toggle = async () => {
  await updateTodo(props.todo.id, {
    ...props.todo,
    completed: !props.todo.completed
  })
  emit('updated')
}

const remove = async () => {
  await deleteTodo(props.todo.id)
  emit('updated')
}
</script>

<template>
  <div class="flex gap-2">
    <button
      @click="toggle"
      class="text-sm px-2 py-1 border rounded"
    >
      {{ todo.completed ? 'Undo' : 'Done' }}
    </button>

    <button
      @click="remove"
      class="text-sm px-2 py-1 text-red-600 border border-red-300 rounded"
    >
      Delete
    </button>
  </div>
</template>
