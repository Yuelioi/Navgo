import type { Collection, SiteStats } from './../api/types'
import { reactive } from 'vue'
import { defineStore } from 'pinia'
import type { Category, CollectionsData } from '@/api/types'

import { useStorage } from '@vueuse/core'
import type { Search, ShowSetting, WallpaperSetting } from '@/models'

import { defaultSearchList, defaultShowSetting, defaultWallpaperSetting } from '@/consts'

export const useBasicStore = defineStore('basic', () => {
  const collectionsDatas = ref<Array<CollectionsData>>([])
  const collectionsMap: Ref<Map<string, Collection>> = ref(new Map())
  const collectionsList = ref<Collection[]>([])
  const likeCollectionsList = ref<Collection[]>([])
  const navs = reactive<Category[]>([])

  const dbVersion = useStorage('dbVersion', 1)
  const currentSearchName = useStorage('currentSearch', '百度')
  const theme = useStorage('theme', 'light')
  const token = useStorage('token', '')

  const searchList = useStorage<Search[]>('searchList', defaultSearchList)
  const showSetting = useStorage<ShowSetting>('showSetting', defaultShowSetting)
  const wallpaperSetting = useStorage<WallpaperSetting>('wallpaperSetting', defaultWallpaperSetting)

  const currentWallpaper = ref('')
  const localWallpaper = ref('')
  const isScrollDown = ref(false)
  const siteStats = ref<SiteStats>()

  return {
    token,
    dbVersion,
    localWallpaper,
    currentWallpaper,
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
    searchList,
    wallpaperSetting
  }
})
