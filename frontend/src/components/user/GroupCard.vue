<template>
  <div class="p-2 md:p-4 cursor-pointer" v-for="collection in collections.slice(0, max)">
    <!-- 单个链接卡片 -->
    <div class="shadow-lg backdrop-blur-md px-4 md:py-4 group hover:ring-2 ring-primary rounded-md">
      <div class="h-20 md:h-28 flex flex-col">
        <div class="flex flex-row flex-1 items-center relative" @click="open(collection.link)">
          <div class="avatar static size-12">
            <div class="h-full rounded-xl">
              <img
                :alt="collection.title"
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
              <span class="font-bold line-clamp-1">{{ collection.title }}</span>
            </div>

            <!-- 手机端 -->
            <div
              class="size-12 md:hidden mr-0 flex justify-center rounded-full items-center border"
              @click.stop="like(collection)">
              <span class="icon-[lucide--star] bg-slate-950 size-5"></span>
            </div>

            <!-- PC端描述 -->
            <div
              class="tooltip static tooltip-bottom hidden md:block"
              :data-tip="collection.description">
              <div class="flex text-left text-sm opacity-35">
                <span class="line-clamp-1">{{ collection.description }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="border border-neutral-content/50 hidden md:block my-2"></div>

        <!-- PC端 -->
        <div class="h-6 md:flex items-center hidden">
          <div class="" v-for="tag in collection.tags">
            <div class="badge badge-ghost opacity-50">{{ tag }}</div>
          </div>

          <div class="ml-auto space-x-2">
            <span
              v-if="collection.like"
              class="icon-[line-md--star-filled] bg-warning size-5"
              @click="like(collection)"></span>
            <span v-else class="icon-[lucide--star] size-5" @click="like(collection)"></span>
            <!-- <span
              class="icon-[lucide--square-arrow-right] size-5"
              @click="
                router.push({
                  name: 'site',
                  params: { id: collection.cid }
                })
              "></span> -->
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="col-start-1 col-span-full text-center w-full my-4"
    v-if="collections.length > max"
    @click="max = max + limit">
    <div class="btn btn-outline btn-sm">显示更多</div>
  </div>
</template>

<script setup lang="ts">
import { open } from '@/utils'
import router from '@/router'
import { db } from '@/db/db'
import type { Collection } from '@/api'

import { imageLoadError, isMobileDevice } from '@/utils'

const store = useBasicStore()

let limit = 5

if (!isMobileDevice()) {
  limit = limit * 2
}
const max = ref(limit)

const { likeCollectionsList } = storeToRefs(store)

const props = defineProps<{
  collections: Collection[]
}>()

async function like(collection: Collection) {
  const id = likeCollectionsList.value.findIndex((ele: Collection) => {
    return ele.cid === collection.cid
  })

  if (id == -1) {
    // 将 collection 转换为可序列化的对象
    const serializableCollection = {
      id: collection.cid,
      cid: collection.cid,
      title: collection.title,
      description: collection.description,
      link: collection.link,
      favicon: collection.favicon,
      like: true
    }
    collection.like = true

    likeCollectionsList.value.push(collection)

    await db.addData('likes', serializableCollection)
  } else {
    likeCollectionsList.value.splice(id, 1)
    collection.like = false
    await db.deleteData('likes', collection.cid)
  }
}
</script>
<style scoped>
.tooltip:before {
  --tooltip-color: rgb(55 55 55);
}
</style>
