<template>
  <main class="flex w-screen h-screen bg-base-200">
    <div>
      <!-- 主界面左侧栏 -->
      <SideBar></SideBar>
    </div>
    <div class="right flex flex-col flex-1 h-full">
      <!-- 主界面顶部 -->
      <HeaderArea class="h-12" />
      <!-- 主界面内容区域 -->
      <div class="flex-1">
        <router-view></router-view>
      </div>
      <FooterArea class="mb-6" />
    </div>
  </main>
</template>

<script setup lang="ts"></script>
