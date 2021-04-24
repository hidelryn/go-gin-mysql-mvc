package libs

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 공용 - 컨트롤러에 넘어온 파라미터를 해당 arg 구조체에 파싱, 값이 존재 하지 않으면 실패 응답
// arg - 파싱할 데이터를 담을 구조체
func RequestParams(c *gin.Context, arg interface{}) error {
	if err := c.ShouldBindJSON(arg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": err.Error(),
		})
		return err
	}
	return nil
}

// 공용 - JSON으로 응답 하는 함수
// err - 에러 매개 변수
// datas - 응답할 구조체 데이터들
func ResponseJSON(c *gin.Context, err error, datas ...interface{}) {
	code := 200
	message := "SUCCESS"
	results := map[string]interface{}{}
	if err != nil {
		code = 500
		message = err.Error()
	} else {
		for _, data := range datas {
			t := reflect.TypeOf(data)
			currentModel := t.Elem().String() // Model의 이름을 가져 온다 EX) *Models.User
			modelName := strings.Split(currentModel, ".")[1]
			results[modelName] = data
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"results": results,
	})
}

// 타임 관련 리턴 함수
// 첫 리턴은 ms로 된 타임스탬프
// 두번째 리턴은 yyyymmdd
func GetTime() (uint64, uint64) {
	init := time.Now().UTC()
	ts := init.UnixNano() / int64(time.Millisecond)
	yyyymmdd, _ := strconv.ParseInt(init.Format("20060102"), 10, 64)
	return uint64(ts), uint64(yyyymmdd)
}

// uint32로 구성된 고유 아이디 생성 - 구글에서 만든 패키지 사용
func GetUUID() uint32 {
	return uuid.New().ID()
}
