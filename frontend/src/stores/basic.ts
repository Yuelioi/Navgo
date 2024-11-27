import type { Collection, SiteStats } from './../api/types'
import { reactive } from 'vue'
import { defineStore } from 'pinia'
import type { Category, CollectionsData } from '@/api/types'

import { useStorage } from '@vueuse/core'
import type { Search } from '@/models/search'

import { defaultSearchList } from '@/consts/search'
import type { ShowSetting } from '@/models/show'
import { defaultShowSetting } from '@/consts/show'

export const useBasicStore = defineStore('basic', () => {
  const collectionsDatas = ref<Array<CollectionsData>>([])
  const collectionsMap: Ref<Map<string, Collection>> = ref(new Map())
  const collectionsList = ref<Collection[]>([])
  const likeCollectionsList = ref<Collection[]>([])
  const navs = reactive<Category[]>([])

  const dbVersion = useStorage('currentSearch', 1)
  const currentSearchName = useStorage('currentSearch', '百度')
  const theme = useStorage('theme', 'light')

  const token = useStorage('token', '')

  const showSetting = useStorage<ShowSetting>('showSetting', defaultShowSetting)

  const isAdmin = useStorage('isAdmin', false)
  const opacity = useStorage('opacity', 0.5)

  const isScrollDown = ref(false)

  const siteStats = ref<SiteStats>()

  const searchList = useStorage<Search[]>('searchList', defaultSearchList)

  return {
    opacity,
    token,
    isAdmin,
    dbVersion,
    navs,
    collectionsDatas,
    collectionsList,
    collectionsMap,
    showSetting,
    theme,
    isScrollDown,
    currentSearchName,
    likeCollectionsList,
    siteStats,
    searchList
  }
})
