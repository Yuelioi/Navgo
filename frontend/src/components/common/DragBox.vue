<template>
  <!-- 图标 -->
  <div
    class="flex border rounded-lg input-bordered items-center select-none w-full focus-within:outline-info"
    contenteditable="true"
    @paste.prevent="handlePaste">
    <div class="flex items-center justify-center w-full relative h-full min-h-48">
      <div class="absolute inset-3 left-4 w-fit">
        <div class="flex items-center justify-center space-x-2">
          <span class="select-none">{{ title }}</span>
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
        <input type="file" @change="onFileChange" ref="fileInput" class="hidden" accept="image/*" />
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
</template>

<script setup lang="ts">
import { imageLoadError } from '@/utils'

let icon = defineModel<File | null>('icon', { default: null, required: true })
defineProps({
  title: {
    type: String,
    required: false,
    default: '图标'
  }
})

const fileInput = useTemplateRef<HTMLInputElement | null>('fileInput')
const iconRef = useTemplateRef<HTMLImageElement>('iconRef')

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
</script>
