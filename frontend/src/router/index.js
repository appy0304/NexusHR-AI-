import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const routes = [
  {
    path: '/login',
    name: 'login',
    component: () => import('../pages/LoginPage.vue'),
    meta: { guest: true },
  },
  {
    path: '/',
    name: 'dashboard',
    component: () => import('../pages/DashboardPage.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/users',
    name: 'users',
    component: () => import('../pages/EmployeeDirectory.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/analytics',
    name: 'analytics',
    component: () => import('../pages/AnalyticsPage.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/leaves',
    name: 'leaves',
    component: () => import('../pages/LeavePage.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/ai-chat',
    name: 'ai-chat',
    component: () => import('../pages/AiChat.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/knowledge-base',
    name: 'knowledge-base',
    component: () => import('../pages/KnowledgeBase.vue'),
    meta: { requiresAuth: true },
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Navigation guard — require authentication
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  // Load token from localStorage on page refresh
  if (!authStore.accessToken) {
    authStore.loadFromStorage()
  }

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
  } else if (to.meta.guest && authStore.isAuthenticated && to.path === '/login') {
    next('/')
  } else {
    next()
  }
})

export default router
