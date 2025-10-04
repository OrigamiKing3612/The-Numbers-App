<template>
    <div class="number-display">
        {{ number }}
    </div>
    <div class="message" v-if="number != ''">
        #{{ number }} is currently<br>
        <span :class="{'in': checked_status === 'in', 'out': checked_status === 'out'}">checked {{ checked_status }}</span>
    </div>
</template>

<script setup lang="ts">
const number = defineModel<string>({ required: true })
const checked_status = defineModel<'out' | 'in'>('checked_status', { required: true })
</script>

<style scoped>
.number-display {
    border: 2px solid var(--tertiary);
    margin: 20px 0;
    padding: 10px 20px;
    font-size: 48px;
    font-weight: bolder;
    text-align: center;
    align-content: center;
    width: 200px;
    height: 80px;
    border-radius: var(--border-radius);
    color: var(--dark-text);
}

.message {
    font-size: 24px;
    font-weight: bold;
    align-content: center;
    text-align: center;
    height: 70px;
    width: 300px;
    margin-bottom: 15px;
    color: var(--dark-text);

.in {
    color: var(--primary-alt);
}

.out {
    color: var(--primart);
}
}
</style>
