<template>
  <div class="pb-20">
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">My Tasks</h1>
        <p class="text-gray-500 mt-1">Manage your daily goals</p>
      </div>
      <button 
        @click="isModalOpen = true"
        class="btn-primary flex items-center gap-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
        </svg>
        Add Task
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="pending" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-black"></div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-50 text-red-600 p-4 rounded-xl text-center border border-red-100">
      Failed to load todos. Please try again.
    </div>

    <div v-else class="space-y-8">
      <!-- Ongoing Tasks -->
      <section>
        <h2 class="text-xl font-semibold text-gray-800 mb-4 flex items-center gap-2">
          <span class="w-2 h-2 rounded-full bg-black"></span>
          Ongoing
          <span class="text-sm font-normal text-gray-500 ml-2">({{ ongoingTodos.length }})</span>
        </h2>
        
        <div v-if="ongoingTodos.length === 0" class="glass-panel p-8 rounded-xl text-center text-gray-500">
          No ongoing tasks. Add one to get started!
        </div>

        <div v-else class="grid gap-4">
          <div 
            v-for="todo in ongoingTodos" 
            :key="todo.id"
            class="glass-panel p-4 rounded-xl flex items-center justify-between group hover:border-gray-400 transition-colors"
          >
            <div class="flex items-center gap-4">
              <button 
                @click="toggleComplete(todo)"
                class="w-6 h-6 rounded-full border-2 border-gray-400 hover:border-black hover:bg-gray-100 transition-all flex items-center justify-center"
                title="Mark as complete"
              >
              </button>
              <span class="text-gray-900">{{ todo.title }}</span>
            </div>
            
            <button 
              @click="deleteTodo(todo.id)"
              class="text-gray-400 hover:text-red-600 opacity-0 group-hover:opacity-100 transition-all p-2"
              title="Delete task"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </div>
      </section>

      <!-- Completed Tasks -->
      <section v-if="completedTodos.length > 0">
        <h2 class="text-xl font-semibold text-gray-400 mb-4 flex items-center gap-2">
          <span class="w-2 h-2 rounded-full bg-gray-300"></span>
          Completed
          <span class="text-sm font-normal text-gray-400 ml-2">({{ completedTodos.length }})</span>
        </h2>
        
        <div class="grid gap-4">
          <div 
            v-for="todo in completedTodos" 
            :key="todo.id"
            class="glass-panel p-4 rounded-xl flex items-center justify-between bg-gray-50 border-gray-100"
          >
            <div class="flex items-center gap-4">
              <button 
                @click="toggleComplete(todo)"
                class="w-6 h-6 rounded-full border-2 border-gray-400 bg-gray-400 flex items-center justify-center text-white hover:bg-gray-500 hover:border-gray-500 transition-colors"
                title="Mark as incomplete"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
              </button>
              <span class="text-gray-400 line-through">{{ todo.title }}</span>
            </div>
            
             <button 
              @click="deleteTodo(todo.id)"
              class="text-gray-300 hover:text-red-600 transition-colors p-2"
              title="Delete task"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </div>
      </section>
    </div>

    <AddTodoModal 
      :is-open="isModalOpen" 
      @close="isModalOpen = false"
      @added="refresh"
    />
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'auth'
})

const { token } = useAuth()
const config = useRuntimeConfig()
const isModalOpen = ref(false)

const { data: todos, pending, error, refresh } = await useFetch(`${config.public.apiBase}/todos`, {
  headers: {
    Authorization: `Bearer ${token.value}`
  }
})

// Ensure todos is an array in case of null response
const safeTodos = computed(() => todos.value || [])

const ongoingTodos = computed(() => safeTodos.value.filter(t => !t.completed))
const completedTodos = computed(() => safeTodos.value.filter(t => t.completed))

const toggleComplete = async (todo) => {
  // Optimistic update could go here, but let's stick to simple refresh for now or local mutation
  // Local mutation for snappiness, then api call
  
  // Note: Nuxt useFetch result is reactive, but better not mutate it directly if we want consistent state.
  // We'll call API then refresh.
  
  try {
    await useFetch(`${config.public.apiBase}/todos/${todo.id}`, {
      method: 'PUT',
      headers: { Authorization: `Bearer ${token.value}` },
      body: {
        title: todo.title,
        completed: !todo.completed
      }
    })
    refresh()
  } catch (e) {
    console.error('Failed to update todo', e)
  }
}

const deleteTodo = async (id) => {
  if (!confirm('Are you sure you want to delete this task?')) return
  
  try {
    await useFetch(`${config.public.apiBase}/todos/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${token.value}` }
    })
    refresh()
  } catch (e) {
    console.error('Failed to delete todo', e)
  }
}
</script>
