<template>
  <div class="h-full grid grid-cols-one justify-center select-none mb-8">
    <div class="pt-12 overflow-y-scroll">
      <div role="alert" class="alert alert-info flex my-8 rounded-none">
        <span class="icon-[lucide--info] size-5"></span>
        <div class="flex w-full">
          <span class="font-bold transition-transform">注意</span>
          <span class="ml-8">
            个人收藏只会存在本地浏览器,请勿删除当前页面浏览器数据 (虽然一般也不会删除)
          </span>
        </div>
      </div>

      <div class="bg-base-200">
        <div class="p-4">
          <div class="flex w-full gap-4 items-center">
            <div class="btn btn-neutral btn-sm btn-ghost" @click="addLike">
              <span class="icon-[lucide--square-plus] size-5"></span>
            </div>
            <div class="btn btn-primary btn-sm btn-ghost" @click="saveCollections()">
              <span class="icon-[lucide--save] size-6"></span>
            </div>
            <div
              class="ml-auto flex items-center space-x-2 mr-4 bg-accent rounded-md px-2 py-1"
              v-if="isModify">
              <div class=""><span class="icon-[lucide--speech] size-5"></span></div>
              <div class="font-bold">尚未保存</div>
            </div>
          </div>
        </div>
      </div>
      <div class="border-2 border-base-100"></div>
      <table class="table bg-base-200 rounded-none">
        <!-- head -->
        <thead>
          <tr class="font-bold">
            <th class="font-bold">图标</th>
            <th class="font-bold">标题</th>
            <th class="font-bold">链接</th>
            <th class="font-bold">描述</th>
            <th class="font-bold">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(collection, index) in likeCollectionsList">
            <td>
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
            </td>
            <td>
              <input type="text" class="input w-full" v-model="collection.title" />
            </td>
            <td>
              <input type="text" class="input w-full" v-model="collection.link" @invalid="" />
            </td>
            <td><textarea type="text" class="input w-full" v-model="collection.description" /></td>
            <th class="w-min">
              <div class="flex max-w-fit space-x-1">
                <button class="btn btn-neutral btn-sm btn-ghost" @click="removeCollection(index)">
                  <span class="icon-[lucide--square-x] size-6"></span>
                </button>
              </div>
            </th>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type Collection } from '@/api'
import { db } from '@/db/db'

const isModify = ref(false)

const store = useBasicStore()
const { likeCollectionsList } = storeToRefs(store)

function imageLoadError(event: Event) {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}

function removeCollection(id: number) {
  likeCollectionsList.value.splice(id, 1)
  isModify.value = true
}

async function saveCollections() {
  await db.clearData('likes')

  likeCollectionsList.value.forEach(async (collection: Collection) => {
    const parsedUrl = new URL(collection.link)

    if (parsedUrl.hostname != '') {
      const serializableCollection = {
        id: parsedUrl.hostname,
        cid: parsedUrl.hostname,
        title: collection.title,
        description: collection.description,
        proxy: collection.proxy,
        link: collection.link,
        favicon: collection.favicon
      }

      await db.addData('likes', serializableCollection)
    }
  })
  isModify.value = false
}

function addLike() {
  if (likeCollectionsList.value.length < 20) {
    likeCollectionsList.value.push({
      cid: '',
      title: '',
      link: ''
    })
    isModify.value = true
  } else {
    Message({ message: '最多添加20个收藏喔!' })
  }
}
</script>
