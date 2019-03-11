./node_modules/ganache-cli/cli.js \
    --accounts=20 --defaultBalanceEther=100000000 \
    --port=7545  --networkId=777 \
    --gasLimit=0xfffffffffff --gasPrice=0x01 \
    --deterministic --mnemonic="ganache-side" \
    --db="./side-ethereum" \
    | tee ganache-side.log
    # --debug
