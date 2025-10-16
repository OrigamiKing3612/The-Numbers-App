<template>
    <div class="container">
        <NumberDisplay v-model="number" v-model:checked_status="checked_status" />
        <ButtonGrid v-model="number" />
        <Button @click="dothething" class="checkin-button">
            Check in
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

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.checkin-button {
    margin-top: 30px;
    user-select: none;
    width: 100%;
    height: 80px;
}
</style>
