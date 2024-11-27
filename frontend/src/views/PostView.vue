<template>
  <div class="container">
    <div class="flex h-full">
      <div class="grid w-full p-6 space-y-2 bg-base-100 rounded-lg shadow-lg">
        <div class="">投稿</div>
        <div class="divider my-0"></div>

        <div class="">
          <div class="join input input-bordered gap-2 w-full">
            <div class="flex items-center justify-center pr-2 space-x-2">
              <span class="text-sm select-none">网址</span>
              <span class="icon-[lucide--link-2] size-5"></span>
            </div>
            <div class="divider"></div>
            <input
              type="text"
              class="grow input-bordered"
              placeholder="https://"
              v-model="form.link" />
          </div>
          <div v-if="errorFields?.link" class="text-warning !my-1">
            {{ errorFields.link[0].message }}
          </div>
        </div>

        <div class="">
          <div class="join flex input input-bordered gap-2 w-full">
            <div class="flex items-center justify-center pr-2 space-x-2">
              <span class="text-sm select-none">名称</span>
              <span class="icon-[lucide--case-lower] size-5"></span>
            </div>

            <input
              type="text"
              class="grow input-bordered"
              placeholder="Name"
              v-model="form.title" />
          </div>
          <div v-if="errorFields?.title" class="text-warning !my-1">
            {{ errorFields.title[0].message }}
          </div>
        </div>

        <!-- 图标 -->
        <div
          class="flex border rounded-lg input-bordered items-center select-none w-full focus-within:outline-info"
          contenteditable="true"
          @paste.prevent="handlePaste">
          <div class="flex items-center justify-center w-full relative h-full min-h-48">
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
              <button
                class="h-full inline-flex items-center space-x-4 btn-wide"
                @click="openFileInput">
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

            <div class="indicator size-36 rounded-md flex" v-if="icon !== null">
              <div class="avatar justify-center">
                <div class="rounded-xl p-2 border border-dashed border-base-content">
                  <img ref="iconRef" @error="imageLoadError" />
                  <div class="indicator-item">
                    <span class="btn btn-sm icon-[lucide--circle-x]" @click="icon = null"></span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <textarea
          class="textarea w-full textarea-bordered"
          placeholder="网址介绍"
          v-model="form.description"></textarea>
        <div v-if="errorFields?.description" class="text-warning !my-1">
          {{ errorFields.description[0].message }}
        </div>

        <div class="flex justify-between w-full">
          <div class="tooltip tooltip-bottom" data-tip="自动补全信息">
            <button class="btn btn-sm btn-primary" @click="queryMeta">一键填写</button>
          </div>
          <button class="btn btn-sm btn-primary" :disabled="!pass" @click="handleSubmit">
            提交
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { addCollection } from '@/api'
import { net } from '@/logic'
import { useAsyncValidator } from '@vueuse/integrations/useAsyncValidator'
import type { Rules } from 'async-validator'
import { imageLoadError } from '@/utils'
const fileInput = useTemplateRef<HTMLInputElement | null>('fileInput')
const iconRef = useTemplateRef<HTMLImageElement>('iconRef')

const form = reactive({ link: '', title: '', description: '' })
const rules: Rules = {
  link: {
    type: 'url',
    min: 5,
    required: true
  },
  title: {
    type: 'string',
    min: 2,
    required: true
  },
  description: [
    {
      type: 'string',
      max: 200,
      required: true
    }
  ]
}

const { pass, errorFields } = useAsyncValidator(form, rules)

const icon = ref<File | null>(null)

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
    form.link = form.link || data.title
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

  const config = {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }

  const result = await addCollection(formData, config)
  if (result.data.code > 0) {
    Message({ message: '提交成功' })
  } else {
    Message({ message: result.data.msg, type: 'warn' })
  }
}
</script>
