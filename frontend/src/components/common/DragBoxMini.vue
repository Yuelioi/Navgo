<template>
  <!-- 图标 -->
  <div
    class="flex select-none w-full focus-within:outline-none"
    contenteditable="true"
    @paste.prevent="handlePaste">
    <div class="flex items-center justify-center w-full relative h-full">
      <!-- 上传图片 -->
      <div
        class="flex w-full h-full overflow-hidden items-center justify-center"
        @drop.prevent="handleDrag"
        v-if="icon === null">
        <button
          aria-label="上传"
          class="h-full inline-flex items-center space-x-4"
          @click="openFileInput">
          <span class="icon-[lucide--arrow-big-up-dash] size-8"></span>
        </button>
        <input
          type="file"
          @change="onFileChange"
          ref="fileInput"
          class="hidden focus-within:border-none"
          accept="image/*" />
      </div>

      <div class="indicator group flex" v-if="icon !== null">
        <div class="avatar justify-center">
          <div class="">
            <img ref="iconRef" alt="图标" @error="imageLoadError" />

            <span
              class="hidden group-hover:block center text-warning size-6 icon-[lucide--circle-x]"
              @click="icon = null"></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { imageLoadError } from '@/utils'

let icon = defineModel<File | null>('icon', { default: null, required: true })

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
