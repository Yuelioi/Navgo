<template>
  <div class="flex flex-wrap">
    <div class="basis-1/6 p-4 cursor-pointer" v-for="collection in group.collections">
      <!-- 单个链接卡片 -->
      <div class="bg-base-300 shadow-md p-4 hover:ring-2 hover:scale-105 rounded-md">
        <div class="h-28 flex flex-col">
          <div
            class="flex flex-row flex-1 overflow-hidden items-center relative"
            @click="open(collection.link)">
            <img
              :src="'https://cdn.yuelili.com/nav/icons/' + collection.cid + '.png'"
              class="h-full rounded-full" />

            <div class="pl-4">
              <div class="font-bold">{{ collection.title }}</div>
              <div class="line-clamp-1">{{ collection.description }}</div>
            </div>
          </div>

          <div class="border my-2"></div>

          <div class="h-6 flex items-center">
            <div class="badge badge-ghost" v-if="collection.proxy">需要代理</div>
            <div class="" v-for="tag in collection.tags">
              <div class="badge badge-ghost">{{ tag }}</div>
            </div>
            <div
              class="ml-auto tooltip tooltip-bottom"
              data-tip="查看详情"
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
</script>
