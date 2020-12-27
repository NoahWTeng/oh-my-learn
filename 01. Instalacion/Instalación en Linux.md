

## 1. Instalación en Linux (Ubuntu)

------

### 1.1 Actualizar sistema Linux

```
sudo apt-get update 
sudo apt-get upgrade
```

![sudo-apt-update-upgrade](/home/jirohero/Desktop/Goland/Img/sudo-apt-update-upgrade.png)

### 1.2 Descarga del paquete instalación de Go para Linux

![go-download-windows-amd64](/home/jirohero/Desktop/Goland/Img/go-download-windows-amd64.png)

### 1.3  Extracción binario GoLand

Ahora, necesitamos extraer el tarball usando el comando tar en un directorio de nuestra elección.

Usaremos**`/usr/local/`**.

```
// -C significa moverse al directorio dato
// -xzf son 3 indicadores, -x para extraer, -z para gzip, -f archivo.

sudo tar -C /usr/local/ -xzf go1.13.5.linux-amd64.tar.gz
```

### 1.3 Configuración ruta de Go

```
echo $PATH
```

![gopath-is-not-set-1](/home/jirohero/Desktop/Goland/Img/gopath-is-not-set-1.png)

Vemos que aun no esta configurada.

```
sudo vim $HOME/.profile
```

```
// ./profile
export PATH=$PATH:/usr/local/go/bin
// guardamos
```

```
// aplicamos cambios en el sistema
source .profile
```

Por ultimo verificamos Go

```
$ go version
go version go1.15.6 linux/amd64
```

