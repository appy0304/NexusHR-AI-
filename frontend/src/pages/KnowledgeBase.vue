<template>
  <DashboardLayout>
    <section class="max-w-4xl mx-auto space-y-6">
      <header class="mb-8">
        <p class="text-sm uppercase tracking-[0.35em] text-cyan-300">AI Administration</p>
        <h1 class="text-3xl font-semibold text-white">Knowledge Base Upload</h1>
        <p class="mt-2 text-slate-400">Upload HR policies, leave rules, and company guidelines to the AI Vector Database.</p>
      </header>

      <div class="bg-white/5 border border-white/10 rounded-3xl p-6 backdrop-blur-xl shadow-soft">
        <form @submit.prevent="submitPolicy" class="space-y-5">
          <!-- Document ID -->
          <div>
            <label class="block text-sm font-medium text-slate-300 mb-2">Document Title / ID</label>
            <input v-model="form.id" type="text" required placeholder="e.g., leave_policy_2026"
                   class="w-full bg-[#0f172a] border border-slate-700/50 rounded-xl px-4 py-3 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-500/50 transition">
            <p class="mt-1 text-xs text-slate-500">A unique identifier for this piece of information.</p>
          </div>

          <!-- Document Text -->
          <div>
            <label class="block text-sm font-medium text-slate-300 mb-2">Policy Content</label>
            <textarea v-model="form.text" required rows="10" placeholder="Paste the exact text of the policy here..."
                      class="w-full bg-[#0f172a] border border-slate-700/50 rounded-xl px-4 py-3 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-cyan-500/50 transition resize-y"></textarea>
          </div>

          <!-- Status Message -->
          <div v-if="status.message" :class="['p-4 rounded-xl text-sm border', status.type === 'success' ? 'bg-emerald-500/10 border-emerald-500/20 text-emerald-400' : 'bg-rose-500/10 border-rose-500/20 text-rose-400']">
            {{ status.message }}
          </div>

          <!-- Submit Button -->
          <div class="flex justify-end pt-4">
            <button type="submit" :disabled="loading"
                    class="rounded-xl bg-gradient-to-r from-violet-500 to-cyan-400 px-6 py-3 text-sm font-semibold text-white disabled:opacity-50 hover:shadow-glow transition flex items-center gap-2">
              <span v-if="loading" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></span>
              {{ loading ? 'Uploading to Pinecone...' : 'Upload to Knowledge Base' }}
            </button>
          </div>
        </form>
      </div>
    </section>
  </DashboardLayout>
</template>

<script setup>
import { ref, reactive } from 'vue'
import DashboardLayout from '../layouts/DashboardLayout.vue'
import { aiAPI } from '../services/aiService'

const loading = ref(false)
const status = reactive({ message: '', type: '' })

const form = reactive({
  id: '',
  text: ''
})

const submitPolicy = async () => {
  if (!form.id.trim() || !form.text.trim()) return

  loading.value = true
  status.message = ''
  
  try {
    const response = await aiAPI.upload(form.id, form.text)
    status.type = 'success'
    status.message = response.data.message || 'Successfully uploaded to the vector database!'
    form.id = ''
    form.text = ''
  } catch (error) {
    status.type = 'error'
    status.message = error.response?.data?.message || 'Failed to upload document. Please check the backend connection.'
  } finally {
    loading.value = false
  }
}
</script>
