package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/api/internal/middleware"
)

func generateMetabaseToken(numero uint) (string, error) {
	// Chave secreta do Metabase
	metabaseSecretKey := "3eedf2c1dd3fbe24b86f789918b092e2eadf5c3dffef5b8987b46cf5c37be175"

	// Configuração do payload
	payload := jwt.MapClaims{
		"resource": map[string]interface{}{
			"dashboard": numero,
		},
		"params": map[string]interface{}{},
		"exp":    time.Now().Add(time.Minute * 10).Unix(), // 10 minutos de expiração
	}

	// Criação do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(metabaseSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func generateMetabaseIframeURL(c *gin.Context) {
	numeroStr := c.Param("dash")
	// Geração do token
	numero, err := strconv.ParseUint(numeroStr, 10, 64)
	token, err := generateMetabaseToken(uint(numero))
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao gerar token"})
		return
	}

	// Construção da URL do iframe
	iframeURL := fmt.Sprintf("%s/embed/dashboard/%s#bordered=true&titled=true", "http://localhost:3000", token)

	// Retorno da URL do iframe
	c.JSON(200, gin.H{"iframe_url": iframeURL})
}

func main() {
	// Criação do roteador Gin
	r := gin.Default()
	r.Use(middleware.CorsMiddleware)
	// Rota para obter a URL do iframe Metabase
	r.GET(":dash/metabase-iframe-url", generateMetabaseIframeURL)

	// Inicialização do servidor
	r.Run(":8080")
}
