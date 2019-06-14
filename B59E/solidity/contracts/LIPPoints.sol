pragma solidity ^0.5.0;
 
import "./lib/SafeMath.sol";
import "./lib/Ownership.sol";
import './lib/ERC20ProInterface.sol';

contract LIPPoints is ERC20ProInterface, Ownership {
    using SafeMath for uint;
 
    string public symbol;
    string public  name;
    uint public decimals;
    uint _totalSupply;
    mapping(address => uint) balances;
    mapping(address => mapping(address => uint)) allowed;
    
    // Constructor Initiate
    constructor() public payable {
        symbol = "LIP";
        name = "LinePoints";
        decimals = 18;
        _totalSupply = 1000 * 10**uint(decimals);
        
        // Release all the token to contract address
        balances[address(this)] = _totalSupply;
        emit Transfer(address(0), address(this), _totalSupply);
    }
    
    function totalSupply() public view returns (uint) {
        return _totalSupply.sub(balances[address(0)]);
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
        require(balances[address(this)] > tokens);
        require(tokens > 0);
        
        balances[address(this)] = balances[address(this)].sub(tokens);
        balances[to] = balances[to].add(tokens);
        emit Transfer(address(this), to, tokens);
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
    
    
    // Fallback Function
    function () external payable {
        revert();
    }
}