<template><div></div></template>

<script setup lang="ts">
const _themes = ['light', 'dark']
const { switchTheme } = useTheme(_themes)

onMounted(async () => {
    // 加载配置
})

onUnmounted(() => {})
</script>
