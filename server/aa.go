package server

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html"
	"net/http"
	"os/exec"
	"strings"
	"webDemo/common"
)

type Info struct {
	addr string
	host string
}

func SayHello(c *gin.Context) {

	//requestDate := make(map[string]string)
	//c.ShouldBindJSON(&requestDate)
	//s := `
	//  _______   ______    _______         ______   _____  ____
	// /       | /      \  /       |______ /      \ /     \/    \
	///$$$$$$$/ /$$$$$$  |/$$$$$$$//      |$$$$$$  |$$$$$$ $$$$  |
	//$$      \ $$    $$ |$$ |     $$$$$$/ /    $$ |$$ | $$ | $$ |
	// $$$$$$  |$$$$$$$$/ $$ \_____       /$$$$$$$ |$$ | $$ | $$ |
	///     $$/ $$       |$$       |      $$    $$ |$$ | $$ | $$ |
	//$$$$$$$/   $$$$$$$/  $$$$$$$/        $$$$$$$/ $$/  $$/  $$/
	//`
	//
	//c.String(http.StatusOK, s)

	value, _ := c.Get("aaa")
	fmt.Println(value)
	c.String(http.StatusOK, "string(value)")
}

func SqlI1(c *gin.Context) {
	address := c.Query("address")
	sql := "select * from order where address = " + address
	exec := common.DB.Raw(sql)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI2(c *gin.Context) {
	address := c.Query("address")
	sqlMap := make(map[string]string)
	sqlMap["addr"] = address
	sql := "select * from order where address = " + sqlMap["addr"]
	exec := common.DB.Raw(sql)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI3(c *gin.Context) {
	address := c.Query("address")
	sqlMap := make(map[string]string)
	sqlMap["addr"] = address
	sql := "select * from order where address = " + sqlMap["addr1"]
	exec := common.DB.Raw(sql)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI4(c *gin.Context) {
	address := c.Query("address")
	sqlMap := make(map[string]string)
	sqlMap["addr"] = address
	sqlMap["addr"] = "address"
	sql := "select * from order where address = " + sqlMap["addr"]
	exec := common.DB.Raw(sql)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI5(c *gin.Context) {
	address := c.Query("address")
	info := Info{
		addr: address,
		host: "test",
	}
	sql := "select * from order where address = " + info.addr
	exec := common.DB.Raw(sql)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI6(c *gin.Context) {
	address := c.Query("address")
	info := Info{
		addr: address,
		host: "test",
	}
	sql := "select * from order where address = " + info.host
	exec := common.DB.Raw(sql)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI7(c *gin.Context) {
	address := c.Query("address")
	info := Info{
		addr: address,
		host: "test",
	}
	info.addr = "addr"
	sql := "select * from order where address = " + info.addr
	exec := common.DB.Raw(sql)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI8(c *gin.Context) {
	address := c.Query("address")
	if "aaa" == "aaa" {
		sql := "select * from order where address = " + address
		exec := common.DB.Raw(sql)
		rows, err := exec.Rows()
		fmt.Println(rows)
		fmt.Println(err)
	}
}

func SqlI9(c *gin.Context) {
	value_rce := c.Request.FormValue("host")

	if value_rce != "" {
		host := html.EscapeString(value_rce)
		out, _ := exec.Command("bash", "-c", "ping -c 5 "+host).CombinedOutput()
		fmt.Println(out)
		//fmt.Println(strings.Replace(string(out), "\n", "\\n", -1))
		//結果をjsonにして送信
		//outは[]byteなのでキャストしてからエスケープしてまたキャスト
		//msg := []byte(fmt.Sprintf(`{"msg": "%s"}`, strings.Replace(string(out), "\n", "\\n", -1)))
		//
		//w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		//w.WriteHeader(http.StatusOK)
		//w.Write(msg)
	}
}

func SqlI10(c *gin.Context) {
	value_rce := c.Request.FormValue("host")

	if value_rce != "" {
		out, _ := exec.Command("ping", "-c 5", value_rce).CombinedOutput()
		fmt.Println(out)
	}
}

func SqlI11(c *gin.Context) {
	address := c.Query("address")
	var b1 strings.Builder
	b1.WriteString("select * from order where address = ")
	b1.WriteString(address)
	exec := common.DB.Raw(b1.String())
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI12(c *gin.Context) {
	address := c.Query("address")
	ss := fmt.Sprintf("%s%s", "select * from order where address = ", address)
	exec := common.DB.Raw(ss)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI13(c *gin.Context) {
	address := c.Query("address")
	var str []string = []string{"select * from order where address = ", address}
	s3 := strings.Join(str, "")
	exec := common.DB.Raw(s3)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI14(c *gin.Context) {
	address := c.Query("address")

	var bt bytes.Buffer
	bt.WriteString("select * from order where address = ")
	bt.WriteString(address)
	//获得拼接后的字符串
	s3 := bt.String()

	exec := common.DB.Raw(s3)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI15(c *gin.Context) {
	address := c.Query("address")
	var b1 strings.Builder
	b1.WriteString("select * from order where address = ")
	b1.WriteString(address)
	b1.Reset()
	exec := common.DB.Raw(b1.String())
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI16(c *gin.Context) {
	address := c.Query("address")

	var bt bytes.Buffer
	bt.WriteString("select * from order where address = ")
	bt.WriteString(address)
	//获得拼接后的字符串
	bt.Reset()
	s3 := bt.String()

	exec := common.DB.Raw(s3)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}

func SqlI17(c *gin.Context) {
	address := c.Query("address")

	var bt bytes.Buffer
	bt.WriteString("select * from order where address = ")
	bt.WriteString(address)
	//获得拼接后的字符串
	bt.Reset()
	s3 := bt.String()

	exec := common.DB.Raw(s3)
	rows, err := exec.Rows()
	fmt.Println(rows)
	fmt.Println(err)
}
