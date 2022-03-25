package tlsLayer

import (
	"crypto/tls"
	"log"
	"net"
	"unsafe"

	"github.com/hahahrfool/v2ray_simple/utils"
	utls "github.com/refraction-networking/utls"
)

// 关于utls的简单分析，可参考
//https://github.com/hahahrfool/v2ray_simple/discussions/7

type Client struct {
	tlsConfig  *tls.Config
	uTlsConfig *utls.Config
	use_uTls   bool
	alpnList   []string
}

func NewClient(host string, insecure bool, use_uTls bool, alpnList []string) *Client {

	c := &Client{
		use_uTls: use_uTls,
	}

	if use_uTls {

		c.uTlsConfig = &utls.Config{
			InsecureSkipVerify: insecure,
			ServerName:         host,
			NextProtos:         c.alpnList,
		}
		if utils.CanLogInfo() {
			log.Println("using utls and Chrome fingerprint for", host)
		}
	} else {
		c.tlsConfig = &tls.Config{
			InsecureSkipVerify: insecure,
			ServerName:         host,
			NextProtos:         c.alpnList,
		}

	}

	return c
}

func (c *Client) Handshake(underlay net.Conn) (tlsConn *Conn, err error) {

	if c.use_uTls {
		utlsConn := utls.UClient(underlay, c.uTlsConfig, utls.HelloChrome_Auto)
		err = utlsConn.Handshake()
		if err != nil {
			return
		}
		tlsConn = &Conn{
			Conn:           utlsConn,
			ptr:            unsafe.Pointer(utlsConn.Conn),
			tlsPackageType: utlsPackage,
		}

	} else {
		officialConn := tls.Client(underlay, c.tlsConfig)
		err = officialConn.Handshake()
		if err != nil {
			return
		}

		tlsConn = &Conn{
			Conn:           officialConn,
			ptr:            unsafe.Pointer(officialConn),
			tlsPackageType: official,
		}

	}
	return
}
