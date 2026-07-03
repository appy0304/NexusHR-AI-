<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4 backdrop-blur-sm">
    <div class="w-full max-w-md rounded-3xl border border-white/10 bg-slate-950 p-6 shadow-2xl">
      <h3 class="text-xl font-semibold text-white">{{ isEditing ? 'Edit Employee' : 'Create Employee' }}</h3>
<p class="mt-1 text-sm text-slate-300">Fill in the employee details below.</p>

      <form class="mt-6 space-y-4" @submit.prevent="submitForm">
      <div>
  <label class="mb-1 block text-sm text-slate-300">Employee ID</label>
  <input v-model="form.employeeId" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">First Name</label>
  <input v-model="form.firstName" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">Last Name</label>
  <input v-model="form.lastName" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">Email</label>
  <input v-model="form.email" type="email" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">Department</label>
  <input v-model="form.department" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">Designation</label>
  <input v-model="form.designation" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">Salary</label>
  <input v-model.number="form.salary" type="number" min="0" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">Status</label>
  <select v-model="form.employmentStatus" class="w-full rounded-2xl border border-white/10 bg-slate-900 px-4 py-3 text-white outline-none">
    <option value="active">Active</option>
    <option value="inactive">Inactive</option>
    <option value="terminated">Terminated</option>
    <option value="onboarding">Onboarding</option>
    <option value="leave">Leave</option>
  </select>
</div>
<div>
  <label class="mb-1 block text-sm text-slate-300">Skills</label>
  <input v-model="form.skillsText" placeholder="Go, MongoDB, Vue" class="w-full rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-3 text-white outline-none" />
</div>

        <p v-if="errorText" class="text-sm text-rose-300">{{ errorText }}</p>

        <div class="flex justify-end gap-3 pt-2">
          <button type="button" class="rounded-2xl border border-white/10 bg-white/[0.06] px-4 py-2 text-white" @click="emit('close')">Cancel</button>
          <button type="submit" class="rounded-2xl bg-gradient-to-r from-violet-500 to-cyan-400 px-4 py-2 font-semibold text-white">
            {{ isEditing ? 'Update' : 'Create' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { computed, reactive, ref } from 'vue'
import { useEmployeesStore } from '../../stores/employees'
import { useEmployees } from '../../composables/useEmployees'

const emit = defineEmits(['close'])
const store = useEmployeesStore()
const { createEmployee, updateEmployee } = useEmployees()

const selected = store.selectedEmployee
const isEditing = computed(() => Boolean(selected))
const errorText = ref('')

const form = reactive({
  employeeId: selected?.employeeId || '',
  firstName: selected?.firstName || '',
  lastName: selected?.lastName || '',
  email: selected?.email || '',
  phone: selected?.phone || '',
  department: selected?.department || '',
  designation: selected?.designation || '',
  salary: selected?.salary || 0,
  employmentStatus: selected?.employmentStatus || 'active',
  skillsText: selected?.skills?.join(', ') || '',
})

const validate = () => {
  if (!form.firstName.trim()) return 'First name is required.'
  if (!form.lastName.trim()) return 'Last name is required.'
  if (!form.email.trim()) return 'Email is required.'
  if (!form.department.trim()) return 'Department is required.'
  if (!form.designation.trim()) return 'Designation is required.'
  if (Number(form.salary) < 0) return 'Salary cannot be negative.'
  return ''
}

const buildPayload = () => ({
  employeeId: form.employeeId.trim(),
  firstName: form.firstName.trim(),
  lastName: form.lastName.trim(),
  email: form.email.trim().toLowerCase(),
  phone: form.phone.trim(),
  department: form.department.trim(),
  designation: form.designation.trim(),
  salary: Number(form.salary),
  employmentStatus: form.employmentStatus,
  skills: form.skillsText
    .split(',')
    .map((skill) => skill.trim())
    .filter(Boolean),
  dateOfBirth: selected?.dateOfBirth || '2000-01-01T00:00:00Z',
  joiningDate: selected?.joiningDate || new Date().toISOString(),
  address: selected?.address || {
    street: '',
    city: 'Pune',
    state: 'Maharashtra',
    zipCode: '411001',
    country: 'India',
  },
  emergencyContact: selected?.emergencyContact || {
    name: 'Not Provided',
    relationship: 'Not Provided',
    phone: '0000000000',
    email: '',
  },
})

const submitForm = async () => {
  errorText.value = validate()
  if (errorText.value) return

  try {
    const payload = buildPayload()

    if (isEditing.value) {
      await updateEmployee(store.selectedEmployee.id || store.selectedEmployee._id, payload)
      store.setToast('Employee updated successfully', 'success')
    } else {
      await createEmployee(payload)
      store.setToast('Employee created successfully', 'success')
    }

    emit('close')
    store.closeModal()
  } catch (e) {
    errorText.value = e.message || 'Something went wrong.'
  }
}
</script>