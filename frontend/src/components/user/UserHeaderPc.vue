<template>
  <div class="navbar">
    <div class="navbar-start space-x-2 ml-4">
      <div class="flex items-center">
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
            <div
              class="flex bg-base-100 flex-col p-6 rounded-lg h-80 w-96 space-y-4 ring-1 ring-primary">
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
        <div class="" @click.stop="showLoginWindow = true" v-if="isAdmin">
          <span class="icon-[lucide--user-round] size-6"></span>
        </div>

        <div class="" @click="open('https://github.com/Yuelioi/Navgo')">
          <span class="icon-[lucide--github] size-6"></span>
        </div>

        <div class="">
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

    <VDialog title="登录" v-model:show="showLoginWindow">
      <div class="bg-base-100 w-96 p-1">
        <label class="input input-bordered flex items-center gap-2">
          <input type="text" class="grow" placeholder="Search" />
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 16 16"
            fill="currentColor"
            class="h-4 w-4 opacity-70">
            <path
              fill-rule="evenodd"
              d="M9.965 11.026a5 5 0 1 1 1.06-1.06l2.755 2.754a.75.75 0 1 1-1.06 1.06l-2.755-2.754ZM10.5 7a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0Z"
              clip-rule="evenodd" />
          </svg>
        </label>
        <label class="input input-bordered flex items-center gap-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 16 16"
            fill="currentColor"
            class="h-4 w-4 opacity-70">
            <path
              d="M2.5 3A1.5 1.5 0 0 0 1 4.5v.793c.026.009.051.02.076.032L7.674 8.51c.206.1.446.1.652 0l6.598-3.185A.755.755 0 0 1 15 5.293V4.5A1.5 1.5 0 0 0 13.5 3h-11Z" />
            <path
              d="M15 6.954 8.978 9.86a2.25 2.25 0 0 1-1.956 0L1 6.954V11.5A1.5 1.5 0 0 0 2.5 13h11a1.5 1.5 0 0 0 1.5-1.5V6.954Z" />
          </svg>
          <input type="text" class="grow" placeholder="Email" />
        </label>
        <label class="input input-bordered flex items-center gap-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 16 16"
            fill="currentColor"
            class="h-4 w-4 opacity-70">
            <path
              d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6ZM12.735 14c.618 0 1.093-.561.872-1.139a6.002 6.002 0 0 0-11.215 0c-.22.578.254 1.139.872 1.139h9.47Z" />
          </svg>
          <input type="text" class="grow" placeholder="Username" />
        </label>
        <label class="input input-bordered flex items-center gap-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 16 16"
            fill="currentColor"
            class="h-4 w-4 opacity-70">
            <path
              fill-rule="evenodd"
              d="M14 6a4 4 0 0 1-4.899 3.899l-1.955 1.955a.5.5 0 0 1-.353.146H5v1.5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2.293a.5.5 0 0 1 .146-.353l3.955-3.955A4 4 0 1 1 14 6Zm-4-2a.75.75 0 0 0 0 1.5.5.5 0 0 1 .5.5.75.75 0 0 0 1.5 0 2 2 0 0 0-2-2Z"
              clip-rule="evenodd" />
          </svg>
          <input type="password" class="grow" value="password" />
        </label>
      </div>
    </VDialog>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'
import { open } from '@/utils'
import { VDialog } from '@/plugins/dialog'

const store = useBasicStore()
const { theme, isAdmin } = storeToRefs(store)

const showHotkeyWindow = ref(false)
const showLoginWindow = ref(false)

const { switchTheme } = useTheme()

function changeTheme() {
  theme.value = theme.value == 'light' ? 'dark' : 'light'
  switchTheme(theme.value)
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
