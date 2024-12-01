<template>
  <div class="wave shadow-lg flex-col h-full hidden md:flex">
    <!-- 侧栏logo -->

    <ul
      v-if="showSetting.siderBar"
      @click="showSetting.siderBar = false"
      class="menu h-20 rounded-box lg:menu-lg">
      <li class="mt-auto">
        <a>
          <span class="icon-[lucide--chevrons-left] size-5" v-if="!showSetting.siderBar"></span>
          <span class="icon-[lucide--chevrons-right] size-5" v-else></span>
          <div>月离导航</div>
        </a>
      </li>
    </ul>

    <!-- 侧栏分类列表 -->
    <div class="w-56 flex-1 relative flex flex-col" v-if="showSetting.siderBar">
      <template v-if="isAdmin && finish">
        <AsyncAdminSideBar></AsyncAdminSideBar>
      </template>
      <template v-if="!isAdmin && finish">
        <AsyncUserSideBar></AsyncUserSideBar>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
const store = useBasicStore()
const { showSetting } = storeToRefs(store)

import { useAdminStatus } from '@/hooks/useAdminStatus'
const { isAdmin, finish } = useAdminStatus()

const AsyncAdminSideBar = defineAsyncComponent({
  loader: () => import('../admin/AdminSideBar.vue')
})
const AsyncUserSideBar = defineAsyncComponent({
  loader: () => import('../user/UserSideBar.vue')
})
</script>

<style>
.wave {
  background: radial-gradient(
    circle,
    rgba(148, 233, 255, 0.2) 0%,
    rgba(200, 236, 255, 0.3) 40%,
    rgba(255, 255, 255, 0.2) 100%
  );
}
</style>
