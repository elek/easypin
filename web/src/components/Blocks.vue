<script setup>
import axios from "axios";
import {defineProps, ref} from 'vue'

var blocks = ref()
var transactions = ref()
axios.get("/api/v0/block/all").then(async function (res) {
  blocks.value = res.data
})
</script>

<template>
  <div>
    <h2>Recently pinned IPFS root blocks</h2>
    <main role="main">
      <table class="table">
        <thead>
        <tr>
          <th scope="col">Cid</th>
          <th scope="col">Pinned at</th>
          <th scope="col">Period (days)</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="block in blocks">
          <td>
            <router-link :to="'/block/'+block.Hash">{{ block.Hash }}</router-link>
          </td>
          <td>{{ block.Pinned }}</td>
          <td>{{ block.Expiry }}</td>
          <td>{{ block.ValidDays }}</td>
        </tr>
        </tbody>
      </table>
    </main>
  </div>
</template>

<style scoped>
</style>
