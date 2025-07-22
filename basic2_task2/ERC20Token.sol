// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ERC20Token {
    //余额
    mapping(address => uint256) private _balances;
    // 部署地址
    address  public  _owner;
    // 授权列表
    mapping(address account => mapping(address spender => uint256)) private _allowances;

    uint256 private _totalSupply;
    string private _nick;
    string private _token;

    event Transfer(address to, uint256 value);
    event Approval(address owner, address spender,uint256 value);
    event LogAddress(address one, address two);

    error ErrInvalidAddr(address addr);
    error ErrInsufficientBalance(address addr);

    constructor(string memory _name,string memory _symbol)  {
        _nick = _name;
        _token = _symbol;
        _owner = msg.sender;

    }

    // 检查账户是否为0
    modifier checkIsOwner(address owner) {
        require(owner != _owner, "not is owner address");
        _;
    }

    function decimals() public view virtual returns (uint8) {
        return 6;
    }

    function name() public view returns (string memory) {
        return _nick;
    }

    function symbol() public view returns (string memory) {
        return _token;
    }

    // 查询余额
    function balanceOf(address account) external view returns (uint256){
        return _balances[account];
    }

    // 转帐
    function transfer(address to, uint256 value) external returns (bool){
        return updateBalanceOf(_owner, to, value);
    }

    // 授权
    function approve(address spender, uint256 value) external returns (bool){
        _allowances[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }

    // 查询授权额度
    function allowance(address owner, address spender) external view returns (uint256){
        uint256 value =  _allowances[owner][spender];
        return value;
    }

    // 授权转帐 (转帐不减少授权额度)
    function transferFrom(address from, address to, uint256 value) external returns (bool){
        // 检查授权额度
        uint256 allowanceBalance = _allowances[from][_owner];
        // 检查余额
        if (allowanceBalance < value) {
            revert ErrInsufficientBalance(from);
        }

        uint256 balance = _balances[from];
        if (balance < value) {
            revert ErrInsufficientBalance(from);
        }

        return updateBalanceOf(from, to, value);
    }

    // 更新余额
    function updateBalanceOf(address from, address to, uint256 value) internal returns (bool) {
        if (to == address(0) || from == address(0)){
            revert ErrInvalidAddr(address(0));
        }

        uint256 fromBalance = _balances[from];
        if (fromBalance < value) {
            revert ErrInsufficientBalance(from);
        }
        _balances[from] -= value;
        _balances[to] += value;
        emit Transfer(to, value);
        return true;
    }

    function mine(address owner,uint256 value) external returns (bool) {
        emit LogAddress(owner, _owner);
        require(owner == _owner, "not is owner address");
        _totalSupply += value;
        _balances[owner] += value;
        return true;
    }
}