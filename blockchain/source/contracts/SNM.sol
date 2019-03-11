pragma solidity ^0.4.11;

library SafeMath {
    function mul(uint a, uint b) internal pure returns (uint) {
        uint c = a * b;
        assert(a == 0 || c / a == b);
        return c;
    }

    function div(uint a, uint b) internal pure returns (uint) {
        assert(b > 0);
        uint c = a / b;
        assert(a == b * c + a % b);
        return c;
    }

    function sub(uint a, uint b) internal pure returns (uint) {
        assert(b <= a);
        return a - b;
    }

    function add(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        assert(c >= a);
        return c;
    }

    function max64(uint64 a, uint64 b) internal pure returns (uint64) {
        return a >= b ? a : b;
    }

    function min64(uint64 a, uint64 b) internal pure returns (uint64) {
        return a < b ? a : b;
    }

    function max256(uint256 a, uint256 b) internal pure returns (uint256) {
        return a >= b ? a : b;
    }

    function min256(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }
}


contract ERC20Basic {
    uint public totalSupply;
    function balanceOf(address who) public view returns (uint);
    function transfer(address to, uint value) public returns (bool success);
    event Transfer(address indexed from, address indexed to, uint value);
}

contract BasicToken is ERC20Basic {
    using SafeMath for uint;

    mapping(address => uint) balances;

    modifier onlyPayloadSize(uint size) {
        require(msg.data.length >= size + 4, "BasicToken require msg.data.length >= size + 4");
        _;
    }

    function transfer(address _to, uint _value) public onlyPayloadSize(2 * 32) returns (bool success) {
        balances[msg.sender] = balances[msg.sender].sub(_value);
        balances[_to] = balances[_to].add(_value);
        emit Transfer(msg.sender, _to, _value);
        return true;
    }

    function balanceOf(address _owner) public view returns (uint balance) {
        return balances[_owner];
    }

}

contract ERC20 is ERC20Basic {
    function allowance(address owner, address spender) public view returns (uint);
    function transferFrom(address from, address to, uint value) public returns (bool success);
    function approve(address spender, uint value) public;
    event Approval(address indexed owner, address indexed spender, uint value);
}

contract StandardToken is BasicToken, ERC20 {

    mapping (address => mapping (address => uint)) allowed;

    function transferFrom(address _from, address _to, uint _value) public returns (bool success) {
        uint _allowance = allowed[_from][msg.sender];
        balances[_to] = balances[_to].add(_value);
        balances[_from] = balances[_from].sub(_value);
        allowed[_from][msg.sender] = _allowance.sub(_value);
        emit Transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint _value) public {
        allowed[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
    }

    function allowance(address _owner, address _spender) public view returns (uint remaining) {
        return allowed[_owner][_spender];
    }

}

contract SNM  is StandardToken {
    string public name = "NAC Token";
    string public symbol = "NAC";
    uint public decimals = 18;
    uint constant TOKEN_LIMIT = 444 * 1e6 * 1e18;
    address public ico;

    // We block token transfers until ICO is finished.
    bool public tokensAreFrozen = false;

    constructor() public {
        uint totalSupply = 40 * 1e6 * 1e18;
        balances[msg.sender] = totalSupply;
    }

    // constructor() public {
    //     address owner = msg.sender;
    //     uint totalSupply = 444 * 1e6 * 1e18;
    //     balances[msg.sender] = totalSupply;
    // }

    // Mint few tokens and transefer them to some address.
    function mint(address _holder, uint _value) external {
        // require(msg.sender == ico);
        require(_value != 0, "SNM.mint require _value != 0");
        require(totalSupply + _value <= TOKEN_LIMIT, "SNM.mint require totalSupply + _value <= TOKEN_LIMIT");

        balances[_holder] += _value;
        totalSupply += _value;
        emit Transfer(0x0, _holder, _value);
    }


    // Allow token transfer.
    function defrost() external pure{
    // NOTICE: no longer need.
    //   require(msg.sender == ico);
    //   tokensAreFrozen = false;
    }

    function transfer(address _to, uint _value) public returns (bool success) {
        // require(!tokensAreFrozen);
        super.transfer(_to, _value);
        return true;
    }


    function transferFrom(address _from, address _to, uint _value) public returns (bool success) {
        // require(!tokensAreFrozen);
        super.transferFrom(_from, _to, _value);
        return true;
    }


    function approve(address _spender, uint _value) public {
        // require(!tokensAreFrozen);
        super.approve(_spender, _value);
    }
}
