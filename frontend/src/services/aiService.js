import api from './api'

export const aiAPI = {
    ask: (query, userId) => api.post('/ai/ask', { query, userId }),
    upload: (id, text) => api.post('/upload', { id, text })
}
