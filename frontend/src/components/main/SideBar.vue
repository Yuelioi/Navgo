<template>
  <Transition name="slide-fade">
    <div v-if="showNav" class="w-48 bg-base-100 h-full">
      <div class="top h-12"></div>
      <ul class="menu menu-lg rounded-box">
        <li v-for="nav in navs">
          <a>
            <span class="icon-[lucide--carrot]"></span>
            {{ nav.title }}
          </a>
        </li>
      </ul>
    </div>
  </Transition>
</template>

<script setup lang="ts">
const store = useBasicStore()
const { navs, showNav } = storeToRefs(store)
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
