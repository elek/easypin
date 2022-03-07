<script setup>
import {computed, inject, onMounted, ref} from 'vue'
import {BigNumber, ethers} from "ethers";


const address = inject("$address")
const amount = ref("")

const cfg = inject("$cfg")
const tokenAbi = [
  "function transfer(address to, uint amount) returns (bool)",
  "function allowance(address owner, address spender) view returns (uint256)",
  "function approve(address spender, uint256 amount) returns (bool)"
];
const bridgeAbi = [
  "function deposit(uint256 amount) returns (bool)"
]


var deposit = function () {
  var web3 = new ethers.providers.Web3Provider(window.ethereum)
  const bridge = new ethers.Contract(cfg["here.token"], tokenAbi, web3.getSigner())
  bridge.transfer(cfg["here.bridge"], BigNumber.from(amount.value).mul(BigNumber.from(10).pow(18)))
}

</script>

<template>
  <form class="form-signin">
    <h1 class="h3 mb-3 font-weight-normal">Pin IPFS CID</h1>
    <label for="inputCid" class="sr-only">CID</label>
    <input v-model="amount" type="text" id="inputCid" class="mb-3 form-control"
           placeholder="The CID of the IPFS file to be pinned"
           required autofocus>
    <button type="button"
            class="btn btn-lg btn-warning btn-block"
            @click="deposit()">PIN
    </button>
  </form>
</template>

<style scoped>
a {
  color: #42b983;
}

</style>
