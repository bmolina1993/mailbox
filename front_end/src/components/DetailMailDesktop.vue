<script setup>
import { inject, ref } from "vue";
import { getDate, getFirstLetter } from "../utils";

// ------------------------
// injeccion datos globales
// ------------------------
const dataModal = inject("dataModal");

// ---------
// variables
// ---------
const isActiveToList = ref(false);

// ---------
// funciones
// ---------
const toggleToList = () => (isActiveToList.value = !isActiveToList.value);
</script>

<template>
  <section
    class="mt-5 hidden h-[90vh] w-1/2 flex-col rounded-tl-lg rounded-tr-lg bg-darkPrimary text-white lg:flex"
  >
    <header
      class="flex gap-x-1 rounded-tl-lg rounded-tr-lg bg-darkPrimary py-2 px-5"
    >
      <p
        class="overflow-x-scroll overscroll-x-contain whitespace-nowrap py-1 scrollbar-thin scrollbar-track-darkSecondary scrollbar-thumb-darkPrimary scrollbar-track-rounded-full scrollbar-thumb-rounded-full"
      >
        {{ dataModal.data.subject }}
      </p>
    </header>
    <main
      class="mt-3 flex h-full flex-col gap-y-3 overflow-y-scroll px-5 pb-10"
      :class="'lg:overflow-y-scroll lg:overscroll-y-contain lg:scrollbar-thin lg:scrollbar-track-darkSecondary lg:scrollbar-thumb-darkPrimary lg:scrollbar-track-rounded-full lg:scrollbar-thumb-rounded-full'"
    >
      <!-- info mail -->
      <div>
        <div class="flex gap-x-2">
          <!-- circulo de correo -->
          <div
            :class="
              dataModal.data.from &&
              'flex h-12 w-12 shrink-0 items-center justify-center rounded-full bg-tertiary text-xl text-black'
            "
          >
            {{ getFirstLetter(dataModal.data.from) }}
          </div>

          <div class="w-full">
            <p>
              {{ dataModal.data.date && getDate(dataModal.data.date, "long") }}
            </p>
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
</style>
