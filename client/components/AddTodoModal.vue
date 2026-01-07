<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
    <!-- Backdrop -->
    <div 
      class="absolute inset-0 bg-gray-900/50 backdrop-blur-sm transition-opacity"
      @click="$emit('close')"
    ></div>

    <!-- Modal Content -->
    <div class="glass-panel w-full max-w-md p-6 rounded-2xl relative z-10 transform transition-all scale-100 opacity-100 bg-white shadow-xl">
      <h3 class="text-2xl font-bold text-gray-900 mb-6">Add New Task</h3>
      
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">Task Title</label>
          <input 
            v-model="title" 
            type="text" 
            required
            class="input-field"
            placeholder="What needs to be done?"
            ref="inputRef"
          />
        </div>

        <div class="flex items-center justify-end gap-3 mt-8">
          <button 
            type="button" 
            @click="$emit('close')"
            class="px-4 py-2 text-gray-600 hover:text-gray-900 transition-colors"
          >
            Cancel
          </button>
          <button 
            type="submit" 
            class="btn-primary"
            :disabled="loading"
          >
            {{ loading ? 'Adding...' : 'Add Task' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  isOpen: Boolean
})

const emit = defineEmits(['close', 'added'])
const title = ref('')
const loading = ref(false)
const inputRef = ref(null)
const { token } = useAuth()
const config = useRuntimeConfig()

// Focus input when modal opens
watch(() => props.isOpen, (newVal) => {
  if (newVal) {
    setTimeout(() => inputRef.value?.focus(), 100)
    title.value = ''
  }
})

const handleSubmit = async () => {
  if (!title.value.trim()) return

  loading.value = true
  try {
    const { error } = await useFetch(`${config.public.apiBase}/todos`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token.value}`
      },
      body: {
        title: title.value
      }
    })

    if (error.value) throw error.value
    
    emit('added')
    emit('close')
    title.value = ''
  } catch (e) {
    console.error('Failed to add todo', e)
    // Optionally show error toast
  } finally {
    loading.value = false
  }
}
</script>
