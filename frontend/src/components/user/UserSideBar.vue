<template>
  <Transition name="slide-fade">
    <div class="w-full h-full">
      <ul class="menu rounded-box">
        <li v-for="(nav, index) in navs" @click="scrollToSection(nav.cid)">
          <a class="">
            <span :class="icons[index]" class="size-5"></span>
            <span v-if="!showSetting.siderBar" class="text-base">{{ nav.title }}</span>
          </a>
        </li>
      </ul>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import router from '@/router'
import { icons } from '@/stores/icons'

const store = useBasicStore()
const { navs, showSetting } = storeToRefs(store)

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
