./node_modules/ganache-cli/cli.js \
    --accounts=20 --defaultBalanceEther=100000000 \
    --port=8545  --networkId=666 \
    --gasLimit=0xfffffffffff --gasPrice=0x01 \
    --deterministic --mnemonic="ganache-side" \
    --db="./main-ethereum" \
    | tee ganache-main.log
    # --debug
