<template>
  <aside class="glass-card w-72 p-4 dark:bg-slate-800/50">

    <div class="mb-8 flex items-center gap-3">
      <div class="h-12 w-12 rounded-2xl bg-gradient-to-br from-violet-500 to-cyan-400"></div>
      <div>
        <p class="text-xs uppercase tracking-[0.35em] text-cyan-300">SaaS</p>
        <h3 class="text-lg font-semibold text-white">Admin Hub</h3>
      </div>
    </div>
    <nav class="space-y-2 text-sm text-slate-200">
      <router-link to="/"
        class="block rounded-2xl bg-white/[0.08] px-4 py-3 hover:bg-white/[0.12] transition-colors">Dashboard</router-link>
      <router-link to="/users"
        class="block rounded-2xl px-4 py-3 hover:bg-white/[0.12] transition-colors">Users</router-link>
      <router-link to="/analytics"
        class="block rounded-2xl px-4 py-3 hover:bg-white/[0.12] transition-colors">Analytics</router-link>
      <router-link to="/leaves"
        class="block rounded-2xl px-4 py-3 hover:bg-white/[0.12] transition-colors">Leaves</router-link>
      <router-link to="/ai-chat" class="block rounded-2xl px-4 py-3 hover:bg-white/[0.12] transition-colors">AI
        Assistant</router-link>
      <router-link to="/knowledge-base" class="block rounded-2xl px-4 py-3 hover:bg-white/[0.12] transition-colors">Knowledge Base</router-link>

    </nav>
  </aside>
</template>