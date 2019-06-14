import Web3 from 'web3'
import {address, ABI} from './contract/PlatformContract'

let getPlatformContract = async function () {
    let web3 = await new Web3(window.web3.currentProvider)
    let PlatformContract = await web3.eth.contract(ABI)
    let Contract = PlatformContract.at(address)
    console.log(PlatformContract)
  	console.log(Contract)
    return Contract
}

export default getPlatformContract