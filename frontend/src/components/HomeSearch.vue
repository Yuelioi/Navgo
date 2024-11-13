<template>
  <div class="flex justify-center">
    <div class="w-1/3">
      <div role=" tablist" class="tabs tabs-boxed my-6">
        <a
          class="tab [--tab-bg:red]"
          :class="{ 'tab-active': searchItem.active }"
          v-for="(searchItem, current) in searchList"
          @click="switchSearch(current)">
          {{ searchItem.name }}
        </a>
      </div>

      <label
        class="input input-bordered flex items-center transition ease-in-out hover:-translate-y-1 hover:scale-105">
        <input
          type="text"
          class="grow"
          :placeholder="currentSearch?.placeholder"
          v-model="searchValue" />
        <div class="btn btn-sm btn-primary" @click="open(currentSearch?.url + searchValue)">
          搜索
          <span class="icon-[lucide--search]"></span>
        </div>
      </label>
    </div>
  </div>
</template>

<script setup lang="ts">
import { open } from '@/utils'

interface Search {
  name: string
  url: string
  active?: boolean
  placeholder: string
}

const currentSearch = ref<Search>()
const searchValue = ref('')

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

function switchSearch(id: number) {
  searchList.value.forEach((ele) => {
    ele.active = false
  })

  searchList.value[id].active = true
  currentSearch.value = searchList.value[id]
}

onMounted(() => {
  currentSearch.value = searchList.value[0]
})
</script>
