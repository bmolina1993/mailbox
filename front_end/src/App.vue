<script setup>
import { onBeforeMount, provide, reactive, ref } from "vue";
import { Searcher, Mails } from "./components/";
import { useFetch, fetchRandomUser } from "./utils/";
import { iconUser } from "./components/img/";

// ---------
// variables
// ---------
let dataAPI = reactive({ data: [] });
let dataSearcher = reactive({ data: [] });
let dataUserSelected = reactive({ data: [] });
let dataUserFolderSelected = reactive({ data: [] });
let dataRandomUser = reactive({ data: [] });
let dataModal = reactive({ data: {} });
const showModal = ref(false);
const showModalMenu = ref(false);
const srcUser = ref(iconUser);

// ---------
// funciones
// ---------
onBeforeMount(async () => {
  dataRandomUser.data = await fetchRandomUser();
  dataAPI.data = await useFetch();

  //data [buscador] inicia con datos API
  const dataFiltered = dataAPI.data.filter((item) => item.index == "allen-p");

  dataUserSelected.data = dataFiltered;
  dataUserFolderSelected.data = dataFiltered;
  dataSearcher.data = dataFiltered;

  // ---------------------------------
  //agrega nombre usuario a randomUser
  // ---------------------------------
  const auxUser = [];
  dataAPI.data.forEach((item) => {
    auxUser.push(item.index);
  });

  //quita duplicados
  //se queda con lista de nombre usuarios de API mail
  const user = [...new Set(auxUser)];

  //ordena datos
  user.sort((a, z) => a.localeCompare(z));

  //agrega nombre a randomUser por indice [user]
  user.forEach((item, idx) => {
    dataRandomUser.data[idx].name = item;
  });
});

// --------------
// datos globales
// --------------
provide("dataAPI", dataAPI);
provide("dataSearcher", dataSearcher);
provide("dataUserSelected", dataUserSelected);
provide("dataUserFolderSelected", dataUserFolderSelected);
provide("dataModal", dataModal);
provide("dataRandomUser", dataRandomUser);
provide("showModal", showModal);
provide("showModalMenu", showModalMenu);
provide("srcUser", srcUser);
</script>

<template>
  <header class="mt-3">
    <Searcher />
  </header>

  <main class="mt-3">
    <Mails />
  </main>
</template>
