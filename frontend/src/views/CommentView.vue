<template>
  <div class="h-full grid grid-cols-one justify-center select-none">
    <div class="p-6 my-8 w-full bg-base-200">
      <div class="flex flex-col">
        <div class="bg-base-300 p-4" v-for="comment in comments">
          <div class="flex rounded-lg">
            <div class="avatar static size-8">
              <div class="h-full rounded-xl">
                <div class="rounded-full flex h-full items-center justify-center bg-primary/70">
                  <span class="font-bold text-lg">{{ comment.nickname[0] }}</span>
                </div>
              </div>
            </div>

            <div class="">{{ comment.nickname }}</div>
            <div class="ml-auto text-sm">{{ comment.date }}</div>
          </div>
          <div class="mt-2">{{ comment.content }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { comments as getComments, type Comment } from '@/api'

const comments = ref<Comment[]>([])

onMounted(async () => {
  const data = await getComments()

  comments.value = data['data']['data']['comments']
})
</script>
