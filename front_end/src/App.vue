<script setup>
import { onBeforeMount, provide, reactive, ref } from "vue";
import { Searcher, Mails, FolderTree, DetailMailDesktop } from "./components";
import { useFetch, fetchRandomUser, proxyToObject } from "./utils";
import { iconUser } from "./components/img";

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
const isActiveUser = ref("");
const isActiveFolder = ref("");

// ---------
// funciones
// ---------
onBeforeMount(async () => {
  dataRandomUser.data = await fetchRandomUser();
  dataAPI.data = await useFetch();

  //data [buscador] inicia con datos API
  const userFiltered = dataAPI.data.sort((a, z) =>
    a.index.localeCompare(z.index)
  )[0].index;

  const dataFiltered = dataAPI.data.filter(
    (item) => item.index == userFiltered
  );

  dataUserSelected.data = dataFiltered;
  dataUserFolderSelected.data = dataFiltered;
  dataSearcher.data = dataFiltered;
  isActiveUser.value = dataFiltered[0].index;

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
provide("isActiveUser", isActiveUser);
provide("isActiveFolder", isActiveFolder);
</script>

<template>
  <header
    class="mt-3 lg:mt-0 lg:flex lg:items-center lg:gap-x-2.5 lg:bg-darkPrimary lg:py-3"
  >
    <h1
      class="hidden lg:block lg:w-1/6 lg:pr-2.5 lg:text-xl lg:font-bold lg:text-white"
    >
      Mailbox connect
    </h1>
    <Searcher />

    <!-- usuarios -->
    <figure
      class="hidden w-1/2 gap-x-2.5 overflow-y-hidden overflow-x-scroll overscroll-x-contain whitespace-nowrap py-2.5 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full lg:flex"
    >
      <img
        v-for="item in dataRandomUser.data"
        :src="item.picture"
        :id="item.name"
        @click="getUser"
        class="scale-105 cursor-pointer rounded-full border-2 border-transparent"
        :class="{ activeUser: item.name == isActiveUser }"
      />
    </figure>
  </header>

  <main class="mt-3 lg:relative lg:mt-0">
    <!-- para ingresar color en barra lateral izquierdo, ya que por por padding queda con bg madre -->
    <div
      id="leftBGMain"
      class="hidden lg:absolute lg:block lg:h-full lg:w-5 lg:bg-darkPrimary"
    ></div>
    <!-- contenido principal, con division de bloques en porcentajes como el header -->
    <section id="mainBG" class="lg:flex lg:gap-x-2.5 lg:px-5">
      <!-- barra lateral izquierdo, usuario y carpetas -->
      <div class="hidden lg:block lg:w-1/6 lg:bg-darkPrimary lg:text-white">
        <!-- nombre usuario seleccionado -->
        <div class="flex gap-x-2.5 pt-3 pl-5 pb-2.5">
          <img class="w-6" :src="iconUser" />
          {{ isActiveUser }}
        </div>

        <FolderTree />
      </div>

      <!-- lista de correos -->
      <Mails />

      <!-- detalle correo -->
      <DetailMailDesktop />
    </section>
  </main>
</template>
