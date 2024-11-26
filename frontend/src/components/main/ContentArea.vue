<template>
  <!-- 主界面内容区域 -->
  <!-- 需要给个锚点`anchor` 用于返回顶部 -->
  <!-- pb-20 给手机端菜单预留空间 -->
  <div
    :class="{ hidden: !isCollapsedNav && isMobileDevice() }"
    class="anchor flex-1 w-full flex flex-col overflow-y-scroll pb-20 md:pb-0">
    <router-view class="content"></router-view>
    <FooterArea />
  </div>
</template>

<script setup lang="ts">
import { isMobileDevice } from '@/utils'

const store = useBasicStore()
const { isCollapsedNav } = storeToRefs(store)
</script>
