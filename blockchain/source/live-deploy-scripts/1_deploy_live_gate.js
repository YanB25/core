let GateMultisig = artifacts.require('./MultiSigWallet.sol');
let GateKeeperLive = artifacts.require('./SimpleGatekeeperWithLimitLive.sol');
let SNMMaster = artifacts.require('./SNM.sol');

let MSOwners = [
    '0x5c865774723bf00895b3620700998906e58085fe',
    '0x29f3ea08889d7cc9f9d609850ff26b65a0315469',

    '0xdaec8F2cDf27aD3DF5438E5244aE206c5FcF7fCd',
    '0xd9a43e16e78c86cf7b525c305f8e72723e0fab5e',
    '0x72cb2a9AD34aa126fC02b7d32413725A1B478888',
    '0x1f50Be5cbFBFBF3aBD889e17cb77D31dA2Bd7227',
    '0xe062C67207F7E478a93EF9BEA39535d8EfFAE3cE',
    '0x5fa359a9137cc5ac2a85d701ce2991cab5dcd538',
    '0x7aa5237e0f999a9853a9cc8c56093220142ce48e',
    '0xd43f262536e916a4a807d27080092f190e25d774',
    '0xdd8422eed7fe5f85ea8058d273d3f5c17ef41d1c',

];

let MSRequired = 1;
// let freezingTime = 60 * 15;
let freezingTime = 0;
// let SNMMasterchainAddress = '0x983f6d60db79ea8ca4eb9968c6aff8cfa04b3c63';
// let actualGasPrice = 3000000000;
let actualGasPrice = 0;

module.exports = function (deployer, network) {
    deployer.then(async () => { // eslint-disable-line promise/catch-or-return
        if (network === 'master') {
            // 0) deploy `SNM Master`
            await deployer.deploy(SNMMaster, { gasPrice: actualGasPrice });
            let token = await SNMMaster.deployed();

            // 1) deploy `GatekeeperLive` multisig
            await deployer.deploy(GateMultisig, MSOwners, MSRequired, { gasPrice: actualGasPrice });
            let multisig = await GateMultisig.deployed();

            // 2) deploy Live Gatekeeper
            await deployer.deploy(GateKeeperLive, token.address, freezingTime, { gasPrice: actualGasPrice });
            let gk = await GateKeeperLive.deployed();

            // 2.1) add keeper with 100k limit for testing
            await gk.ChangeKeeperLimit('0x29603ad75d6c9d04282b4b046c31527e5ed5b3ac', 100000 * 1e18, { gasPrice: actualGasPrice }); // eslint-disable-line max-len
            await gk.ChangeKeeperLimit('0xed8d41b14113f2f545e879939b0dfc5639838f3b', 100000 * 1e18, { gasPrice: actualGasPrice }); // eslint-disable-line max-len
            // await gk.ChangeKeeperLimit('0xad0636e65ac989168a3b29ac82b38f3086bd35f8', 100000 * 1e18, { gasPrice: actualGasPrice }); // eslint-disable-line max-len
            await gk.ChangeKeeperLimit(gk.address, 100000 * 1e18, { gasPrice: actualGasPrice }); // eslint-disable-line max-len

            // 3) transfer Live Gatekeeper ownership to `GatekeeperLive` multisig
            await gk.transferOwnership(multisig.address, { gasPrice: actualGasPrice });
        }
    });
};
