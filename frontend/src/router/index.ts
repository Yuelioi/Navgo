import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import { loadData } from '@/stores/data'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/site/:id',
      name: 'site',
      component: () => import('../views/SiteView.vue')
    }
  ]
})

router.beforeEach(async (to, from, next) => {
  try {
    await loadData() // 等待数据加载完成
    next() // 数据加载完成后再进入页面
  } catch (error) {
    next('/loading') // 如果加载失败，可以重定向到loading页面
  }
})

export default router
