<template>
  <div class="container py-8">
    <div class="p-4 rounded-lg shadow-lg flex flex-col">
      <div class="flex items-center space-x-4">
        <!-- 分类筛选 -->
        <div class="dropdown dropdown-bottom">
          <div tabindex="0" role="button" class="btn m-1">
            {{ currentCategory?.title || '分类筛选' }}
          </div>
          <ul
            tabindex="0"
            class="menu flex-col flex-nowrap w-56 dropdown-content bg-base-100 rounded-box z-[11] p-2 shadow max-h-96 overflow-y-auto">
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
          <input type="text" class="grow" placeholder="Search" v-model="searchValue" />
          <span class="icon-[lucide--search]"></span>
        </label>
      </div>

      <div class="divider"></div>

      <div class="">
        <div class="overflow-x-auto">
          <table class="table rounded-none table-sm table-pin-rows">
            <!-- head -->
            <thead>
              <tr class="font-bold">
                <th class="text-center">索引</th>
                <th class="">图标</th>
                <th class="">标题</th>
                <th class="">链接</th>
                <th class="">描述</th>
                <th class="">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(collection, index) in filteredCollections">
                <td class="text-center">{{ index + 1 }}</td>
                <td>
                  <div class="avatar static size-12 group">
                    <div class="h-full rounded-xl">
                      <img
                        :src="'https://cdn.yuelili.com/nav/icons/' + collection?.cid + '.png'"
                        @error="imageLoadError"
                        class="h-full rounded-full" />
                      <div class="btn flex p-0 btn-secondary" @click="uploadImage(collection)">
                        <span
                          class="hidden group-hover:block icon-[lucide--arrow-big-up-dash] size-8"></span>
                        <span class="group-hover:hidden text-sm">
                          暂无
                          <br />
                          图片
                        </span>
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
                <td>
                  <textarea
                    type="text"
                    class="input w-full overflow-hidden"
                    v-model="collection.description" />
                </td>
                <th class="w-min">
                  <div class="flex max-w-fit space-x-3">
                    <button
                      class="btn btn-sm btn-square btn-info btn-outline"
                      @click="removeCollection()">
                      <span class="icon-[lucide--save] size-6"></span>
                    </button>
                    <button
                      class="btn btn-sm btn-square btn-error btn-outline"
                      @click="removeCollection()">
                      <span class="icon-[lucide--square-x] size-6"></span>
                    </button>
                  </div>
                </th>
              </tr>
            </tbody>
          </table>

          <div class="join">
            <button class="join-item btn" @click="currentPage = 1">
              <span class="icon-[lucide--chevron-first]"></span>
            </button>
            <button class="join-item btn" @click="currentPage = Math.max(1, currentPage - 1)">
              <span class="icon-[lucide--chevron-left]"></span>
            </button>
            <input
              :key="currentPage"
              class="join-item btn btn-square"
              type="radio"
              name="options"
              checked
              :aria-label="currentPage.toString()" />

            <button
              class="join-item btn"
              @click="currentPage = Math.min(patination, currentPage + 1)">
              <span class="icon-[lucide--chevron-right]"></span>
            </button>
            <button class="join-item btn" @click="currentPage = patination">
              <span class="icon-[lucide--chevron-last]"></span>
            </button>
            <label
              class="join-item input input-bordered flex items-center gap-2"
              @click.stop="jump"
              @keyup="jumpKey($event)">
              <input type="text" class="grow" placeholder="跳转" v-model="searchPage" />
              <span class="opacity-50">/{{ patination }}</span>
              <span class="icon-[lucide--corner-down-left]"></span>
            </label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Collection, CollectionsData } from '@/api/types'
import * as api from '@/api/index'
import { imageLoadError } from '@/utils'

import { useFuse } from '@vueuse/integrations/useFuse'
import type { UseFuseOptions } from '@vueuse/integrations/useFuse'

const searchValue = ref('')

const collectionsDatas = ref<Array<CollectionsData>>([])
const collections: Ref<Collection[]> = ref([])

const pageSize = 20
const currentPage = ref(1)
const searchPage = ref(1)

interface Nav {
  title: string
  cid: string
  path: string
  parent?: Nav
  children?: Nav[]
}

const currentCategory = ref<Nav>()

const filteredFullCollections: Ref<Collection[]> = computed(() => {
  if (searchValue.value) {
    const tmpCollections = collections.value.filter((ele) => {
      if (!currentCategory.value?.title) {
        return ele
      }
      if (ele.category?.path === currentCategory.value?.path) {
        return ele
      }
    })

    currentPage.value = 1

    const { results } = useFuse(searchValue.value, tmpCollections, options)
    return results.value.map((result) => result.item)
  } else {
    if (!currentCategory.value?.path) {
      return collections.value
    }

    currentPage.value = 1

    return collections.value.filter((ele) => {
      if (ele.category?.path === currentCategory.value?.path) {
        return ele
      }
    })
  }
})

const filteredCollections: Ref<Collection[]> = computed(() => {
  return filteredFullCollections.value.slice(
    (currentPage.value - 1) * pageSize,
    currentPage.value * pageSize
  )
})

const patination: Ref<number> = computed(() => {
  return ((filteredFullCollections.value.length / pageSize) | 0) + 1
})

const navs = reactive<Nav[]>([])

const options = computed<UseFuseOptions<Collection>>(() => ({
  fuseOptions: {
    keys: ['link', 'title', 'description'],
    includeScore: true,
    shouldSort: true,
    threshold: 0.2
  },
  resultLimit: 0
}))

function validatePage() {
  const x = Number(searchPage.value)
  if (x >= 1 && x <= patination.value) {
    currentPage.value = x
  }
}

function jump() {
  validatePage()
}
function jumpKey(event: KeyboardEvent) {
  if (event.key === 'Enter') {
    validatePage()
  }
}

function uploadImage(collection: Collection) {}
function removeCollection() {}

onMounted(async () => {
  const resp = await api.collections()
  collectionsDatas.value = resp.data.data.datas

  collectionsDatas.value.forEach((ele) => {
    // 设置第一组group可见
    ele.groups[0].show = true

    // 初始化导航菜单
    const parent: Nav = {
      title: ele.category.title,
      cid: ele.category.cid || '',
      path: ele.category.path || '',
      children: []
    }

    ele.groups.forEach((group) => {
      // 填充导航菜单
      if (parent.children) {
        parent.children.push({
          title: group.category.title,
          cid: group.category.cid || '',
          path: group.category.path || '',
          parent: parent
        })
      }

      // 填充导航
      group.collections.forEach((item) => {
        collections.value.push(item)
      })
    })

    navs.push(parent)
  })
})
</script>
