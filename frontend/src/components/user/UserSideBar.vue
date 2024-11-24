<template>
  <Transition name="slide-fade">
    <div class="bg-base-100 float-left w-full h-full">
      <ul class="menu menu-lg rounded-box">
        <li v-for="(nav, index) in navs" @click="scrollToSection(nav.cid)">
          <a class="">
            <span :class="icons[index]"></span>
            <span v-if="!isCollapsedNav">{{ nav.title }}</span>
          </a>
        </li>
      </ul>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import router from '@/router'

const store = useBasicStore()
const { navs, isCollapsedNav } = storeToRefs(store)

const icons = [
  'icon-[lucide--box]',
  'icon-[lucide--sparkle]',
  'icon-[lucide--youtube]',
  'icon-[lucide--square-code]',
  'icon-[lucide--leaf]',
  'icon-[lucide--briefcase-business]',
  'icon-[lucide--plane]'
]

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
