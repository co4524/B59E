const address = "0xB1438a0fa2c652ce33367D1bA59979403a798DbE"

const ABI = [
    {
      "inputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor",
      "signature": "constructor"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_title",
          "type": "string"
        },
        {
          "name": "_description",
          "type": "string"
        },
        {
          "name": "_imgHash",
          "type": "string"
        },
        {
          "name": "_minorityProofHash",
          "type": "string"
        },
        {
          "name": "_targetFund",
          "type": "uint256"
        }
      ],
      "name": "addMission",
      "outputs": [
        {
          "name": "success",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function",
      "signature": "0xd2e66654"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "BCPAddr",
          "type": "address"
        },
        {
          "name": "_missionID",
          "type": "uint256"
        },
        {
          "name": "tokens",
          "type": "uint256"
        }
      ],
      "name": "donateMission",
      "outputs": [
        {
          "name": "success",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function",
      "signature": "0x1e238b68"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "BCPAddr",
          "type": "address"
        },
        {
          "name": "LIPAddr",
          "type": "address"
        }
      ],
      "name": "closeMission",
      "outputs": [
        {
          "name": "success",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function",
      "signature": "0x3cb1fa7a"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getMission",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "uint256"
        },
        {
          "name": "",
          "type": "uint256"
        },
        {
          "name": "",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function",
      "signature": "0x130e0906"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_missionID",
          "type": "uint256"
        }
      ],
      "name": "getMissionById",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "string"
        },
        {
          "name": "",
          "type": "uint256"
        },
        {
          "name": "",
          "type": "uint256"
        },
        {
          "name": "",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function",
      "signature": "0x16cc67e5"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_missionID",
          "type": "uint256"
        }
      ],
      "name": "escrowInfo",
      "outputs": [
        {
          "name": "",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function",
      "signature": "0x83560525"
    }
]

export {address, ABI}