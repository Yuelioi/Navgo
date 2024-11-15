<template>
  <div class="overflow-hidden h-full">
    <div class="sticky top-1/4">
      <div class="absolute right-16">
        <div class="card card-compact bg-base-100 p-4 w-64 shadow-xl">
          <div class="card-body">
            <h2 class="card-title">使用说明</h2>
            <ul class="steps steps-vertical">
              <li class="step step-accent font-bold">输入网址</li>
              <li class="step step-accent font-bold">单击一键填写</li>
              <li class="step step-accent font-bold">上传图片(非必要)</li>
              <li class="step step-accent font-bold">提交</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <div class="flex w-full h-full items-center justify-center">
      <div class="flex flex-col space-y-4 self-center w-1/4 items-center justify-center">
        <div class="join input input-bordered gap-2 w-full hover:border-info">
          <div class="flex border-r-2 items-center justify-center pr-2 space-x-2">
            <span class="text-sm select-none">网址</span>
            <span class="icon-[lucide--link-2] size-5"></span>
          </div>
          <div class="divider"></div>
          <input type="text" class="grow input-bordered" placeholder="https://" v-model="link" />
        </div>

        <div class="join input input-bordered gap-2 w-full hover:border-info">
          <div class="flex border-r-2 items-center justify-center pr-2 space-x-2">
            <span class="text-sm select-none">名称</span>
            <span class="icon-[lucide--case-lower] size-5"></span>
          </div>
          <div class="divider"></div>
          <input type="text" class="grow input-bordered" placeholder="Name" v-model="name" />
        </div>

        <div class="join input input-bordered gap-2 w-full hover:border-info select-none">
          <div class="flex border-r-2 items-center justify-center pr-3 space-x-2">
            <span class="text-sm select-none">代理</span>
            <span class="icon-[lucide--utility-pole]"></span>
          </div>
          <div class="divider"></div>
          <div class="flex flex-1 items-center select-none">
            <span class="text-sm opacity-45">如果需要科学上网, 请打勾</span>

            <input type="checkbox" checked class="ml-auto checkbox" v-model="proxy" />
          </div>
        </div>

        <!-- 图标 -->
        <div
          class="flex border rounded-lg input-bordered select-none w-full hover:border-info"
          contenteditable="true"
          @paste.prevent="handlePaste">
          <div class="flex items-center justify-center w-full h-36 relative">
            <div class="absolute inset-3 left-4 w-fit">
              <div class="flex items-center justify-center space-x-2">
                <span class="select-none">图标</span>
                <span class="icon-[lucide--image]"></span>
              </div>
            </div>

            <!-- 上传图片 -->
            <div
              class="flex w-full h-full items-center justify-center"
              @drop.prevent="handleDrag"
              v-if="icon === null">
              <button class="h-full inline-flex items-center space-x-4" @click="openFileInput">
                <span class="select-none">拖拽/粘贴/上传</span>
                <span class="icon-[lucide--arrow-big-up-dash] size-8"></span>
              </button>
              <input
                type="file"
                @change="onFileChange"
                ref="fileInput"
                class="hidden"
                accept="image/*" />
            </div>

            <div
              class="indicator border size-24 rounded-md border-base-content border-dashed"
              v-if="icon !== null">
              <div class="indicator-item">
                <span class="btn btn-sm icon-[lucide--circle-x]" @click="icon = null"></span>
              </div>
              <div class="avatar p-2">
                <div class="rounded-xl">
                  <img ref="iconRef" @error="imageLoadError" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <textarea
          class="textarea w-full textarea-bordered hover:border-info"
          placeholder="网址介绍"
          v-model="description"></textarea>

        <div class="flex justify-between w-full">
          <div class="tooltip tooltip-bottom" data-tip="自动补全信息">
            <button class="btn btn-sm btn-primary" @click="queryMeta">一键填写</button>
          </div>
          <button class="btn btn-sm btn-primary">提交</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { net } from '@/logic'

const fileInput = useTemplateRef<HTMLInputElement | null>('fileInput')
const iconRef = useTemplateRef<HTMLImageElement>('iconRef')

const link = ref('')
const name = ref('')
const proxy = ref(false)
const icon = ref<File | null>(null)
const description = ref('')

function imageLoadError(event: Event) {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
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
  const data = await net(link.value)
  if (data) {
    Message({ message: '获取成功...' })
    name.value = name.value || data.title
    description.value = description.value || data.description || ''
    proxy.value = data.proxy || false
  } else {
    Message({ message: '获取失败...', type: 'warn' })
  }
}
</script>
