package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once // 한번만 호출
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service 추가예정
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		// 사용할 CRUD 함수 등록
		// userRouterInstance.router.engine.POST("/", userRouterInstance.create)
		// userRouterInstance.router.engine.PUT("/", userRouterInstance.update)
		// userRouterInstance.router.engine.DELETE("/", userRouterInstance.delete)
		// userRouterInstance.router.engine.GET("/", userRouterInstance.get)
		// 이게 너무 기니까... v0.5.1  네트워크의 레지스터 함수로 변경
		router.regPOST("/", userRouterInstance.create)
		router.regGET("/", userRouterInstance.get)
		router.regUPDATE("/", userRouterInstance.update)
		router.regDELETE("/", userRouterInstance.delete)

	})

	return userRouterInstance
}

// CRUD
// 매개변수에 반드시 gin.Context가 들어가야함
// http.HandleFunc("/", fff)
// func fff(w http.ResponseWriter, r *http.Request) 와 동일
func (u *userRouter) create(c *gin.Context) {
	fmt.Println("POST 메서드입니다")
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("GET 메서드입니다")

}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("PUT 메서드입니다")

}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("DELETE 메서드입니다")

}
