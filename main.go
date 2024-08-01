// リスト3.1
/* 次のコマンドで実行できます。
   （Windowsの場合、sudo は不要です）
  $ sudo go run server.go

 * ほかのウェブサーバが動いているときは、すぐに終了してしまいます。
 * 起動しておいてブラウザに http://localhost/ を入力すると "404 page not found" が表示されます。
   それ以上の機能はありません。
	 終了するには control+C を押してください。

なお、次のコマンドを実行すると01simplest という実行ファイルができます。
  $ go build
その後で次のコマンドを実行するとサーバが起動します。
 sudo ./01simplest

*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello"))
}

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})

	fmt.Println("finished")
	// handler := MyHandler{}

	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: &handler,
	// }

	// server.ListenAndServeTLS("cert.pem", "key.pem")
}
