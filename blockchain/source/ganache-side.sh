./node_modules/ganache-cli/cli.js \
    --accounts=20 --defaultBalanceEther=100000000 \
    --port=7545 --gasLimit=0xffffff --networkId=777 \
    --deterministic --mnemonic="ganache-side" \
    --db="./side-ethereum" \
    | tee ganache-side.log
    # --debug
