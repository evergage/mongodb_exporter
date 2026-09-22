package shared

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net"
	"net/url"
	"path/filepath"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCARotationTLSConnection(t *testing.T) {
	oldCA, _, oldPEM := makeCertificate(t, nil, nil, "old-ca", true, false)
	newCA, newCAKey, newPEM := makeCertificate(t, nil, nil, "new-ca", true, false)
	_, _, serverPEM := makeCertificate(t, newCA, newCAKey, "127.0.0.1", false, false)
	_, _, clientPEM := makeCertificate(t, newCA, newCAKey, "mongodb-exporter", false, true)

	roots := x509.NewCertPool()
	roots.AddCert(oldCA)
	roots.AddCert(newCA)
	serverCert, err := tls.X509KeyPair(serverPEM, serverPEM)
	if err != nil {
		t.Fatal(err)
	}
	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{serverCert}, ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: roots, MinVersion: tls.VersionTLS12})
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	done := make(chan struct{})
	defer close(done)
	go serveMongoTLS(listener, done)

	dir := t.TempDir()
	caPath := filepath.Join(dir, "ca-chain.pem")
	clientPath := filepath.Join(dir, "client.pem")
	if err := ioutil.WriteFile(caPath, append(oldPEM, newPEM...), 0600); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(clientPath, clientPEM, 0600); err != nil {
		t.Fatal(err)
	}
	uri := fmt.Sprintf("mongodb://%s/?tls=true&tlsCAFile=%s&tlsCertificateKeyFile=%s&serverSelectionTimeoutMS=3000", listener.Addr().String(), url.QueryEscape(caPath), url.QueryEscape(clientPath))
	client := MongoClient(&MongoSessionOpts{URI: uri})
	if client == nil {
		t.Fatal("MongoClient rejected the rotated CA bundle")
	}
	defer client.Disconnect(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatal(err)
	}
}

func makeCertificate(t *testing.T, parent *x509.Certificate, parentKey *rsa.PrivateKey, name string, isCA, client bool) (*x509.Certificate, *rsa.PrivateKey, []byte) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{SerialNumber: serial, Subject: pkix.Name{CommonName: name}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment, IsCA: isCA, BasicConstraintsValid: true}
	if isCA {
		tmpl.KeyUsage |= x509.KeyUsageCertSign
		tmpl.ExtKeyUsage = nil
	} else if client {
		tmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}
	} else {
		tmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}
		tmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	}
	if parent == nil {
		parent = tmpl
		parentKey = key
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, parent, &key.PublicKey, parentKey)
	if err != nil {
		t.Fatal(err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	return cert, key, append(certPEM, keyPEM...)
}

func serveMongoTLS(listener net.Listener, done <-chan struct{}) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go func() {
			defer conn.Close()
			for {
				header := make([]byte, 16)
				if _, err := io.ReadFull(conn, header); err != nil {
					return
				}
				length := int(binary.LittleEndian.Uint32(header))
				if length < 16 || length > 16<<20 {
					return
				}
				body := make([]byte, length-16)
				if _, err := io.ReadFull(conn, body); err != nil {
					return
				}
				requestID := int32(binary.LittleEndian.Uint32(header[4:8]))
				op := int32(binary.LittleEndian.Uint32(header[12:16]))
				doc, _ := bson.Marshal(bson.D{{Key: "ok", Value: 1.0}, {Key: "ismaster", Value: true}, {Key: "isWritablePrimary", Value: true}, {Key: "maxWireVersion", Value: 13}, {Key: "minWireVersion", Value: 0}, {Key: "logicalSessionTimeoutMinutes", Value: 30}})
				var reply []byte
				if op == 2004 {
					reply = make([]byte, 36+len(doc))
					binary.LittleEndian.PutUint32(reply[0:4], uint32(len(reply)))
					binary.LittleEndian.PutUint32(reply[8:12], uint32(requestID))
					binary.LittleEndian.PutUint32(reply[12:16], 1)
					binary.LittleEndian.PutUint32(reply[32:36], 1)
					copy(reply[36:], doc)
				} else {
					reply = make([]byte, 21+len(doc))
					binary.LittleEndian.PutUint32(reply[0:4], uint32(len(reply)))
					binary.LittleEndian.PutUint32(reply[8:12], uint32(requestID))
					binary.LittleEndian.PutUint32(reply[12:16], 2013)
					reply[20] = 0
					copy(reply[21:], doc)
				}
				if _, err := conn.Write(reply); err != nil {
					return
				}
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}
}
