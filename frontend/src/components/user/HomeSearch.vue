<template>
  <div class="flex w-full justify-center">
    <div class="xl:w-1/2 w-full flex flex-col items-center">
      <div role="tablist" class="w-full tabs tabs-boxed rounded-none mt-6 bg-transparent">
        <a
          class="tab [--tab-bg:#0f172a] dark:[--tab-bg:#646872]"
          :class="{ 'tab-active': searchItem.active }"
          v-for="searchItem in searchList"
          role="tab"
          @click="switchSearch(searchItem.name)">
          {{ searchItem.name }}
        </a>
      </div>

      <label
        class="input ring-0 input-bordered w-full flex border-base-content items-center transition ease-in-out">
        <input
          type="text"
          class="grow ring-0 !outline-none"
          :placeholder="currentSearch?.placeholder"
          v-model="searchValue"
          ref="searchRef"
          @keyup="handleEnter($event)"
          @input="handleSearchChange" />

        <span
          class="icon-[lucide--x] size-6 mx-3"
          v-if="searchValue.length > 0"
          @click="searchValue = ''"></span>
        <div class="btn btn-sm btn-primary" @click="open(currentSearch?.url + searchValue)">
          搜索
          <span class="icon-[lucide--search]"></span>
        </div>
      </label>

      <div
        class="relative w-full"
        v-if="searchValue.length > 0 && searchResults.length > 0 && currentSearch?.name == '站内'">
        <div class="absolute flex w-full items-center justify-center">
          <div class="h-96 z-10 w-full overflow-y-scroll flex flex-col space-y-4">
            <div class="m-4" v-for="item in searchResults" @click="open(item.item.link)">
              <div class="p-4 rounded-lg flex items-center space-x-3 bg-base-100/70 drop-shadow-md">
                <div class=""><span class="icon-[lucide--menu] size-5"></span></div>
                <div class="flex flex-col flex-1">
                  <div class="flex justify-between w-full">
                    <span class="font-bold text-lg">{{ item.item.title }}</span>
                    <div class="avatar static size-8">
                      <div class="h-full rounded-xl">
                        <img
                          :alt="item.item.title"
                          loading="lazy"
                          height="320"
                          :src="'https://cdn.yuelili.com/nav/icons/' + item.item.cid + '.png'"
                          @error="imageLoadError"
                          class="h-full rounded-full" />
                        <div
                          class="rounded-full flex h-full items-center justify-center bg-purple-200">
                          <span class="font-bold text-lg">{{ item.item.title[0] }}</span>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="line-clamp-1 w-4/5 opacity-75">{{ item.item.description }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Collection } from '@/api'
import { open } from '@/utils'
import { useFuse } from '@vueuse/integrations/useFuse'
import type { UseFuseOptions } from '@vueuse/integrations/useFuse'
import type { FuseResult } from 'fuse.js'
import { imageLoadError } from '@/utils'
import type { Search } from '@/models'
const store = useBasicStore()
const { currentSearchName, collectionsList, searchList } = storeToRefs(store)

const showSearch = ref(false)

const currentSearch = ref<Search>()
const searchValue = ref('')

const searchResults = ref<FuseResult<Collection>[]>([])
const searchRef = useTemplateRef<HTMLInputElement>('searchRef')

const options = computed<UseFuseOptions<Collection>>(() => ({
  fuseOptions: {
    keys: ['cid', 'title', 'description'],
    includeScore: true,
    shouldSort: true,
    threshold: 0.2
  },
  resultLimit: 20
}))

function switchSearch(name: string) {
  searchValue.value = ''
  searchList.value.forEach((ele: Search) => {
    ele.active = false
    if (ele.name == name) {
      currentSearch.value = ele
      ele.active = true
      currentSearchName.value = ele.name
    }
  })
}

function handleSearchChange() {
  if (searchValue.value !== '' && currentSearch.value?.name === '站内') {
    // @ts-ignore
    const { results } = useFuse(searchValue.value, collectionsList.value, options)
    searchResults.value = results.value
    showSearch.value = true
  } else {
    searchResults.value = []
  }
}

function handleEnter(event: KeyboardEvent) {
  if (event.key == 'Enter') {
    if (currentSearch.value) {
      open(currentSearch.value.url + searchValue.value)
    }
  }
}

// 处理快捷键
function handleKeyDown(event: KeyboardEvent) {
  // Tab , 切换搜索引擎并选择
  if (event.key === 'Tab' && searchValue.value.length === 0) {
    event.preventDefault()
    const idx = searchList.value.findIndex((ele: Search) => {
      return ele.name === currentSearch.value?.name
    })

    const nextID = (idx + 1) % searchList.value.length
    switchSearch(searchList.value[nextID].name)
    searchRef.value?.select()
  }

  // Esc 清除搜索
  if (event.key === 'Escape') {
    event.preventDefault()
    searchValue.value = ''
  }
}

onMounted(() => {
  if (currentSearchName.value) {
    switchSearch(currentSearchName.value)
  } else {
    switchSearch('百度')
  }
  document.addEventListener('keydown', handleKeyDown)
})
</script>

<style scoped>
.tab-active {
  color: #fff;
}
</style>
