<template>
  <div
    class="text-neutral-content p-2 w-full flex items-center bg-gradient-to-r from-neutral to-transparent rounded-t-md">
    <span class="icon-[lucide--star] size-5 ml-4 mr-2"></span>
    <span class="font-bold text-lg">{{ title }}</span>
  </div>
</template>

<script setup lang="ts">
const { title } = defineProps(['title'])
</script>
