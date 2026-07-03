<template>
  <div class="fixed bottom-4 right-4 z-50 flex w-80 flex-col gap-3">
    <Toast v-if="store.toast" :message="store.toast.message" :type="store.toast.type" />
  </div>
</template>

<script setup>
import Toast from './Toast.vue'
import { useEmployeesStore } from '../../stores/employees'

const store = useEmployeesStore()
</script>