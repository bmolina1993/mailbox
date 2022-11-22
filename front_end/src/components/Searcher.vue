<script setup>
import { inject } from "vue";
import { iconMenu } from "./img/";
import { proxyToObject } from "../utils/";

const dataAPI = inject("dataAPI");
const dataSearcher = inject("dataSearcher");

/*
filtra data por los siguientes campos:
  . body
  . from
  . subject
  . index = usuario
  . folder
*/
const onChangeInput = (event) => {
  const value = event.target.value.toLowerCase();

  const dataFiltered = [...dataAPI.data].filter(
    (item) =>
      item.body.toLowerCase().includes(value) ||
      item.from.toLowerCase().includes(value) ||
      item.subject.toLowerCase().includes(value) ||
      item.index.toLowerCase().includes(value) ||
      item.folder.toLowerCase().includes(value)
  );

  dataSearcher.data = proxyToObject(dataFiltered);
};
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
        @input="onChangeInput"
      />
    </div>
    <div class="h-6 w-6 shrink-0 rounded-full bg-gray-300"></div>
  </div>
</template>

<style scoped></style>
