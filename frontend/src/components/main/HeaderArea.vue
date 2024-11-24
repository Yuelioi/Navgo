<template>
  <div class="md:h-20">
    <template v-if="isAdmin">
      <AdminHeader></AdminHeader>
    </template>
    <template v-else>
      <UserHeader></UserHeader>
    </template>
  </div>
</template>

<script setup lang="ts">
const store = useBasicStore()
const { isAdmin } = storeToRefs(store)
</script>
