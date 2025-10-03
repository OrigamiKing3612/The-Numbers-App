<template>
    <div class="container">
        <h1>The Numbers App</h1>
        <div class="button-grid">
            <div v-for="n in 9" :key="n">
                <button @click="dothething" class="keypad-button">
                    {{ n }}
                </button>
            </div>
            <button @click="dothething" class="keypad-button darker clear-button">
                Clear
            </button>
            <button @click="dothething" class="keypad-button">
                0
            </Button>
            <button @click="dothething" class="keypad-button darker">
                <IconDelete />
            </button>
        </div>
        <Button @click="dothething">
            Check In
        </Button>
    </div>
</template>

<script setup lang="ts">
const dothething = () => {
    alert('You checked in!')
}
</script>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.keypad-button {
    border: none;
    background: var(--keypad-background);
    border-radius: var(--border-radius);
    padding: 30px;
    color: var(--dark-text);
    font-weight: bold;
    font-size: 24px;
    margin: 8px;
}

.button-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    align-items: center;
    justify-items: center;
    /* width: 100%; */
}

.darker {
    background: var(--background);
}

.clear-button {
    font-size: 18px;
}

</style>
