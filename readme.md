### Modul Oluşturmak

```sh
go mod init [module_name]
```

### Modul Ekleme

```sh
# go geom modulunu projeye ekler
go get github.com/twpayne/go-geom
```

### Modulu kullanma

```go
package main

import "github.com/twpayne/go-geom"

func main() {
	//geom modulu altındaki newMultiPolygon fonksiyonun cagırır
	geom.NewMultiPolygon()
}

```

### Kullanılmayan Modulu Kaldırma

```sh
go mod tidy
```

### Docker Kurulumu

https://docs.docker.com/desktop/install/ubuntu/

- Launch Docker Desktop

```sh
systemctl --user start docker-desktop
```

- Get Props

```sh
docker compose version
Docker Compose version v2.17.3

docker --version
Docker version 23.0.5, build bc4487a

docker version
Client: Docker Engine - Community
 Cloud integration: v1.0.31
 Version:           23.0.5
 API version:       1.42
```
