# Instalación del Entorno y Hyperledger Fabric

Este proyecto proporciona un script de instalación automatizado para configurar las dependencias necesarias y el entorno de Hyperledger Fabric en un sistema basado en Linux.

## Requisitos Previos

Antes de ejecutar el script, asegúrate de tener permisos de administrador (`sudo`) y conexión a Internet estable.

## Instalación

### 1. Clonar el Repositorio
```bash
 git clone <URL_DEL_REPOSITORIO>
 cd <NOMBRE_DEL_REPOSITORIO>
```

### 2. Ejecutar el Script de Instalación
El script `install-env.sh` instalará las dependencias necesarias, incluyendo:

- Git
- cURL
- Docker y Docker Compose
- Go (opcional)
- jq (opcional)
- Hyperledger Fabric (binaries, imágenes Docker y samples)

Para iniciar la instalación, ejecuta:

```bash
chmod +x install-env.sh
./install-env.sh
```

Si el sistema requiere permisos de superusuario, usa:

```bash
sudo ./install-env.sh
```

## Verificación de Instalación
Una vez finalizada la instalación, verifica que los componentes de Hyperledger Fabric están instalados correctamente:

```bash
fabric-samples/bin/configtxgen --version
```

Si el comando devuelve una versión válida, la instalación se ha realizado con éxito.

## Notas
- Si Docker no se inicia automáticamente, puedes ejecutarlo manualmente con:
  ```bash
  sudo systemctl start docker
  ```
- Para evitar usar `sudo` en los comandos de Docker, agrega tu usuario al grupo `docker` y reinicia la sesión:
  ```bash
  sudo usermod -aG docker $USER
  ```

## Contacto
Si encuentras problemas durante la instalación, revisa la documentación oficial de Hyperledger Fabric o consulta con la comunidad.

