<script setup>
import { onBeforeMount, provide, reactive, ref } from "vue";
import { Searcher, Mails } from "./components/";
import { useFetch, fetchRandomUser } from "./utils/";
import { iconUser } from "./components/img/";

let dataAPI = reactive({ data: [] });
let dataSearcher = reactive({ data: [] });
let dataRandomUser = reactive({ data: [] });
const showModalMenu = ref(false);
const srcUser = ref(iconUser);

onBeforeMount(async () => {
  dataRandomUser.data = await fetchRandomUser();
  dataAPI.data = await useFetch();

  //data buscador inicia con set de datos API
  dataSearcher.data = dataAPI.data;

  // ---------------------------------
  //agrega nombre usuario a randomUser
  // ---------------------------------
  const auxUser = [];
  //const auxFolder = [];
  dataAPI.data.forEach((item) => {
    auxUser.push(item.index);
    /*
    if (
      item.folder == "inbox" ||
      item.folder == "sent_items" ||
      item.folder == "deleted_items" ||
      item.folder == "all_documents"
    ) {
      auxFolder.push(item.folder);
    }
     */
  });

  //quita duplicados
  //se queda con lista de nombre usuarios de API mail
  const user = [...new Set(auxUser)];
  //const folder = [...new Set(auxFolder)];

  //ordena datos
  user.sort((a, z) => a.localeCompare(z));
  //folder.sort((a, z) => a.localeCompare(z));

  //agrega nombre a randomUser por indice [user]
  user.forEach((item, idx) => {
    dataRandomUser.data[idx].name = item;
  });

  //console.log("user:", user);
  //console.log("folder:", folder);
});

//datos globales
provide("dataAPI", dataAPI);
provide("dataSearcher", dataSearcher);
provide("dataRandomUser", dataRandomUser);
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

<style scoped></style>
