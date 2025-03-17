#!/bin/bash

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

# Instalar Git
check_install git git

# Instalar cURL
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

# Instalar Go (opcional)
if ! command -v go &> /dev/null; then
    echo "Go no está instalado. Instalando..."
    sudo apt-get update
    sudo apt-get install -y golang-go
else
    echo "Go ya está instalado."
fi

# Instalar jq (opcional)
check_install jq jq

# Descargar e instalar Hyperledger Fabric
if [ ! -f "install-fabric.sh" ]; then
    echo "Descargando script de instalación de Hyperledger Fabric..."
    curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh
    chmod +x install-fabric.sh
fi

# Ejecutar el script de instalación con todos los componentes
./install-fabric.sh docker samples binary

echo "Instalación completada."
