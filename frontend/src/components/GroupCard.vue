<template>
  <div class="flex flex-wrap relative">
    <div class="basis-1/5 p-4 cursor-pointer" v-for="collection in group.collections">
      <!-- 单个链接卡片 -->
      <div
        class="bg-base-300 shadow-md p-4 group hover:ring-2 ring-primary hover:scale-105 rounded-md">
        <div class="h-28 flex flex-col">
          <div class="flex flex-row flex-1 items-center relative" @click="open(collection.link)">
            <div class="avatar static size-12">
              <div class="h-full rounded-xl">
                <img
                  :src="'https://cdn.yuelili.com/nav/icons/' + collection.cid + '.png'"
                  @error="imageLoadError"
                  class="h-full rounded-full" />
                <div class="rounded-full flex h-full items-center justify-center bg-purple-200">
                  <span class="font-bold text-lg">{{ collection.title[0] }}</span>
                </div>
              </div>
            </div>

            <div class="pl-4 flex-1">
              <div class="items-center flex space-x-2">
                <span class="font-bold">{{ collection.title }}</span>
                <span class="items-center group-hover:flex hidden space-x-1">
                  <span class="icon-[ep--view] size-3"></span>
                  <span class="text-sm">{{ collection.view }}</span>
                </span>
              </div>

              <div class="tooltip static tooltip-bottom" :data-tip="collection.description">
                <div class="flex text-left">
                  <span class="line-clamp-1">{{ collection.description }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="border my-2"></div>

          <div class="h-6 flex items-center">
            <div class="badge badge-ghost" v-if="collection.proxy">需要代理</div>
            <div class="" v-for="tag in collection.tags">
              <div class="badge badge-ghost">{{ tag }}</div>
            </div>
            <div
              class="ml-auto"
              @click="
                router.push({
                  name: 'site',
                  params: { id: collection.cid }
                })
              ">
              <span class="icon-[lucide--square-arrow-right] size-5"></span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { open } from '@/utils'
import router from '@/router'

const props = defineProps<{
  group: any
}>()

function imageLoadError(event: Event) {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}
</script>
<style scoped>
.tooltip:before {
  --tooltip-color: rgb(55 55 55);
}
</style>
