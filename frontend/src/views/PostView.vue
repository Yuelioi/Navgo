<template>
  <div class="container">
    <div class="flex h-full">
      <div
        class="grid w-full p-6 space-y-2 backdrop-blur-md z-10 bg-base-100/50 rounded-lg shadow-lg">
        <div class="">投稿</div>
        <div class="divider my-0"></div>

        <div class="">
          <div class="join bg-transparent input input-bordered gap-2 w-full">
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
        </div>

        <div class="">
          <div class="join bg-transparent flex input input-bordered gap-2 w-full">
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
        </div>

        <DragBox v-model:icon="icon"></DragBox>

        <textarea
          class="textarea w-full textarea-bordered bg-transparent"
          placeholder="网址介绍"
          v-model="form.description"></textarea>

        <div class="flex justify-between w-full">
          <div class="tooltip tooltip-bottom" data-tip="自动补全信息">
            <button class="btn btn-sm btn-primary" @click="queryMeta">一键填写</button>
          </div>
          <button class="btn btn-sm btn-primary" @click="handleSubmit">提交</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { addCollection } from '@/api'
import { net } from '@/logic'

const form = reactive({ link: '', title: '', description: '' })

const icon = ref<File | null>(null)

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
