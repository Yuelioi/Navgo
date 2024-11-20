import type { Collection, SiteStats } from './../api/types'
import { reactive } from 'vue'
import { defineStore } from 'pinia'
import type { Category, CollectionsData } from '@/api/types'

import { useStorage } from '@vueuse/core'

export const useBasicStore = defineStore('basic', () => {
  const collectionsDatas = ref<Array<CollectionsData>>([])
  const collectionsMap: Ref<Map<string, Collection>> = ref(new Map())
  const collectionsList = ref<Collection[]>([])
  const likeCollectionsList = ref<Collection[]>([])
  const navs = reactive<Category[]>([])

  const dbVersion = useStorage('currentSearch', 1)
  const currentSearchName = useStorage('currentSearch', '百度')
  const theme = useStorage('theme', 'light')
  const collapseNav = useStorage('collapseNav', false)
  const isAdmin = useStorage('isAdmin', false)
  const showMyCollection = useStorage('showMyCollection', true)
  const showWebCollection = useStorage('showWebCollection', true)
  const showFooter = useStorage('showFooter', true)

  const isScrollDown = ref(false)

  const siteStats = ref<SiteStats>()

  return {
    isAdmin,
    dbVersion,
    navs,
    collectionsDatas,
    collectionsList,
    collectionsMap,
    collapseNav,
    theme,
    isScrollDown,
    currentSearchName,
    likeCollectionsList,
    showMyCollection,
    showWebCollection,
    siteStats,
    showFooter
  }
})
