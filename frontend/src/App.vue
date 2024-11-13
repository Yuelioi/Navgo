<template>
  <main class="flex w-screen h-screen bg-base-100 text-base-content">
    <div class="shadow-lg">
      <!-- 主界面左侧栏 -->
      <SideBar></SideBar>
    </div>
    <div class="right flex flex-col flex-1 h-full overflow-hidden">
      <!-- 主界面顶部 -->
      <HeaderArea class="h-12 shadow-md" />
      <!-- 主界面内容区域 -->
      <div class="flex-1 overflow-y-scroll">
        <router-view></router-view>
      </div>
      <FooterArea class="mb-6" />
    </div>
  </main>
</template>

<script setup lang="ts"></script>
