// SPDX-License-Identifier: MIT

// 1. Deploy Mocks when we are on local anvil chain
// 2. Keep track of contract addresses across different chains
// Sepolia ETH/USD - 0x694AA1769357215DE4FAC081bf1f309aDC325306
// Mainnet ETH/USD - 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419

import {Script} from "forge-std/Script.sol";
import {MockV3Aggregator} from "test/mock/MockV3Aggregator.sol";

pragma solidity ^0.8.18;

contract HelperConfig is Script {
    // If we are on local anvil chain, we deploy mocks
    // Otherwise, grab the existing address from the live network

    NetworkConfig public activeNetworkConfig;
    uint8 public constant DECIMALS = 8;
    int256 public constant INITIAL_PRICE = 2000e8;

    struct NetworkConfig {
        address priceFeed; // ETH/USD price feed address
    }

    constructor() {
        if (block.chainid == 11155111) {
            activeNetworkConfig = getSepoliaEthConfig();
        } else if (block.chainid == 1) {
            activeNetworkConfig = getMainnetEthConfig();
        } else {
            activeNetworkConfig = getOrCreateAnvilEthConfig();
        }
    }

    function getSepoliaEthConfig() public pure returns (NetworkConfig memory) {
        // price feed address
        // vrf address, gas price
        NetworkConfig memory sepoliaConfig = NetworkConfig({priceFeed: 0x694AA1769357215DE4FAC081bf1f309aDC325306});
        return sepoliaConfig;
    }

    function getMainnetEthConfig() public pure returns (NetworkConfig memory) {
        // price feed address
        // vrf address, gas price
        NetworkConfig memory mainnetConfig = NetworkConfig({priceFeed: 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419});
        return mainnetConfig;
    }

    function getOrCreateAnvilEthConfig() public returns (NetworkConfig memory) {
        // price feed address

        // 1. Deploy a mock contract
        // 2. Return the mock address

        // qn :- why can't the function be pure due to the presence of vm?

        if (activeNetworkConfig.priceFeed != address(0)) {
            return activeNetworkConfig;
        }

        vm.startBroadcast();
        MockV3Aggregator mockPriceFeed = new MockV3Aggregator(DECIMALS, INITIAL_PRICE);
        vm.stopBroadcast();

        NetworkConfig memory anvilConfig = NetworkConfig({priceFeed: address(mockPriceFeed)});
        return anvilConfig;
    }
}
