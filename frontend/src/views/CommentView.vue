<template>
  <div class="h-full pb-32 md:pb-12 container">
    <div class="p-6 my-8 w-full h-full bg-base-200 rounded-lg">
      <div class="flex flex-col h-full space-y-8 justify-between">
        <!-- 评论区 -->
        <div class="flex flex-col space-y-4 overflow-y-scroll">
          <div class="bg-base-300 p-4 rounded-xl" v-for="comment in comments">
            <div class="flex">
              <div class="avatar static size-8">
                <div class="h-full rounded-xl">
                  <div class="rounded-full flex h-full items-center justify-center bg-primary/70">
                    <span class="font-bold text-lg">{{ comment.nickname[0] }}</span>
                  </div>
                </div>
              </div>

              <div class="ml-2">{{ comment.nickname }}</div>
              <div class="ml-auto text-sm">{{ comment.date }}</div>
              <div class="divider"></div>
            </div>
            <div class="mt-2">{{ comment.content }}</div>
          </div>
        </div>

        <!-- 发布 -->
        <div class="w-full flex flex-col">
          <div class="">
            <textarea
              class="textarea textarea-bordered w-full"
              placeholder="提交建议或者反馈"
              v-model="form.content"></textarea>

            <div v-if="errorFields?.content" class="text-warning top-full !my-1">
              {{ errorFields.content[0].message }}
            </div>
          </div>

          <div class="flex items-center mt-2">
            <label
              for="text"
              class="input input-bordered flex items-center input-sm gap-2 max-w-xs">
              <span class="select-none text-nowrap">昵称</span>
              <input type="text" placeholder="张三" class="" v-model="form.nickname" />
            </label>
            <button class="ml-auto btn btn-primary btn-sm" :disabled="!pass" @click="submitComment">
              提交
            </button>
          </div>

          <div v-if="errorFields?.nickname" class="text-warning !my-1">
            {{ errorFields.nickname[0].message }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { addComment, comments as getComments, type Comment } from '@/api'

import type { Rules } from 'async-validator'

import { useAsyncValidator } from '@vueuse/integrations/useAsyncValidator'

const form = reactive({ nickname: '', content: '' })
const rules: Rules = {
  nickname: {
    type: 'string',
    min: 2,
    required: true
  },

  content: [
    {
      type: 'string',
      min: 2,
      max: 200,
      required: true
    }
  ]
}

const { pass, errorFields } = useAsyncValidator(form, rules)

const comments = ref<Comment[]>([])

async function submitComment() {
  const resp = await addComment(form)
  if (resp.data.code >= 0) {
    Message({ message: '发布成功' })
    const data = await getComments()
    comments.value = data['data']['data']['comments']
  } else {
    Message({ message: resp.data.msg, type: 'warn' })
  }
}

onMounted(async () => {
  const resp = await getComments()
  if (resp.data.code >= 0) {
    comments.value = resp['data']['data']['comments']
  }
})
</script>
