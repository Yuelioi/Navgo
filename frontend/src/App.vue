<template>
  <main class="flex w-screen h-screen text-base-content bg-base-100">
    <div>
      <!-- 主界面左侧栏 -->
      <SideArea></SideArea>
    </div>
    <div class="right relative flex flex-col flex-1 h-full">
      <!-- 主界面顶部 -->
      <HeaderArea />
      <ContentArea></ContentArea>
    </div>
  </main>
</template>
