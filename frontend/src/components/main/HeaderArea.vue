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
        <button class="btn btn-ghost btn-sm">
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
          <!-- this hidden checkbox controls the state -->
          <input
            type="checkbox"
            class="theme-controller"
            value="synthwave"
            @click="changeTheme()" />

          <!-- sun icon -->
          <span class="icon-[lucide--sun] swap-off size-6 fill-current"></span>

          <!-- moon icon -->
          <span class="icon-[lucide--moon] swap-on size-6 fill-current"></span>
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const store = useBasicStore()
const { collapseNav, theme } = storeToRefs(store)

const { switchTheme } = useTheme()

function changeTheme() {
  theme.value = theme.value == 'light' ? 'dark' : 'light'
  switchTheme(theme.value)
}
</script>
