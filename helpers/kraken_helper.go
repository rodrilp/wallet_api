package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
	"os"
)

    func getKrakenSignature(urlPath string, values url.Values, secret []byte) string {
          
        sha := sha256.New()
        sha.Write([]byte(values.Get("nonce") + values.Encode()))
        shasum := sha.Sum(nil)

        mac := hmac.New(sha512.New, secret)
        mac.Write(append([]byte(urlPath), shasum...))
        macsum := mac.Sum(nil)
        return base64.StdEncoding.EncodeToString(macsum)
    }

    func GetSignature(uri string, payload url.Values) string  {
		// API-Sing parameter is generated with the following calc
		// HMAC-SHA512 of (URI path + SHA256(nonce + POST data)) and base64 decoded secret API key

        apiSecret := os.Getenv("KRAKEN_PRIVATE_KEY")
		b64DecodedSecret, _ := base64.StdEncoding.DecodeString(apiSecret)

		// nonce should be always an ioncrease value
		// One very nice approach is to use the timestamp
		//  this can never be a lower value than a previous value
		// nonce := time.Now().UnixMilli()
		// it is passed in the payload

        signature := getKrakenSignature(uri, payload, b64DecodedSecret)

		return signature
    }