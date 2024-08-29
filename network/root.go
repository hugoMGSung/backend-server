package network

import (
	"github.com/gin-gonic/gin"
)

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engine: gin.New(),
	}

	// newUserRouter를 두번 호출해도 한번만 동작
	// r.engine.Run(":8080")
	newUserRouter(r)
	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}

// userRouterInstance.router.engine.POST("/", userRouterInstance.create)
//
//	userRouterInstance.router.engine.PUT("/", userRouterInstance.update)
//	userRouterInstance.router.engine.DELETE("/", userRouterInstance.delete)
//	userRouterInstance.router.engine.GET("/", userRouterInstance.get)
//
// v0.5.1 레지스터 함수 등록
func (n *Network) regPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) regGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)

}

func (n *Network) regUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)

}

func (n *Network) regDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)

}
