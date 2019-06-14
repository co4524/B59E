pragma solidity ^0.5.0;

contract Escrow {
    
    uint missionID;
    uint amount;
    address owner;
    
    constructor(uint _missionID) public payable {
        owner = msg.sender;
        missionID = _missionID;
    }
    
    function escrowInfo() public view returns(address) {
        return (address(this));
    }   
}