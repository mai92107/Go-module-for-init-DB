package init

import (
	constVar "batchLog/const"
	"batchLog/model"
	"encoding/json"
	"fmt"
	"os"
)

// LoadEnvFromJSON 讀取 JSON 檔案並設置環境變數
func LoadEnvFromJSON(filename string) error {
	// 打開 JSON 檔案
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("❌ 無法開啟 JSON 檔案: %v", err)
	}
	defer file.Close()

	var machine model.Machine
	// 解析 JSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&machine); err != nil {
		return fmt.Errorf("❌ 解析 JSON 失敗: %v", err)
	}

	if machine.DB.Status{
		constVar.DBConst = machine.DB
		fmt.Println("reading config data successfully from file ...")
	}
	

	return nil
}