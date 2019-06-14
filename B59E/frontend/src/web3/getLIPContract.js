import Web3 from 'web3'
import {address, ABI} from './contract/LIPContract'

let getLIPContract = async function () {
    let web3 = await new Web3(window.web3.currentProvider)
    let LIPContract = await web3.eth.contract(ABI)
    let Contract = LIPContract.at(address)
    console.log(LIPContract)
  	console.log(Contract)
    return Contract
}

export default getLIPContract
