import { defineStore } from 'pinia'
import { authAPI } from '../services/auth'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    accessToken: null,
    isAuthenticated: false,
  }),

  actions: {
    async login(email, password) {
      const response = await authAPI.login(email, password)
      this.accessToken = response.accessToken
      this.user = response.user
      this.isAuthenticated = true
      localStorage.setItem('accessToken', response.accessToken)
    },

    logout() {
      this.accessToken = null
      this.user = null
      this.isAuthenticated = false
      localStorage.removeItem('accessToken')
    },

    loadFromStorage() {
      const token = localStorage.getItem('accessToken')
      if (token) {
        this.accessToken = token
        // Parse user info from token (base64 decode)
        try {
          const payload = JSON.parse(atob(token.split('.')[1]))
          this.user = {
            id: payload.user_id,
            email: payload.email,
            role: payload.role,
          }
          this.isAuthenticated = true
        } catch {
          this.logout()
        }
      }
    },

    hasRole(role) {
      return this.user && this.user.role === role
    },

    hasAnyRole(roles) {
      return this.user && roles.includes(this.user.role)
    },
  },
})
