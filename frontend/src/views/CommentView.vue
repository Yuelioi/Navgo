<template>
  <div class="container">
    <div
      class="p-6 w-full h-full bg-base-100/50 backdrop-blur-md shadow-md rounded-lg z-10 relative">
      <div class="flex flex-col h-full space-y-8 justify-between">
        <!-- 评论区 -->
        <div class="flex flex-col space-y-4 overflow-y-scroll">
          <div class="bg-base-100/70 p-4 rounded-xl" v-for="comment in comments">
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
              class="textarea textarea-bordered bg-base-100/70 w-full"
              placeholder="提交建议或者反馈"
              v-model="form.content"></textarea>
          </div>

          <div class="flex items-center mt-2">
            <label
              for="text"
              class="input bg-base-100/70 input-bordered flex items-center input-sm gap-2 max-w-xs">
              <span class="select-none text-nowrap">昵称</span>
              <input type="text" placeholder="张三" class="" v-model="form.nickname" />
            </label>
            <button class="ml-auto btn btn-primary btn-sm" @click="submitComment">提交</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { addComment, comments as getComments, type Comment } from '@/api'

const form = reactive({ nickname: '', content: '' })

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
