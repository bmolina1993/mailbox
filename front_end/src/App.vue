<script setup>
import { onBeforeMount, onMounted, provide, reactive } from "vue";
import { Searcher, Mails } from "./components/";
import { useFetch } from "./utils/fetch";

let dataAPI = reactive({ data: [] });

onBeforeMount(async () => {
  dataAPI.data = await useFetch();
});

provide("dataAPI", dataAPI);
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
