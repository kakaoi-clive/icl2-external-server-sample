package api

import (
	"encoding/json"
	"icl2-external-server-sample/internal/token"

	"fmt"
	"net/http"
	"strings"
)

const _port = ":8080"

// StartServer 기본적인 http server를 시작하는 함수
func StartServer() {
	http.HandleFunc("/api/generate", generateToken)
	http.HandleFunc("/api/validate", validateToken)

	go func() {
		err := http.ListenAndServe(_port, nil)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Printf("iCL 2.0 External Auth Sample Server Stared at %s\n", _port)
}

// generateToken 토큰을 생성하는 Handler (GET METHOD)
func generateToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/generate Entered.")

	switch r.Method {
	case "GET":
		t, err := token.GenerateToken()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(nil)
		} else {
			fmt.Printf("token generate ok, generated token : %s\n", t)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(fmt.Sprintf("{\"token\": \"%s\"}", t)))
		}
	default:
		fmt.Println("request method is not GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write(nil)
	}

}

type ResponseValidateToken struct {
	Admin bool `json:"admin"`
}

// validateToken 토큰을 확인하는 Handler (GET METHOD)
func validateToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/validate Entered.")

	switch r.Method {
	case "GET":
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			fmt.Println("no authorization header")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(nil)
			return
		}

		userToken := strings.TrimPrefix(authHeader, "Bearer ")
		if authHeader == userToken {
			fmt.Println("no bearer type in authorization header")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(nil)
			return
		}

		t, err := token.ValidateToken(userToken)
		if err != nil {
			fmt.Printf("validate token fail : %s\n", err.Error())
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write(nil)
		} else {
			fmt.Printf("validate ok, user's UUID : %s\n", t.TokenUUID)
			w.WriteHeader(http.StatusOK)
			// 검증한 token이 admin token 인 경우, {"admin" : true} 설정 필요
			rvt := &ResponseValidateToken{
				Admin: false,
			}
			res, err := json.Marshal(rvt)
			if err != nil {
				fmt.Printf("Marshal Response fail : %s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write(nil)
			}
			_, _ = w.Write(res)
		}
	default:
		fmt.Println("request method is not GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write(nil)
	}

}
