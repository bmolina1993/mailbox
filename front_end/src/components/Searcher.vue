<script setup>
import { iconMenu } from "./img/";
</script>

<template>
  <div
    class="flex justify-between gap-x-1 rounded-full bg-darkSecondary px-2 py-1"
  >
    <div class="flex w-full gap-x-1">
      <img :src="iconMenu" />
      <input
        type="text"
        name="searchId"
        id="searchId"
        class="h-6 w-full bg-transparent text-white outline-0"
        placeholder="Buscar..."
      />
    </div>
    <div class="h-6 w-6 shrink-0 rounded-full bg-gray-300"></div>
  </div>
</template>

<style scoped></style>
