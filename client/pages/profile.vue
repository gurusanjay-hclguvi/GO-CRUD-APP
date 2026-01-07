<template>
  <div class="max-w-2xl mx-auto py-12">
    <div class="glass-panel p-8 rounded-2xl relative overflow-hidden">
      <div class="flex flex-col sm:flex-row items-center gap-8">
        <!-- Avatar Placeholder -->
        <div class="w-32 h-32 rounded-full bg-black flex items-center justify-center shadow-lg">
          <span class="text-4xl font-bold text-white uppercase">{{ user?.name?.charAt(0) || 'U' }}</span>
        </div>
        
        <div class="text-center sm:text-left space-y-2">
          <h1 class="text-3xl font-bold text-gray-900">{{ user?.name }}</h1>
          <p class="text-gray-500 flex items-center justify-center sm:justify-start gap-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
              <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
              <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
            </svg>
            {{ user?.email }}
          </p>
          <div class="pt-4">
             <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-gray-100 text-gray-800 border border-gray-200">
               Active Account
             </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'auth'
})

const { user } = useAuth()
</script>
