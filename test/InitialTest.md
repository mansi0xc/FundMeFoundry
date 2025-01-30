// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {Test, console} from "forge-std/Test.sol";
import {FundMe} from "../src/FundMe.sol";
import {DeployFundMe} from "../script/DeployFundMe.s.sol";
import {PriceConverter} from "../src/PriceConverter.sol";
import {AggregatorV3Interface} from "./../lib/chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol"; 

contract FundeMeTest is Test{
    FundMe fundMe;
    function setUp() external{
        // us -> FundMeTest -> FundMe
        // fundMe = new FundMe(0x694AA1769357215DE4FAC081bf1f309aDC325306);
        DeployFundMe deployFundMe = new DeployFundMe();
        fundMe = deployFundMe.run();
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

}