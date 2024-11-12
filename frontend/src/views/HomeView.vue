<template>
  <main class="my-8">
    <!-- 公告 -->
    <div class=""></div>

    <!-- 导航 -->
    <div class="flex w-full flex-col space-y-8">
      <!-- 顶级 -->
      <div class="card bg-base-100 w-4/5 self-center shadow-xl" v-for="data in collectionsDatas">
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
                  :aria-label="group.category.title" />
                <div role="tabpanel" class="tab-content bg-base-100 my-6">
                  <div class="flex flex-wrap">
                    <div class="basis-1/6 p-4" v-for="collection in group.collections">
                      <div class="bg-base-100 shadow-xl p-4">
                        <div class="h-28 flex flex-col">
                          <div class="flex flex-row flex-1 overflow-hidden items-center relative">
                            <img :src="collection.thumbnail" class="h-full rounded-full" />

                            <div class="pl-4">
                              <div class="font-bold">{{ collection.title }}</div>
                              <div class="line-clamp-1">{{ collection.description }}</div>
                            </div>
                          </div>

                          <div class="border my-2"></div>

                          <div class="h-6">
                            <div class="badge badge-ghost" v-if="collection.proxy">需要代理</div>
                            <div class="" v-for="tag in collection.tags">
                              <div class="badge badge-ghost">{{ tag }}</div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
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
</script>
