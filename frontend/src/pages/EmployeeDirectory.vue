<template>
  <DashboardLayout>
        <section class="space-y-8 dark:text-slate-200">

      <!-- Header -->
      <header class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
        <div>
          <p class="text-sm uppercase tracking-[0.35em] text-cyan-300">Employee Management</p>
          <h1 class="text-3xl font-semibold text-white lg:text-4xl">Employee Directory</h1>
          <p class="mt-2 text-slate-300">Browse, search, and manage all employees in the organization.</p>
        </div>
        <button
          class="rounded-2xl bg-gradient-to-r from-violet-500 via-fuchsia-500 to-cyan-400 px-5 py-3 text-sm font-semibold text-white shadow-glow transition hover:-translate-y-0.5"
          @click="openCreateModal"
        >
          + Create Employee
        </button>
      </header>

      <!-- Stats Cards -->
      <div class="grid gap-6 md:grid-cols-3">
        <StatsCard label="Total Employees" :value="totalCount" icon="Users" />
        <StatsCard label="Active Employees" :value="activeCount" icon="TrendingUp" />
        <StatsCard label="Departments" :value="departmentCount" icon="Sparkles" />
      </div>

      <!-- Employee Table -->
      <EmployeeTable @edit="openEditModal" @delete="confirmDelete" />
    </section>

    <!-- Modals -->
    <EmployeeModal v-if="store.isModalOpen" @close="store.closeModal" />
    <DeleteConfirmModal v-if="store.isDeleteOpen" @close="store.closeDeleteModal" />
    <ToastContainer />
  </DashboardLayout>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import DashboardLayout from '../layouts/DashboardLayout.vue'
import StatsCard from '../components/dashboard/StatsCard.vue'
import EmployeeTable from '../components/dashboard/EmployeeTable.vue'
import EmployeeModal from '../components/dashboard/EmployeeModal.vue'
import DeleteConfirmModal from '../components/dashboard/DeleteConfirmModal.vue'
import ToastContainer from '../components/ui/ToastContainer.vue'
import { useEmployeesStore } from '../stores/employees'
import { useEmployees } from '../composables/useEmployees'

const store = useEmployeesStore()
const { employees, loadEmployees } = useEmployees()

// Computed stats
const totalCount = computed(() => employees.value.length)
const activeCount = computed(() =>
  employees.value.filter((e) => e.employmentStatus === 'active').length
)
const departmentCount = computed(() =>
  new Set(employees.value.map((e) => e.department).filter(Boolean)).size
)

// Modal handlers
const openCreateModal = () => {
  store.setSelectedEmployee(null)
  store.openModal()
}

const openEditModal = (employee) => {
  store.setSelectedEmployee(employee)
  store.openModal()
}

const confirmDelete = (employee) => {
  store.openDeleteModal(employee)
}

// Load data on page mount
onMounted(async () => {
  await loadEmployees()
  store.setEmployees(employees.value)
})
</script>
