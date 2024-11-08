import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/modules/auth/store/auth'
import AuthLayout from '@/modules/auth/layout/AuthLayout.vue'
import LoginPage from '@/modules/auth/pages/LoginPage.vue'
import RegisterPage from '@/modules/auth/pages/RegisterPage.vue'
import NotFound404 from '@/modules/common/pages/NotFound404.vue'
import ProjectLayout from '@/modules/projects/layouts/ProjectLayout.vue'
import ProjectDetailPage from '@/modules/projects/pages/ProjectDetailPage.vue'
import ProjectFormPage from '@/modules/projects/pages/ProjectFormPage.vue'
import ProjectListPage from '@/modules/projects/pages/ProjectListPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/projects',
      name: 'projectLayout',
      component: ProjectLayout,
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'projectList',
          component: ProjectListPage,
        },
        {
          path: ':id',
          name: 'projectDetail',
          component: ProjectDetailPage,
        },
        {
          path: 'new',
          name: 'projectNew',
          component: ProjectFormPage,
        },
        {
          path: ':id/edit',
          name: 'projectEdit',
          component: ProjectFormPage,
          props: true,
        },
      ],
    },
    {
      path: '/auth',
      redirect: { name: 'login' },
      component: AuthLayout,
      children: [
        {
          path: '/login',
          name: 'login',
          component: LoginPage,
        },
        {
          path: '/register',
          name: 'register',
          component: RegisterPage,
        },
      ],
    },
    {
      path: '/:pathMatch(.*)*',
      component: NotFound404,
    },
  ],
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ name: 'login' })
  } else if (to.name === 'login' && authStore.isAuthenticated) {
    next({ name: 'projectLayout' })
  } else {
    next()
  }
})

export default router
