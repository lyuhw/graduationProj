package main

import (
	"database/sql"
	"fmt"
	"frontBackProject/controller"
	"frontBackProject/server"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type Request struct {
	Value string `json:"value"`
}

type Response struct {
	Data string `json:"data"`
}

var historyBlock server.Block
var historyChain []server.Block

func printAllBlock(c *gin.Context) {
	var result string
	for _, block := range historyChain {
		result += fmt.Sprintf("Index: %d\n", block.Index)
		result += fmt.Sprintf("Timestamp: %s\n", block.Timestamp)
		result += fmt.Sprintf("Data: %s\n", block.Data)
		result += fmt.Sprintf("PrevHash: %s\n", block.PrevHash)
		result += fmt.Sprintf("Hash: %s\n", block.Hash)
		result += "---------------------------------\n"
	}
	c.String(http.StatusOK, result)
}
func printBlock(b server.Block) {
	fmt.Printf("Index: %d\n", b.Index)
	fmt.Printf("Timestamp: %s\n", b.Timestamp)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("PrevHash: %s\n", b.PrevHash)
	fmt.Printf("Hash: %s\n", b.Hash)
	fmt.Println("-------------------------------")
}
func main() {

	// 连接MySQL数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/awpos")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	historyChain = append(historyChain, server.AddBlock("initChain", historyChain))
	historyBlock = server.CreateGenesisBlock(historyChain)
	// 执行SQL语句清空表
	_, err = db.Exec("TRUNCATE TABLE historyblock")
	if err != nil {
		panic(err.Error())
	}

	insertBlock(db, historyBlock)

	r := gin.Default()
	r.LoadHTMLGlob("frontend/*")

	r.GET("/select", func(c *gin.Context) {
		selectIndex := controller.SelectWalletSlice[len(controller.SelectWalletSlice)-1]
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":               "selectedWallet",
			"selectedWalletIndex": selectIndex})
	})

	r.POST("/getData", func(c *gin.Context) {
		var request Request
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 将输入的值包装成Response实例
		response := Response{request.Value}
		c.JSON(http.StatusOK, response)
	})

	//--------------------------------------
	r.POST("/result", func(c *gin.Context) {
		inputValue := c.PostForm("input")
		setThreshold, err := strconv.Atoi(inputValue)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid input")
			return
		}

		if setThreshold <= server.DataDetect() {
			c.String(http.StatusOK, "ATTACK!!!")
			//response := Response{Data: "ATTACK!!!"}
			//c.JSON(http.StatusOK, response)
			attackData := ""
			switch server.DataDetect() {
			case 1:
				attackData = server.BadForDao()
			case 10:
				attackData = server.CoinAgeAttack()
			case 11:
				attackData = server.CoinAgeAttack() + server.BadForDao()
			default:
				attackData = "Not a SetAttack!"
			}
			historyBlock = server.CreateBlock(attackData, historyChain)

			historyChain = append(historyChain, historyBlock)
			//historyChain = append(historyChain, server.AddBlock(attackData, historyChain))

			//server.PrintAllBlock(historyChain)
			printBlock(historyBlock)
			insertBlock(db, historyBlock)
		} else {
			c.String(http.StatusOK, "There NO Attack")
			//response := Response{Data: "There No Attack"}
			//c.JSON(http.StatusOK, response)
		}
	})

	//查看所有历史report
	r.GET("/blocks", printAllBlock)

	//设置404界面
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found2222")
	})
	r.Run(":8080")

}

func insertBlock(db *sql.DB, historyBlock server.Block) {
	// 构建插入语句
	query := "INSERT INTO historyblock(blockindex, timestamp, data, prevhash, hash) VALUES(?, ?, ?, ?, ?)"

	// 执行插入操作
	_, err := db.Exec(query, historyBlock.Index, historyBlock.Timestamp, historyBlock.Data, historyBlock.PrevHash, historyBlock.Hash)
	if err != nil {
		panic(err)
	}
	fmt.Println("-------------------------------")
	fmt.Println("History Block Insert Success!!!")
	fmt.Println("-------------------------------")
}
