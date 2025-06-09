# $ forge inspect FundMe storageLayout

╭-------------------------+--------------------------------+------+--------+-------+-----------------------╮
| Name                    | Type                           | Slot | Offset | Bytes | Contract              |
+==========================================================================================================+
| s_addressToAmountFunded | mapping(address => uint256)    | 0    | 0      | 32    | src/FundMe.sol:FundMe |
|-------------------------+--------------------------------+------+--------+-------+-----------------------|
| s_funders               | address[]                      | 1    | 0      | 32    | src/FundMe.sol:FundMe |
|-------------------------+--------------------------------+------+--------+-------+-----------------------|
| s_priceFeed             | contract AggregatorV3Interface | 2    | 0      | 20    | src/FundMe.sol:FundMe |
╰-------------------------+--------------------------------+------+--------+-------+-----------------------╯

**Steps:**
# forge script DeployFundMe --rpc-url http://127.0.0.1:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast`

# cast code <fund me contract address>

copy result in - https://etherscan.io/opcode-tool

[1] PUSH1 0x80
[3] PUSH1 0x40
[4] MSTORE
[6] PUSH1 0x04
[7] CALLDATASIZE
[8] LT
[11] PUSH2 0x008a
[12] JUMPI
[14] PUSH1 0x00
[15] CALLDATALOAD
[17] PUSH1 0xe0
[18] SHR
[19] DUP1
[24] PUSH4 0x893d20e8
[25] GT
[28] PUSH2 0x0059

[...]

[2353] PUSH28 0xe0228f06e7bc545e39c1f22cacc71d8a4341037cfdf74d6d9364736f

we changed the visibility from `public` to `external`. `external` is more gas efficient
