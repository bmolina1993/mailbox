<script setup>
import { inject, ref } from "vue";
import { iconMenu } from "./img/";
import { proxyToObject } from "../utils/";

const dataAPI = inject("dataAPI");
const dataSearcher = inject("dataSearcher");
const dataUserSelected = inject("dataUserSelected");
const showModalMenu = inject("showModalMenu");
const srcUser = inject("srcUser");

const toggleModalMenu = () => {
  //abre modal
  showModalMenu.value = !showModalMenu.value;

  //quita scroll-y en body cuando se habre modal
  document.querySelector("body").style.overflowY = "hidden";
};

/*
filtra data por los siguientes campos:
  . body
  . from
  . subject
  . folder
*/
const onChangeInput = (event) => {
  const value = event.target.value.toLowerCase();

  //const dataFiltered = [...dataAPI.data].filter(
  const dataFiltered = [...dataUserSelected.data].filter(
    (item) =>
      item.body.toLowerCase().includes(value) ||
      item.from.toLowerCase().includes(value) ||
      item.subject.toLowerCase().includes(value) ||
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
      <img :src="iconMenu" @click="toggleModalMenu" class="cursor-pointer" />
      <input
        type="text"
        name="searchId"
        id="searchId"
        class="h-6 w-full bg-transparent text-white outline-0"
        placeholder="Buscar..."
        @input="onChangeInput"
      />
    </div>
    <img :src="srcUser" class="h-6 w-6 rounded-full" />
  </div>
</template>

<style scoped></style>
