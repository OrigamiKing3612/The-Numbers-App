<template>
    <div class="button-grid">
        <div v-for="n in 9" :key="n">
            <button @click="addNumber(n)" class="keypad-button">
                {{ n }}
            </button>
        </div>
        <button @click="number = ''" class="keypad-button darker clear-button">
            clear
        </button>
        <button @click="addNumber(0)" class="keypad-button">
            0
        </button>
        <button @click="number = number.slice(0, -1)" class="keypad-button darker delete-button">
            <IconDelete style="width: 20px;" />
        </button>
    </div>
</template>

<script setup lang="ts">
const addNumber = (n: number) => {
    if (number.value.length >= 4) return
    number.value = number.value + `${n}`
}

const number = defineModel<string>({ required: true })
</script>

<style scoped>
.keypad-button {
    display: flex;
    justify-content: center;
    align-items: center;
    border: none;
    background: var(--keypad-background);
    border-radius: var(--border-radius);
    padding: 10px 40px;
    color: var(--dark-text);
    font-weight: bold;
    font-size: 48px;
    height: 100%;
}

.button-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    align-items: stretch;
    justify-items: stretch;
    gap: 15px;
}

.darker {
    background: var(--background);
    color: var(--text)
}

.clear-button {
    font-size: 10px;
}
</style>
