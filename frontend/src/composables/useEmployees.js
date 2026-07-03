import { ref } from 'vue'
import { employeeService } from '../services/employeeService'

// Single shared state — created ONCE when this module is first imported
const sharedEmployees = ref([])
const sharedLoading = ref(false)

export function useEmployees() {
  const loadEmployees = async () => {
    sharedLoading.value = true
    try {
      sharedEmployees.value = await employeeService.getEmployees()
    } finally {
      sharedLoading.value = false
    }
  }

  const createEmployee = async (payload) => {
    const result = await employeeService.createEmployee(payload)
    await loadEmployees()
    return result
  }

  const updateEmployee = async (id, payload) => {
    const result = await employeeService.updateEmployee(id, payload)
    await loadEmployees()
    return result
  }

  const deleteEmployee = async (id) => {
    const result = await employeeService.deleteEmployee(id)
    await loadEmployees()
    return result
  }

  return {
    employees: sharedEmployees,
    loading: sharedLoading,
    loadEmployees,
    createEmployee,
    updateEmployee,
    deleteEmployee,
  }
}
