<template>
   <DashboardLayout>
  <div class="space-y-6">
    <h1 class="text-2xl font-bold text-white">Leave Management</h1>

    <!-- Leave Balance Card -->
    <div class="glass-card p-6">
      <h2 class="text-lg font-semibold text-white mb-4">Leave Balance</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div class="bg-white/5 rounded-2xl p-4">
          <p class="text-sm text-slate-400">Total Allocated</p>
          <p class="text-2xl font-bold text-white">{{ leaveBalance?.totalAllocated || 0 }}</p>
        </div>
        <div class="bg-white/5 rounded-2xl p-4">
          <p class="text-sm text-slate-400">Used</p>
          <p class="text-2xl font-bold text-cyan-400">{{ leaveBalance?.used || 0 }}</p>
        </div>
        <div class="bg-white/5 rounded-2xl p-4">
          <p class="text-sm text-slate-400">Remaining</p>
          <p class="text-2xl font-bold text-green-400">{{ leaveBalance?.remaining || 0 }}</p>
        </div>
      </div>
    </div>

    <!-- Apply for Leave Form -->
    <div class="glass-card p-6">
      <h2 class="text-lg font-semibold text-white mb-4">Apply for Leave</h2>
      <form @submit.prevent="handleApplyLeave" class="space-y-4">
        <div>
          <label class="block text-sm text-slate-300 mb-2">Leave Type</label>
          <select v-model="form.leaveType" class="w-full rounded-2xl border border-white/10 bg-slate-900/70 px-4 py-2 text-white">
            <option value="annual">Annual Leave</option>
            <option value="sick">Sick Leave</option>
            <option value="maternity">Maternity Leave</option>
            <option value="paternity">Paternity Leave</option>
            <option value="unpaid">Unpaid Leave</option>
            <option value="comp_off">Comp Off</option>
          </select>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-300 mb-2">Start Date</label>
            <input v-model="form.startDate" type="date" class="w-full rounded-2xl border border-white/10 bg-slate-900/70 px-4 py-2 text-white" />
          </div>
          <div>
            <label class="block text-sm text-slate-300 mb-2">End Date</label>
            <input v-model="form.endDate" type="date" class="w-full rounded-2xl border border-white/10 bg-slate-900/70 px-4 py-2 text-white" />
          </div>
        </div>
        <div>
          <label class="block text-sm text-slate-300 mb-2">Reason</label>
          <textarea v-model="form.reason" rows="3" class="w-full rounded-2xl border border-white/10 bg-slate-900/70 px-4 py-2 text-white"></textarea>
        </div>
        <button type="submit" class="rounded-2xl bg-gradient-to-r from-violet-500 to-cyan-400 px-6 py-2 text-sm font-semibold text-white">
          Submit Request
        </button>
      </form>
    </div>

    <!-- Leave Requests Table -->
    <div class="glass-card p-6">
      <h2 class="text-lg font-semibold text-white mb-4">My Leave Requests</h2>
      <table class="w-full">
        <thead>
          <tr class="border-b border-white/10">
            <th class="pb-2 text-left text-sm text-slate-300">Type</th>
            <th class="pb-2 text-left text-sm text-slate-300">Start Date</th>
            <th class="pb-2 text-left text-sm text-slate-300">End Date</th>
            <th class="pb-2 text-left text-sm text-slate-300">Days</th>
            <th class="pb-2 text-left text-sm text-slate-300">Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="leave in leaveStore.leaves" :key="leave.id" class="border-b border-white/5">
            <td class="py-3 text-sm text-white">{{ leave.leaveType }}</td>
            <td class="py-3 text-sm text-slate-300">{{ formatDate(leave.startDate) }}</td>
            <td class="py-3 text-sm text-slate-300">{{ formatDate(leave.endDate) }}</td>
            <td class="py-3 text-sm text-slate-300">{{ leave.days }}</td>
            <td class="py-3">
              <span :class="getStatusClass(leave.status)" class="px-3 py-1 rounded-full text-xs font-semibold">
                {{ leave.status }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
  </DashboardLayout>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useLeaveStore } from '../stores/leave'
import DashboardLayout from '../layouts/DashboardLayout.vue'

const leaveStore = useLeaveStore()

const form = ref({
  leaveType: 'annual',
  startDate: '',
  endDate: '',
  reason: '',
})

const handleApplyLeave = async () => {
  try {
    await leaveStore.createLeave({
      employeeId: 'CURRENT_USER_ID', // Replace with actual user ID
      leaveType: form.value.leaveType,
      startDate: form.value.startDate,
      endDate: form.value.endDate,
      reason: form.value.reason,
    })
    // Reset form
    form.value = { leaveType: 'annual', startDate: '', endDate: '', reason: '' }
  } catch (err) {
    console.error('Failed to create leave request:', err)
  }
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString()
}

const getStatusClass = (status) => {
  const classes = {
    pending: 'bg-yellow-500/20 text-yellow-300',
    approved: 'bg-green-500/20 text-green-300',
    rejected: 'bg-red-500/20 text-red-300',
    cancelled: 'bg-slate-500/20 text-slate-300',
  }
  return classes[status] || ''
}
</script>