pragma solidity ^0.6.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract Sender is ReentrancyGuard {
    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    uint64 constant MAX_UINT64 = 18_446_744_073_709_551_615;


    function sendAssets(
        address asset,
        uint256 amount
    ) public nonReentrant {

        uint8 decimals = ERC20(asset).decimals();
        uint256 balanceBefore = IERC20(asset).balanceOf(address(this));
        IERC20(asset).safeTransferFrom(msg.sender, address(this), amount);
        uint256 balanceAfter = IERC20(asset).balanceOf(address(this));
    }

}
