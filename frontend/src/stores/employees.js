import { defineStore } from 'pinia'

export const useEmployeesStore = defineStore('employees', {
  state: () => ({
    employees: [],
    loading: false,
    query: '',
    currentPage: 1,
    pageSize: 6,
    selectedEmployee: null,
    isModalOpen: false,
    isDeleteOpen: false,
    toast: null,
  }),
  actions: {
    setEmployees(list) {
      this.employees = list
    },
    setQuery(value) {
      this.query = value
    },
    setSelectedEmployee(employee) {
      this.selectedEmployee = employee
    },
    openModal() {
      this.isModalOpen = true
    },
    closeModal() {
      this.isModalOpen = false
      this.selectedEmployee = null
    },
    openDeleteModal(employee) {
      this.selectedEmployee = employee
      this.isDeleteOpen = true
    },
    closeDeleteModal() {
      this.isDeleteOpen = false
      this.selectedEmployee = null
    },
    setToast(message, type = 'success') {
      this.toast = { message, type }
      setTimeout(() => {
        this.toast = null
      }, 2500)
    },
  },
})
