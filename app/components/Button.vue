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
    background: var(--background);
    color: var(--text);
    border-radius: var(--border-radius);
    font-size: 32px;
    padding: 16px 64px;
    font-weight: bold;
    box-shadow: inset 0px -2px 2.5px #fff5;
    border: none;
}
</style>
