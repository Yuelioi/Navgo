import type { Collection } from './../api/types'
import { reactive } from 'vue'
import { defineStore } from 'pinia'
import type { Category, CollectionsData } from '@/api/types'

import { useStorage } from '@vueuse/core'

export const useBasicStore = defineStore('basic', () => {
  const collectionsDatas = ref<Array<CollectionsData>>([])
  const collectionsMap: Ref<Map<string, Collection>> = ref(new Map())
  const collectionsList = ref<Collection[]>([])
  const navs = reactive<Category[]>([])

  const currentSearchName = useStorage('currentSearch', '百度')
  const theme = useStorage('theme', 'light')
  const collapseNav = useStorage('collapseNav', false)
  const isScrollDown = ref(false)

  return {
    navs,
    collectionsDatas,
    collectionsList,
    collectionsMap,
    collapseNav,
    theme,
    isScrollDown,
    currentSearchName
  }
})
