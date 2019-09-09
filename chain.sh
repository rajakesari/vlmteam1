echo "Installing vlmteam1 chaincode to peer0.manufacturer.vlmteam1.com..."

# install chaincode
# Install code on manufacturer peer
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlmteam1.com/users/Admin@manufacturer.vlmteam1.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlmteam1.com:7051" cli peer chaincode install -n vlmteam1cc -v 2.1 -p github.com/vlmteam1/go -l golang

echo "Installed vlmteam1 chaincode to peer0.manufacturer.vlmteam1.com"

echo "Installing vlmteam1 chaincode to peer0.customer.vlmteam1.com...."

# Install code on customer peer
docker exec -e "CORE_PEER_LOCALMSPID=CustomerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.vlmteam1.com/users/Admin@customer.vlmteam1.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.vlmteam1.com:7051" cli peer chaincode install -n vlmteam1cc -v 2.1 -p github.com/vlmteam1/go -l golang

echo "Installed vlmteam1 chaincode to peer0.customer.vlmteam1.com"

echo "Installing vlmteam1 chaincode to peer0.RTA.vlmteam1.com..."
# Install code on RTA peer
docker exec -e "CORE_PEER_LOCALMSPID=RTAMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/RTA.vlmteam1.com/users/Admin@RTA.vlmteam1.com/msp" -e "CORE_PEER_ADDRESS=peer0.RTA.vlmteam1.com:7051" cli peer chaincode install -n vlmteam1cc -v 2.1 -p github.com/vlmteam1/go -l golang

sleep 5

echo "Installed vlmteam1 chaincode to peer0.RTA.vlmteam1.com"


echo "Installing vlmteam1 chaincode to peer0.Scrapper.vlmteam1.com..."
# Install code on Scrapper peer
docker exec -e "CORE_PEER_LOCALMSPID=ScrapperMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/scrapper.vlmteam1.com/users/Admin@scrapper.vlmteam1.com/msp" -e "CORE_PEER_ADDRESS=peer0.scrapper.vlmteam1.com:7051" cli peer chaincode install -n vlmteam1cc -v 2.1 -p github.com/vlmteam1/go -l golang

sleep 5

echo "Installed vlmteam1 chaincode to peer0.scrapper.vlmteam1.com"

echo "Instantiating vlmteam1 chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.vlmteam1.com/users/Admin@manufacturer.vlmteam1.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.vlmteam1.com:7051" cli peer chaincode instantiate -o orderer.vlmteam1.com:7050 -C tfbcchannel -n vlmteam1cc -l golang -v 2.1 -c '{"Args":[""]}' -P "OR ('ManufacturerMSP.member','CustomerMSP.member','RTAMSP.member','ScrapperMSP.member')"

echo "Instantiated vlmteam1 chaincode."

echo "Following is the docker network....."

