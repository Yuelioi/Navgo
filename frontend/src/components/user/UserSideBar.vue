<template>
  <Transition name="slide-fade">
    <div class="w-full h-full flex flex-col">
      <ul class="menu rounded-box h-full lg:menu-lg">
        <li v-for="(nav, index) in navs" @click="scrollToSection(nav.cid)">
          <a class="">
            <span :class="icons[index]" class="size-5"></span>
            <span class="text-base lg:text-lg">{{ nav.title }}</span>
          </a>
        </li>

        <div class="mt-auto"></div>

        <li v-if="token">
          <a
            @click="
              router.push({
                name: 'allNavs'
              })
            ">
            <span class="icon-[lucide--activity] size-6"></span>
            <span>后台</span>
          </a>
        </li>
        <li>
          <a @click.stop="showLoginWindow = true">
            <span class="icon-[lucide--user-round] size-6"></span>
            <span>登录</span>
          </a>
        </li>

        <li>
          <a href="https://github.com/Yuelioi/Navgo" target="_blank">
            <span class="icon-[lucide--github] size-6"></span>
            <span>Github</span>
          </a>
        </li>

        <VDialog title="登录" v-model:show="showLoginWindow">
          <div class="p-4">
            <div class="w-96 p-2 space-y-4">
              <label class="input input-bordered flex items-center gap-2">
                <span class="icon-[lucide--user]"></span>
                <input type="text" class="grow" placeholder="Username" v-model="form.username" />
              </label>
              <label class="input input-bordered flex items-center gap-2">
                <span class="icon-[lucide--key-round]"></span>
                <input type="password" class="grow" v-model="form.password" />
              </label>

              <input
                type="button"
                class="w-full grow btn btn-sm btn-primary"
                value="登录"
                @click="login" />
            </div>
          </div>
        </VDialog>
      </ul>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import router from '@/router'
import { icons } from '@/stores/icons'

const showLoginWindow = ref(false)
const store = useBasicStore()
const { navs, token } = storeToRefs(store)

import { VDialog } from '@/plugins/dialog'
import { auth, checkToken, type User } from '@/api'

const form = reactive<User>({
  iD: 0,
  createdAt: '',
  updatedAt: '',
  username: '',
  password: ''
})

function scrollToSection(cid?: string) {
  if (router.currentRoute.value.name != 'home') {
    router.push({
      name: 'home'
    })
  }

  if (cid) {
    const element = document.getElementById(cid)
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' })
    }
  }
}

async function login() {
  if (form) {
    const resp = await auth(form)
    token.value = resp.data.data.token
    showLoginWindow.value = false
    Message({ message: '登录成功' })
  }
}

onMounted(async () => {
  const resp = await checkToken(token.value)
  if (resp.data.code < 0) {
    token.value = ''
  }
})
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
