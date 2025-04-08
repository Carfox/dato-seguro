#!/bin/bash

sudo apt-get update && sudo apt-get upgrade -y
# Función para verificar si un programa está instalado
check_install() {
    if ! command -v $1 &> /dev/null; then
        echo "$1 no está instalado. Instalando..."
        sudo apt-get update
        sudo apt-get install -y $2
    else
        echo "$1 ya está instalado."
    fi
}

check_install git git
check_install curl curl

# Instalar Docker y Docker Compose
if ! command -v docker &> /dev/null; then
    echo "Docker no está instalado. Instalando..."
    sudo apt-get update
    sudo apt-get -y install docker.io docker-compose
    sudo systemctl start docker
    sudo systemctl enable docker
    sudo usermod -aG docker $USER
    echo "Docker instalado correctamente. Reinicia la sesión para aplicar los cambios de grupo."
else
    echo "Docker ya está instalado."
fi

# Instalar Go
if ! command -v go &> /dev/null; then
    echo "Go no está instalado. Instalando..."
    sudo apt-get update
    sudo apt-get install -y golang-go
else
    echo "Go ya está instalado."
fi

check_install jq jq

# Descargar e instalar Hyperledger Fabric
if [ ! -f "install-fabric.sh" ]; then
    echo "Descargando script de instalación de Hyperledger Fabric..."
    curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh
    chmod +x install-fabric.sh
fi

# Ejecutar el script de instalación con todos los componentes
read -p "¿Deseas instalar Hyperledger Fabric? (s/n): " respuesta

if [[ "$respuesta" == "s" || "$respuesta" == "S" ]]; then
    echo "Instalando Hyperledger Fabric..."
    ./install-fabric.sh docker samples binary
    echo "Instalación completada."
else
    echo "Saltando la instalación de Hyperledger Fabric."
fi

# Continuar con el siguiente paso
echo "Instalación finalizada."


# Habilitar comandos de Hyperledger-Fabric a cli. 
echo 'export PATH=$PATH:$PWD/fabric-samples/bin' >> ~/.profile
source ~/.profile