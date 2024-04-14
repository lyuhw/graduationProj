package test

import (
	"encoding/json"
	"frontBackProject/server"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func init() {
	time.Sleep(20 * time.Second)
}

type Controller struct {
	SelectWalletSlice []int
}

func TestSelectWalletIndex(t *testing.T) {
	// 设置模拟数据
	controller := &Controller{
		SelectWalletSlice: []int{1, 2, 3},
	}

	// 发送GET请求
	resp, err := http.Get("http://localhost:8080/select")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Title               string `json:"title"`
		SelectedWalletIndex int    `json:"selectedWalletIndex"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// 断言结果
	expectedIndex := controller.SelectWalletSlice[len(controller.SelectWalletSlice)-1]
	if result.SelectedWalletIndex != expectedIndex {
		t.Errorf("Expected selected wallet index to be %d, but got %d", expectedIndex, result.SelectedWalletIndex)
	}
}

func TestGetData(t *testing.T) {
	// 设置模拟数据
	value := "test"

	// 发送POST请求
	resp, err := http.Post("http://localhost:8080/getData", "application/json", strings.NewReader(`{"value":"`+value+`"}`))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// 断言结果
	if result.Data != value {
		t.Errorf("Expected data to be %s, but got %s", value, result.Data)
	}
}

func TestHandleAttackData(t *testing.T) {
	// 设置模拟数据
	inputValue := "5"
	setThreshold, err := strconv.Atoi(inputValue)
	if err != nil {
		t.Fatalf("Failed to convert input value to integer: %v", err)
	}

	// 发送POST请求
	resp, err := http.PostForm("http://localhost:8080/result", url.Values{"input": {inputValue}})
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// 断言结果
	if setThreshold <= server.DataDetect() {
		expectedMessage := "ATTACK!!!"
		if result.Message != expectedMessage {
			t.Errorf("Expected message to be %s, but got %s", expectedMessage, result.Message)
		}
	} else {
		expectedMessage := "There NO Attack"
		if result.Message != expectedMessage {
			t.Errorf("Expected message to be %s, but got %s", expectedMessage, result.Message)
		}
	}
}

func TestViewAllBlocks(t *testing.T) {
	// 发送GET请求
	resp, err := http.Get("http://localhost:8080/blocks")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Blocks []server.Block `json:"blocks"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// 断言结果
	if len(result.Blocks) == 0 {
		t.Error("Expected to have at least one block in the response")
	}
}

func TestNotFoundPage(t *testing.T) {
	// 发送GET请求
	resp, err := http.Get("http://localhost:8080/notfound")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// 断言结果
	expectedMessage := "404 not found2222"
	if result.Message != expectedMessage {
		t.Errorf("Expected message to be %s, but got %s", expectedMessage, result.Message)
	}
}
