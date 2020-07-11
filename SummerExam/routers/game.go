package routers

import (
	"Advance/SummerExam/util/response"
	"github.com/gin-gonic/gin"
	"time"
)

func ChessboardInit() [15][15]int {
	var chessboard [15][15]int
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			chessboard[i][j] = 0
		}
	}
	return chessboard
} //初始化棋盘，未下棋位置数据为0

func Game(g *gin.Context) {
	chessboard := ChessboardInit()
	chessOne := make(chan [2]int)
	chessTwo := make(chan [2]int)
	var row, col int
	var a, b int
	go func() {
		for {
			select {
			case gamerOne := <-chessOne:
				row = gamerOne[0]
				col = gamerOne[1]
				chessboard[row][col] = 1
				a++
			case gamerTwo := <-chessTwo:
				row = gamerTwo[0]
				col = gamerTwo[1]
				chessboard[row][col] = 2
				b++
			default:
				time.Sleep(time.Second * 45)
				if a > b {
					msg := "一号玩家已超时，系统自动判负，二号玩家胜利！"
					response.OK(g, msg)
				} else {
					msg := "二号玩家已超时，系统自动判负，一号玩家胜利！"
					response.OK(g, msg)
				}
			}
		}
	}()
	row1 := g.GetInt("row1")
	col1 := g.GetInt("col1")
	one := [2]int{row1, col1}
	row2 := g.GetInt("row2")
	col2 := g.GetInt("col2")
	two := [2]int{row2, col2}
	chessOne <- one
	chessTwo <- two
} //利用go协程与select轮询下棋，若时间超出45未落棋则自动判负
