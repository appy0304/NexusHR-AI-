<template>
  <article class="glass-card p-5 transition duration-300 hover:-translate-y-1 hover:shadow-glow">
    <div class="flex items-start justify-between">
      <div>
        <p class="text-sm text-slate-300">{{ label }}</p>
        <h3 class="mt-2 text-3xl font-semibold text-white">{{ value }}</h3>
      </div>
      <div class="rounded-2xl bg-white/[0.08] p-3 text-cyan-300">
        <component :is="iconMap[icon]" class="h-5 w-5" />
      </div>
    </div>
  </article>
</template>

<script setup>
import { Users, TrendingUp, Sparkles } from 'lucide-vue-next'

const props = defineProps({
  label: String,
  value: [String, Number],
  icon: String,
})

const iconMap = { Users, TrendingUp, Sparkles }
</script>