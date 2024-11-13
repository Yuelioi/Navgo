import type { Collection } from './../api/types'
import { reactive } from 'vue'
import { defineStore } from 'pinia'
import type { Category, CollectionsData } from '@/api/types'

import { useStorage } from '@vueuse/core'

export const useBasicStore = defineStore('basic', () => {
  const collectionsDatas = ref<Array<CollectionsData>>([])
  const collections: Ref<Map<string, Collection>> = ref(new Map())
  const navs = reactive<Category[]>([])

  const theme = useStorage('theme', 'light')
  const collapseNav = useStorage('collapseNav', false)
  const isScrollDown = ref(false)

  return { navs, collectionsDatas, collections, collapseNav, theme, isScrollDown }
})
