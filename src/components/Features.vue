<script setup lang="ts">
</script>

<template>
  <section class="py-20 bg-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <h2 class="text-4xl font-bold text-gray-900 mb-4">核心功能</h2>
        <p class="text-xl text-gray-600 max-w-2xl mx-auto">
          Vue 3 提供了强大的开发工具和现代化的架构
        </p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="bg-gray-50 rounded-2xl p-8 hover:shadow-xl transition-all transform hover:-translate-y-2 border border-gray-100">
          <div class="w-14 h-14 bg-gradient-to-r from-blue-500 to-indigo-500 rounded-xl flex items-center justify-center mb-6">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-3">Composition API</h3>
          <p class="text-gray-600">
            使用更灵活的代码组织方式，实现更好的逻辑复用和代码可维护性
          </p>
        </div>
        <div class="bg-gray-50 rounded-2xl p-8 hover:shadow-xl transition-all transform hover:-translate-y-2 border border-gray-100">
          <div class="w-14 h-14 bg-gradient-to-r from-purple-500 to-pink-500 rounded-xl flex items-center justify-center mb-6">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-3">TypeScript</h3>
          <p class="text-gray-600">
            完整的 TypeScript 支持，提供强大的类型检查和智能提示
          </p>
        </div>
        <div class="bg-gray-50 rounded-2xl p-8 hover:shadow-xl transition-all transform hover:-translate-y-2 border border-gray-100">
          <div class="w-14 h-14 bg-gradient-to-r from-green-500 to-teal-500 rounded-xl flex items-center justify-center mb-6">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-3">性能优化</h3>
          <p class="text-gray-600">
            虚拟 DOM 改进，更小的包体积，更快的渲染速度
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
