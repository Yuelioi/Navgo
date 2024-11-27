<template>
  <div class="shadow-lg md:flex flex-col h-full hidden bg-base-100">
    <!-- 侧栏logo -->
    <div class="h-20 flex items-center justify-center text-lg">
      <div
        class="font-bold"
        v-if="showSetting.siderBar"
        :class="{ hidden: showSetting.siderBar, 'md:block': showSetting.siderBar }">
        L
      </div>
      <div class="font-bold" v-else>月离导航</div>
    </div>

    <!-- 侧栏分类列表 -->
    <div
      class="w-56 flex-1 relative flex flex-col"
      :class="{
        'max-w-max': showSetting.siderBar,
        'md:w-full': showSetting.siderBar,
        hidden: showSetting.siderBar,
        'md:flex': showSetting.siderBar
      }">
      <template v-if="isAdmin">
        <AdminSideBar></AdminSideBar>
      </template>
      <template v-else>
        <UserSideBar></UserSideBar>
      </template>

      <ul class="menu rounded-box" @click="showSetting.siderBar = !showSetting.siderBar">
        <li>
          <a class="">
            <span class="icon-[lucide--chevrons-left] size-5" v-if="!showSetting.siderBar"></span>
            <span class="icon-[lucide--chevrons-right] size-5" v-else></span>
            <span v-if="!showSetting.siderBar" class="text-base">折叠</span>
          </a>
        </li>
      </ul>
    </div>

    <!-- 手机端折叠按钮 -->
    <div
      class="btn btn-square fixed md:hidden top-3/4 z-50"
      @click="showSetting.siderBar = !showSetting.siderBar"
      v-if="showSetting.siderBar">
      <span class="icon-[lucide--chevrons-right] size-6"></span>
    </div>
  </div>
</template>

<script setup lang="ts">
const store = useBasicStore()
const { isAdmin, showSetting } = storeToRefs(store)
</script>
