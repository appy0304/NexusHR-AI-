import api from './api'

const unwrap = (response) => {
  const payload = response?.data?.data ?? response?.data ?? response

  if (Array.isArray(payload)) {
    return payload
  }

  if (payload && typeof payload === 'object' && 'data' in payload) {
    if (Array.isArray(payload.data)) {
      return payload.data
    }
    if (payload.data == null) {
      return []
    }
  }

  return payload?.data ?? payload
}

export const employeeService = {
  async getEmployees(params = {}) {
    const { data } = await api.get('/employees', { params })
    return unwrap(data)
  },

  async getEmployee(id) {
    const { data } = await api.get(`/employees/${id}`)
    return unwrap(data)
  },

  async createEmployee(payload) {
    const { data } = await api.post('/employees', payload)
    return unwrap(data)
  },

  async updateEmployee(id, payload) {
    const { data } = await api.put(`/employees/${id}`, payload)
    return unwrap(data)
  },

  async deleteEmployee(id) {
    const { data } = await api.delete(`/employees/${id}`)
    return unwrap(data)
  },
}