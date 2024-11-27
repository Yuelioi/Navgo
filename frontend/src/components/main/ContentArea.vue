<template>
  <!-- 主界面内容区域 -->
  <!-- 需要给个锚点`anchor` 用于返回顶部 -->
  <!-- pb-20 给手机端菜单预留空间 -->
  <div
    :class="{ hidden: !showSetting.siderBar && isMobileDevice() }"
    class="anchor flex-1 w-full flex flex-col md:pb-0 overflow-y-scroll">
    <router-view
      class="content py-8 lg:py-12"
      :class="{ 'mb-4': showSetting.footer }"></router-view>
    <FooterArea />
    <!-- <img
      src="/wallpaper.png"
      alt=""
      class="w-full h-full fixed pointer-events-none"
      :style="{ opacity: opacity }" /> -->
  </div>
</template>

<script setup lang="ts">
import { isMobileDevice } from '@/utils'

const store = useBasicStore()
const { showSetting, opacity } = storeToRefs(store)
</script>
