import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

export function useAdminStatus() {
  const route = useRoute()
  const isAdmin = ref(false)

  // 更新管理员状态
  function updateAdminStatus() {
    isAdmin.value = route.fullPath.startsWith('/admin')
  }

  // 监听路由变化
  watch(
    route,
    () => {
      updateAdminStatus()
    },
    { immediate: true }
  )

  return { isAdmin }
}
