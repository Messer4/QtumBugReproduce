const Sender = artifacts.require("Sender");
const ERC20 = artifacts.require("ERC20PresetMinterPauser");

module.exports = async function (deployer) {
    let sender = await Sender.deployed();
    let token = await ERC20.deployed();

    console.log("Token:", token.address);

    // Create example ERC20 and mint a generous amount of it.
    await token.mint("0x7926223070547d2d15b2ef5e7383e541c338ffe9", "1000000000000000000");

    await token.approve(sender.address, "1000000000000000000");
};
