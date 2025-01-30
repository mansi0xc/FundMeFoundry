// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract fundCrowd{
    /* Errors */
    error aurora__DeadlineMustBeInTheFuture();
    error aurora__NotOwner();
    error aurora__GoalNotReached();
    error aurora__GoalReached();
    error aurora__DeadlineReached();
    error aurora__TransferFailed();

    struct Funders{
        address funder;
        uint256 amount;
        string name;
    }

    struct Cause{
        uint256 id;
        string title;
        string name;
        address owner;
        string image;
        string about;
        uint256 goal;
        uint256 deadline;
        uint256 raised;
        Funders[] funders;
        bool isWithdrawn;
    }

    mapping(uint256=>Cause) public allCauses;

    uint256 public s_count;

    modifier OnlyOwner(uint256 causeID){
        Cause storage cause = allCauses[causeID];
        if (cause.owner != msg.sender){
            revert aurora__NotOwner();
        }
        _;
    }

    // creating a new cause
    function createCause(
        string memory _title,
        string memory _name,
        address _owner,
        string memory _image,
        string memory _about,
        uint256 _goal,
        uint256 _deadline
    ) public {

        if (block.timestamp > _deadline){
            revert aurora__DeadlineMustBeInTheFuture();
        }

        Cause storage cause = allCauses[s_count];

        cause.id = s_count;
        cause.title = _title;
        cause.name = _name;
        cause.owner = _owner;
        cause.image = _image;
        cause.about = _about;
        cause.goal = _goal;
        cause.deadline = _deadline;

        s_count++;
    }

    // funding a cause
    function fundCause(uint256 causeID, string memory _name) public payable{
        uint256 _amount = msg.value;
        Cause storage cause = allCauses[causeID];

        if (block.timestamp > cause.deadline){
            revert aurora__DeadlineReached();
        }
        if (cause.goal == cause.raised){
            revert aurora__GoalReached();
        }
        cause.funders.push(Funders(msg.sender, _amount, _name));
        cause.raised += _amount;
    }

    // withdraw funds
    function withdraw(uint256 causeID) public OnlyOwner(causeID){
        Cause storage cause = allCauses[causeID];

        if (cause.raised < cause.goal){
            revert aurora__GoalNotReached();
        }

        // Transfer the raised funds to the owner
        (bool success, ) = payable(msg.sender).call{value: cause.raised}("");
        if (!success) {
            revert aurora__TransferFailed();
        }

        cause.isWithdrawn = true;
    }

    // Getter functions

    function getCauses() public view returns(Cause[] memory){
        uint256 activeCausesCount = 0;
        uint256 totalCauses = s_count;

        for (uint256 i = 0; i<totalCauses; i++){
            if (!allCauses[i].isWithdrawn){
                activeCausesCount++;
            }
        }

        Cause[] memory activeCauses = new Cause[](activeCausesCount);
        uint256 index=0;

        for (uint256 i=0; i<totalCauses; i++){
            if (!allCauses[i].isWithdrawn){
                activeCauses[index] = allCauses[i];
            }   
        }

        return activeCauses;
    }
    
    function getMyCause(address _owner) public view returns(Cause[] memory){
        uint256 count=0;
        uint256 totalCauses=s_count;

        // count the total causes of the given address
        for (uint256 i=0; i<totalCauses; i++){
            if (allCauses[i].owner == _owner){
                count++;
            }
        }

        Cause[] memory ownerCauses = new Cause[](count);
        uint256 index=0;

        for(uint256 i=0; i<totalCauses; i++){
            if (allCauses[i].owner == _owner){
                Cause memory cause = allCauses[i];
                // cause.id = i;
                ownerCauses[index] = cause;
                index++;
            }
        }

        return ownerCauses;
    }

    function getContributedCauses(address _funder) public view returns(Cause[] memory){
        uint256 count=0;
        uint256 totalCauses=s_count;

        // Count the number of causes this address has contributed to
        for (uint256 i = 0; i < totalCauses; i++) {
            Cause storage cause = allCauses[i];
            for (uint256 j = 0; j < cause.funders.length; j++) {
                if (cause.funders[j].funder == _funder) {
                    count++;
                    break; // Move to the next cause once a contribution is found
                }
            }
        }

        // Initialize the array with the counted size
        Cause[] memory contributedCauses = new Cause[](count);
        uint256 index = 0;

        // Populate the array with causes contributed by _funder
        for (uint256 i = 0; i < totalCauses; i++) {
            Cause storage campaign = allCauses[i];
            for (uint256 j = 0; j < campaign.funders.length; j++) {
                if (campaign.funders[j].funder == _funder) {
                    contributedCauses[index] = campaign;
                    index++;
                    break;
                }
            }
        }

        return contributedCauses;
    }
}