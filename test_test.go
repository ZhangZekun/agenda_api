package agenda

import (
	"testing"
	"errors"
)

func Division(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为0")
    }

    return a / b, nil
}

func Test_Division_1(t *testing.T) {
    if i, e := Division(6, 2); i != 3 || e != nil {        //测试函数
        t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
    } else {
        t.Log("第次测试通过了") //记录一些你期望记录的信息
    }
}

func Test_Division_2(t *testing.T) {
    t.Error("就是不通过")
}