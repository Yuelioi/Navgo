<template>
  <!-- 主界面内容区域 -->
  <!-- 需要给个锚点`anchor` 用于返回顶部 -->
  <!-- pb-20 给手机端菜单预留空间 -->
  <div
    :class="{ hidden: !showSetting.siderBar && isMobileDevice() }"
    class="anchor flex-1 w-full flex flex-col md:pb-0 overflow-y-scroll">
    <router-view
      class="content md:py-12 pb-20"
      :class="{ 'mb-4': showSetting.footer }"></router-view>
    <FooterArea v-if="!isAdmin" />
    <img
      v-if="wallpaperSetting.show && currentWallpaper"
      :src="currentWallpaper"
      alt="Loading..."
      class="w-full h-full fixed pointer-events-none object-cover"
      :style="{ opacity: wallpaperSetting.opacity }" />

    <img
      @click="scrollToTop"
      src="https://cdn.yuelili.com/web/assets/christmas.gif"
      class="absolute hidden md:size-32 bottom-12 right-8 z-20"
      alt="" />
  </div>
</template>

<script setup lang="ts">
import { wallpaper } from '@/api'
import { db } from '@/db/db'
import { isMobileDevice } from '@/utils'
import { computedAsync } from '@vueuse/core'

const store = useBasicStore()
const { showSetting, wallpaperSetting, localWallpaper } = storeToRefs(store)

const currentWallpaper = computedAsync(async () => {
  if (wallpaperSetting.value.useLocal) {
    return localWallpaper.value
  } else {
    return '' + (await wallpaper())['data']['data']['id']
  }
})

async function loadLocalWallpaper() {
  if (wallpaperSetting.value.useLocal) {
    const data = await db.queryData('wallpaper', 'wallpaper')

    if (data) {
      localWallpaper.value = data['data']
    }
  }
}

function scrollToTop() {
  document.querySelector('.anchor')?.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}

import { useAdminStatus } from '@/hooks/useAdminStatus'
const { isAdmin } = useAdminStatus()

onMounted(() => {
  loadLocalWallpaper()
})
</script>
