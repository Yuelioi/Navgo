<template>
  <div class="navbar bg-base-200">
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
        <button class="btn btn-ghost btn-sm" @click="router.push('post')">
          <div class="tooltip tooltip-bottom" data-tip="投稿">
            <span class="icon-[lucide--send] size-5"></span>
          </div>
        </button>
      </div>
    </div>

    <div class="navbar-end">
      <div class="flex items-center space-x-4 mr-4">
        <div class="">这是一句话</div>

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
</template>

<script setup lang="ts">
import router from '@/router'

const store = useBasicStore()
const { collapseNav, theme } = storeToRefs(store)

const { switchTheme } = useTheme()

function changeTheme() {
  theme.value = theme.value == 'light' ? 'dark' : 'light'
  switchTheme(theme.value)
}
</script>
