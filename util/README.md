PACKAGE DOCUMENTATION

package util
    import "./util/"

    +build windows

CONSTANTS

const (
    MAXSIZE = 500
    PORT    = 3000
)

VARIABLES

var (
    ADDR = [4]byte{10, 8, 0, 1}
)

FUNCTIONS

func Dial() (s.Handle, s.SockaddrInet4, error)
    Etablie une connexion via socket sur le serveur à l'aide des librairies
    POSIX ou WSA.

func SplitFile(path string) (packets [][]byte, err error)
    Cette fonction découpe un fichier en tableau de buffer de taille
    MAXSIZE.

func Write(sd s.Handle, sa s.SockaddrInet4, b []byte) error


