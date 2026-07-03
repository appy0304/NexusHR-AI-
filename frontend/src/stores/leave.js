import { defineStore } from 'pinia'
import { leaveAPI } from '../services/leaveService'

export const useLeaveStore = defineStore('leave', {
    state: () => ({
        leaves: [],
        currentLeave: null,
        leaveBalance: null,
        loading: false,
        error: null,
    }),

    getters: {
        pendingLeaves: (state) => state.leaves.filter((l) => l.status === 'pending'),
        approvedLeaves: (state) => state.leaves.filter((l) => l.status === 'approved'),
        rejectedLeaves: (state) => state.leaves.filter((l) => l.status === 'rejected'),
    },

    actions: {
        async createLeave(data) {
            this.loading = true
            this.error = null
            try {
                const response = await leaveAPI.create(data)
                return response.data
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async fetchLeaves(params = {}) {
            this.loading = true
            this.error = null
            try {
                const response = await leaveAPI.getAll(params)
                this.leaves = response.data?.data?.items || []
                return this.leaves
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async fetchLeaveById(id) {
            this.loading = true
            this.error = null
            try {
                const response = await leaveAPI.getById(id)
                this.currentLeave = response.data?.data
                return this.currentLeave
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async updateLeave(id, data) {
            this.loading = true
            this.error = null
            try {
                const response = await leaveAPI.update(id, data)
                return response.data
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },

        async fetchLeaveBalance(employeeId, year) {
            this.loading = true
            this.error = null
            try {
                const response = await leaveAPI.getBalance(employeeId, year)
                this.leaveBalance = response.data?.data
                return this.leaveBalance
            } catch (err) {
                this.error = err.message
                throw err
            } finally {
                this.loading = false
            }
        },
    },
})