# util
--
    import "github.com/adrien3d/gobox/util"

+build !windows

+build windows

## Usage

```go
const (
	MAXSIZE = 500
	PORT    = 3000
)
```

```go
var (
	ADDR = [4]byte{10, 8, 0, 1}
)
```

#### func  Dial

```go
func Dial() (s.Handle, s.SockaddrInet4, error)
```
Etablie une connexion via socket sur le serveur à l'aide des librairies POSIX ou
WSA.

#### func  SplitFile

```go
func SplitFile(path string) (packets [][]byte, err error)
```
Cette fonction découpe un fichier en tableau de buffer de taille MAXSIZE.

#### func  Write

```go
func Write(sd s.Handle, sa s.SockaddrInet4, b []byte) error
```
