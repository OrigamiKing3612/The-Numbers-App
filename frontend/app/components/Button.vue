<template>
    <button class="button" @click="$emit('click')">
        <slot />
    </button>
</template>

<script setup lang="ts">
defineEmits<{
    (e: 'click'): void
}>()
</script>

<style scoped>
.button {
    background: var(--gradient-primary);
    color: var(--text);
    border-radius: var(--border-radius);
    font-size: 32px;
    padding: 16px 64px;
    font-weight: bold;
    box-shadow: var(--shadow-outer-primary), var(--shadow-inner-primary);
    border: none;
}
</style>
