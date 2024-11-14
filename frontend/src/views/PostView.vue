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
        <div class="join input input-bordered gap-2 w-full">
          <div class="flex border-r-2 items-center justify-center pr-2 space-x-2">
            <span class="text-sm">网址</span>
            <span class="icon-[lucide--link-2] size-5"></span>
          </div>
          <div class="divider"></div>
          <input type="text" class="grow input-bordered" placeholder="https://" />
        </div>

        <div class="join input input-bordered gap-2 w-full">
          <div class="flex border-r-2 items-center justify-center pr-2 space-x-2">
            <span class="text-sm">名称</span>
            <span class="icon-[lucide--case-lower] size-5"></span>
          </div>
          <div class="divider"></div>
          <input type="text" class="grow input-bordered" placeholder="Name" />
        </div>

        <div class="join input input-bordered gap-2 w-full">
          <div class="flex border-r-2 items-center justify-center pr-3 space-x-2">
            <span class="text-sm">国家</span>
            <span class="icon-[lucide--flag] size-4"></span>
          </div>
          <div class="divider"></div>
          <input type="text" class="grow input-bordered" placeholder="Country" />
        </div>
        <div class="join input input-bordered gap-2 w-full">
          <div class="flex border-r-2 items-center justify-center pr-3 space-x-2">
            <span class="text-sm">代理</span>
            <span class="icon-[lucide--utility-pole]"></span>
          </div>
          <div class="divider"></div>
          <div class="flex flex-1 items-center">
            <span class="text-sm opacity-45">如果需要科学上网, 请打勾</span>

            <input type="checkbox" checked class="ml-auto checkbox" />
          </div>
        </div>

        <!-- 图标 -->
        <div class="flex border rounded-lg input-bordered w-full">
          <div class="flex items-center justify-center w-full h-24 relative">
            <div class="absolute inset-3 left-4 w-fit">
              <div class="flex items-center justify-center space-x-2">
                <span>图标</span>
                <span class="icon-[lucide--image]"></span>
              </div>
            </div>
            <div class="absolute right-4 top-3">
              <div class="flex btn btn-accent btn-square btn-sm">
                <span class="icon-[lucide--x]"></span>
              </div>
            </div>

            <div class="flex items-center justify-center" v-if="icon == ''">
              <button class="btn btn-ghost btn-outline">
                <span class="icon-[lucide--arrow-big-up-dash] size-8"></span>
              </button>
            </div>
            <div class="" v-else>
              <div class="avatar">
                <div class="rounded-xl">
                  <img :src="icon" @error="imageLoadError" />
                </div>
              </div>
            </div>
          </div>
        </div>
        <textarea class="textarea w-full textarea-bordered" placeholder="网址介绍"></textarea>

        <div class="flex justify-between w-full">
          <div class="tooltip tooltip-bottom" data-tip="自动补全信息">
            <button class="btn btn-sm btn-primary">一键填写</button>
          </div>
          <button class="btn btn-sm btn-primary">提交</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const icon = ref('http://nav.yuelili.com/wp-content/uploads/82cf7-www.mojidict.com.png')
// const icon = ref('')

function imageLoadError(event: Event) {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}
</script>
