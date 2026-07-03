<template>
  <DashboardLayout>
    <section class="space-y-8">
      <!-- Header -->
      <header>
        <p class="text-sm uppercase tracking-[0.35em] text-cyan-300">Insights</p>
        <h1 class="text-3xl font-semibold text-white lg:text-4xl">Analytics Dashboard</h1>
        <p class="mt-2 text-slate-300">Monitor workforce metrics and organizational trends.</p>
      </header>

      <!-- Main Stats -->
      <div class="grid gap-6 md:grid-cols-4">
        <StatsCard label="Total Employees" :value="totalCount" icon="Users" />
        <StatsCard label="Active Employees" :value="activeCount" icon="TrendingUp" />
        <StatsCard label="Departments" :value="departmentCount" icon="Sparkles" />
        <StatsCard label="New This Month" :value="newThisMonth" icon="TrendingUp" />
      </div>

      <!-- Department Breakdown -->
      <div class="glass-card p-6">
        <h3 class="text-lg font-semibold text-white mb-4">Department Breakdown</h3>
        <div v-if="departmentData.length === 0" class="text-slate-400">No data available yet.</div>
        <div v-else class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <div
            v-for="dept in departmentData"
            :key="dept.name"
            class="rounded-2xl border border-white/10 bg-white/[0.04] p-4"
          >
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-white">{{ dept.name }}</span>
              <span class="text-lg font-semibold text-cyan-300">{{ dept.count }}</span>
            </div>
            <div class="mt-2 h-2 rounded-full bg-white/10">
              <div
                class="h-2 rounded-full bg-gradient-to-r from-violet-500 to-cyan-400"
                :style="{ width: dept.percentage + '%' }"
              ></div>
            </div>
            <p class="mt-1 text-xs text-slate-400">{{ dept.percentage.toFixed(1) }}% of workforce</p>
          </div>
        </div>
      </div>

      <!-- Status Distribution -->
      <div class="grid gap-6 md:grid-cols-2">
        <div class="glass-card p-6">
          <h3 class="text-lg font-semibold text-white mb-4">Employment Status</h3>
          <div v-if="statusData.length === 0" class="text-slate-400">No data available yet.</div>
          <div v-else class="space-y-3">
            <div v-for="status in statusData" :key="status.name" class="flex items-center justify-between">
              <span class="text-sm text-slate-300">{{ status.name }}</span>
              <span class="text-sm font-semibold text-white">{{ status.count }}</span>
            </div>
          </div>
        </div>

        <div class="glass-card p-6">
          <h3 class="text-lg font-semibold text-white mb-4">Top Skills</h3>
          <div v-if="topSkills.length === 0" class="text-slate-400">No skills data available yet.</div>
          <div v-else class="flex flex-wrap gap-2">
            <span
              v-for="skill in topSkills"
              :key="skill"
              class="rounded-full bg-white/[0.08] px-3 py-1 text-xs text-cyan-200"
            >
              {{ skill }}
            </span>
          </div>
        </div>
      </div>

      <!-- Phase 5 Coming Soon Banner -->
      <div class="glass-card border border-violet-500/30 p-6">
        <div class="flex items-start gap-4">
          <div class="rounded-2xl bg-violet-500/20 p-3 text-violet-300">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <div>
            <h4 class="text-lg font-semibold text-white">Advanced Analytics Coming in Phase 5</h4>
            <p class="mt-1 text-sm text-slate-300">
              Department distribution charts, leave statistics, attrition rate, salary distribution, and AI-powered insights will be added here.
              The backend will use MongoDB aggregation pipelines for real-time dashboard metrics.
            </p>
          </div>
        </div>
      </div>
    </section>
  </DashboardLayout>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import DashboardLayout from '../layouts/DashboardLayout.vue'
import StatsCard from '../components/dashboard/StatsCard.vue'
import { useEmployeesStore } from '../stores/employees'
import { useEmployees } from '../composables/useEmployees'

const store = useEmployeesStore()
const { employees, loadEmployees } = useEmployees()

// Basic stats
const totalCount = computed(() => employees.value.length)
const activeCount = computed(() =>
  employees.value.filter((e) => e.employmentStatus === 'active').length
)
const departmentCount = computed(() =>
  new Set(employees.value.map((e) => e.department).filter(Boolean)).size
)
const newThisMonth = computed(() => {
  const now = new Date()
  const currentMonth = now.getMonth()
  const currentYear = now.getFullYear()
  return employees.value.filter((e) => {
    if (!e.joiningDate) return false
    const d = new Date(e.joiningDate)
    return d.getMonth() === currentMonth && d.getFullYear() === currentYear
  }).length
})

// Department breakdown with percentages
const departmentData = computed(() => {
  const deptMap = {}
  employees.value.forEach((e) => {
    if (!e.department) return
    deptMap[e.department] = (deptMap[e.department] || 0) + 1
  })
  const total = Object.values(deptMap).reduce((sum, c) => sum + c, 0) || 1
  return Object.entries(deptMap)
    .map(([name, count]) => ({
      name,
      count,
      percentage: (count / total) * 100
    }))
    .sort((a, b) => b.count - a.count)
})

// Status distribution
const statusData = computed(() => {
  const statusMap = {}
  employees.value.forEach((e) => {
    if (!e.employmentStatus) return
    statusMap[e.employmentStatus] = (statusMap[e.employmentStatus] || 0) + 1
  })
  return Object.entries(statusMap)
    .map(([name, count]) => ({ name, count }))
    .sort((a, b) => b.count - a.count)
})

// Top skills across all employees
const topSkills = computed(() => {
  const skillCount = {}
  employees.value.forEach((e) => {
    if (!e.skills || !Array.isArray(e.skills)) return
    e.skills.forEach((s) => {
      skillCount[s] = (skillCount[s] || 0) + 1
    })
  })
  return Object.entries(skillCount)
    .sort((a, b) => b[1] - a[1])
    .slice(0, 15)
    .map(([skill]) => skill)
})

// Load data on page mount
onMounted(async () => {
  await loadEmployees()
  store.setEmployees(employees.value)
})
</script>
