./node_modules/ganache-cli/cli.js \
    --accounts=20 --defaultBalanceEther=100000000 \
    --port=8545 --gasLimit=0xffffff --networkId=666 \
    --deterministic --mnemonic="ganache-side" \
    --db="./main-ethereum" \
    | tee ganache-main.log
    # --debug
