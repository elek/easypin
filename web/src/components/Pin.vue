<script setup>
import {computed, inject, onMounted, ref} from 'vue'
import {BigNumber, ethers} from "ethers";
import axios from "axios"

const cid = ref("")
const token = ref("loading")
const state = ref("init")
const account = ref("")
var message = ref("")

var tokenAddress = ""
var pinAddress = ""
var allowance = ref(BigNumber.from(0))
var price = BigNumber.from("10000000")


const provider = new ethers.providers.Web3Provider(window.ethereum)


const tokenAbi = [
  "function transfer(address to, uint amount) returns (bool)",
  "function allowance(address owner, address spender) view returns (uint256)",
  "function approve(address spender, uint256 amount) returns (bool)"];

var pinAbi = [
  "function pin(string cid,uint amount) returns ()"
]

var allowed = function () {
  return allowance.value.gt(price)
}

var checkAllowance = function () {
  var signer = provider.getSigner(account.value)
  const contract = new ethers.Contract(tokenAddress, tokenAbi, signer)
  contract.allowance(account.value, pinAddress).then(function (res) {
    allowance.value = res
  })
}

var setAccount = function (a) {
  account.value = a;
  axios.get("/api/v0/pin/config").then(async function (config) {
    tokenAddress = config.data.TokenContract.replace(/0x/g, "");
    pinAddress = config.data.PinContract.replace(/0x/g, "");
    checkAllowance()
  })
}


var approve = function () {
  var signer = provider.getSigner(account.value)
  const contract = new ethers.Contract(tokenAddress, tokenAbi, signer)
  contract.approve(account.value, pinAddress).then(function (res) {
    checkAllowance()
  })
}


var pin = function () {
  if (cid.value === "") {
    message.value = "Please define a cid"
    return
  }
  var signer = provider.getSigner(account.value)
  const contract = new ethers.Contract(pinAddress, pinAbi, signer)
  contract.pin(cid.value, price).then(function (res) {
  })
}

provider.send("eth_accounts", []).then(function (res) {
  setAccount(res[0])
})

var connect = function () {
  provider.send("eth_requestAccounts", []).then(function (res) {
    setAccount(res[0])
  })
}


</script>

<template>
  <form class="form-signin">
    <h1 class="h3 mb-3 font-weight-normal">Pin your IPFS CID {{ allowance }} </h1>

    <p>{{ message }}</p>

    <input v-model="cid" type="text" id="inputAmount" class="mb-3 form-control" placeholder="IPFS CID"
           required autofocus>

    <button v-if="!account" type="button"
            class="btn btn-lg btn-warning btn-block"
            @click="connect()">Connect to metamask
    </button>

    <div v-if="account">
      <button v-if="allowed()" type="button"
              class="btn btn-lg btn-warning btn-block"
              @click="pin()">Pin It!
      </button>

      <button type="button"
              v-if="!allowed()"
              class="btn btn-lg btn-warning btn-block"
              @click="approve()">Authorize STORJ tokens
      </button>
    </div>

    <p class="mt-5 mb-3 text-muted">Pinning costs n STORJ tokens / megabytes / days</p>
  </form>
</template>

<style scoped>
</style>
