<script setup>
import { inject } from "vue";
import { iconMenu, iconLupa } from "./img/";
import { proxyToObject } from "../utils/";

// ------------------------
// injeccion datos globales
// ------------------------
const dataSearcher = inject("dataSearcher");
const dataUserFolderSelected = inject("dataUserFolderSelected");
const showModalMenu = inject("showModalMenu");
const srcUser = inject("srcUser");

// ---------
// funciones
// ---------
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
*/
const onChangeInput = (event) => {
  const value = event.target.value.toLowerCase();

  const dataFiltered = [...dataUserFolderSelected.data].filter(
    (item) =>
      item.body.toLowerCase().includes(value) ||
      item.from.toLowerCase().includes(value) ||
      item.subject.toLowerCase().includes(value)
  );

  dataSearcher.data = proxyToObject(dataFiltered);
};
</script>

<template>
  <div
    class="flex justify-between gap-x-1 rounded-full bg-darkSecondary px-2 py-1 lg:h-9 lg:w-1/3 lg:rounded-t lg:rounded-b lg:py-1.5"
  >
    <div class="flex w-full gap-x-1">
      <img
        class="w-6 cursor-pointer lg:hidden"
        :src="iconMenu"
        @click="toggleModalMenu"
      />
      <img class="hidden lg:block lg:w-4.5" :src="iconLupa" />
      <input
        type="text"
        name="searchId"
        id="searchId"
        class="h-6 w-full bg-transparent text-white outline-0"
        placeholder="Buscar..."
        @input="onChangeInput"
      />
    </div>
    <img :src="srcUser" class="h-6 w-6 rounded-full lg:hidden" />
  </div>
</template>
