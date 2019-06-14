import Web3 from 'web3'
import axios from 'axios'
import {address, ABI} from './contract/BCPContract'

const url = "http://127.0.0.1:8081/abi/bcp"

let getBCPContract = async function () {
    let web3 = await new Web3(window.web3.currentProvider)
    let BCPContract = await web3.eth.contract(ABI)
    let Contract = BCPContract.at(address)
    console.log(BCPContract)
    console.log(Contract)
    return Contract
    /*
    axios.get(url)
    .then(async (response) => {
      let BCPContract = await web3.eth.contract(ABI)
      let Contract = BCPContract.at(address)
      console.log(BCPContract)
      console.log(Contract)
      return Contract
    })
    .catch((err) => {
      console.log(err)
    })
    */
}

export default getBCPContract
