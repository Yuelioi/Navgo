import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import SiteView from '../views/SiteView.vue'
import PostView from '@/views/PostView.vue'
import LikeView from '@/views/LikeView.vue'
import CommentView from '@/views/CommentView.vue'

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
      component: SiteView
    },
    {
      path: '/post',
      name: 'post',
      component: PostView
    },
    {
      path: '/like',
      name: 'like',
      component: LikeView
    },
    {
      path: '/comment',
      name: 'comment',
      component: CommentView
    }
  ]
})

router.beforeEach(async (to, from, next) => {
  try {
    await loadData()

    next()
  } catch (error) {}
})

export default router
