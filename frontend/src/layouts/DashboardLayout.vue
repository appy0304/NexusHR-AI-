<template>
  <div class="min-h-screen bg-[radial-gradient(circle_at_top,_#111827_0%,_#030712_45%,_#020617_100%)] text-slate-100">
    <div class="mx-auto flex min-h-screen max-w-7xl gap-6 p-4 lg:p-6">
      <Sidebar class="hidden lg:block" />
      <main class="flex-1 rounded-3xl border border-white/10 bg-white/[0.06] p-4 shadow-soft backdrop-blur-xl lg:p-6">
        <Navbar />
        <div class="mt-6">
          <slot />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import Navbar from '../components/layout/Navbar.vue'
import Sidebar from '../components/layout/Sidebar.vue'
</script>