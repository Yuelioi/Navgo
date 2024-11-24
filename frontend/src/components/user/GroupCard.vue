<template>
  <div class="p-2 md:p-4 cursor-pointer" v-for="collection in collections">
    <!-- 单个链接卡片 -->
    <div
      class="bg-base-100 shadow-md px-4 md:py-4 group hover:ring-2 ring-primary hover:scale-105 rounded-md">
      <div class="h-20 md:h-28 flex flex-col">
        <div class="flex flex-row flex-1 items-center relative" @click="open(collection.link)">
          <div class="avatar static size-12">
            <div class="h-full rounded-xl">
              <img
                :src="'https://cdn.yuelili.com/nav/icons/' + collection.cid + '.png'"
                @error="imageLoadError"
                class="h-full rounded-full" />
              <div class="rounded-full flex h-full items-center justify-center bg-primary/70">
                <span class="font-bold text-lg">{{ collection.title[0] }}</span>
              </div>
            </div>
          </div>

          <div class="pl-4 flex-1 flex md:flex-col justify-between">
            <div class="items-center flex space-x-2">
              <span class="font-bold">{{ collection.title }}</span>
            </div>

            <!-- 手机端 -->
            <div
              class="size-12 md:hidden mr-0 flex justify-center rounded-full items-center border"
              @click.stop="like(collection)">
              <span class="icon-[lucide--star] size-5"></span>
            </div>

            <div
              class="tooltip static tooltip-bottom hidden md:block"
              :data-tip="collection.description">
              <div class="flex text-left">
                <span class="line-clamp-1">{{ collection.description }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="border hidden md:block my-2"></div>

        <!-- PC端 -->
        <div class="h-6 md:flex items-center hidden">
          <div class="" v-for="tag in collection.tags">
            <div class="badge badge-ghost">{{ tag }}</div>
          </div>

          <div class="ml-auto space-x-2">
            <span class="icon-[lucide--star] size-5" @click="like(collection)"></span>
            <span
              v-if="isAdmin"
              class="icon-[lucide--square-arrow-right] size-5"
              @click="
                router.push({
                  name: 'site',
                  params: { id: collection.cid }
                })
              "></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { open } from '@/utils'
import router from '@/router'
import { db } from '@/db/db'
import type { Collection } from '@/api'

import { imageLoadError } from '@/utils'

const store = useBasicStore()

const { likeCollectionsList, isAdmin } = storeToRefs(store)

const props = defineProps<{
  collections: Collection[]
}>()

async function like(collection: Collection) {
  const id = likeCollectionsList.value.findIndex((ele) => {
    return ele.cid === collection.cid
  })
  if (id == -1) {
    likeCollectionsList.value.push(collection)

    // 将 collection 转换为可序列化的对象
    const serializableCollection = {
      id: collection.cid,
      cid: collection.cid,
      title: collection.title,
      description: collection.description,
      link: collection.link,
      favicon: collection.favicon
    }

    await db.addData('likes', serializableCollection)
  } else {
    likeCollectionsList.value.splice(id, 1)
    await db.deleteData('likes', collection.cid)
  }
}
</script>
<style scoped>
.tooltip:before {
  --tooltip-color: rgb(55 55 55);
}
</style>
