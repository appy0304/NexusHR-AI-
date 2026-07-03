<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-950 via-slate-900 to-slate-950 px-4">
    <div class="glass-card w-full max-w-md p-8">
      <div class="text-center mb-8">
        <div class="mx-auto h-16 w-16 rounded-2xl bg-gradient-to-br from-violet-500 to-cyan-400 mb-4"></div>
        <h1 class="text-2xl font-semibold text-white">Employee Management Platform</h1>
        <p class="text-sm text-slate-400 mt-2">Sign in to your account</p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-5">
        <div>
          <label class="block text-xs uppercase tracking-[0.35em] text-cyan-300 mb-2">Email</label>
          <input
            v-model="email"
            type="email"
            placeholder="admin@company.com"
            class="w-full rounded-2xl border border-white/10 bg-slate-900/70 px-4 py-3 text-sm text-white outline-none placeholder:text-slate-500 focus:border-cyan-400 transition"
            required
          />
        </div>

        <div>
          <label class="block text-xs uppercase tracking-[0.35em] text-cyan-300 mb-2">Password</label>
          <input
            v-model="password"
            type="password"
            placeholder="Enter your password"
            class="w-full rounded-2xl border border-white/10 bg-slate-900/70 px-4 py-3 text-sm text-white outline-none placeholder:text-slate-500 focus:border-cyan-400 transition"
            required
          />
        </div>

        <p v-if="error" class="text-rose-400 text-sm bg-rose-500/10 rounded-xl p-3">{{ error }}</p>

        <button
          type="submit"
          class="w-full rounded-2xl bg-gradient-to-r from-violet-500 via-fuchsia-500 to-cyan-400 py-3 text-sm font-semibold text-white shadow-glow transition hover:-translate-y-0.5 disabled:opacity-50"
          :disabled="loading"
        >
          {{ loading ? 'Signing in...' : 'Sign In' }}
        </button>
      </form>

      <p class="mt-6 text-center text-xs text-slate-500">
        Default admin: admin@company.com / Admin@123
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
  loading.value = true
  error.value = ''

  try {
    await authStore.login(email.value, password.value)
    router.push('/')
  } catch (err) {
    error.value = err.message || 'Login failed. Please check your credentials.'
  } finally {
    loading.value = false
  }
}
</script>
