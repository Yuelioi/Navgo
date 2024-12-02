<template>
  <div class="w-full flex p-8 h-full space-x-4">
    <!-- 内容区 -->
    <div
      class="flex p-8 flex-col flex-1 space-y-4 bg-base-100/50 backdrop-blur-md z-10 rounded-md border">
      <div class="">投稿</div>
      <div class="divider my-0"></div>

      <div class="">
        <div class="join input input-bordered gap-2 flex">
          <div class="flex items-center justify-center pr-2 space-x-2">
            <span class="text-sm select-none">网址</span>
            <span class="icon-[lucide--link-2] size-5"></span>
          </div>
          <input
            type="text"
            class="grow input-bordered"
            placeholder="https://"
            v-model="form.link" />
        </div>
      </div>
      <!-- 名称 -->
      <div class="">
        <div class="join input input-bordered gap-2 flex">
          <div class="flex items-center justify-center pr-2 space-x-2">
            <span class="text-sm select-none">名称</span>
            <span class="icon-[lucide--case-lower] size-5"></span>
          </div>

          <input type="text" class="grow input-bordered" placeholder="Name" v-model="form.title" />
        </div>
      </div>
      <!-- 分类 -->
      <div class="">
        <div class="join input input-bordered gap-2">
          <div class="flex items-center justify-center pr-2 space-x-2">
            <span class="text-sm select-none">分类</span>
            <span class="icon-[lucide--case-lower] size-5"></span>
          </div>

          <div class="dropdown flex dropdown-right">
            <input
              tabindex="0"
              role="button"
              type="text"
              class=""
              placeholder="暂无分类"
              v-model="form.category" />

            <ul
              tabindex="0"
              class="menu ml-6 flex-col flex-nowrap w-56 dropdown-content rounded-box z-[11] p-2 shadow max-h-96 overflow-y-auto">
              <li>
                <a class="text-nowrap" @click="resetCategory">全部</a>
              </li>

              <li v-for="nav in navs">
                <details>
                  <summary class="text-nowrap">{{ nav.title }}</summary>
                  <ul>
                    <li v-for="child in nav.children" @click.stop="changeCategory(child)">
                      <a class="text-nowrap">{{ child.full_title }}</a>
                    </li>
                  </ul>
                </details>
              </li>
            </ul>
          </div>
        </div>
      </div>
      <!-- 标签 -->
      <div class="">
        <div class="join input input-bordered gap-2 flex">
          <div class="flex items-center justify-center pr-2 space-x-2">
            <span class="text-sm select-none">标签</span>
            <span class="icon-[lucide--link-2] size-5"></span>
          </div>
          <label class="flex relative items-center">
            <div class="space-x-2">
              <div class="badge badge-primary" v-for="(tag, index) in form.tags">
                {{ tag }}
                <span class="icon-[lucide--x]" @click="form.tags.splice(index, 1)"></span>
              </div>
            </div>
            <input
              type="text"
              class="grow pl-4 input-bordered self-center"
              placeholder="按回车创建标签"
              v-model="currentTag"
              @keyup="addTag($event)" />
          </label>
        </div>
      </div>

      <textarea
        class="textarea w-full textarea-bordered input"
        placeholder="网址介绍"
        v-model="form.description"></textarea>
    </div>

    <div class="h-full w-64 flex flex-col space-y-2">
      <div class="border bg-base-100/50 backdrop-blur-md z-10 rounded-md p-4 flex flex-col">
        <div class="">发布设置</div>
        <div class="divider my-1"></div>
        <div class="flex">
          <div class="tooltip tooltip-bottom" data-tip="根据网址自动补全信息">
            <button class="btn btn-sm btn-outline" @click="queryMeta" aria-label="一键填写">
              一键填写
            </button>
          </div>
          <div class="flex ml-auto justify-between">
            <button class="btn btn-sm btn-primary ml-auto" aria-label="提交" @click="handleSubmit">
              提交
            </button>
          </div>
        </div>
      </div>
      <div class="border bg-base-100/50 backdrop-blur-md z-10 rounded-md p-4 flex flex-col">
        <div class="">图标设置</div>
        <div class="divider my-1"></div>

        <DragBox v-model:icon="icon" :title="'图标'"></DragBox>
      </div>

      <!-- 图标 -->
    </div>
  </div>
</template>

<script setup lang="ts">
import * as api from '@/api/index'
import { addCollection, net } from '@/api'
const navs = reactive<api.Nav[]>([])

interface Form {
  link: string
  title: string
  category: string
  description: string
  tags: string[]
}

const currentTag = ref('')
const icon = ref<File | null>(null)
const currentCategory = ref<api.Nav>({
  title: '',
  full_title: ''
})
const form = reactive<Form>({ link: '', title: '', description: '', category: '', tags: [] })

const iconRef = useTemplateRef<HTMLImageElement>('iconRef')

function addTag(event: KeyboardEvent) {
  if (event.key == 'Enter') {
    if (currentTag.value) {
      if (!(form.tags.indexOf(currentTag.value) + 1)) {
        form.tags.push(currentTag.value)
        currentTag.value = ''
      } else {
        Message({ message: '标签已存在' })
      }
    }
  }
}

// 监控icon变动
watch(icon, () => {
  if (icon.value) {
    const reader = new FileReader()
    reader.onload = (e) => {
      if (iconRef.value) {
        iconRef.value.src = e.target?.result as string
      }
    }
    reader.readAsDataURL(icon.value)
  }
})

// 获取网站元数据
async function queryMeta() {
  Message({ message: '获取中...' })
  const resp = await net(form.link)
  if (resp) {
    Message({ message: '获取成功...' })
    form.title = form.title || resp.data.data.title
    form.description = form.description || resp.data.data.description || ''
  } else {
    Message({ message: '获取失败...', type: 'warn' })
  }
}

async function handleSubmit() {
  console.log(currentCategory.value.path)

  const result = await addCollection({
    cid: form.link,
    title: form.title,
    link: form.link,
    description: form.description,
    favicon: icon.value,
    category_path: currentCategory.value.path,
    tags: JSON.stringify(form.tags)
  })
  if (result && result.data.code > 0) {
    Message({ message: '提交成功' })
    form.title = ''
    form.link = ''
    form.description = ''
    icon.value = null
  } else {
    Message({ message: result.data.msg, type: 'warn' })
  }
}

// 分类设置
function resetCategory() {
  form.category = '暂无分类'
  currentCategory.value = {
    cid: 'all',
    title: '暂无分类',
    full_title: ''
  }
}
function changeCategory(nav: Nav) {
  currentCategory.value = nav
  form.category = currentCategory.value.full_title
}

onMounted(async () => {
  // TODO 获取navs接口
  const resp = await api.navs()
  if (resp) {
    Object.assign(navs, resp.data.data.navs)
  }
})
</script>
