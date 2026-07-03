<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 p-4 backdrop-blur-sm">
    <div class="w-full max-w-md rounded-3xl border border-rose-400/20 bg-slate-950 p-6 shadow-2xl">
      <h3 class="text-xl font-semibold text-white">Delete employee?</h3>
      <p class="mt-2 text-sm text-slate-300">
        This will permanently remove {{ store.selectedEmployee?.firstName || 'this employee' }}.
      </p>
      <div class="mt-6 flex justify-end gap-3">
        <button class="rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-2 text-white"
          @click="emit('close')">Cancel</button>
        <button class="rounded-2xl bg-rose-500 px-4 py-2 font-semibold text-white"
          @click="confirmDelete">Delete</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useEmployees } from '../../composables/useEmployees'
import { useEmployeesStore } from '../../stores/employees'

const emit = defineEmits(['close'])
const store = useEmployeesStore()
const { deleteEmployee } = useEmployees()

const confirmDelete = async () => {
  try {
    await deleteEmployee(store.selectedEmployee.id || store.selectedEmployee._id)
    store.setToast('Employee deleted successfully', 'success')
    emit('close')
    store.closeDeleteModal()
  } catch (e) {
    store.setToast(e.message || 'Delete failed', 'error')
  }
}
</script>