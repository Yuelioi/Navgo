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
        <button class="btn btn-ghost btn-sm" @click="router.push({ name: 'setting' })">
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
        <div class="" @click.stop="showLoginWindow = true">
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
      <div class="p-4">
        <div class="w-96 p-2 space-y-4">
          <label class="input input-bordered flex items-center gap-2">
            <span class="icon-[lucide--user]"></span>
            <input type="text" class="grow" placeholder="Username" />
          </label>
          <label class="input input-bordered flex items-center gap-2">
            <span class="icon-[lucide--key-round]"></span>
            <input type="password" class="grow" value="password" />
          </label>

          <input
            type="button"
            class="w-full grow btn btn-sm btn-primary"
            value="登录"
            @click="login" />
        </div>
      </div>
    </VDialog>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'
import { open } from '@/utils'
import { VDialog } from '@/plugins/dialog'

const store = useBasicStore()
const { theme } = storeToRefs(store)

const showHotkeyWindow = ref(false)
const showLoginWindow = ref(false)

const { switchTheme } = useTheme()

function changeTheme() {
  theme.value = theme.value == 'light' ? 'dark' : 'light'
  switchTheme(theme.value)
}

function login() {}
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
