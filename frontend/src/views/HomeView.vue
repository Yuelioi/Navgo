<template>
  <main class="my-8 flex w-full flex-col space-y-8 overflow-hidden">
    <div class="w-4/5 self-center">
      <!-- 公告 -->
      <div class=""></div>

      <!-- 自定义链接 -->

      <!-- 搜索 -->
      <HomeSearch class="mt-8 mb-16"></HomeSearch>

      <!-- 导航 -->
      <div
        class="card bg-base-200 shadow-md hover:shadow-lg my-8"
        :id="data.category.cid"
        v-for="data in collectionsDatas">
        <div class="card-body">
          <div class="flex flex-col w-full">
            <div class="card-title">{{ data.category.title }}</div>

            <div role="tablist" class="tabs tabs-bordered">
              <template v-for="(group, groupIndex) in data.groups">
                <input
                  type="radio"
                  :name="data.category.cid"
                  role="tab"
                  class="tab text-nowrap"
                  :checked="groupIndex === 0"
                  :aria-label="group.category.title"
                  @click="group.show = true" />
                <div role="tabpanel" class="tab-content my-6">
                  <KeepAlive>
                    <AsyncGroupCard
                      v-if="group.show"
                      :data="data"
                      :group="group"
                      :group-index="groupIndex"></AsyncGroupCard>
                  </KeepAlive>
                </div>
              </template>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
const store = useBasicStore()
const { collectionsDatas } = storeToRefs(store)

const AsyncGroupCard = defineAsyncComponent({
  loader: () => import('../components/GroupCard.vue')
})
</script>
