<script setup>
import { inject, ref } from "vue";
import { iconRecibidos, iconEnviados, iconPapelera, iconTodos } from "./img";
import { proxyToObject } from "../utils";

// ------------------------
// injeccion datos globales
// ------------------------
const dataUserSelected = inject("dataUserSelected");
const dataSearcher = inject("dataSearcher");
const dataUserFolderSelected = inject("dataUserFolderSelected");
const isActiveFolder = inject("isActiveFolder");

// ---------
// funciones
// ---------
//filtra por carpeta seleccionada
const filterByFolder = (event) => {
  const folder = event.target.id;

  //activa seleccion de carpeta
  isActiveFolder.value = folder;

  const dataFiltered = [...dataUserSelected.data].filter((item) =>
    item.folder.includes(folder)
  );

  dataSearcher.data = proxyToObject(dataFiltered);
  dataUserFolderSelected.data = proxyToObject(dataFiltered);
};
</script>

<template>
  <!-- lista de carpetas -->
  <section class="px-5">
    <div
      @click="filterByFolder"
      id="inbox"
      :class="{ activeFolder: isActiveFolder == 'inbox' }"
      class="min- flex h-10 cursor-pointer items-center gap-x-2.5 rounded hover:bg-darkSecondary"
    >
      <img class="ml-2.5 w-6" :src="iconRecibidos" />
      Recibidos
    </div>
    <div
      @click="filterByFolder"
      id="sent_items"
      :class="{ activeFolder: isActiveFolder == 'sent_items' }"
      class="flex h-10 cursor-pointer items-center gap-x-2.5 rounded hover:bg-darkSecondary"
    >
      <img class="ml-2.5 w-6" :src="iconEnviados" />
      Enviados
    </div>
    <div
      @click="filterByFolder"
      id="deleted_items"
      :class="{ activeFolder: isActiveFolder == 'deleted_items' }"
      class="flex h-10 cursor-pointer items-center gap-x-2.5 rounded hover:bg-darkSecondary"
    >
      <img class="ml-2.5 w-6" :src="iconPapelera" />
      Papelera
    </div>
    <div
      @click="filterByFolder"
      id=""
      :class="{ activeFolder: isActiveFolder == '' }"
      class="flex h-10 cursor-pointer items-center gap-x-2.5 rounded hover:bg-darkSecondary"
    >
      <img class="ml-2.5 w-6" :src="iconTodos" />
      Todos
    </div>
  </section>
</template>
