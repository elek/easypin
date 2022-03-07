import {createApp, reactive} from 'vue'
import {ethers} from "ethers";
import App from './App.vue'
import axios from "axios"

axios.get("/api/v0/pin/config").then(async function (config) {
    let cfg = config.data

    var accounts = await window.ethereum.request({method: 'eth_requestAccounts'});
    const web3 = new ethers.providers.Web3Provider(window.ethereum)
    var currentNetwork = await web3.getNetwork()
    window.ethereum.on("chainChanged", function () {
        window.location.reload();
    })
    var address = await web3.getSigner().getAddress()
    let app = createApp(App);
    app.provide("$address", address)
    app.provide("$cfg", cfg)
    app.mount('#app');

})


