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
                      <a class="text-nowrap">{{ child.title }}</a>
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
            <button class="btn btn-sm btn-outline" @click="queryMeta">一键填写</button>
          </div>
          <div class="flex ml-auto justify-between">
            <button class="btn btn-sm btn-primary ml-auto" @click="handleSubmit">提交</button>
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
import { addCollection } from '@/api'
import { net } from '@/logic'
import { imageLoadError } from '@/utils'
const navs = reactive<Nav[]>([])

interface Nav {
  title: string
  cid: string
  path: string
  parent?: Nav
  tags: string[]
  children?: Nav[]
}

interface Form {
  link: string
  title: string
  category: string
  description: string
  tags: string[]
}

const currentCategory = ref<Nav>()

const fileInput = useTemplateRef<HTMLInputElement | null>('fileInput')
const iconRef = useTemplateRef<HTMLImageElement>('iconRef')

const currentTag = ref('')

const form = reactive<Form>({ link: '', title: '', description: '', category: '', tags: [] })

const icon = ref<File | null>(null)

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

// 处理拖拽图片事件
function handleDrag(event: DragEvent) {
  const file = event.dataTransfer?.files?.[0]
  handleFile(file)
}
// 处理粘贴图片事件
function handlePaste(event: ClipboardEvent) {
  const items = event.clipboardData?.items
  if (items && items.length > 0) {
    for (let item of items) {
      if (item.type.startsWith('image/')) {
        const file = item.getAsFile()
        handleFile(file)
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

// 显示
function handleFile(file: File | null | undefined) {
  if (file) {
    icon.value = file
  }
}

// 模拟打开对话框
function openFileInput() {
  if (fileInput.value) {
    fileInput.value.click()
  }
}

// 检测打开文件
function onFileChange(event: Event) {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  handleFile(file)
}

// 获取网站元数据
async function queryMeta() {
  Message({ message: '获取中...' })
  const data = await net(form.link)
  if (data) {
    Message({ message: '获取成功...' })
    form.title = form.title || data.title
    form.description = form.description || data.description || ''
  } else {
    Message({ message: '获取失败...', type: 'warn' })
  }
}

async function handleSubmit() {
  const formData = new FormData()
  // 添加表单字段
  formData.append('title', form.title)
  formData.append('link', form.link)
  formData.append('description', form.description)

  if (icon.value) {
    formData.append('favicon', icon.value, icon.value?.name)
  }

  const result = await addCollection(formData)
  if (result.data.code > 0) {
    Message({ message: '提交成功' })
  } else {
    Message({ message: result.data.msg, type: 'warn' })
  }
}

// 分类设置
function resetCategory() {
  form.category = '暂无分类'
  currentCategory.value = undefined
}
function changeCategory(nav: Nav) {
  currentCategory.value = nav
  form.category = currentCategory.value
    ? currentCategory.value?.parent?.title + '/' + currentCategory.value?.title
    : '暂无分类'
}

onMounted(async () => {
  // TODO 获取navs接口
  const resp = await api.collections()

  resp.data.data.datas.forEach((ele) => {
    // 设置第一组group可见
    ele.groups[0].show = true

    // 初始化导航菜单
    const parent: Nav = {
      title: ele.category.title,
      cid: ele.category.cid || '',
      path: ele.category.path || '',
      children: [],
      tags: []
    }

    ele.groups.forEach((group) => {
      // 填充导航菜单
      if (parent.children) {
        parent.children.push({
          title: group.category.title,
          cid: group.category.cid || '',
          path: group.category.path || '',
          parent: parent,
          tags: []
        })
      }
    })

    navs.push(parent)
  })
})
</script>
