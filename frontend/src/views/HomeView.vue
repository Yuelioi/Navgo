<template>
  <div class="flex flex-col relative container z-10">
    <div class="self-center flex-1 my-8 w-full">
      <!-- 公告 -->
      <div
        role="alert"
        class="alert flex w-full backdrop-blur-sm shadow-md border-1 border-opacity-25 bg-base-200/20 overflow-hidden"
        v-if="currentAnnounce && showSetting.announce">
        <span class="icon-[lucide--info] size-5"></span>
        <Transition name="goon" mode="out-in">
          <div class="flex w-full" :key="currentAnnounce.title">
            <span class="font-bold transition-transform">
              {{ currentAnnounce.title }}
            </span>
            <span class="ml-8">{{ currentAnnounce.content }}</span>
            <span class="ml-auto">{{ currentAnnounce.date }}</span>
          </div>
        </Transition>
      </div>

      <!-- 搜索 -->
      <HomeSearch class="mt-12 md:mt-24 mb-16 z-50 relative"></HomeSearch>

      <!-- 导航 -->
      <div class="w-full flex flex-col lg:space-y-16">
        <!-- 个人收藏 -->
        <div
          class="card rounded-md shadow-md hover:shadow-lg my-8"
          id="love"
          v-if="showSetting.likes">
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
        <div
          class="card rounded-md shadow-md hover:shadow-lg"
          :id="data.category.cid"
          v-if="showSetting.collections"
          v-for="(data, index) in collectionsDatas">
          <!-- 分类标题 -->
          <CatTitle :title="data.category.title" :icon="icons[index]"></CatTitle>

          <div class="card-body group px-1 py-1 md:px-4">
            <div role="tablist" class="tabs tabs-bordered">
              <template v-for="(group, groupIndex) in data.groups">
                <input
                  type="radio"
                  :name="data.category.cid"
                  role="tab"
                  class="tab first:ml-4 px-1 sm:px-4 text-nowrap"
                  :checked="groupIndex === 0"
                  :aria-label="group.category.title"
                  @click="group.show = true" />
                <div role="tabpanel" class="tab-content my-1" v-if="group.show">
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
import { getAnnounces } from '@/logic'

import { icons } from '@/stores/icons'

const store = useBasicStore()
const { collectionsDatas, likeCollectionsList, showSetting } = storeToRefs(store)

const AsyncGroupCard = defineAsyncComponent({
  loader: () => import('../components/user/GroupCard.vue')
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

onMounted(async () => {
  const data = await getAnnounces()

  announces.value = data['announces']
  currentAnnounce.value = announces.value[0]
  setInterval(() => {
    const currentIndex = announces.value.indexOf(currentAnnounce.value)
    const nextIndex = (currentIndex + 1) % announces.value.length
    currentAnnounce.value = announces.value[nextIndex]
  }, 5000)
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
