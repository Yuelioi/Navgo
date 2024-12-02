<template>
  <div class="container py-8 z-10">
    <div class="p-4 backdrop-blur-md rounded-lg shadow-lg flex flex-col">
      <div class="flex w-full items-center space-x-4">
        <!-- 分类筛选 -->
        <div class="dropdown dropdown-bottom">
          <div
            tabindex="0"
            role="button"
            aria-label="分类筛选"
            class="btn bg-base-100/50 backdrop-blur-sm m-1">
            {{ currentCategory?.full_title || '分类筛选' }}
          </div>
          <ul
            tabindex="0"
            class="menu flex-col bg-base-100 flex-nowrap w-56 dropdown-content rounded-box z-[11] p-2 shadow max-h-96 overflow-y-auto">
            <li>
              <a class="text-nowrap" @click="currentCategory = undefined">全部</a>
            </li>

            <li v-for="nav in navs">
              <details>
                <summary class="text-nowrap">{{ nav.title }}</summary>
                <ul>
                  <li v-for="child in nav.children" @click.stop="currentCategory = child">
                    <a class="text-nowrap">{{ child.title }}</a>
                  </li>
                </ul>
              </details>
            </li>
          </ul>
        </div>
        <label class="input input-bordered flex items-center gap-2">
          <input
            type="text"
            class="grow"
            placeholder="Search"
            v-model="collectionFilter.keyword"
            @keydown="fetchKeyword" />
          <span class="icon-[lucide--search]"></span>
        </label>
      </div>

      <div class="divider"></div>

      <div class="">
        <div class="overflow-x-auto">
          <table class="table rounded-none table-sm table-pin-rows">
            <!-- head -->
            <thead>
              <tr class="font-bold bg-base-100/20">
                <th class="text-center">索引</th>
                <th class="">图标</th>
                <th class="">标题</th>
                <th class="">链接</th>
                <th class="">描述</th>
                <th class="">标签</th>
                <th class="">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(collection, index) in categoryFilters?.collections">
                <td class="text-center">{{ index + 1 }}</td>
                <td>
                  <div class="avatar static size-12 group">
                    <div class="h-full rounded-xl">
                      <img
                        :alt="collection.title"
                        :src="'https://cdn.yuelili.com/nav/icons/' + collection?.cid + '.png'"
                        @error="imageLoadError"
                        class="h-full rounded-full" />
                      <DragBoxMini v-model:icon="collection.favicon"></DragBoxMini>
                    </div>
                  </div>
                </td>
                <td>
                  <input type="text" class="input w-full" v-model="collection.title" />
                </td>
                <td>
                  <input type="text" class="input w-full" v-model="collection.link" @invalid="" />
                </td>
                <td>
                  <textarea
                    type="text"
                    class="input w-full overflow-hidden"
                    v-model="collection.description" />
                </td>
                <td class="group w-48 overflow-x-hidden">
                  <div class="join items-center">
                    <div class="group-hover:hidden space-x-1 space-y-1">
                      <div class="badge badge-primary" v-for="tag in collection.tags">
                        {{ tag }}
                      </div>
                    </div>
                    <input
                      type="text"
                      class="input w-full input-md hidden group-hover:block"
                      v-model="collection.tags"
                      @change="updateTags(collection)" />
                  </div>
                </td>
                <th class="w-min">
                  <div class="flex max-w-fit space-x-3">
                    <button
                      aria-label="保存"
                      class="btn btn-sm btn-square btn-outline shadow-xl"
                      @click="saveCollection(collection)">
                      <span class="icon-[lucide--save] size-6"></span>
                    </button>
                    <button
                      aria-label="移除"
                      class="btn btn-sm btn-square btn-outline"
                      @click="removeCollection(collection)">
                      <span class="icon-[lucide--square-x] size-6"></span>
                    </button>
                  </div>
                </th>
              </tr>
            </tbody>
          </table>

          <div class="join">
            <button
              aria-label="首页"
              class="join-item btn bg-base-100/50 backdrop-blur-md"
              @click="collectionFilter.page = 1">
              <span class="icon-[lucide--chevron-first]"></span>
            </button>
            <button
              aria-label="上一首页"
              class="join-item btn bg-base-100/50"
              @click="collectionFilter.page = Math.max(1, collectionFilter.page - 1)">
              <span class="icon-[lucide--chevron-left]"></span>
            </button>
            <input
              :key="collectionFilter.page"
              class="join-item btn btn-square"
              type="radio"
              name="options"
              checked
              :aria-label="collectionFilter.page.toString()" />

            <button
              aria-label="下一页"
              class="join-item btn bg-base-100/50"
              @click="
                collectionFilter.page = Math.min(
                  categoryFilters.totalPages,
                  collectionFilter.page + 1
                )
              ">
              <span class="icon-[lucide--chevron-right]"></span>
            </button>
            <button
              class="join-item btn bg-base-100/50"
              @click="collectionFilter.page = categoryFilters.totalPages">
              <span class="icon-[lucide--chevron-last]"></span>
            </button>
            <label
              aria-label="尾页"
              class="join-item input border-none leading-4 bg-base-100/50 input-bordered flex items-center gap-2"
              @keyup="jumpKey($event)">
              <input type="text" class="grow" placeholder="跳转" v-model="searchPage" />
              <span class="opacity-50">/{{ categoryFilters.totalPages }}</span>
              <span class="icon-[lucide--corner-down-left]"></span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Collection } from '@/api/types'
import * as api from '@/api/index'
import { imageLoadError } from '@/utils'
import { VAlert } from '@/plugins/alert'

const collectionFilter = reactive<api.CollectionFilter>({
  categories: [],
  page: 1,
  page_size: 20,
  keyword: ''
})

const searchPage = ref(1)

const currentCategory = ref<api.Nav>()
const categoryFilters = ref<api.CollectionsResponse>({
  count: 1,
  totalPages: 1,
  collections: []
})

const navs = reactive<api.Nav[]>([])

function validatePage() {
  const x = Number(searchPage.value)
  if (x >= 1 && x <= categoryFilters.value.totalPages && collectionFilter.page != x) {
    collectionFilter.page = x
  }
}

function updateTags(collection: Collection) {
  collection.tags = collection.tags?.toString().split(',')
}

function jump() {
  validatePage()
}
function jumpKey(event: KeyboardEvent) {
  if (event.key === 'Enter') {
    validatePage()
  }
}

async function fetchCollections() {
  // 重置分页, 会自动触发更新
  if (collectionFilter.page != 1) {
    collectionFilter.page = 1
    return
  }

  // 正常更新
  const resp = await api.filteredCollections(collectionFilter)
  categoryFilters.value = resp.data.data
}

async function fetchKeyword(event: KeyboardEvent) {
  if (event.key == 'Enter') {
    await fetchCollections()
  }
}

watch(
  () => currentCategory.value,
  async (x, y) => {
    // 分类变化
    if (x != y) {
      collectionFilter.categories = [currentCategory.value?.path as string]
    }

    await fetchCollections()
  }
)

watch(
  () => [collectionFilter.page],
  async () => {
    const resp = await api.filteredCollections(collectionFilter)
    categoryFilters.value = resp.data.data
  }
)

async function saveCollection(collection: Collection) {
  const resp = await api.updateCollection({
    ...collection,
    category_path: collection.category.path,
    favicon: collection.favicon,
    tags: JSON.stringify(collection.tags)
  })
  if (resp.data.code >= 0) {
    Message({ message: '修改成功' })
  } else {
    Message({ message: resp.data.msg, type: 'wanning' })
  }
}
function removeCollection(collection: Collection) {
  // 或者使用自定义设置
  VAlert({ alert: '确认要删除吗' }).then(async (confirmed) => {
    if (confirmed) {
      await api.deleteCollection({ id: collection.link })
    }
  })
}

onMounted(async () => {
  const resp = await api.filteredCollections({
    page: 1,
    keyword: ''
  })

  categoryFilters.value = resp.data.data

  const resp2 = await api.navs()
  Object.assign(navs, resp2.data.data.navs)
})
</script>

<style scoped>
tr:hover {
  background: var(--fallback-b2, oklch(var(--b2) / 0.5));
}
</style>
