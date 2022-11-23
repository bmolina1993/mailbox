<script setup>
import { onBeforeMount, provide, reactive, ref } from "vue";
import { Searcher, Mails } from "./components/";
import { useFetch, fetchRandomUser } from "./utils/";
import { iconUser } from "./components/img/";

let dataAPI = reactive({ data: [] });
let dataSearcher = reactive({ data: [] });
let dataUserSelected = reactive({ data: [] });
let dataRandomUser = reactive({ data: [] });
const showModalMenu = ref(false);
const srcUser = ref(iconUser);

onBeforeMount(async () => {
  dataRandomUser.data = await fetchRandomUser();
  dataAPI.data = await useFetch();

  //data [buscador] inicia con datos API
  //dataSearcher.data = dataAPI.data.filter((item) => item.index == "allen-p");
  dataUserSelected.data = dataAPI.data.filter(
    (item) => item.index == "allen-p"
  );
  dataSearcher.data = dataUserSelected.data;

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

//datos globales
provide("dataAPI", dataAPI);
provide("dataSearcher", dataSearcher);
provide("dataUserSelected", dataUserSelected);
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
