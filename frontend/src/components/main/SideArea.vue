<template>
  <div class="shadow-lg flex flex-col h-full bg-base-100">
    <!-- 侧栏logo -->
    <div class="h-20 flex items-center justify-center text-lg">
      <div
        class="font-bold"
        v-if="isCollapsedNav"
        :class="{ hidden: isCollapsedNav, 'md:block': isCollapsedNav }">
        L
      </div>
      <div class="font-bold" v-else>月离导航</div>
    </div>

    <!-- 侧栏分类列表 -->
    <div
      class="w-56 flex-1 relative flex flex-col"
      :class="{
        'max-w-max': isCollapsedNav,
        'md:w-full': isCollapsedNav,
        hidden: isCollapsedNav,
        'md:flex': isCollapsedNav
      }">
      <template v-if="isAdmin">
        <AdminSideBar></AdminSideBar>
      </template>
      <template v-else>
        <UserSideBar></UserSideBar>
      </template>

      <ul class="menu rounded-box" @click="isCollapsedNav = !isCollapsedNav">
        <li>
          <a class="">
            <span class="icon-[lucide--chevrons-left] size-5" v-if="!isCollapsedNav"></span>
            <span class="icon-[lucide--chevrons-right] size-5" v-else></span>
            <span v-if="!isCollapsedNav" class="text-base">折叠菜单</span>
          </a>
        </li>
      </ul>
    </div>

    <!-- 手机端折叠按钮 -->
    <div
      class="btn btn-square fixed md:hidden top-3/4 z-50"
      @click="isCollapsedNav = !isCollapsedNav"
      v-if="isCollapsedNav">
      <span class="icon-[lucide--chevrons-right] size-6"></span>
    </div>
  </div>
</template>

<script setup lang="ts">
const store = useBasicStore()
const { isAdmin, isCollapsedNav } = storeToRefs(store)
</script>
