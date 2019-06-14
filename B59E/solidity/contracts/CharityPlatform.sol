pragma solidity ^0.5.0;

import "./lib/SafeMath.sol";
import "./lib/ERC20ProInterface.sol";
import "./Escrow.sol";


contract CharityPlatform{
    using SafeMath for uint;
    
    uint internal missionIndex;
    uint internal exchangeRate;
    mapping(uint => address) missionIdxList; // missionID -> missionOwnerAddr
    mapping(address => Mission) missonList; // missionOwnerAddr -> Mission
    mapping(address => uint) ownerToMissionId; // ownerAddr -> missionID
    mapping(address => bool) appliedMissionList; // ownerAddr -> bool
    
    mapping(uint => address) escrowMissionList; // missionID -> escrowAddr 
    
    struct Mission {
        uint id;
        string title;
        string description;
        string imgHash;
        string minorityHash;
        uint targetFund;
        uint fundedAmount;
        address payable ownerAddr;
        mapping(address => uint) ledger;
    }
    
    constructor() public {
        missionIndex = 0;
        exchangeRate = 2;
    }
    
    // Add New Mission
    function addMission(
        string memory _title, 
        string memory _description,
        string memory _imgHash,
        string memory _minorityProofHash,
        uint _targetFund) public returns (bool success) {
        // Check whether msg.sender has had already applied for mission
        require(appliedMissionList[msg.sender] != true);
        
        // Create New Mission
        missionIndex ++;
        Mission memory mission = Mission(missionIndex, _title, _description, _imgHash,  _minorityProofHash, _targetFund, 0, msg.sender);
        missonList[msg.sender] = mission;
        missionIdxList[missionIndex] = msg.sender;
        ownerToMissionId[msg.sender] = missionIndex;
        appliedMissionList[msg.sender] = true;

        // Create Mission Escrow 
        Escrow escrow = new Escrow(missionIndex);
        escrowMissionList[missionIndex] = address(escrow);
        return true;
    }
    

    // Donate Mission 
    function donateMission(address BCPAddr, uint _missionID, uint tokens) public returns (bool success) {
        require(missionIdxList[_missionID] != msg.sender);
        require(missonList[missionIdxList[_missionID]].targetFund - missonList[missionIdxList[_missionID]].fundedAmount >= tokens);
        
        // Donation
        require(ERC20ProInterface(BCPAddr).balanceOf(msg.sender) >= tokens);  // Donator's BCP >= tokens
        // Transfer BCP from Donator to MissionContractAddress
        ERC20ProInterface(BCPAddr).transferFromContract(msg.sender, escrowMissionList[_missionID], tokens); 
        
        missonList[missionIdxList[_missionID]].ledger[msg.sender] = tokens;
        missonList[missionIdxList[_missionID]].fundedAmount = missonList[missionIdxList[_missionID]].fundedAmount.add(tokens);   
        return true;
    }


    function closeMission(address BCPAddr, address LIPAddr) public returns (bool success) {
        require(missionIdxList[ownerToMissionId[msg.sender]] == msg.sender); // only missionOwner can close mission
        
        address escrowAddr = escrowMissionList[ownerToMissionId[msg.sender]];
        uint escrowTotalBCP = ERC20ProInterface(BCPAddr).balanceOf(escrowAddr);
        ERC20ProInterface(BCPAddr).transferFromContract(escrowAddr, BCPAddr, escrowTotalBCP); 
        
       // missonList[missionIdxList[_missionID]].ledger[msg.sender] = tokens;
        missonList[msg.sender].fundedAmount = missonList[msg.sender].fundedAmount.sub(escrowTotalBCP); 
        // Exchage to LIP from BCP.ContractAddress to missionOwner
        require(ERC20ProInterface(LIPAddr).balanceOf(BCPAddr) >= escrowTotalBCP.div(exchangeRate)); // require BCP have enough LIP to exchange
        ERC20ProInterface(LIPAddr).transferFromContract(BCPAddr, msg.sender, escrowTotalBCP.div(exchangeRate)); 
        missonList[msg.sender].fundedAmount = missonList[msg.sender].fundedAmount.sub(0); 
        return true;
    }



    // Information Function
    function getMission() public view returns (uint, string memory, string memory, string memory, string memory, uint, uint, address){
        Mission memory mission = missonList[msg.sender];
        return (mission.id, mission.title, mission.description, mission.imgHash, mission.minorityHash, mission.targetFund, mission.fundedAmount, mission.ownerAddr);
    }
    
    function getMissionById(uint _missionID) public view returns (uint, string memory, string memory, string memory, string memory, uint, uint, address) {
        Mission memory mission = missonList[missionIdxList[_missionID]];
        return (mission.id, mission.title, mission.description, mission.imgHash, mission.minorityHash, mission.targetFund, mission.fundedAmount, mission.ownerAddr);
    }
    
    function escrowInfo(uint _missionID) public view returns(address) {
        return Escrow(escrowMissionList[_missionID]).escrowInfo();
    }
}
