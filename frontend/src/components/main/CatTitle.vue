<template>
  <div
    class="text-neutral-content p-2 w-full flex items-center bg-gradient-to-r from-neutral to-transparent rounded-t-md">
    <span class="size-5 ml-4 mr-2" :class="icon"></span>
    <span class="font-bold text-lg">{{ title }}</span>
  </div>
</template>

<script setup lang="ts">
const { title, icon } = defineProps(['title', 'icon'])
</script>
