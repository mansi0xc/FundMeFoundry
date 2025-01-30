// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {Test, console} from "forge-std/Test.sol";
import {FundMe} from "../src/FundMe.sol";
import {DeployFundMe} from "../script/DeployFundMe.s.sol";
import {PriceConverter} from "../src/PriceConverter.sol";
import {AggregatorV3Interface} from "./../lib/chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract FundeMeTest is Test {
    FundMe fundMe;
    address USER = makeAddr("user");
    uint256 constant SEND_VALUE = 0.1 ether;
    uint256 constant STARTING_VALUE = 100 ether;
    uint256 constant GAS_PRICE = 1;

    function setUp() external {
        // us -> FundMeTest -> FundMe
        // fundMe = new FundMe(0x694AA1769357215DE4FAC081bf1f309aDC325306);
        DeployFundMe deployFundMe = new DeployFundMe();
        fundMe = deployFundMe.run();
        vm.deal(USER, STARTING_VALUE);
    }

    function testMinUSDisFive() public view {
        assertEq(fundMe.MINIMUM_USD(), 5e18);
    }

    function testGetVersion() public view {
        uint256 version = fundMe.getVersion();
        console.log(version);
        // AggregatorV3Interface ghh = AggregatorV3Interface(0x694AA1769357215DE4FAC081bf1f309aDC325306);
        // AggregatorV3Interface ghh = AggregatorV3Interface(0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419);
        // uint256 rate = PriceConverter.getConversionRate(1, ghh);
        // console.log(rate);
    }

    function testFundFailsWithoutEnoughETH() public {
        vm.expectRevert(); // next line should revert / assert(this txn fails/reverts)
        fundMe.fund();
    }

    function testFundUpdatesFundedDataStructure() public {
        vm.prank(USER); // the next txn / actions will be carried out by the address USER
        fundMe.fund{value: SEND_VALUE}();
        uint256 amountFunded = fundMe.getAddressToAmountFunded(USER);
        assertEq(amountFunded, SEND_VALUE);
    }

    function testAddsFunderToArrayOfFunders() public {
        vm.prank(USER); // the previous txn / action is resetted
        fundMe.fund{value: SEND_VALUE}();

        address funder = fundMe.getFunder(0); // that;s why index will be 0 here and not 1
        assertEq(funder, USER);
    }

    modifier funded() {
        vm.prank(USER);
        fundMe.fund{value: SEND_VALUE}();
        _;
    }

    function testOnlyOwnerCanWithdraw() public funded {
        vm.prank(USER);
        vm.expectRevert();
        fundMe.withdraw();
    }

    function testWithdrawWithASingleFunder() public funded {
        // Arrange
        uint256 startingOwnerBalance = fundMe.getOwner().balance;
        uint256 startingFundMeBalance = address(fundMe).balance;

        // Act
        vm.prank(fundMe.getOwner());
        fundMe.withdraw();

        // Assert
        uint256 endingOwnerBalance = fundMe.getOwner().balance;
        uint256 endingFundMeBalance = address(fundMe).balance;
        assertEq(endingFundMeBalance, 0);
        assertEq(startingFundMeBalance + startingOwnerBalance, endingOwnerBalance);
    }

    function testWithdrawFromMultipleFunders() public funded {
        // Arrange
        uint160 numbersOfFunders = 10;
        uint160 startingFunderIndex = 1;
        for (uint160 i = startingFunderIndex; i < numbersOfFunders; i++) {
            // vm.prank new address + vm.deal new address = hoax(address, value);
            hoax(address(i), SEND_VALUE);
            fundMe.fund{value: SEND_VALUE}();
        }

        uint256 startingOwnerBalance = fundMe.getOwner().balance;
        uint256 startingFundMeBalance = address(fundMe).balance;

        /* Looking closer at the `testWithdrawFromASingleFunder` one can observe that we found out the 
        initial balances, then we called a transaction and then we asserted that 
        `startingFundMeBalance + startingOwnerBalance` matches the expected balance, 
        but inside that test we called `withdraw` which should have cost gas. 
        Why didn't the gas we paid affect our balances? Simple, for testing purposes the Anvil gas price 
        is defaulted to `0` (different from what we talked about above in the case of Ethereum mainnet
         where the gas price was around `7 gwei`), so it wouldn't interfere with our testing.*/

        // Act

        uint256 gasStart = gasleft(); // 1000
        vm.txGasPrice(GAS_PRICE);
        vm.startPrank(fundMe.getOwner()); // c:200
        fundMe.withdraw(); // should have spent gas
        vm.stopPrank();

        uint256 gasEnd = gasleft();
        uint256 gasUsed = (gasStart - gasEnd) * tx.gasprice;
        console.log(gasUsed);

        // Assert
        assert(address(fundMe).balance == 0);
        assert(startingFundMeBalance + startingOwnerBalance == fundMe.getOwner().balance);
    }

    /*
    We start by declaring the total number of funders. 
    Then we declare that the `startingFunderIndex` is 1. 
    You see that both these variables are defined as `uint160` and not our usual `uint256`. 
    Down the road, we will use the `startingFunderIndex` as an address. 
    If we look at the definition of an [address]
    (https://docs.soliditylang.org/en/latest/types.html#address) we see that it holds `a 20 byte value` 
    and that `explicit conversions to and from address are allowed for 
    uint160, integer literals, bytes20 and contract types`. 
    Having the index already in `uint160` will save us from casting it when we need to convert it into 
    an address.

    We start our index from `1` because it's not advised to user `address(0)` 
    in this way. `address(0)` has a special regime and should not be pranked.
    */

    function testWithdrawFromMultipleFundersCheaper() public funded {
        uint160 numberOfFunders = 10;
        uint160 startingFunderIndex = 1;
        for (uint160 i = startingFunderIndex; i < numberOfFunders + startingFunderIndex; i++) {
            // we get hoax from stdcheats
            // prank + deal
            hoax(address(i), SEND_VALUE);
            fundMe.fund{value: SEND_VALUE}();
        }

        uint256 startingFundMeBalance = address(fundMe).balance;
        uint256 startingOwnerBalance = fundMe.getOwner().balance;

        vm.startPrank(fundMe.getOwner());
        fundMe.cheaperWithdraw();
        vm.stopPrank();

        assert(address(fundMe).balance == 0);
        assert(startingFundMeBalance + startingOwnerBalance == fundMe.getOwner().balance);
        assert((numberOfFunders + 1) * SEND_VALUE == fundMe.getOwner().balance - startingOwnerBalance);
    }
}
