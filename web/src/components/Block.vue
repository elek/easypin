<script setup>
import axios from "axios";
import {ref} from 'vue'

const props = defineProps({
  hash: {
    type: String,
    required: true
  }
});

console.log(props.hash)
var block = ref()
axios.get("/api/v0/block/" + props.hash).then(async function (res) {
  block.value = res.data
})
</script>

<template>
  <div>
    <h1>{{ block.Hash }}</h1>
    <main role="main">
      <table class="table">
        <tbody>
        <tr>
          <td>Gateway link</td>
          <td><a :href="'https://gateway.ipfs.io/ipfs/QmcUkP3BMDkKNp2V6FTojXqgCMWdpDFQb74uphQVdWpi9Z/' + block.Hash">ipfs.io</a></td>
        </tr>
        <tr>
          <td>Pinned at</td>
          <td>{{ block.Expiry }}</td>
        </tr>
        <tr>
          <td>Expiry</td>
          <td>{{ block.Expiry }}</td>
        </tr>
        </tbody>
      </table>
    </main>
  </div>
</template>

<style scoped>
</style>
