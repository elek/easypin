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
var transactions = ref()
axios.get("/api/v0/block/" + props.hash).then(async function (res) {
  block.value = res.data
})
axios.get("/api/v0/pin/"+props.hash).then(async function (res) {
  transactions.value = res.data
})
</script>

<template>
  <div>
    <h2>{{ block.Hash }}</h2>
    <main role="main">
      <table class="table">
        <tbody>
        <tr>
          <td>Gateway link</td>
          <td><a :href="'https://gateway.ipfs.io/ipfs/QmcUkP3BMDkKNp2V6FTojXqgCMWdpDFQb74uphQVdWpi9Z/' + block.Hash">ipfs.io</a></td>
        </tr>
        <tr>
          <td>Pinned at</td>
          <td>{{ block.Pinned }}</td>
        </tr>
        <tr>
          <td>Days</td>
          <td>{{ block.ValidDays }}</td>
        </tr>
        </tbody>
      </table>
      <h3>Related transactions</h3>
      <table class="table">
        <thead>
        <tr>
          <th scope="col">Transaction</th>
          <th scope="col">Index</th>
          <th scope="col">Cid</th>
          <th scope="col">Raw amount</th>
          <th scope="col">Processed</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="tx in transactions">
          <td><a :href="'https://rinkeby.etherscan.io/tx/'+tx.Transaction">{{ tx.Transaction }}</a></td>
          <td>{{ tx.LogIndex }}</td>
          <td>{{ tx.Cid }}</td>
          <td>{{ tx.Amount }}</td>
          <td>{{ tx.Processed }}</td>
        </tr>
        </tbody>
      </table>
    </main>
  </div>
</template>

<style scoped>
</style>
