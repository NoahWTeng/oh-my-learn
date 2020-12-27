

## 1. Instalación en Mac

------

### 1.1 Descarga del paquete instalación de Go para Mac

Ir al pagina oficial de GoLang https://golang.org/dl/ y descarga el paquete de instalación.

![go-mac-os-installer-package](/home/jirohero/Desktop/Goland/Img/go-mac-os-installer-package.png)

### 1.2 Instalación 

Arranca el paquete de instalación y sigue estos pasos.

![go-installer-package-1](/home/jirohero/Desktop/Goland/Img/go-installer-package-1.png)

![go-installer-package-2](/home/jirohero/Desktop/Goland/Img/go-installer-package-2.png)

![go-installer-package-3-1024x719](/home/jirohero/Desktop/Goland/Img/go-installer-package-3-1024x719.png)

> ruta de instalación:  **` /usr/local/go/`**.
>
> ruta de lanzamiento: **`/usr/local/go/bin/go `**

### 1.3 Verificación de la instalación

Abrimos terminal y lanzamos el siguiente comando.

```
$ go version
go version go1.15.6 darwin/amd64
```

------

*En el caso de que no se reconozca el comando, deberíamos agregar el PATH.

Abrimos la terminal y lanzamos este comando.

```
export PATH=/usr/local/go/bin:$PATH
```

Si queremos hacer permanente el comando, porque ahora mismo con el comando que lanzamos solo servirá para una sesión en cuando se cierre la terminal y abramos otra nueva deberíamos lanzar otra vez el comando. 

```
vim ~/.bash_profile
```

```
// ./bash_profile
export PATH=/usr/local/go/bin:$PATH
```

