<template>
  <div class="flex w-full justify-center">
    <div class="w-1/2 flex flex-col items-center">
      <div role="tablist" class="w-full tabs tabs-lifted rounded-none mt-6">
        <a
          class="tab [--tab-bg:#0f172a] dark:[--tab-bg:#a3a7af]"
          :class="{ 'tab-active': searchItem.active }"
          v-for="searchItem in searchList"
          @click="switchSearch(searchItem.name)">
          {{ searchItem.name }}
        </a>
      </div>

      <label
        class="input ring-0 input-bordered rounded-none w-full flex items-center transition ease-in-out">
        <input
          type="text"
          class="grow ring-0 !outline-none"
          :placeholder="currentSearch?.placeholder"
          v-model="searchValue"
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

      <div class="relative w-full" v-if="searchValue.length > 0 && currentSearch?.name == '站内'">
        <div class="absolute flex items-center justify-center">
          <div
            class="h-96 shadow-lg z-10 overflow-y-scroll flex flex-col rounded-b-2xl bg-base-300 space-y-2">
            <div
              class="bg-base-200 rounded-md m-4 hover:inset-1 hover:bg-neutral/30 hover:cursor-pointer"
              v-for="item in result"
              @click="open(item.item.link)">
              <div class="p-4 flex items-center space-x-3">
                <div class=""><span class="icon-[lucide--menu] size-5"></span></div>
                <div class="flex flex-col flex-1">
                  <div class="flex justify-between w-full">
                    <span class="font-bold text-lg">{{ item.item.title }}</span>
                    <div class="avatar static size-8">
                      <div class="h-full rounded-xl">
                        <img
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

const store = useBasicStore()
const { currentSearchName, collectionsList } = storeToRefs(store)

const showSearch = ref(false)

interface Search {
  name: string
  url: string
  active?: boolean
  placeholder: string
}

const currentSearch = ref<Search>()
const searchValue = ref('')

const result = ref<FuseResult<Collection>[]>([])

const searchList = ref([
  {
    name: '百度',
    url: 'https://www.baidu.com/s?wd=',
    active: true,
    placeholder: '百度一下'
  },
  {
    name: '必应',
    url: 'https://www.bing.com/search?q=',
    placeholder: '必应一下'
  },
  {
    name: '谷歌',
    url: 'https://www.google.com/search?q=',
    placeholder: '谷歌一下'
  },
  {
    name: '站内',
    url: '/search?q=',
    placeholder: '站内搜索'
  }
])

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
  searchList.value.forEach((ele) => {
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
    console.log(collectionsList.value)

    const { results } = useFuse(searchValue.value, collectionsList.value, options)
    result.value = results.value
    console.log(results.value)
    showSearch.value = true
  } else {
    result.value = []
  }
}

function imageLoadError(event: Event) {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}

function handleEnter(event: KeyboardEvent) {
  if (event.key == 'Enter') {
    if (currentSearch.value) {
      open(currentSearch.value.url + searchValue.value)
    }
  }
}

onMounted(() => {
  switchSearch(currentSearchName.value)
})
</script>

<style scoped>
.tab-active {
  color: #fff;
}
</style>
