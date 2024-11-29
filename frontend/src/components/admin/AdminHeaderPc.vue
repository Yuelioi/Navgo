<template>
  <div class="navbar">
    <div class="navbar-start space-x-2 ml-4">
      <div class="flex items-center">
        <!-- LOGO + 折叠 -->
        <button
          class="btn btn-ghost btn-sm"
          v-if="!showSetting.siderBar"
          aria-label="折叠"
          @click="showSetting.siderBar = !showSetting.siderBar">
          <div class="tooltip tooltip-bottom" data-tip="折叠">
            <span class="icon-[lucide--menu] size-5"></span>
          </div>
        </button>
        <button
          v-if="!showSetting.siderBar"
          class="btn btn-ghost btn-sm"
          aria-label="主页"
          @click="router.push({ name: 'home' })">
          <div class="tooltip tooltip-bottom text-lg" data-tip="主页">月离导航</div>
        </button>

        <button
          v-if="showSetting.siderBar"
          class="btn btn-ghost btn-sm"
          aria-label="主页"
          @click="router.push({ name: 'home' })">
          <span class="icon-[lucide--home] size-6"></span>
        </button>
      </div>
    </div>

    <div class="navbar-end">
      <div class="flex items-center space-x-4 mr-4">
        <div class="">
          <button class="btn btn-ghost btn-sm" @click="router.push({ name: 'setting' })">
            <div class="tooltip tooltip-bottom" data-tip="设置">
              <span class="icon-[lucide--settings] size-6"></span>
            </div>
          </button>
          <label class="swap swap-rotate">
            <input
              type="checkbox"
              class="theme-controller"
              value="synthwave"
              @click="changeTheme()"
              :checked="theme == 'dark'" />

            <span class="icon-[lucide--sun] size-5 swap-on fill-current"></span>
            <span class="icon-[lucide--moon] size-0.5 swap-off fill-current"></span>
          </label>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import router from '@/router'

const store = useBasicStore()
const { theme, showSetting } = storeToRefs(store)

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
