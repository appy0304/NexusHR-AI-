import api from './api'

export const authAPI = {
  async login(email, password) {
    const { data } = await api.post('/auth/login', { email, password })
    return data.data
  },

  async refresh(refreshToken) {
    const { data } = await api.post('/auth/refresh', { refreshToken })
    return data.data
  },

  async logout(refreshToken) {
    const { data } = await api.post('/auth/logout', { refreshToken })
    return data
  },

  async createAdmin(email, password, name) {
    const { data } = await api.post('/auth/create-admin', { email, password, name })
    return data
  },
}
