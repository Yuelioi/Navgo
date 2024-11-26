import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '@/views/HomeView.vue'

import { loadData } from '@/stores/data'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        loadData: true
      }
    },
    {
      path: '/post',
      name: 'post',
      component: () => import('@/views/PostView.vue'),
      meta: {
        loadData: true
      }
    },

    {
      path: '/comment',
      name: 'comment',
      component: () => import('@/views/CommentView.vue'),
      meta: {
        loadData: true
      }
    },
    {
      path: '/setting',
      name: 'setting',
      component: () => import('@/views/SettingView.vue'),
      meta: {
        loadData: true
      }
    },
    {
      path: '/admin',
      children: [
        {
          path: 'nav',
          children: [
            {
              path: 'all',
              name: 'allNavs',
              component: () => import('@/components/admin/nav/AllNavs.vue')
            },
            {
              path: 'add',
              name: 'addNav',
              component: () => import('@/components/admin/nav/AddNav.vue')
            },
            {
              path: 'category',
              name: 'navCategory',
              component: () => import('@/components/admin/nav/NavCategory.vue')
            }
          ]
        },
        {
          path: 'announce',
          name: 'announce',
          component: () => import('@/components/admin/AdminAnnounce.vue')
        },
        {
          path: 'approval',
          name: 'approval',
          component: () => import('@/components/admin/AdminApproval.vue')
        },
        {
          path: 'comment',
          name: 'adminComment',
          component: () => import('@/components/admin/AdminComment.vue')
        }
      ]
    }
  ]
})

router.beforeEach(async (to, from, next) => {
  try {
    const basicStore = useBasicStore()

    const { isAdmin } = storeToRefs(basicStore)

    if (/^\/admin/.test(to.fullPath)) {
      // TODO 判断是否有权限
      isAdmin.value = true
    } else {
      isAdmin.value = false
    }

    if (to.meta['loadData']) {
      await loadData()
    }

    next()
  } catch (error) {}
})

export default router
