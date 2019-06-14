pragma solidity ^0.5.0;

contract Ownership {
    address public owner;
    address private newOwner;
 
    event OwnershipTransferred(address indexed _from, address indexed _to);
 
    constructor() public {
        owner = msg.sender;
    }
 
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
 
    // Transfer ownership to new address
    function transferOwnership(address _newOwner) public onlyOwner {
        newOwner = _newOwner;
    }
    
    // New owner trigger the function to accept the ownership
    function acceptOwnership() public {
        require(msg.sender == newOwner);
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
        newOwner = address(0);
    }
}