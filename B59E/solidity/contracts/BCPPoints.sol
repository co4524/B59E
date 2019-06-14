pragma solidity ^0.5.0;

import "./lib/SafeMath.sol";
import "./lib/Ownership.sol";
import "./lib/ERC20ProInterface.sol";

contract BCPPoints is ERC20ProInterface, Ownership {
    using SafeMath for uint;
 
    string public symbol;
    string public  name;
    uint public decimals;
    uint _totalSupply;
    mapping(address => uint) balances;
    mapping(address => mapping(address => uint)) allowed;
    
    uint exchangeRate; // 1 LIP -> 2 BCP
    
    
    // Constructor Initiate
    constructor() public payable {
        symbol = "BCP";
        name = "Blockchain Charity Points";
        decimals = 18;
        _totalSupply = 2000 * 10**uint(decimals);
        exchangeRate = 2;

        // Release all the token to contract address
        balances[address(this)] = _totalSupply;
        balances[msg.sender] = _totalSupply;
        emit Transfer(address(0), address(this), _totalSupply);
    }
    
    function totalSupply() public view returns (uint) {
        return _totalSupply;
    }
    
    function balanceOf(address tokenOwner) public view returns (uint balance) {
        return balances[tokenOwner];
    }

    function transfer(address to, uint tokens) public returns (bool success) {
        require(balances[msg.sender] > 0);
        require(tokens > 0);
        
        balances[msg.sender] = balances[msg.sender].sub(tokens);
        balances[to] = balances[to].add(tokens);
        emit Transfer(msg.sender, to, tokens);
        return true;
    }

    function allowance(address tokenOwner, address spender) public view returns (uint remaining) {
        return allowed[tokenOwner][spender];
    }

    function approve(address spender, uint tokens) public returns (bool success) {
        allowed[msg.sender][spender] = tokens;
        emit Approval(msg.sender, spender, tokens);
        return true;
    }

    function transferFrom(address from, address to, uint tokens) public returns (bool success) {
        require(balances[from] > 0);
        require(allowed[from][msg.sender] > 0);
        require(tokens > 0);
        
        balances[from] = balances[from].sub(tokens);
        allowed[from][msg.sender] = allowed[from][msg.sender].sub(tokens);
        balances[to] = balances[to].add(tokens);
        emit Transfer(from, to, tokens);
        return true;
    }

    /*-----------------Additional Function-----------------*/
    function distributePoints(address to, uint tokens) public onlyOwner returns (bool success) {
        require(balances[address(this)] >= tokens);
        require(tokens > 0);
        
        balances[address(this)] = balances[address(this)].sub(tokens);
        balances[to] = balances[to].add(tokens);
        return true;
    }
    
    function transferFromContract(address from, address to, uint tokens) public returns (bool success) {
        require(balances[from] >= tokens);
        require(tokens > 0);
        
        balances[from] = balances[from].sub(tokens);
        balances[to] = balances[to].add(tokens);
        emit Transfer(from, to, tokens);
        return true;   
    }
    
    
    function LIPtoBCP(address LIPAddr, uint LIPtokens) public returns (bool success) {
        require(ERC20ProInterface(LIPAddr).balanceOf(msg.sender) >= LIPtokens);
        require(balances[address(this)] >= LIPtokens*exchangeRate);
        
        // Receive LIP from msg.sender
        ERC20ProInterface(LIPAddr).transferFromContract(msg.sender, address(this), LIPtokens);
        // Exchange BCP from BCPContract to msg.sender
        transferFromContract(address(this), msg.sender, LIPtokens*exchangeRate);
        return true;
    }
    
    
    function BCPtoLIP(address LIPAddr, uint BCPtokens) public returns (bool success) {
        require(BCPtokens % exchangeRate == 0);
        require(balances[msg.sender] >= BCPtokens);
        require(ERC20ProInterface(LIPAddr).balanceOf(address(this)) >= BCPtokens.div(exchangeRate));
    
        // Exchange LIP from BCPContract to msg.sender
        ERC20ProInterface(LIPAddr).transferFromContract(address(this), msg.sender, BCPtokens.div(exchangeRate));
        // Receive BCP from msg.sender
        transferFromContract(msg.sender, address(this), BCPtokens);
        return true;
    }
    

    
    // Fallback Function 
    function () external payable {
        revert();
    }
}