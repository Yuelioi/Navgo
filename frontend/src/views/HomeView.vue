<template>
  <div class="flex flex-col relative container z-10">
    <div class="self-center flex-1 my-8 w-full">
      <!-- 公告 -->
      <div role="alert"
        class="alert flex w-full backdrop-blur-sm shadow-md border-1 overflow-hidden border-opacity-25 bg-base-200/20"
        v-if="currentAnnounce && showSetting.announce">
        <span class="icon-[lucide--info] size-5 hidden md:block"></span>
        <Transition name="goon" mode="out-in">
          <div class="flex w-full" :key="currentAnnounce.content">
            <span class="font-bold transition-transform line-clamp-1">
              {{ currentAnnounce.title }}
            </span>
            <span class="mx-2 md:ml-8 line-clamp-1">{{ currentAnnounce.content }}</span>
            <span class="ml-auto hidden md:block">{{ currentAnnounce.date }}</span>
          </div>
        </Transition>
      </div>

      <!-- 搜索 -->
      <AsyncHomeSearch class="mt-12 md:mt-24 mb-16 z-50 relative"></AsyncHomeSearch>

      <!-- 导航 -->
      <div class="w-full flex flex-col lg:space-y-16">
        <!-- 个人收藏 -->
        <div class="card rounded-md shadow-md hover:shadow-lg my-8" id="love" v-if="showSetting.likes">
          <CatTitle :title="'我的收藏'" :icon="'icon-[lucide--star]'"></CatTitle>

          <div class="card-body md:mt-2 px-1 pb-2 md:pb-4">
            <div class="flex flex-col w-full space-x-4">
              <div class="grid grid-cols-card relative" v-if="likeCollectionsList.length > 0">
                <AsyncGroupCard :collections="likeCollectionsList"></AsyncGroupCard>
              </div>
              <div class="select-none mx-4" v-else>
                <div>
                  还没有收藏哦, 请先收藏网站,或者在
                  <span class="font-bold px-2">"设置>我的收藏"</span>
                  添加
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 网站导航 -->
        <div class="card rounded-md shadow-md hover:shadow-lg" :id="data.category.cid" v-if="showSetting.collections"
          v-for="(data, index) in collectionsDatas">
          <!-- 分类标题 -->
          <CatTitle :title="data.category.title" :icon="icons[index]"></CatTitle>

          <div class="card-body group px-1 py-1 md:px-4">
            <div role="tablist" class="tabs tabs-bordered" aria-label="Tabs">
              <template v-for="(group, groupIndex) in data.groups">
                <input type="radio" :name="data.category.cid" role="tab" class="tab first:ml-4 px-1 sm:px-4 text-nowrap"
                  :id="`${data.category.cid}-tab-${groupIndex}`"
                  :aria-controls="`${data.category.cid}-tabpanel-${groupIndex}`" :tabindex="groupIndex === 0 ? 0 : -1"
                  :checked="groupIndex === 0" :aria-label="group.category.title" @click="group.show = true" />
                <div role="tabpanel" class="tab-content my-1" :id="`${data.category.cid}-tabpanel-${groupIndex}`"
                  :aria-labelledby="`${data.category.cid}-tab-${groupIndex}`" v-if="group.show">
                  <KeepAlive>
                    <div class="grid grid-cols-card relative">
                      <AsyncGroupCard :collections="group.collections"></AsyncGroupCard>
                    </div>
                  </KeepAlive>
                </div>
              </template>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import type { Announce } from '@/api'
import * as api from "@/api"

import { icons } from '@/stores/icons'

const store = useBasicStore()
const { collectionsDatas, likeCollectionsList, showSetting } = storeToRefs(store)

const AsyncGroupCard = defineAsyncComponent({
  loader: () => import('../components/user/GroupCard.vue')
})

const AsyncHomeSearch = defineAsyncComponent({
  loader: () => import('../components/user/HomeSearch.vue')
})

const announces = ref<Announce[]>([])

const currentAnnounce = ref<Announce>({
  iD: 0,
  createdAt: '',
  updatedAt: '',
  title: '',
  content: '',
  date: ''
})

const route = useRoute()
const router = useRouter()

function checkMsg() {
  // 登录失败请求
  const msg = route.query.msg
  if (msg) {
    Message({ message: msg as string })
    router.replace({ query: {} })
  }
}

watch(route, () => {
  checkMsg()
})

onMounted(async () => {
  const resp = await api.announces()
  announces.value = resp.data.data['announces']
  currentAnnounce.value = announces.value[0]
  setInterval(() => {
    const currentIndex = announces.value.indexOf(currentAnnounce.value)
    const nextIndex = (currentIndex + 1) % announces.value.length
    currentAnnounce.value = announces.value[nextIndex]
  }, 5000)
  checkMsg()
})
</script>

<style scoped>
.goon-enter-active,
.goon-leave-active {
  transition: all 1s ease-in-out;
}

.goon-enter-from {
  transform: translateY(40px);
}

.goon-leave-to {
  transform: translateY(-40px);
}
</style>
