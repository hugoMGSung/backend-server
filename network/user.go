package network

import (
	"backend-server/service"
	"backend-server/types"
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

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
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
func (u *userRouter) get(c *gin.Context) {
	fmt.Println("GET 메서드입니다")

	// c.JSON(200, "Succeed!") // 여기를 utils.go 아래 okResponse로 변경
	// u.router.okResponse(c, "OK 테스트입니다")
	//type ApiResponse struct {
	// 	Result      int64  `json:"result"`
	// 	Description string `json:"description"`
	// }
	// u.router.okResponse(c, &types.ApiResponse{
	// 	Result:      1,
	// 	Description: "GET 메서드 성공!",
	// })
	// UserResponse를 사용한다면
	// u.router.okResponse(c, &types.UserResponse{
	// 	ApiResponse: &types.ApiResponse{
	// 		Result:      1,
	// 		Description: "GET 메서드 성공!",
	// 	},
	// 	User: nil,
	// })
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("GET 성공입니다", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("POST 메서드입니다")
	var req types.CreateRequest

	// URL, Query 검증
	// c.ShouldBindUri(); c.ShouldBindQuery() // ...
	// 요청검증
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("POST 바인딩오류!", -1, err.Error()),
		})
	} else if u.userService.Create(req.ToUser()); err != nil {
		u.router.failResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Create POST오류!", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("POST 성공입니다", 1, nil),
		})
	}
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("PUT 메서드입니다")
	var req types.UpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("PUT 바인딩오류!", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("Update PUT오류!", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("PUT 성공입니다", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("DELETE 메서드입니다")
	var req types.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("DELETE 바인딩오류!", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("DELETE 오류!", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("DELETE 성공입니다", 1, nil),
		})
	}
}
