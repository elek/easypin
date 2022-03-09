import logo from './logo.svg';
import './App.css';
import React from 'react';
import { ethers } from 'ethers';
const axios = require('axios').default;

const tokenContractAddress = '0x8098165d982765097E4aa17138816e5b95f9fDb5'; // STORJ by default
const pinContractAddress = '0x069FaE1B18F4c18852b3F07f60f82121F8A9030b'; // Pinning service smart contract
const ipfsHTTPEndpoint = 'http://127.0.0.1:5001/api/v0/add'; // Use for local ipfs node
//const ipfsHTTPEndpoint = 'https://ipfs-demo.dev.storj.io/api/v0/add';


// The ERC-20 Contract ABI, which is a common contract interface
// for tokens (this is the Human-Readable ABI format)
const storjTokenAbi = [
  // Some details about the token
  "function name() view returns (string)",
  "function symbol() view returns (string)",

  // Get the account balance
  "function balanceOf(address) view returns (uint)",

  // Get the allowance
  "function allowance(address owner, address spender) view returns (uint)",

  // Approve token for use
  "function approve(address to, uint amount)",

  // An event triggered whenever anyone transfers to someone else
  //"event Transfer(address indexed from, address indexed to, uint amount)"
];

const pinAbi = [
  // Pin to IPFS
  "function pin(string ipfsHash, uint amount)",
];

function App() {

  const [metamaskDetected, setMetamaskDetected] = React.useState(false);
  const [network, setNetwork] = React.useState();
  const [provider, setProvider] = React.useState(null);
  const [metamaskAccounts, setMetamaskAccounts] = React.useState(false);
  const [approved, setApproved] = React.useState(false);
  const [formOpen, setFormOpen] = React.useState(false);
  const [hashes, setHashes] = React.useState([]);

  const connectMetamask = async () => {
    console.log("connecting metamask");
    // MetaMask requires requesting permission to connect users accounts
    try {
      setMetamaskAccounts(await provider.send("eth_requestAccounts", []));
    } catch (error) {
      console.log("ERROR");
      console.error(error);
    }
  }

  const approveSomething = async () => {
    // Get signer
    const signer = provider.getSigner();

    // Initiate Contract
    const storjTokenContract = new ethers.Contract(tokenContractAddress, storjTokenAbi, signer);

    // Check if there are already tokens approved
    const numberApproved = await storjTokenContract.allowance(metamaskAccounts[0], pinContractAddress);
    const numberApprovedFormatted = ethers.utils.formatUnits(numberApproved, 8); // 8 is hardcoded as the # decimals for the ERC20 Token
    setApproved(numberApprovedFormatted);

    // Approve spend of tokens (STORJ) by the Pin contract
    const approveStorjToken = await storjTokenContract.approve(pinContractAddress, 1000000000);
    console.log(approveStorjToken);
  }

  const pinSomething = async (event) => {
    event.preventDefault();
    const hash = event.target.ipfsHash.value;
    // Get signer
    const signer = provider.getSigner();

    // Initiate Contract
    const pinContract = new ethers.Contract(pinContractAddress, pinAbi, signer);

    // Submit the transaction to the Pin contract
    const submitPin = await pinContract.pin(hash, 100000000);
    console.log(submitPin);
    closeForm();
  }

  const initPin = async (event) => {
    event.preventDefault();
    const file = event.target.pinFile.value; // This is wrong
    const bodyFormData = new FormData();
    bodyFormData.append('file', file);
    try {
      const pinData = await axios({
        method: 'POST',
        url: ipfsHTTPEndpoint,
        data: bodyFormData,
        headers: { "Content-Type": "multipart/form-data" }
      });
      const newHash = pinData.data.Hash;
      console.log(pinData);
      setHashes([...hashes, newHash])

    } catch (error) {
      console.error(error);
    }
  }

  const openForm = () => {
    setFormOpen(true);
  }

  const closeForm = () => {
    setFormOpen(false);
  }


  React.useEffect(async () => {
    if (window.ethereum && window.ethereum.isConnected) {
      setMetamaskDetected(true);
      console.log("Ethereum provider (Metamask) detected");

      // A Web3Provider wraps a standard Web3 provider, which is
      // what MetaMask injects as window.ethereum into each page
      const newProvider = new ethers.providers.Web3Provider(window.ethereum);
      setProvider(newProvider);
    } else {
      console.error("Ethereum provider (Metamask) not detected");
    }
	}, [])


  React.useEffect(async () => {
    if (provider) {
      await provider.getNetwork().then((info) => {
        setNetwork(info);
      })
    } 
  }, [provider])

  return (
    <div>
      <h1>Step 1: Initial IPFS pin</h1>
      <h3>This doesn't work yet | To create initial IPFS pin see <a href="https://ipfs-demo.dev.storj.io/" target="_blank">https://ipfs-demo.dev.storj.io/</a></h3>
      
      <form onSubmit={initPin}>
        <input type="file" id="pinFile" name="filename"/>
        <input type="submit" value="Submit"/>
      </form>
      <br/>
      <p>IPFS Hashes</p>
      <ul id="hashes">
      {hashes.map((hash) => <li><p>Hash: <a href={'https://'+hash+'.ipfs.dweb.link'} target="_blank">{hash}</a></p></li>)}
      </ul>
      <hr/>
      <h1>Step 2: IPFS Pinning Backed By Storj via Smart Contract</h1>
      {!metamaskDetected &&
        <p>Metamask Not Detected</p>
      }
      {!metamaskAccounts &&
        <button onClick={connectMetamask} disabled={!metamaskDetected}>Connect Metamask</button>
      }
      
      {metamaskAccounts &&
        <>
          <p><strong>Network info</strong></p>
          <ul>
            <li><p>Chain ID: {network.chainId}</p></li>
            <li><p>Name: {network.name}</p></li>
          </ul>
          <button onClick={approveSomething}>Approve</button>
          <button onClick={openForm} disabled={!approved}>Pin</button>
        </>
      }
      {formOpen &&
        <form onSubmit={pinSomething}>
          <label>IPFS Hash:</label><br/>
          <input type="text" id="ipfsHash" name="ipfsHash"/><br/>
          <input type="submit" value="Submit"/>
        </form> 
      }
      {approved &&
        <>
        <p><strong>Token Approvals</strong></p>
          <ul>
            <li><p># STORJ Approved: {approved}</p></li>
          </ul>
        </>
      }
    </div>
  );
}

export default App;
