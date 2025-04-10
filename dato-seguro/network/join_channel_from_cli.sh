#!/bin/bash
#Este script se debe ejecutar dentro del cli de hyperledger fabric, luego de haber ingresado la primera org (registrocivil.gob.ec), para unir las otras organizaciones al canal.
# Este script se utiliza para unir las organizaciones al canal de Hyperledger Fabric.
CHANNEL_NAME=datoseguro

# Unir AntMSP
echo "🛰️  Uniéndose al canal como AntMSP..."
export CORE_PEER_LOCALMSPID="CneMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/cne.gob.ec/peers/peer0.cne.gob.ec/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/cne.gob.ec/users/Admin@cne.gob.ec/msp
export CORE_PEER_ADDRESS=peer0.cne.gob.ec:7051

peer channel join -b datoseguro.block
echo "✅ CneMSP unido al canal."

# Unir AntMSP
echo "🛰️  Uniéndose al canal como AntMSP..."
export CORE_PEER_LOCALMSPID="AntMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ant.gob.ec/peers/peer0.ant.gob.ec/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/ant.gob.ec/users/Admin@ant.gob.ec/msp
export CORE_PEER_ADDRESS=peer0.ant.gob.ec:7051

peer channel join -b datoseguro.block
echo "✅ AntMSP unido al canal."

# Unir DinardarpMSP
echo "🛰️  Uniéndose al canal como DinardarpMSP..."
export CORE_PEER_LOCALMSPID="DinardarpMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dinardarp.gob.ec/peers/peer0.dinardarp.gob.ec/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dinardarp.gob.ec/users/Admin@dinardarp.gob.ec/msp
export CORE_PEER_ADDRESS=peer0.dinardarp.gob.ec:7051

peer channel join -b datoseguro.block
echo "✅ DinardarpMSP unido al canal."
