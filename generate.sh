rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile VLMOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile TFBCOrgChannel -outputCreateChannelTx ./config/tfbcchannel.tx -channelID tfbcchannel
