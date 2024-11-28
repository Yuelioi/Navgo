<template>
  <Transition name="slide-fade">
    <div class="w-full h-full flex flex-col">
      <ul class="menu rounded-box h-full lg:menu-lg">
        <li v-for="(nav, index) in navs" @click="scrollToSection(nav.cid)">
          <a class="">
            <span :class="icons[index]" class="size-5"></span>
            <span class="text-base lg:text-lg">{{ nav.title }}</span>
          </a>
        </li>

        <li class="mt-auto">
          <a>
            <span class="icon-[lucide--github] size-6"></span>
            <span>Github</span>
          </a>
        </li>
        <li>
          <a @click.stop="showLoginWindow = true">
            <span class="icon-[lucide--user-round] size-6"></span>
            <span>登录</span>
          </a>
        </li>

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
      </ul>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import router from '@/router'
import { icons } from '@/stores/icons'

const showLoginWindow = ref(false)
const store = useBasicStore()
const { navs } = storeToRefs(store)

import { VDialog } from '@/plugins/dialog'

function scrollToSection(cid?: string) {
  if (router.currentRoute.value.name != 'home') {
    router.push({
      name: 'home'
    })
  }

  if (cid) {
    const element = document.getElementById(cid)
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' })
    }
  }
}

function login() {}
</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.3s ease-out;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(-20px);
  opacity: 0;
}
</style>
