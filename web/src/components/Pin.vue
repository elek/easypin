<script setup>
import {computed, ref} from 'vue'
import {BigNumber, ethers} from "ethers";
import axios from "axios"

const cid = ref("")
const token = ref("loading")
const state = ref("init")
const account = ref("")
var message = ref("")

var tokenAddress = ""
var pinAddress = ""
var duration = ref(1)
var allowance = ref(BigNumber.from(0))
var balance = ref(BigNumber.from(0))

var defaultAllowance = BigNumber.from("10000000000")

const provider = new ethers.providers.Web3Provider(window.ethereum)

const tokenAbi = [
  "function transfer(address to, uint amount) returns (bool)",
  "function allowance(address owner, address spender) view returns (uint256)",
  "function approve(address spender, uint256 amount) returns (bool)",
  "function balanceOf(address acc) view returns (uint256)"];

var pinAbi = [
  "function pin(string cid,uint amount) returns ()"
]

var price = function (d) {
  if (d.value === 200) {
    return BigNumber.from("2000000000")
  } else {
    return BigNumber.from(10000000).mul(d.value)
  }
}
var allowed = function () {
  return allowance.value.gt(price(duration))
}

var checkAllowance = function () {
  var signer = provider.getSigner(account.value)
  const contract = new ethers.Contract(tokenAddress, tokenAbi, signer)
  contract.allowance(account.value, pinAddress).then(function (res) {
    allowance.value = res
  })
}

var checkBalance = function () {
  var signer = provider.getSigner(account.value)
  const contract = new ethers.Contract(tokenAddress, tokenAbi, signer)
  contract.balanceOf(account.value).then(function (res) {
    balance.value = res
  })
}

var setAccount = function (a) {
  account.value = a;
  axios.get("/api/v0/pin/config").then(async function (config) {
    tokenAddress = config.data.TokenContract.replace(/0x/g, "");
    pinAddress = config.data.PinContract.replace(/0x/g, "");
    checkAllowance()
    checkBalance()
  })
}


var approve = function () {
  var signer = provider.getSigner(account.value)
  const contract = new ethers.Contract(tokenAddress, tokenAbi, signer)
  contract.approve(pinAddress, defaultAllowance).then(function (res) {
    checkAllowance()
  })
}


var pin = function () {
  if (cid.value === "") {
    message.value = "Please define a hash"
    return
  }
  var signer = provider.getSigner(account.value)
  const contract = new ethers.Contract(pinAddress, pinAbi, signer)
  contract.pin(cid.value, price(duration)).then(function (res) {
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

var tokenStr = function (a) {
  var out = ""
  var raw = a.toString()
  var dec = 8
  for (var i = raw.length - 1; i >= 0; i--) {
    if (dec === 0 && out.length > 0) {
      out = "." + out
    }

    if (out.length > 0 || dec <= 0 || raw[i] !== "0") {
      out = raw[i] + out
    }
    dec--
  }
  if (dec >= 0 && out.length > 0) {
    out = "0." + out
  }
  if (out.length == 0) {
    out = "0"
  }
  return out
}

var p = computed(() => {
  return price(duration)
})
</script>

<template>
  <form class="form-signin">
    <h1 class="h3 mb-3 font-weight-normal">Pin your IPFS hash</h1>

    <p>{{ message }}</p>

    <input v-model="cid" type="text" id="inputAmount" class="mb-3 form-control" placeholder="IPFS hash"
           required autofocus>

    <div class="form-check form-check-inline">
      <input class="form-check-input" type="radio" name="duration" id="1y" value="1" v-model="duration">
      <label class="form-check-label" for="1y">1 year</label>
    </div>
    <div class="form-check form-check-inline">
      <input class="form-check-input" type="radio" name="duration" id="2y" value="2" v-model="duration">
      <label class="form-check-label" for="2y">2 years</label>
    </div>
    <div class="form-check form-check-inline">
      <input class="form-check-input" type="radio" name="duration" id="forever" value="200" v-model="duration">
      <label class="form-check-label" for="forever">forever</label>
    </div>

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


    <p class="mt-5 mb-3 text-muted">Pinning costs {{ tokenStr(p) }} STORJ + gas fee</p>
    <div v-if="account">
      <p>Address: {{ account }}</p>
      <p v-if="balance">Balance: {{ tokenStr(balance) }}</p>
      <p v-if="allowance">Allowance: {{ tokenStr(allowance) }}</p>
    </div>
  </form>
</template>

<style scoped>
</style>
