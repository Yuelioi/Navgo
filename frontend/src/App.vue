<template>
  <main class="flex w-screen h-screen bg-base-100 text-base-content">
    <div class="shadow-lg">
      <!-- 主界面左侧栏 -->
      <SideBar></SideBar>
    </div>
    <div class="right relative flex flex-col flex-1 h-full">
      <!-- 主界面顶部 -->
      <HeaderArea class="h-20 shadow-md" />
      <!-- 主界面内容区域 -->
      <div class="flex-1 w-full flex flex-col overflow-y-scroll">
        <router-view class="content"></router-view>
        <FooterArea />
      </div>
    </div>
  </main>
</template>

<style scoped>
.content {
  background-image: linear-gradient(
      #d9d9d9 1px,
      transparent 1px,
      transparent calc(100% - 1px),
      #d9d9d9 calc(100% - 1px)
    ),
    linear-gradient(
      90deg,
      #d9d9d9 1px,
      transparent 1px,
      transparent calc(100% - 1px),
      #d9d9d9 calc(100% - 1px)
    );

  background-size: 25px 25px;
}
[data-theme='dark'] .content {
  background-image: linear-gradient(
      #3d3d3d 1px,
      transparent 1px,
      transparent calc(100% - 1px),
      #3d3d3d calc(100% - 1px)
    ),
    linear-gradient(
      90deg,
      #3d3d3d 1px,
      transparent 1px,
      transparent calc(100% - 1px),
      #3d3d3d calc(100% - 1px)
    );

  background-size: 25px 25px;
}
</style>
