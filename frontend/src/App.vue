<template>
  <main class="flex w-screen h-screen bg-base-100 text-base-content">
    <div>
      <!-- 主界面左侧栏 -->
      <SideArea></SideArea>
    </div>
    <div class="right relative flex flex-col flex-1 h-full">
      <!-- 主界面顶部 -->
      <HeaderArea />
      <!-- 主界面内容区域 -->
      <!-- 需要给个锚点`anchor` 用于返回顶部 -->
      <!-- pb-20 给手机端菜单预留空间 -->
      <div class="anchor flex-1 w-full flex flex-col overflow-y-scroll pb-20 md:pb-0">
        <router-view class="content"></router-view>
        <FooterArea />
      </div>
    </div>
  </main>
</template>
