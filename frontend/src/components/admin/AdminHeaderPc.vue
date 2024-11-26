<template>
  <div class="navbar">
    <div class="navbar-start space-x-2 ml-4">
      <div class="flex items-center">
        <button class="btn btn-ghost btn-sm" @click="router.push({ name: 'setting' })">
          <div class="tooltip tooltip-bottom" data-tip="设置">
            <span class="icon-[lucide--settings] size-5"></span>
          </div>
        </button>
      </div>
    </div>

    <div class="navbar-end">
      <div class="flex items-center space-x-4 mr-4">
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
  </div>
</template>

<script setup lang="ts">
import router from '@/router'

const store = useBasicStore()
const { theme } = storeToRefs(store)

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
