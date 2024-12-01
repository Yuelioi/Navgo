<template>
  <div class="container flex flex-col space-y-8">
    <div
      role="alert"
      class="alert bg-info/75 text-info-content border-none backdrop-blur-md z-10 flex rounded-md">
      <span class="icon-[lucide--info] size-5 hidden md:block"></span>
      <div class="flex w-full">
        <span class="hidden md:block font-bold transition-transform">注意</span>
        <span class="md:ml-8 text-left">个人收藏只会存在本地浏览器,请勿删除当前页面浏览器数据</span>
      </div>
    </div>

    <div
      class="flex z-10 backdrop-blur-md w-full flex-col p-6 space-y-4 bg-base-100/50 border border-neutral-content/50 rounded-md shadow-md">
      <div class="font-bold"><span>Tips</span></div>
      <div class="divider"></div>
      <div class="flex items-center w-full justify-between">
        <kbd class="kbd">Tab</kbd>
        <div class="kbd">切换搜索引擎</div>
      </div>
      <div class="flex items-center w-full justify-between">
        <kbd class="kbd">Esc</kbd>
        <div class="kbd">删除搜索关键词</div>
      </div>
    </div>

    <div
      class="flex backdrop-blur-md z-10 w-full flex-col p-6 space-y-4 bg-base-100/50 border border-neutral-content/50 rounded-md shadow-md">
      <div class="font-bold flex items-center justify-between">
        <span class="font-bold">显示设置</span>
      </div>
      <div class="divider"></div>

      <div class="flex flex-col">
        <div class="flex md:hidden items-center w-full justify-between">
          <span class="font-bold">夜间模式</span>
          <label class="label cursor-pointer">
            <input
              type="checkbox"
              class="toggle"
              value="synthwave"
              @click="changeTheme()"
              :checked="theme == 'dark'" />
          </label>
        </div>

        <div class="flex items-center w-full justify-between">
          <span>显示我的收藏</span>
          <div class="form-control">
            <label class="label cursor-pointer">
              <input type="checkbox" class="toggle" v-model="showSetting.likes" />
            </label>
          </div>
        </div>
        <div class="flex items-center w-full justify-between">
          <span>显示网站收藏</span>
          <div class="form-control">
            <label class="label cursor-pointer">
              <input type="checkbox" class="toggle" v-model="showSetting.collections" />
            </label>
          </div>
        </div>

        <div class="flex items-center w-full justify-between">
          <span>显示公告</span>
          <div class="form-control">
            <label class="label cursor-pointer">
              <input type="checkbox" class="toggle" v-model="showSetting.announce" />
            </label>
          </div>
        </div>
        <div class="flex items-center w-full justify-between">
          <span>显示页脚</span>
          <div class="form-control">
            <label class="label cursor-pointer">
              <input type="checkbox" class="toggle" v-model="showSetting.footer" />
            </label>
          </div>
        </div>
      </div>
    </div>

    <div
      class="flex backdrop-blur-md z-10 w-full flex-col p-6 space-y-4 bg-base-100/50 border border-neutral-content/50 rounded-md shadow-md">
      <div class="font-bold flex items-center justify-between">
        <span class="font-bold">壁纸设置</span>
      </div>
      <div class="divider"></div>
      <div class="flex items-center w-full justify-between">
        <span>显示壁纸</span>
        <div class="form-control">
          <label class="label cursor-pointer">
            <input type="checkbox" class="toggle" v-model="wallpaperSetting.show" />
          </label>
        </div>
      </div>
      <div class="flex items-center w-full justify-between">
        <span>使用本地壁纸</span>
        <div class="form-control">
          <label class="label cursor-pointer">
            <input
              type="checkbox"
              class="toggle"
              v-model="wallpaperSetting.useLocal"
              @change="loadLocalWallpaper" />
          </label>
        </div>
      </div>

      <div class="flex items-center w-full justify-between" v-if="wallpaperSetting.show">
        <span>壁纸不透明度</span>
        <div class="form-control">
          <label class="label cursor-pointer">
            <input
              type="range"
              min="0"
              max="100"
              step="10"
              class="range"
              @input="wallpaperSetting.opacity = opacityPercentage / 100"
              v-model="opacityPercentage" />
          </label>
        </div>
      </div>
      <DragBox :title="'壁纸'" v-model:icon="icon" v-if="wallpaperSetting.useLocal"></DragBox>
    </div>

    <div
      class="hidden backdrop-blur-md z-10 md:flex w-full flex-col p-6 space-y-4 bg-base-100/50 border border-neutral-content/50 rounded-md shadow-md">
      <div class="flex items-center">
        <span class="font-bold">搜索设置</span>

        <div class="ml-auto flex gap-4">
          <div class="btn btn-primary btn-sm" @click="addSearch">添加</div>
          <div class="btn btn-sm" @click="resetSearch">重置</div>
        </div>
      </div>
      <div class="divider"></div>
      <table class="table">
        <thead>
          <tr>
            <th class="hidden md:block">序号</th>
            <th>名称</th>
            <th>搜索链接</th>
            <th>占位符</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <template v-for="(search, index) in searchList" :key="search.name">
            <tr class="group">
              <th class="hidden md:block">{{ index + 1 }}</th>
              <td>
                <input
                  type="text"
                  class="w-full min-w-[13vw] input input-bordered"
                  v-model.lazy="search.name" />
              </td>
              <td>
                <input
                  type="text"
                  class="w-full min-w-[20vw] input input-bordered"
                  v-model="search.url" />
              </td>
              <td>
                <input
                  type="text"
                  class="w-full min-w-[20vw] input input-bordered"
                  v-model.lazy="search.placeholder" />
              </td>
              <td width="200">
                <div class="flex md:hidden group-hover:flex gap-2">
                  <button
                    class="btn btn-sm hidden md:block"
                    @click="swap(searchList, index, index - 1)"
                    :disabled="index == 0">
                    <span class="icon-[lucide--arrow-up]"></span>
                  </button>
                  <button
                    class="btn btn-sm hidden md:block"
                    @click="swap(searchList, index, index + 1)"
                    :disabled="index == searchList.length - 1">
                    <span class="icon-[lucide--arrow-down]"></span>
                  </button>
                  <button class="btn btn-error btn-sm" @click="removeSearch(index)">删除</button>
                </div>
              </td>
            </tr>
          </template>
        </tbody>
      </table>
    </div>

    <div
      class="hidden z-10 md:flex shadow-md border border-neutral-content/50 rounded-md bg-base-100/50 w-full flex-col p-6 space-y-4">
      <div class="flex items-center justify-between">
        <span class="font-bold">收藏设置</span>
        <div class="flex gap-4 items-center">
          <div class="btn btn-primary btn-sm" @click="addLike">添加</div>
          <div class="btn btn-success text-success-content btn-sm" @click="saveCollections()">
            保存
          </div>
          <div
            class="ml-auto flex items-center space-x-2 mr-4 text-warning-content bg-warning rounded-md px-2 py-1"
            v-if="isModify">
            <span class="icon-[lucide--triangle-alert]"></span>
            <div class="font-bold">尚未保存</div>
          </div>
        </div>
      </div>
      <div class="divider"></div>

      <table class="table rounded-none">
        <thead>
          <tr class="font-bold">
            <th class="font-bold hidden md:block">图标</th>
            <th class="font-bold">标题</th>
            <th class="font-bold">链接</th>
            <th class="font-bold">描述</th>
            <th class="font-bold">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(collection, index) in likeCollectionsList"
            class="group"
            :key="collection.link">
            <td class="hidden md:block">
              <div class="avatar static size-12">
                <div class="h-full rounded-xl">
                  <img
                    :alt="collection.title"
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
            <td>
              <textarea
                type="text"
                class="input w-full overflow-y-hidden"
                v-model="collection.description" />
            </td>
            <td width="md:200">
              <div class="md:flex hidden group-hover:flex space-x-2">
                <!-- TODO 允许排序 但是需要加个order属性 -->
                <!-- <button
                  class="btn btn-sm"
                  @click="swap(likeCollectionsList, index, index - 1)"
                  :disabled="index == 0">
                  <span class="icon-[lucide--arrow-up]"></span>
                </button>
                <button
                  class="btn btn-sm"
                  @click="swap(likeCollectionsList, index, index + 1)"
                  :disabled="index == likeCollectionsList.length - 1">
                  <span class="icon-[lucide--arrow-down]"></span>
                </button> -->
                <button class="btn ml-auto btn-error btn-sm" @click="removeCollection(index)">
                  删除
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type Collection } from '@/api'
import { db } from '@/db/db'
import { imageLoadError } from '@/utils'
import { defaultSearchList } from '@/consts'
import { useCloned } from '@vueuse/core'

const { switchTheme } = useTheme()

const icon = ref<File | null>(null)
const isModify = ref(false)
const opacityPercentage = ref(0)

const store = useBasicStore()
const {
  likeCollectionsList,
  localWallpaper,
  currentWallpaper,
  showSetting,
  searchList,
  theme,
  wallpaperSetting
} = storeToRefs(store)

opacityPercentage.value = wallpaperSetting.value.opacity * 100

watch(icon, async (newFile) => {
  if (newFile) {
    const reader = new FileReader()

    // 定义读取操作完成时的回调函数
    reader.onload = (event) => {
      // event.target.result包含了读取的结果
      const dataUrl = event.target?.result
      db.addData('wallpaper', {
        id: 'wallpaper',
        data: dataUrl
      })
      localWallpaper.value = dataUrl as string
      currentWallpaper.value = dataUrl as string
    }

    // 开始读取文件，这里我们希望得到一个data URL
    reader.readAsDataURL(newFile)
  } else {
    localWallpaper.value = ''
    await db.deleteData('wallpaper', 'wallpaper')
  }
})

function removeCollection(id: number) {
  likeCollectionsList.value.splice(id, 1)
  isModify.value = true
}

function changeTheme() {
  theme.value = theme.value == 'light' ? 'dark' : 'light'
  switchTheme(theme.value)
}

function resetSearch() {
  const { cloned } = useCloned(defaultSearchList)
  searchList.value = cloned.value
}

function removeSearch(id: number) {
  if (searchList.value.length > 1) {
    searchList.value.splice(id, 1)
    Message({ message: '删除成功' })
  } else {
    Message({ message: '请保留至少一个搜索项' })
  }
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
        link: collection.link,
        favicon: collection.favicon
      }

      await db.addData('likes', serializableCollection)
    }
  })
  isModify.value = false
}

function swap(arr: any, sourceIndex: number, targetIndex: number) {
  if (
    sourceIndex < 0 ||
    sourceIndex >= arr.length ||
    targetIndex < 0 ||
    targetIndex >= arr.length
  ) {
    return
  }
  const temp = arr[sourceIndex]
  arr[sourceIndex] = arr[targetIndex]
  arr[targetIndex] = temp
}

function addLike() {
  if (likeCollectionsList.value.length < 20) {
    likeCollectionsList.value.push({
      cid: '',
      title: '',
      link: '',
      category: { title: '' },
      like: false,
      description: ''
    })
    isModify.value = true
  } else {
    Message({ message: '最多添加20个收藏喔!' })
  }
}
function addSearch() {
  if (searchList.value.length < 7) {
    searchList.value.push({
      name: '',
      url: '',
      placeholder: ''
    })
  } else {
    Message({ message: '最多添加7个搜索喔!' })
  }
}

async function loadLocalWallpaper() {
  if (wallpaperSetting.value.useLocal) {
    const data = await db.queryData('wallpaper', 'wallpaper')

    if (data) {
      localWallpaper.value = data['data']

      const [mediaType, base64Data] = data['data'].split(';base64,')
      const mimeType = mediaType.split(':')[1]

      fetch(`data:${mimeType};base64,${base64Data}`)
        .then((res) => res.blob())
        .then((blob) => {
          const file = new File([blob], 'image.png', { type: mimeType })
          icon.value = file
        })
        .catch((err) => console.error('Error converting Data URL to File:', err))
    }
  }
}

onMounted(async () => {
  await loadLocalWallpaper()
})
</script>
