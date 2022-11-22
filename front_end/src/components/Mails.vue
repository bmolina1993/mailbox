<script setup>
import { inject } from "vue";
import { getDate, getFirstLetter } from "../utils/";
const dataAPI = inject("dataAPI");
</script>

<template>
  <div class="flex h-full w-full flex-col gap-y-px gap-x-2.5 px-5 py-2">
    <ul
      style="border: 1px solid"
      class="ulMain flex h-full w-full cursor-pointer flex-col justify-between text-white hover:bg-darkSecondary"
      v-for="item in dataAPI.data"
    >
      <div class="flex w-full justify-between">
        <div class="flex w-5/6 items-center gap-x-2.5">
          <!-- circulo de correo bg-primary -->
          <div
            class="circleMail flex h-8 w-8 shrink-0 items-center justify-center rounded-full text-black"
          >
            {{ getFirstLetter(item.from) }}
          </div>
          <div
            class="overflow-x-scroll overscroll-x-contain whitespace-nowrap py-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
          >
            <li>{{ item.from }}</li>
            <li>{{ item.subject }}</li>
          </div>
        </div>

        <div>
          {{ `${getDate(item.date, "mmm")} ${getDate(item.date, "dd")}` }}
        </div>
      </div>

      <div
        class="overflow-x-scroll overscroll-x-contain whitespace-nowrap py-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
      >
        {{ item.subject }}
      </div>
    </ul>
  </div>
</template>

<style scoped></style>
