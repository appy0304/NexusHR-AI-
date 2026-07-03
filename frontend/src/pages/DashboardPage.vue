<template>
  <DashboardLayout>
    <section class="space-y-8">
      <header class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
        <div>
          <p class="text-sm uppercase tracking-[0.35em] text-cyan-300">Employee Management</p>
<h1 class="text-3xl font-semibold text-white lg:text-4xl">AI Enable employee dashboard</h1>
<p class="mt-2 text-slate-300">Search, create, update, and delete employees with a polished SaaS interface.</p>
        </div>
        <button
          class="rounded-2xl bg-gradient-to-r from-violet-500 via-fuchsia-500 to-cyan-400 px-5 py-3 text-sm font-semibold text-white shadow-glow transition hover:-translate-y-0.5"
          @click="openCreateModal"
        >
          + Create User
        </button>
      </header>

      <div class="grid gap-6 md:grid-cols-3">
        <StatsCard label="Total Employees" :value="employees.length" icon="Users" />
<StatsCard label="Active Employees" :value="activeCount" icon="TrendingUp" />
<StatsCard label="Departments" :value="departmentCount" icon="Sparkles" />
      </div>

      <EmployeeTable @edit="openEditModal" @delete="confirmDelete" />
    </section>

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

const activeCount = computed(() =>
  employees.value.filter((employee) => employee.employmentStatus === 'active').length,
)

const departmentCount = computed(() =>
  new Set(employees.value.map((employee) => employee.department).filter(Boolean)).size,
)

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

onMounted(async () => {
  await loadEmployees()
  store.setEmployees(employees.value)
})
</script>
<!-- 
<script setup>
import { computed, onMounted } from 'vue'
import DashboardLayout from '../layouts/DashboardLayout.vue'
import StatsCard from '../components/dashboard/StatsCard.vue'
import UserTable from '../components/dashboard/EmployeeTable.vue'
import UserModal from '../components/dashboard/EmployeeModal.vue'
import DeleteConfirmModal from '../components/dashboard/DeleteConfirmModal.vue'
import ToastContainer from '../components/ui/ToastContainer.vue'
import { useUsersStore } from '../stores/employees.js'
import { useUsers } from '../composables/useEmployees.js'

const store = useUsersStore()
const { users, loadUsers } = useUsers()

const averageAge = computed(() => {
  if (!users.value.length) return 0
  return Math.round(users.value.reduce((sum, user) => sum + Number(user.age || 0), 0) / users.value.length)
})

const openCreateModal = () => {
  store.setSelectedUser(null)
  store.openModal()
}

const openEditModal = (user) => {
  store.setSelectedUser(user)
  store.openModal()
}

const confirmDelete = (user) => {
  store.openDeleteModal(user)
}

onMounted(async () => {
  await loadUsers()
  store.setUsers(users.value)
})
</script> -->
