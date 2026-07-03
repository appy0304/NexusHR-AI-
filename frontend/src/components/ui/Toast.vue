<template>
  <div :class="['rounded-2xl border px-4 py-3 text-sm shadow-xl', toneClass]">
    {{ message }}
  </div>
</template>

<script setup>
const props = defineProps({ message: String, type: String })

const toneClass = props.type === 'error'
  ? 'border-rose-400/20 bg-rose-500/10 text-rose-100'
  : 'border-emerald-400/20 bg-emerald-500/10 text-emerald-100'
</script>