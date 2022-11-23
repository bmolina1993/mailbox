<script setup>
import { inject, reactive, ref } from "vue";
import { getDate, getFirstLetter, proxyToObject } from "../utils/";
import { iconArrowLeft } from "./img/";

const dataAPI = inject("dataAPI");
const dataSearcher = inject("dataSearcher");
const showModalMenu = inject("showModalMenu");
const dataRandomUser = inject("dataRandomUser");
const srcUser = inject("srcUser");

const showModal = ref(false);
const isActiveToList = ref(false);
const isActiveUser = ref("");
let dataModal = reactive({ data: {} });

const closeModal = () => {
  //quita scroll-y en body cuando se habre modal
  document.querySelector("body").style.overflowY = "";

  //cierra modal
  showModal.value = !showModal.value;
};

const toggle = (data) => {
  //save data modal
  dataModal.data = data;
  //abre modal
  showModal.value = !showModal.value;

  //quita scroll-y en body cuando se habre modal
  document.querySelector("body").style.overflowY = "hidden";
};

const toggleModalMenu = () => {
  //cierra modal
  showModalMenu.value = !showModalMenu.value;

  //quita scroll-y en body cuando se habre modal
  document.querySelector("body").style.overflowY = "";
};

const toggleToList = () => (isActiveToList.value = !isActiveToList.value);

//seleccion de perfil en menu
const getUser = (event) => {
  const user = event.target.id;

  const dataFiltered = [...dataAPI.data].filter((item) =>
    item.index.includes(user)
  );

  dataSearcher.data = proxyToObject(dataFiltered);

  //activa seleccion de usuario
  isActiveUser.value = user;

  //guarda ruta img de usuario para [searcher]
  srcUser.value = event.target.src;
};
</script>

<template>
  <!-- modal detalle correo -->
  <teleport to="#app">
    <section
      id="modal"
      v-show="showModal"
      class="fixed top-0 flex h-screen w-screen flex-col bg-gray-900 text-white"
    >
      <header class="flex gap-x-1 bg-darkSecondary py-2 px-5">
        <img class="cursor-pointer" @click="closeModal" :src="iconArrowLeft" />
        <p
          class="overflow-x-scroll overscroll-x-contain whitespace-nowrap py-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
        >
          {{ dataModal.data.subject }}
        </p>
      </header>
      <main
        class="mt-3 flex h-full flex-col gap-y-3 overflow-y-scroll px-5 pb-10"
      >
        <!-- info mail -->
        <div>
          <div class="flex gap-x-2">
            <!-- circulo de correo -->
            <div
              class="flex h-12 w-12 shrink-0 items-center justify-center rounded-full bg-tertiary text-xl text-black"
            >
              {{ getFirstLetter(dataModal.data.from) }}
            </div>

            <div class="w-full">
              <p>{{ getDate(dataModal.data.date, "long") }}</p>
              <p
                class="overflow-x-scroll overscroll-x-contain whitespace-nowrap py-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
              >
                {{ dataModal.data.from }}
              </p>
            </div>
          </div>
          <!--
            se agrega padding-left(pl-14) contemplando
            width(w-12) de circulo correo y gap (gap-x-2) del padre
          -->
          <ul
            @click="toggleToList"
            class="pl-14"
            :class="{ activeToList: isActiveToList }"
          >
            <li
              class="listTo overflow-x-scroll overscroll-x-contain whitespace-nowrap pb-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
              v-for="itemTo in dataModal.data.to"
            >
              {{ itemTo }}
            </li>
          </ul>
        </div>

        <!-- contenido correo -->
        <p class="whitespace-pre-line break-words">
          {{ dataModal.data.body }}
        </p>
      </main>
    </section>
  </teleport>

  <!-- modal menu latereal -->
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
      </div>

      <!-- menu lateral derecho -->
      <div
        @click="toggleModalMenu"
        class="h-full w-1/5 cursor-pointer bg-menu"
      ></div>
    </aside>
  </teleport>

  <!-- lista correo -->
  <div class="flex h-full w-full flex-col gap-y-px gap-x-2.5 py-2">
    <ul
      class="ulMain flex h-full w-full cursor-pointer flex-col justify-between px-5 text-white hover:bg-darkSecondary"
      v-for="item in dataSearcher?.data"
      @click="toggle(proxyToObject(item))"
    >
      <div class="flex w-full justify-between">
        <div class="flex w-5/6 items-center gap-x-2.5">
          <!-- circulo de correo -->
          <div
            class="circleMail flex h-8 w-8 shrink-0 items-center justify-center rounded-full text-black"
          >
            {{ getFirstLetter(item?.from) }}
          </div>
          <div
            class="overflow-x-scroll overscroll-x-contain whitespace-nowrap py-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
          >
            <li>{{ item?.from }}</li>
            <li>{{ item?.subject }}</li>
          </div>
        </div>

        <div>
          {{ `${getDate(item?.date, "mmm")} ${getDate(item?.date, "dd")}` }}
        </div>
      </div>

      <div
        class="overflow-x-scroll overscroll-x-contain whitespace-nowrap py-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
      >
        {{ item?.body }}
      </div>
    </ul>
  </div>
</template>

<style scoped>
/* agrega [To] al 1er item de lista */
.listTo:nth-child(1)::before {
  content: "To: ";
}
.listTo:nth-child(1n + 2)::before {
  content: "To: ";
  color: transparent;
}

/* pointer al 1er [To] */
.listTo:nth-child(1) {
  cursor: pointer;
}

/* a partir del 2do [To] oculta dato  */
.listTo:nth-child(1n + 2) {
  display: none;
}

/* activa lista [To] */
.activeToList .listTo:nth-child(1n + 2) {
  display: block;
}

/* usuario seleccionado */
.activeUser {
  /* border: 2px solid #9ccc66; */
  border: 2px solid #7986cc;
  transform: scale(1.05);
}
</style>
