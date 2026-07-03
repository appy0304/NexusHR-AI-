import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useThemeStore = defineStore('theme', () => {
  const mode = ref(localStorage.getItem('theme') || 'light')

  const isDark = computed(() => mode.value === 'dark')

  function toggle() {
    mode.value = mode.value === 'light' ? 'dark' : 'light'
    apply()
  }

  function apply() {
    if (mode.value === 'dark') {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
    localStorage.setItem('theme', mode.value)
  }

  function init() {
    const saved = localStorage.getItem('theme')
    if (saved) {
      mode.value = saved
    } else if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
      mode.value = 'dark'
    }
    apply()
  }

  return { mode, isDark, toggle, apply, init }
})
