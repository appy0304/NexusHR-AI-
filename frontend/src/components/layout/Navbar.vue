<template>
  <nav class="glass-card flex items-center justify-between px-4 py-4 dark:bg-slate-800/50">

    <div>
          <p class="text-xs uppercase tracking-[0.35em] text-cyan-300">Dashboard</p>
      <h2 class="text-xl font-semibold text-white">Employee Command Center</h2>

    </div>
    <div class="flex items-center gap-3">
      <input
        v-model="store.query"
        placeholder="Search employees"
                class="rounded-2xl border border-white/10 bg-slate-900/70 dark:bg-slate-700/70 px-4 py-2 text-sm text-white outline-none ring-0 placeholder:text-slate-400"

         />
         
      <!-- Theme Toggle -->
      <button
        @click="themeStore.toggle()"
        class="rounded-2xl border border-white/10 bg-white/5 px-4 py-2 text-sm text-white hover:bg-white/10 transition-colors flex items-center gap-2"
      >
        <span v-if="themeStore.isDark">☀️ Light</span>
        <span v-else>🌙 Dark</span>
      </button>

      <!-- User Info + Logout -->
      <div v-if="authStore.user" class="flex items-center gap-3">
        <div class="text-right">
          <p class="text-sm font-medium text-white">{{ authStore.user.name || authStore.user.email }}</p>
          <p class="text-xs text-slate-400 capitalize">{{ authStore.user.role }}</p>
        </div>
        <div class="h-9 w-9 rounded-full bg-gradient-to-br from-violet-500 to-cyan-400 flex items-center justify-center text-white text-sm font-semibold">
          {{ getInitials(authStore.user.name || authStore.user.email) }}
        </div>
        <button
          @click="handleLogout"
          class="rounded-2xl border border-rose-500/30 bg-rose-500/10 px-4 py-2 text-sm text-rose-300 hover:bg-rose-500/20 transition-colors"
        >
          Logout
        </button>
      </div>

      <!-- Not Authenticated fallback -->
      <router-link
        to="/login"
        v-else
        class="rounded-2xl bg-gradient-to-r from-violet-500 to-cyan-400 px-4 py-2 text-sm font-semibold text-white"
      >
        Login
      </router-link>
    </div>
  </nav>
</template>

<script setup>
import { useEmployeesStore } from '../../stores/employees'
import { useAuthStore } from '../../stores/auth'
import { useRouter } from 'vue-router'
import { useThemeStore } from '../../stores/theme'


const store = useEmployeesStore()
const authStore = useAuthStore()
const router = useRouter()
const themeStore = useThemeStore()


const getInitials = (name) => {
  if (!name) return 'U'
  return name.split(' ').map(n => n[0]).join('').toUpperCase().slice(0, 2)
}

const handleLogout = async () => {
  try {
    const refreshToken = localStorage.getItem('refreshToken')
    if (refreshToken) {
      await authStore.logout(refreshToken)
    }
  } catch {
    // Ignore logout API errors — still clear local state
  } finally {
    authStore.logout()
    router.push('/login')
  }
}
</script>
