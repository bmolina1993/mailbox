<script setup>
import { onBeforeMount, provide, reactive, ref } from "vue";
import { Searcher, Mails } from "./components/";
import { useFetch, fetchRandomUser } from "./utils/";

let dataAPI = reactive({ data: [] });
let dataSearcher = reactive({ data: [] });
let dataRandomUser = reactive({ data: [] });
const showModalMenu = ref(false);

onBeforeMount(async () => {
  dataRandomUser.data = await fetchRandomUser();
  dataAPI.data = await useFetch();

  //data buscador inicia con set de datos API
  dataSearcher.data = dataAPI.data;
});

provide("dataAPI", dataAPI);
provide("dataSearcher", dataSearcher);
provide("dataRandomUser", dataRandomUser);
provide("showModalMenu", showModalMenu);
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
