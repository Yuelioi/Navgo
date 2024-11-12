import { reactive } from 'vue'
import { defineStore } from 'pinia'
import type { Category, CollectionsData } from '@/api/types'

export const useBasicStore = defineStore('basic', () => {
  const collectionsDatas = ref<Array<CollectionsData>>([])
  const navs = reactive<Category[]>([])
  const showNav = ref(true)

  return { navs, collectionsDatas, showNav }
})
