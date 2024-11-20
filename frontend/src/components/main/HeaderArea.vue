<template>
  <div class="navbar bg-base-200 hidden md:flex">
    <div class="navbar-start space-x-2 ml-4">
      <div class="flex items-center">
        <button class="btn btn-ghost btn-sm" v-if="collapseNav" @click="collapseNav = !collapseNav">
          <span class="icon-[lucide--text] size-5"></span>
        </button>
        <button class="btn btn-ghost btn-sm" v-else @click="collapseNav = !collapseNav">
          <span class="icon-[lucide--menu] size-5"></span>
        </button>
        <button class="btn btn-ghost btn-sm" @click="router.push({ name: 'home' })">
          <div class="tooltip tooltip-bottom" data-tip="主页">
            <span class="icon-[lucide--house] size-5"></span>
          </div>
        </button>
        <button class="btn btn-ghost btn-sm" @click="router.push({ name: 'post' })">
          <div class="tooltip tooltip-bottom" data-tip="投稿">
            <span class="icon-[lucide--send] size-5"></span>
          </div>
        </button>

        <button class="btn btn-ghost btn-sm" @click="router.push({ name: 'comment' })">
          <div class="tooltip tooltip-bottom" data-tip="留言板">
            <span class="icon-[lucide--clipboard-pen] size-5"></span>
          </div>
        </button>
        <button
          class="btn btn-ghost btn-sm"
          @click="router.push({ name: 'setting' })"
          @click.ctrl="isAdmin = !isAdmin">
          <div class="tooltip tooltip-bottom" data-tip="设置">
            <span class="icon-[lucide--settings] size-5"></span>
          </div>
        </button>
        <button
          class="btn btn-ghost btn-sm relative hidden md:block"
          @mousemove="showHotkeyWindow = true"
          @mouseleave="showHotkeyWindow = false">
          <div class="tooltip tooltip-bottom" data-tip="帮助">
            <span class="icon-[lucide--message-circle-question] size-5"></span>
          </div>
        </button>

        <Transition>
          <div class="absolute z-10 top-24" v-if="showHotkeyWindow">
            <div class="flex bg-base-200 flex-col p-6 rounded-lg h-80 w-96 space-y-4 ring-2">
              <div class="font-bold"><span>帮助</span></div>
              <div class="divider"></div>
              <div class="flex items-center w-full justify-between">
                <kbd class="kbd">Tab</kbd>
                <div class="kbd">切换搜索引擎</div>
              </div>
              <div class="flex items-center w-full justify-between">
                <kbd class="kbd">Esc</kbd>
                <div class="kbd">删除搜索关键词</div>
              </div>
              <div class="">单击收藏 / 再次单击取消收藏</div>
            </div>
          </div>
        </Transition>
      </div>
    </div>

    <div class="navbar-end">
      <div class="flex items-center space-x-4 mr-4">
        <div class="hidden lg:block mr-2">这是一句话</div>

        <span
          class="icon-[lucide--github] size-6"
          @click="open('https://github.com/Yuelioi/Navgo')"></span>

        <label class="swap swap-rotate">
          <input
            type="checkbox"
            class="theme-controller"
            value="synthwave"
            @click="changeTheme()"
            :checked="theme == 'dark'" />

          <span class="icon-[lucide--sun] size-6 swap-on fill-current"></span>
          <span class="icon-[lucide--moon] size-6 swap-off fill-current"></span>
        </label>
      </div>
    </div>
  </div>

  <!-- 手机端 -->
  <div class="btm-nav z-10 md:hidden">
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
      <span class="btm-nav-label">返回顶部</span>
    </button>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'
import { open } from '@/utils'

const store = useBasicStore()
const { collapseNav, theme, isAdmin } = storeToRefs(store)

const showHotkeyWindow = ref(false)

const { switchTheme } = useTheme()

function changeTheme() {
  theme.value = theme.value == 'light' ? 'dark' : 'light'
  switchTheme(theme.value)
}

function scrollToTop() {
  document.querySelector('.anchor')?.scrollTo({
    top: 0,
    behavior: 'smooth'
  })
}
</script>

<style scoped>
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
