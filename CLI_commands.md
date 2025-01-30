// This is for compiling and deploying contract
forge compile
anvil
# forge create SimpleStorage --interactive
or forge create SimpleStorage --rpc-url <> --private-key <>

// This is for deploying a contract within a contract?
forge script script/DeploySimpleStorage.s.sol
forge script script/DeploySimpleStorage.s.sol --rpc-url http://127.0.0.1:8545
# forge script script/DeploySimpleStorage.s.sol --rpc-url http://127.0.0.1:8545 --broadcast --private-key <>

//to interact with the code
cast send <contract address> "function_name(<argument type>)" <argument> --rpc-url <> --private-key <>
# example : cast send 0x5FbDB2315678afecb367f032d93F642f64180aa3 "store(uint256)" 123 --rpc-url $RPC_URL --private-key $PRIVATE_KEY
cast call <contract address> "function_name(<argument type>)"
# example : cast call 0x5FbDB2315678afecb367f032d93F642f64180aa3 "retreive()"
cast --to-base <hex value> dec


// to convert hex to dec
cast --to-base <hex value> dec

//to remove history from CLI
history -c
rm .bash_history

//to auto format
either go to the command palette and format it or
forge fmt


// always run source .env to enable and save all changes in the .env file