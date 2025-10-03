<template>
    <div class="container">
        <h1>The Numbers App</h1>
        <NumberDisplay v-model="number" v-model:checked_status="checked_status" />
        <ButtonGrid v-model="number" />
        <Button @click="dothething" class="checkin-button">
            Check In
        </Button>
    </div>
</template>

<script setup lang="ts">
const number = ref<string>("")

const checked_status = ref<'out' | 'in'>("out")
const dothething = () => {
    alert('You checked in!')
}

</script>

<style scoped lan>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.checkin-button {
    margin-top: 30px;
    user-select: none;
    width: 100%;
}
</style>
