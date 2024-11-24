<template>
  <!-- 手机端 -->
  <div class="btm-nav z-10">
    <button class="" @click="router.push({ name: 'home' })">
      <span class="icon-[lucide--house] size-5"></span>
      <span class="btm-nav-label">主页</span>
    </button>
    <button class="" @click="router.push({ name: 'post' })">
      <span class="icon-[lucide--send] size-5"></span>
      <span class="btm-nav-label">投稿</span>
    </button>
    <button class="" @click="router.push({ name: 'comment' })">
      <span class="icon-[lucide--clipboard-pen] size-5"></span>
      <span class="btm-nav-label">讨论</span>
    </button>
    <button class="" @click="router.push({ name: 'setting' })" @click.ctrl="isAdmin = !isAdmin">
      <span class="icon-[lucide--settings] size-5"></span>
      <span class="btm-nav-label">设置</span>
    </button>
    <button class="" @click="scrollToTop">
      <span class="icon-[lucide--arrow-up-from-line] size-5"></span>
      <span class="btm-nav-label">顶部</span>
    </button>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'

const store = useBasicStore()
const { theme, isAdmin } = storeToRefs(store)

const { switchTheme } = useTheme()

function scrollToTop() {
  document.querySelector('.anchor')?.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}
</script>
