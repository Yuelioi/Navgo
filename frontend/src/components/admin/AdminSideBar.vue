<template>
  <Transition name="slide-fade">
    <div class="float-left w-full h-full">
      <ul class="menu" v-for="(nav, index) in adminNavs">
        <!-- 有子级 && 没折叠 -->
        <li v-if="showSetting.siderBar && nav.children">
          <details>
            <summary>
              <span :class="nav.icon" class="size-5"></span>
              <span class="md:text-lg">{{ nav.title }}</span>
            </summary>
            <ul v-if="nav.children">
              <li v-for="(child, index) in nav.children">
                <a
                  class=""
                  @click.stop="
                    router.push({
                      name: child.target
                    })
                  ">
                  <span :class="child.icon"></span>
                  <span>{{ child.title }}</span>
                </a>
              </li>
            </ul>
          </details>
        </li>
        <!-- 没子级 / 折叠 -->
        <li
          v-else
          @click.stop="
            router.push({
              name: nav.target
            })
          ">
          <!-- 没子级 -->
          <a v-if="!nav.children">
            <span :class="nav.icon" class="size-5"></span>
            <span class="md:text-lg">{{ nav.title }}</span>
          </a>
        </li>
      </ul>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import router from '@/router'

const store = useBasicStore()
const { showSetting } = storeToRefs(store)

interface Nav {
  cid: string
  title: string
  icon: string
  target: string
  collapse?: boolean // 父级需要加
  children?: Nav[]
}

const adminNavs = reactive<Nav[]>([
  {
    cid: 'guide',
    title: '导航',
    icon: 'icon-[lucide--navigation]',
    target: 'navs',
    collapse: false,
    children: [
      {
        cid: 'all',
        title: '所有导航',
        icon: 'icon-[lucide--leaf]',
        target: 'allNavs'
      },
      {
        cid: 'add',
        title: '添加导航',
        icon: 'icon-[lucide--square-plus]',
        target: 'addNav'
      },
      {
        cid: 'category',
        title: '导航分类',
        icon: 'icon-[lucide--layout-grid]',
        target: 'navCategory'
      }
    ]
  },
  {
    cid: 'announce',
    title: '公告',
    icon: 'icon-[lucide--bell]',
    target: 'announce'
  },
  {
    cid: 'approval',
    title: '审核',
    icon: 'icon-[lucide--scroll-text]',
    target: 'approval'
  },
  {
    cid: 'comment',
    title: '评论',
    icon: 'icon-[lucide--message-square-more]',
    target: 'adminComment'
  }
])
</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.3s ease-out;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(-20px);
  opacity: 0;
}
</style>
