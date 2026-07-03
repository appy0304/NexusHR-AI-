<template>
  <section class="glass-card overflow-hidden p-4 lg:p-6">
    <div class="mb-4 flex items-center justify-between">
      <div>
        <h3 class="text-xl font-semibold text-white">All Users</h3>
        <p class="text-sm text-slate-300">Polished CRUD experience for your dashboard.</p>
      </div>
    </div>

    <div v-if="loading" class="space-y-3">
      <div v-for="n in 4" :key="n" class="h-14 animate-pulse rounded-2xl bg-white/[0.06]"></div>
    </div>

    <div v-else-if="filteredEmployees.length === 0" class="rounded-3xl border border-dashed border-white/10 p-8 text-center text-slate-300">
      No users found. Create your first one.
    </div>

    <div v-else class="overflow-x-auto">
      <table class="w-full dark:bg-slate-800 dark:text-slate-100">

        <thead class="text-slate-400">
          <tr>
            <th class="px-3 py-3">Employee</th>
<th class="px-3 py-3">Department</th>
<th class="px-3 py-3">Designation</th>
<th class="px-3 py-3">Status</th>
<th class="px-3 py-3">Created</th>
<th class="px-3 py-3 text-right">Actions</th>
          </tr>
        </thead>
        <tbody>
         <tr v-for="employee in paginatedEmployees" :key="employee.id || employee._id" class="border-t border-white/[0.06] transition hover:bg-white/[0.06]">
  <td class="px-3 py-4">
    <p class="font-medium text-white">{{ employee.firstName }} {{ employee.lastName }}</p>
    <p class="text-xs text-slate-400">{{ employee.email }}</p>
  </td>
  <td class="px-3 py-4">{{ employee.department }}</td>
  <td class="px-3 py-4">{{ employee.designation }}</td>
  <td class="px-3 py-4">{{ employee.employmentStatus }}</td>
  <td class="px-3 py-4">{{ formatDate(employee.createdAt) }}</td>
  <td class="px-3 py-4 text-right">
    <button class="mr-2 rounded-xl bg-white/[0.08] px-3 py-2 text-xs text-cyan-200" @click="$emit('edit', employee)">Edit</button>
    <button class="rounded-xl bg-rose-500/15 px-3 py-2 text-xs text-rose-200" @click="$emit('delete', employee)">Delete</button>
  </td>
</tr>
        </tbody>
      </table>
    </div>
  </section>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useEmployeesStore } from '../../stores/employees'
import { useEmployees } from '../../composables/useEmployees'

const store = useEmployeesStore()
const { employees, loading, loadEmployees } = useEmployees()

const filteredEmployees = computed(() => {
  const q = store.query.toLowerCase().trim()
  if (!q) return employees.value

  return employees.value.filter((employee) => {
    const fullName = `${employee.firstName || ''} ${employee.lastName || ''}`.toLowerCase()
    return (
      fullName.includes(q) ||
      employee.email?.toLowerCase().includes(q) ||
      employee.department?.toLowerCase().includes(q) ||
      employee.designation?.toLowerCase().includes(q)
    )
  })
})

const paginatedEmployees = computed(() => {
  const start = (store.currentPage - 1) * store.pageSize
  return filteredEmployees.value.slice(start, start + store.pageSize)
})

const formatDate = (value) => value ? new Date(value).toLocaleString() : '-'

onMounted(async () => {
  await loadEmployees()
  store.setEmployees(employees.value)
})
</script>
