## ⚠️ Nota importante: Preparación del entorno antes de ejecutar comandos

Si estás en una nueva terminal, o si los comandos como `cryptogen`, `configtxgen` o `peer` no funcionan correctamente, asegúrate de configurar el entorno adecuadamente antes de continuar.

### 🔧 Paso 1: Agregar binarios de Fabric al PATH

Desde la raíz donde se encuentra la carpeta `fabric-samples`, ejecuta:

```
~/blockchain/
├── fabric-samples/       # Contiene los binarios, ejemplos oficiales y scripts de Hyperledger Fabric
│   └── bin/              # Contiene herramientas como peer, cryptogen, configtxgen
│
└── dato-seguro/          # Proyecto personalizado con configuración y red de Dato Seguro
    ├── network/          # Carpeta donde vive la red (docker-compose, crypto, scripts)
    └── README.md         # Documentación del proyecto
```

```bash
cd ~/blockchain/
echo 'export PATH=$PATH:$PWD/fabric-samples/bin' >> ~/.profile
source ~/.profile
```

Esto asegura que los binarios (`cryptogen`, `configtxgen`, `peer`, etc.) estén disponibles en la terminal.

---

### 📂 Paso 2: Ejecutar comandos dentro de la red correcta

Si ves un error como:

```bash
cryptogen: error: open ./crypto-config.yaml: no such file or directory
```

Significa que no estás en el directorio donde se encuentra el archivo de configuración. En el caso del proyecto `dato-seguro`, navega a la carpeta `network`:

```bash
cd ~/blockchain/dato-seguro/network
cryptogen generate --config=./crypto-config.yaml
```

✅ Este comando debería ejecutarse correctamente y generar los certificados para las siguientes organizaciones:

- registrocivil.gob.ec
- cne.gob.ec
- ant.gob.ec
- dinardarp.gob.ec

---

### 💡 Recomendación

Incluye estos pasos al iniciar cualquier sesión de desarrollo o pruebas para evitar errores relacionados con rutas o configuraciones no cargad


# Pasos para levantar la red de Dato Seguro

Este documento detalla los pasos necesarios para iniciar la red de Hyperledger Fabric correspondiente al sistema **Dato Seguro**.


---

## 1. Crear los certificados de las organizaciones

```bash
cryptogen generate --config=./crypto-config.yaml
```

---

## 2. Crear el bloque genesis

```bash
export FABRIC_CFG_PATH=$PWD/configtx.yaml
configtxgen -profile OrdererSingleMSPSolo -channelID system-channel -outputBlock ./channel-artifacts/genesis.block
```

---

## 3. Crear la transacción de configuración del canal

```bash
configtxgen -profile DatoSeguroChannel -outputCreateChannelTx channel-artifacts/channel.tx -channelID datoseguro
```

---

## 4. Crear los Anchor Peers de cada organización

```bash
configtxgen -profile DatoSeguroChannel -outputAnchorPeersUpdate ./channel-artifacts/RegistroCivilMSPanchor.tx -channelID datoseguro -asOrg RegistroCivilMSP

configtxgen -profile DatoSeguroChannel -outputAnchorPeersUpdate ./channel-artifacts/CneMSPanchor.tx -channelID datoseguro -asOrg CneMSP

configtxgen -profile DatoSeguroChannel -outputAnchorPeersUpdate ./channel-artifacts/AntMSPanchor.tx -channelID datoseguro -asOrg AntMSP

configtxgen -profile DatoSeguroChannel -outputAnchorPeersUpdate ./channel-artifacts/DinardarpMSPanchor.tx -channelID datoseguro -asOrg DinardarpMSP
```

---

## 5. Levantar la red con Docker Compose

Una vez configuradas las imágenes de los contenedores:

```bash
export VERBOSE=false
export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=datoseguro
CHANNEL_NAME=$CHANNEL_NAME docker-compose -f docker-compose-cli.yaml up -d
```

---

## 6. Crear el canal desde el CLI

Se crea el canal con el orderer:

```bash
export CHANNEL_NAME=datoseguro
peer channel create -o orderer.gob.ec:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls true \
--cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/gob.ec/orderers/orderer.gob.ec/msp/tlscacerts/tlsca.gob.ec-cert.pem
```

---

## 7. Unir a las organizaciones al canal

Desde el CLI para cada organización:

```bash
peer channel join -b datoseguro.block  # Registro Civil por defecto tiene su id
```

Para las demas organizaciones se tiene que agregar con el script de join_channel_from_cli.sh

Este script permite unir de forma sencilla a las organizaciones adicionales (por ejemplo, `AntMSP`, `CneMSP`, `DinardarpMSP`) al canal `datoseguro`.

```bash
chmod +x join_channel_from_cli.sh
./join_channel.sh
```


---

## ✅ Red lista

Una vez que todas las organizaciones se han unido al canal, la red de Dato Seguro está lista para desplegar chaincodes y operar.
