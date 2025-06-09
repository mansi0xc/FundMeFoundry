# forge build
# forge test -vv (for the console.log statements)  
how does the number of v's in the command affect the output?
according to the experiment, more number of v's is giving more info

the above commands were for the following code:
import {Test, console} from "forge-std/Test.sol";

contract FundeMeTest is Test{
    uint256 number = 1;
    function setUp() external{
        number = 2;
    }
    
    function testDemo() public view {
        console.log(number);
        console.log("Hello");
        assertEq(number, 2);
        console.log(number);
    }
}

function setUp() external{
        // us -> FundMeTest -> FundMe
        fundMe = new FundMe();
    }
    
    function testMinUSDisFive() public view {
        assertEq(fundMe.MINIMUM_USD(), 5e18);
    }

function testOwnerIsMsgSender() public view{
        address owner = fundMe.i_owner();  // Store the value in a local variable
        console.log("Owner (i_owner): ", owner);
        console.log("Test Contract Address: ", address(this));
        console.log("msg.sender: ", msg.sender);
        assertEq(fundMe.i_owner(), address(this));
        // assertEq(fundMe.i_owner(), msg.sender);  // will give error
    }

# ################################################ #

# forge test -vvvv --fork-url $SEPOLIA
# forge coverage --fork-url $SEPOLIA

For forked tests:
# forge test <-vvvv>
# forge test --match-test <test function name> or 
**forge test --mt <test function name>**
# forge coverage (is successful only if all tests pass) 
# forge test --summary --detailed

**REMEMBER:** Whenever you have a situation where two or more `vm` cheatcodes come one after the other keep in mind that these would ignore one another. In other words, when we call `vm.expectRevert();` that won't apply to `vm.prank(alice);`, it will apply to the `withdraw` call instead. The same would have worked if these had been reversed. Cheatcodes affect transactions, not other cheatcodes.

# forge inspect FundMe storageLayout

function testWithdrawFromASingleFunder() public funded {
    // Arrange
    uint256 startingFundMeBalance = address(fundMe).balance;
    uint256 startingOwnerBalance = fundMe.getOwner().balance;

    vm.txGasPrice(GAS_PRICE);
    uint256 gasStart = gasleft();

    // Act
    vm.startPrank(fundMe.getOwner());
    fundMe.withdraw();
    vm.stopPrank();

    uint256 gasEnd = gasleft();
    uint256 gasUsed = (gasStart - gasEnd) * tx.gasprice;
    console.log("Withdraw consumed: %d gas", gasUsed);

    // Assert
    uint256 endingFundMeBalance = address(fundMe).balance;
    uint256 endingOwnerBalance = fundMe.getOwner().balance;
    assertEq(endingFundMeBalance, 0);
    assertEq(
        startingFundMeBalance + startingOwnerBalance,
        endingOwnerBalance
    );

}