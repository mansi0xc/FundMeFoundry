// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

import {Script} from "forge-std/Script.sol";
import {FundMe} from "./../src/FundMe.sol";
import {HelperConfig} from "./HelperConfig.s.sol";

contract DeployFundMe is Script {
    function run() external returns (FundMe) {
        // Before start broadcast -> not a real txn, its a simulated one
        HelperConfig helperConfig = new HelperConfig();
        (address ethUSDPriceFeed) = helperConfig.activeNetworkConfig();

        // After start broadcast -> Real txn!
        vm.startBroadcast();
        FundMe fundMe = new FundMe(ethUSDPriceFeed);
        vm.stopBroadcast();
        return fundMe;
    }
}
