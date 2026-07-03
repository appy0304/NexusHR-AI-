<template>
  <DashboardLayout>
    <section class="flex flex-col h-[calc(100vh-180px)]">
      <!-- Header -->
      <header class="mb-4">
          <p class="text-sm uppercase tracking-[0.35em] text-cyan-300">AI Knowledge Base</p>
          <h1 class="text-2xl font-semibold text-white">HR Assistant</h1>
      </header>

      <!-- Chat Messages -->
      <div ref="chatBox" class="flex-1 overflow-y-auto space-y-3 pr-2">
        <div v-for="(msg, i) in messages" :key="i"
             :class="['flex', msg.role === 'user' ? 'justify-end' : 'justify-start']">
          <div :class="['max-w-[75%] rounded-2xl px-4 py-3 text-sm',
                msg.role === 'user'
                  ? 'bg-gradient-to-r from-violet-500 to-cyan-400 text-white'
                  : 'bg-white/10 text-slate-200 border border-white/10']">
            <p class="whitespace-pre-wrap">{{ msg.content }}</p>
            <div v-if="msg.sources?.length" class="mt-2 pt-2 border-t border-white/10 text-xs text-slate-400">
              📄 {{ msg.sources.join(', ') }}
            </div>
          </div>
        </div>
        <div v-if="loading" class="flex justify-start">
          <div class="bg-white/10 rounded-2xl px-4 py-3 text-sm text-slate-400 animate-pulse">
            Thinking...
          </div>
        </div>
      </div>

      <!-- Input -->
      <div class="mt-4 flex gap-2">
        <input v-model="input" @keyup.enter="send" type="text"
               placeholder="Ask about company policies..."
               class="flex-1 rounded-xl bg-white/10 border border-white/10 px-4 py-3 text-sm text-white placeholder-slate-400 outline-none focus:ring-2 focus:ring-cyan-400/50" />
        <button @click="send" :disabled="loading || !input.trim()"
                class="rounded-xl bg-gradient-to-r from-violet-500 to-cyan-400 px-6 py-3 text-sm font-semibold text-white disabled:opacity-40 hover:-translate-y-0.5 transition">
          Send
        </button>
      </div>
    </section>
  </DashboardLayout>
</template>

<script setup>
import { ref, nextTick } from 'vue'
import DashboardLayout from '../layouts/DashboardLayout.vue'
import { aiAPI } from '../services/aiService'

const input = ref('')
const loading = ref(false)
const chatBox = ref(null)
const messages = ref([
  { role: 'ai', content: 'Hello! Ask me anything about company policies, leaves, or HR guidelines.' }
])

const send = async () => {
  const q = input.value.trim()
  if (!q) return

  messages.value.push({ role: 'user', content: q })
  input.value = ''
  loading.value = true
  await scrollDown()

  try {
    const { data } = await aiAPI.ask(q, 'current-user')
    messages.value.push({
      role: 'ai',
      content: data.data.answer,
      sources: data.data.sources
    })
  } catch {
    messages.value.push({ role: 'ai', content: 'Sorry, something went wrong. Please try again.' })
  } finally {
    loading.value = false
    await scrollDown()
  }
}

const scrollDown = async () => {
  await nextTick()
  if (chatBox.value) chatBox.value.scrollTop = chatBox.value.scrollHeight
}
</script>
