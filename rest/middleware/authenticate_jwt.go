package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/util"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthenticated", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")

		if len(headerArr) != 2 {
			http.Error(w, "Unauthenticated", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "Unauthenticated", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)

		newSignature := base64UrlEncoding(hash)
		if newSignature != jwtSignature {
			util.SendData(w, "Access denied, better luck next time", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncoding(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
