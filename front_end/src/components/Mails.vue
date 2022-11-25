<script setup>
import { inject } from "vue";
import { DetailMail, ModalMenu } from "../components";
import { getDate, getFirstLetter, proxyToObject } from "../utils";

// ------------------------
// injeccion datos globales
// ------------------------
const dataSearcher = inject("dataSearcher");
const showModal = inject("showModal");
const dataModal = inject("dataModal");

// ---------
// funciones
// ---------
//accion cuando click sobre un correo
const toggle = (event, data) => {
  const widthScreen = event.view.visualViewport.width;

  //save data modal
  dataModal.data = data;

  //solo ejecutara codigo si la resolucion es >= 1024px
  //contempla media query creada para modo Desktop
  if (widthScreen < 1024) {
    //abre modal
    showModal.value = !showModal.value;

    //quita scroll-y en body cuando se habre modal
    //para no tomar scroll de lista correo
    document.querySelector("body").style.overflowY = "hidden";
  }
};
</script>

<template>
  <!-- Detalle correo -->
  <DetailMail />

  <!-- modal menu latereal -->
  <ModalMenu />

  <!-- lista correo -->
  <div
    class="flex h-full w-full flex-col gap-y-px gap-x-2.5 py-2 lg:mt-5 lg:h-[90vh] lg:w-1/3 lg:overflow-x-hidden lg:rounded-tl-lg lg:rounded-tr-lg lg:bg-darkPrimary lg:py-0"
    :class="'lg:overflow-y-scroll lg:overscroll-y-contain lg:scrollbar-thin lg:scrollbar-track-darkSecondary lg:scrollbar-thumb-darkPrimary lg:scrollbar-track-rounded-full lg:scrollbar-thumb-rounded-full'"
  >
    <ul
      class="ulMain flex h-full w-full cursor-pointer flex-col justify-between py-2.5 px-5 text-white hover:bg-darkSecondary lg:h-auto lg:rounded-tl-lg lg:rounded-tr-lg"
      v-for="item in dataSearcher?.data"
      @click="toggle($event, proxyToObject(item))"
    >
      <div class="flex w-full justify-between">
        <div class="flex w-5/6 items-center gap-x-2.5">
          <!-- circulo de correo -->
          <div
            class="circleMail flex h-8 w-8 shrink-0 items-center justify-center rounded-full text-black"
          >
            {{ getFirstLetter(item?.from) }}
          </div>
          <div class="overflow-hidden whitespace-nowrap">
            <li>{{ item?.from }}</li>
            <li>{{ item?.subject }}</li>
          </div>
        </div>

        <div>
          {{ `${getDate(item?.date, "mmm")} ${getDate(item?.date, "dd")}` }}
        </div>
      </div>

      <div class="overflow-hidden whitespace-nowrap">
        {{ item?.body }}
      </div>
    </ul>
  </div>
</template>
