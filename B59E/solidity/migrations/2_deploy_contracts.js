var BCPPoints = artifacts.require("./BCPPoints.sol");
var LIPPoints = artifacts.require("./LIPPoints.sol");
var Platform  = artifacts.require("./CharityPlatform.sol");

module.exports = function(deployer, network, accounts) {
    deployer.deploy(BCPPoints)
    deployer.deploy(LIPPoints)
    deployer.deploy(Platform)
}
