<template>
  <div
    class="w-screen h-screen flex justify-center items-center bg-gradient-to-r from-[#4766f4] to-[#b6f3c9]"
  >
    <div class="w-96 h-56 relative text-white">
      <img
        class="relative object-cover w-full h-full rounded-xl"
        src="https://i.imgur.com/kGkSg1v.png"
      />
      <div class="w-full absolute px-8 top-8">
        <div class="flex justify-between">
          <div class="">
            <p class="font-light">Name</p>
            <p class="font-medium tracking-widest">Human</p>
          </div>
          <div class="w-14 h-14">
            <img class="object-cover" v-if="logo" :src="logo" />
          </div>
        </div>

        <div class="pt-1">
          <p class="font-light">Card Number</p>
          <input
            v-model="cardNumber"
            type="text"
            class="font-medium tracking-more-wider bg-transparent outline-none"
            placeholder="Add Number Here..."
          />
        </div>
        <div v-if="!tested" class="pt-6">
          <button
            type="button"
            class="bg-transparent border-[1px] rounded-xl px-2 py-0.5"
            @click="checkCardNumber"
          >
            Check
          </button>
        </div>
        <div v-if="isValid == true && !err" class="pt-6 w-full h-full">
          Your Card Number is Valid
        </div>
        <div v-if="isValid == false && tested && !err" class="pt-6">
          Your Card Number is Invalid
        </div>
        <div v-if="err" class="pt-6">{{ err.number }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, type Ref } from "vue";

let cardNumber: Ref<string> = ref("");
let tested: Ref<boolean> = ref(false);
let isValid: Ref<boolean> = ref(false);
let err: Ref<object> = ref({});

let logo = ref("");
const logos = {
  mastercard: "https://i.imgur.com/bbPHJVe.png",
  visa: "https://cdn.freebiesupply.com/logos/large/2x/visa-logo-png-transparent.png",
};

watch(cardNumber, (newCardNumber) => {
  if (newCardNumber.startsWith("51") || newCardNumber.startsWith("52")) {
    logo.value = logos["mastercard"];
  } else if (newCardNumber.startsWith("4")) {
    logo.value = logos["visa"];
  } else {
    logo.value = "";
  }
  tested.value = false;
  isValid.value = false;
  err.value = {};
});

async function checkCardNumber(): Promise<void> {
  try {
    const data = await fetch(
      `http://localhost:4000/validate/${cardNumber.value}`
    );
    if (data.status == 422) {
      err.value = await data.json();
      tested.value = true;
      return;
    }
    const res = await data.json();
    isValid.value = res;
    tested.value = true;
  } catch (er: any) {
    err.value = er;
  }
}
</script>

<style scoped></style>
