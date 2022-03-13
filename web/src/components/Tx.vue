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

    <main role="main">
      <table class="table">
        <thead>
        <tr>
          <th scope="col">Transaction</th>
          <th scope="col">Index</th>
          <th scope="col">Cid</th>
          <th scope="col">Amount</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="tx in transactions">
          <td>{{ tx.Transaction }}</td>
          <td>Otto</td>
          <td>{{ tx.Cid }}</td>
          <td>{{ tx.TokenValue }}</td>
        </tr>
        </tbody>
      </table>
    </main>
  </div>
</template>

<style scoped>
</style>
