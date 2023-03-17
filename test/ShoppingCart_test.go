package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/motion-go/domain"
	"github.com/ribeirosaimon/motion-go/pkg/shoppingcart"
	"github.com/ribeirosaimon/motion-go/test/util"
)

var shoppingCartRouter = func(engine *gin.RouterGroup) {
	shoppingcart.NewShoppingCartRouter(engine.Group("/api/v1"), util.ConnectDatabaseTest)
}

func TestHaveToCreateShoppingCartAndReturnOk(t *testing.T) {
	defer util.RemoveDatabase()
	session, err := util.SignUp(domain.USER, domain.ADMIN)
	if err != nil {
		util.ErrorTest(err.Error())
	}
	fmt.Println(session)
	resp, _, err := util.CreateEngineRequest(http.MethodGet, "/api/v1/shopping-cart/create",
		nil, shoppingCartRouter, session.SessionId)

	if err != nil {
		util.ErrorTest(err.Error())
	}

	if resp.Code != http.StatusCreated {
		util.ErrorTest(fmt.Sprintf("Expected Http status: %d; but is received: %d", http.StatusCreated, resp.Code))
	}

}
