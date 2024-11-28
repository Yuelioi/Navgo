<template>
  <div class="container hidden sticky z-10 lg:block">
    <div class="card rounded-md shadow-lg backdrop-blur-md my-8" v-if="showSetting.footer">
      <div class="card-body flex-row justify-between">
        <div class="left xl:flex flex-col space-y-4 hidden">
          <h2 class="card-title">月离导航</h2>
          <p>站长很懒,所以懒得写介绍</p>
          <p class="text-sm">Copyright © 2024 月离导航 | 皖ICP备19025249号</p>
        </div>

        <div class="right">
          <div class="stats shadow-md bg-base-100/50 backdrop-blur-sm">
            <div class="stat">
              <div class="stat-figure text-primary">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  class="inline-block h-8 w-8 stroke-current">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <div class="stat-title">总浏览量</div>
              <div class="stat-value">{{ siteStats?.total_visitors }}</div>
              <div class="stat-desc"></div>
            </div>

            <div class="stat">
              <div class="stat-figure text-primary">
                <span class="icon-[lucide--link] size-6"></span>
              </div>
              <div class="stat-title">网址数量</div>
              <div class="stat-value">{{ siteStats?.links_count }}</div>
              <div class="stat-desc"></div>
            </div>

            <div class="stat">
              <div class="stat-figure text-primary">
                <span class="icon-[lucide--users-round] size-6"></span>
              </div>
              <div class="stat-title">昨日访客</div>
              <div class="stat-value">{{ siteStats?.last_day_visitors }}</div>
              <div class="stat-desc"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const store = useBasicStore()
let { isScrollDown, siteStats, showSetting } = storeToRefs(store)

const lastScrollTop = ref(0)

const handleScroll = () => {
  const currentScrollTop = window.scrollY || document.documentElement.scrollTop
  // 如果当前滚动位置大于上次滚动位置，说明是向下滚动
  isScrollDown.value = currentScrollTop > lastScrollTop.value
  lastScrollTop.value = currentScrollTop
}

onMounted(() => {
  // 检测滚动

  document.addEventListener('scroll', () => {
    console.log(111)
  })
  window.addEventListener('scroll', () => {
    console.log(111)
  })
})

onBeforeMount(() => {
  document.removeEventListener('scroll', handleScroll)
})
</script>
