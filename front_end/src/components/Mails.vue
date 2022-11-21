<script setup>
import { inject } from "vue";
import { getDate } from "../utils/";
const dataAPI = inject("dataAPI");

//fondo [ul]
// bg-darkPrimary
//class="overflow-x-scroll whitespace-nowrap"
// id="leftInfo"
</script>

<template>
  <div class="flex h-20 w-full flex-col gap-y-2.5 px-5 py-2">
    <ul
      style="border: 1px solid"
      class="flex h-full w-full justify-between text-white"
      v-for="item in dataAPI.data"
    >
      <div class="flex w-full justify-between">
        <div class="flex w-5/6 items-center gap-x-2.5">
          <div class="h-8 w-8 shrink-0 rounded-full bg-primary"></div>

          <div class="overflow-x-scroll whitespace-nowrap">
            <li>{{ item.from }}</li>
            <li>{{ item.subject }}</li>
          </div>
        </div>

        <div>
          {{ `${getDate(item.date, "mmm")} ${getDate(item.date, "dd")}` }}
        </div>
      </div>
    </ul>
  </div>
</template>

<style scoped></style>
