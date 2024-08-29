package network

import "github.com/gin-gonic/gin"

// v0.6.1 utils로 이동
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

// Response 형태 구성 유틸함수
func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

func (n *Network) failResponse(c *gin.Context, result interface{}) {
	c.JSON(400, result)
}
