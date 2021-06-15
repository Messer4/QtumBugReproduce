const Wormhole = artifacts.require("Sender");
const ERC20 = artifacts.require("ERC20PresetMinterPauser");

module.exports = async function (deployer) {
    await deployer.deploy(Wormhole);
    await deployer.deploy(ERC20, "Test Token","TKN");
};
