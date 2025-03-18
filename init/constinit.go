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

	var envVars model.Dbconfig
	// 解析 JSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&envVars); err != nil {
		return fmt.Errorf("❌ 解析 JSON 失敗: %v", err)
	}

	constVar.DBConst = envVars
	fmt.Println("reading config data successfully from file ...")

	return nil
}