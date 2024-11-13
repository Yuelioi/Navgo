<template>
  <div class="">
    <!-- 删除需要权限 -->
    {{ collection?.title }}
  </div>
</template>

<script lang="ts" setup>
import type { Collection } from '@/api'

const store = useBasicStore()
let { collections } = storeToRefs(store)

const collection = ref<Collection>()

const route = useRoute()

const fetchCollectionById = (id: string) => {
  const decodedId = decodeURIComponent(id)
  collection.value = collections.value.get(decodedId)
}

onMounted(() => {
  onMounted(() => {
    console.log('site')
  })
  console.log(collections.value)
  const id = route.params['id']
  console.log(id)

  if (id) {
    fetchCollectionById(String(id))
  }
})

watch(route, (newRoute) => {
  const id = newRoute.params['id']
  if (id) {
    fetchCollectionById(String(id))
  }
})
</script>
