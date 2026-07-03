import api from './api'

export const leaveAPI = {
  // Create a new leave request
  create: (data) => api.post('/leaves', data),

  // Get all leave requests
  getAll: (params) => api.get('/leaves', { params }),

  // Get a single leave request by ID
  getById: (id) => api.get(`/leaves/${id}`),

  // Update a leave request (approve/reject)
  update: (id, data) => api.put(`/leaves/${id}`, data),

  // Get leave balance
  getBalance: (employeeId, year) => api.get('/leaves/balance', { params: { employeeId, year } }),
}