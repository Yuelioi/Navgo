<template>
  <Transition name="slide-fade">
    <div v-if="collapseNav" class="bg-base-200 float-left h-full w-56">
      <div class="top h-12 flex items-center justify-center">这是一个logo</div>
      <ul class="menu menu-lg rounded-box">
        <li v-for="nav in navs" @click="scrollToSection(nav.cid)">
          <a>
            <span class="icon-[lucide--zap]"></span>
            {{ nav.title }}
          </a>
        </li>
      </ul>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import router from '@/router'

const store = useBasicStore()
const { navs, collapseNav } = storeToRefs(store)

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
