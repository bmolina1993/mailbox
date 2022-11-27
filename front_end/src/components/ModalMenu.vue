<script setup>
import { inject } from "vue";
import { iconUser } from "./img";
import { proxyToObject } from "../utils";
import { FolderTree } from "../components";

// ------------------------
// injeccion datos globales
// ------------------------
const showModalMenu = inject("showModalMenu");
const dataRandomUser = inject("dataRandomUser");
const dataUserSelected = inject("dataUserSelected");
const dataSearcher = inject("dataSearcher");
const dataAPI = inject("dataAPI");
const srcUser = inject("srcUser");
const isActiveUser = inject("isActiveUser");
const isActiveFolder = inject("isActiveFolder");

// ---------
// funciones
// ---------
const toggleModalMenu = () => {
  //cierra modal
  showModalMenu.value = !showModalMenu.value;

  //quita scroll-y en body cuando se habre modal
  //para no tomar scroll de lista correo
  document.querySelector("body").style.overflowY = "";
};

//seleccion de perfil en menu
const getUser = (event) => {
  const user = event.target.id;

  const dataFiltered = [...dataAPI.data].filter((item) =>
    item.index.includes(user)
  );

  dataSearcher.data = proxyToObject(dataFiltered);
  dataUserSelected.data = proxyToObject(dataFiltered);

  //activa seleccion de usuario
  isActiveUser.value = user;

  //guarda ruta img de usuario para [searcher]
  srcUser.value = event.target.src;

  //por cada cambio de usuario,
  //deja la carpeta seleccionada [Todos]
  isActiveFolder.value = "";
};
</script>

<template>
  <teleport to="#app">
    <aside
      id="modalMenu"
      v-show="showModalMenu"
      class="fixed top-0 flex h-screen w-screen text-white"
    >
      <!-- menu lateral izq -->
      <div class="h-full w-4/5 bg-darkPrimary pt-3">
        <header class="pb-3 pl-5 text-xl font-bold">Mailbox connect</header>
        <!-- usuarios -->
        <figure
          class="flex gap-x-2.5 overflow-y-hidden overflow-x-scroll overscroll-x-contain whitespace-nowrap border-y py-2.5 pl-5 pb-2.5 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
        >
          <img
            v-for="item in dataRandomUser.data"
            :src="item.picture"
            :id="item.name"
            @click="getUser"
            class="w-12 cursor-pointer rounded-full border-2 border-transparent"
            :class="{ activeUser: item.name == isActiveUser }"
          />
        </figure>

        <!-- nombre usuario seleccionado -->
        <div class="flex gap-x-2.5 pt-3 pl-5 pb-2.5">
          <img class="w-6" :src="iconUser" />
          {{ isActiveUser }}
        </div>

        <!-- lista de carpetas -->
        <FolderTree />
      </div>

      <!-- menu lateral derecho -->
      <div
        @click="toggleModalMenu"
        class="h-full w-1/5 cursor-pointer bg-menu"
      ></div>
    </aside>
  </teleport>
</template>
