<script setup>
import axios from "axios";
import {computed, ref} from 'vue'

var transactions = ref()
axios.get("/api/v0/pin/all").then(async function (res) {
  transactions.value = res.data
})
</script>

<template>
  <div>
    <h2>Recently imported transactions</h2>
    <main role="main">
      <table class="table">
        <thead>
        <tr>
          <th scope="col">Transaction</th>
          <th scope="col">Index</th>
          <th scope="col">Cid</th>
          <th scope="col">Raw amount</th>
          <th scope="col">Processed</th>
          <th scope="col">Error</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="tx in transactions">
          <td><a :href="'https://rinkeby.etherscan.io/tx/'+tx.Transaction">{{ tx.Transaction }}</a></td>
          <td>{{ tx.LogIndex }}</td>
          <td>{{ tx.Cid }}</td>
          <td>{{ tx.Amount }}</td>
          <td>{{ tx.Processed }}</td>
          <td>{{ tx.Error }}</td>
        </tr>
        </tbody>
      </table>
    </main>
  </div>
</template>

<style scoped>
</style>
