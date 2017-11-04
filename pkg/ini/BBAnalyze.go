/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月31日 星期二
 * @ini 基枝分析模型   Base Branch Analyze(BBAnalyze)
 */

package ini

import (
	"fmt"
	"strings"
)

const (
	//正式发布时删除当前的代码
	//OnlyTestPutOut = true
	OnlyTestPutOut = false
)

// CBaseKeys => CBaseValues
// 基枝模型
type BBA struct {
	DataQueue         map[string]interface{}   // 数据队列值
	BranchRunningKeys []string                 // 分支运行值 孪生键值缓存
	BranchRunning     []map[string]interface{} // 分支运行值
	BRCIdx            int                      // 分支当前的索引值
	BRCkey            string                   // 分支当前的map键
	MlValue           interface{}              // 多行抽象类型
	Ini               *Ini                     // Ini 对象
}

//生成基枝分析模型实例
func BBAnalyze(I *Ini) *BBA {
	bba := &BBA{
		DataQueue:         map[string]interface{}{},
		BRCIdx:            -1,
		BRCkey:            "",
		BranchRunning:     []map[string]interface{}{},
		BranchRunningKeys: []string{},
		Ini:               I,
	}
	return bba
}

// 更新基键值， 可能为map/array 键值
func (bba *BBA) UpdateBaseKey(bKey string) {
	bba.BRCkey = bKey
	tmpMap := map[string]interface{}{
		"isInit": false,
	}

	// 行分支不存在缓存时，先环境准备
	if -1 == bba.BRCIdx {
		// 分支
		bba.BranchRunning = []map[string]interface{}{tmpMap}
	} else { // 已经存在值时
		bba.BranchRunning = append(bba.BranchRunning, tmpMap)
	}
	bba.BRCIdx = bba.BRCIdx + 1
	bba.BranchRunningKeys = append(bba.BranchRunningKeys, bKey)
	if OnlyTestPutOut {
		fmt.Println("1). ", bba.BranchRunningKeys, bKey, bba.BRCIdx, "----->1.0P(解析出基枝)", bba.Ini.File.GetLine())
		fmt.Println("    1.a", bba.DataQueue)
	}
}

// 推值送到分支列队
func (bba *BBA) PushQueue(key string, value interface{}) *BBA {
	if bba.BRCkey == "" { // 推送数据到顶层数据中心
		bba.DataQueue[key] = value
		// 同步初始化处理
		bba.BRCIdx = -1
		bba.BranchRunningKeys = []string{}
		bba.BranchRunning = []map[string]interface{}{}
		//} else {
	} else if bba.BRCIdx > -1 {
		isUpdateMk := false
		if OnlyTestPutOut {
			//println(bba.BRCIdx, "<&>", bba.BranchRunning, bba.BranchRunningKeys)
			fmt.Println("2)  2a.", bba.BRCIdx, bba.BRCkey, "<&>(推送值到数据中心/map)", value, key)
			fmt.Println("       2b.", bba.BranchRunning, bba.BranchRunningKeys)
		}
		tmpMap := bba.BranchRunning[bba.BRCIdx]
		if OnlyTestPutOut {
			fmt.Println("       2c.", tmpMap, bba.BranchRunning, bba.BRCIdx, bba.BranchRunningKeys, ".------------>2", bba.Ini.File.GetLine(), value, key)
		}
		if tmpMap["isInit"] == false { // 如果还没有初始化并且，首先是map时
			tmpMap["map"] = map[string]interface{}{
				key: value,
			}
			tmpMap["type"] = "MAP"
			tmpMap["isInit"] = true
			isUpdateMk = true
		} else if tmpMap["isInit"] == true { // 已经初始化
			tMValue := map[string]interface{}{}
			if tmpMap["type"] == "MAP" || tmpMap["type"] == "BOTH" { // 检测到的全部为 map
				tMValue = tmpMap["map"].(map[string]interface{})
				tMValue[key] = value
			} else { // 检测到的为 array -> both
				tmpMap["type"] = "BOTH"
				tMValue = map[string]interface{}{
					key: value,
				}
			}
			//tMValue[key] = value
			tmpMap["map"] = tMValue
			isUpdateMk = true
		}
		if isUpdateMk {
			bba.BranchRunning[bba.BRCIdx] = tmpMap
		}
	}
	return bba
}

// 分行数组值推送
func (bba *BBA) MiltiLineToArray(value string) *BBA {
	// string,	-> string
	//value = strings.TrimSpace(value)
	if strings.LastIndex(value, IniParseSettings["limiter"]) == 0 {
		value = value[0 : len(value)-2]
	}
	if bba.BRCkey != "" {
		isUpdateMk := false
		tmpMap := bba.BranchRunning[bba.BRCIdx]
		if tmpMap["isInit"] == false { // 如果还没有初始化并且，首先是array时
			tmpMap["array"] = []interface{}{value}
			tmpMap["type"] = "ARRAY"
			tmpMap["isInit"] = true
			isUpdateMk = true
		} else if tmpMap["isInit"] == true { // 已经初始化
			tArrayValue := tmpMap["array"].([]interface{})
			if tmpMap["type"] == "ARRAY" || tmpMap["type"] == "BOTH" { // 检测到的全部为 array
				tArrayValue = append(tArrayValue, value)
			} else {
				tmpMap["type"] = "BOTH"
				tArrayValue = []interface{}{value}
			}
			tmpMap["array"] = tArrayValue
			isUpdateMk = true
		}
		if isUpdateMk {
			bba.BranchRunning[bba.BRCIdx] = tmpMap
		}
	}
	return bba
}

// 提交数据到外表 queue
// 当前的基枝遍历完成，去基枝
func (bba *BBA) CommitQueue() bool {
	isSuccess := false
	if bba.BRCIdx > -1 && bba.BRCkey != "" {
		key := bba.BRCkey
		BRCLen := len(bba.BranchRunning) - 1
		//BRCLen := len(bba.BranchRunningKeys) - 1
		tmpMap := bba.BranchRunning[bba.BRCIdx]
		var value interface{}
		value = ""
		if tmpMap["isInit"] == false { // 空数组
			value = []string{}
		} else if tmpMap["isInit"] == true {
			switch tmpMap["type"] {
			case "ARRAY": // 纯数组
				value = tmpMap["array"]
			case "MAP": // 数组 map 类型
				value = tmpMap["map"]
			case "BOTH":
				value = tmpMap["array"]
				value = append(value.([]interface{}), tmpMap["map"])
			}
		}
		//BRCLen = BRCLen - 1
		if BRCLen > 0 {
			if OnlyTestPutOut {
				fmt.Println("       3.a.a", bba.BranchRunningKeys, bba.BranchRunning, bba.DataQueue, "{---->3.0a")
			}

			//BRCLen = BRCLen - 1
			//tmpMap2 := bba.BranchRunning[BRCLen]
			// 此处可优化
			/*
				for k,_:= range tmpMap2{
					bba.BRCkey = k
					break
				}
			*/
			bba.BRCkey = bba.BranchRunningKeys[BRCLen]
			bba.BranchRunning = bba.BranchRunning[0:BRCLen]
			bba.BranchRunningKeys = bba.BranchRunningKeys[0:BRCLen]
			if OnlyTestPutOut {
				fmt.Println("       3.a.b-"+bba.Ini.File.GetLineString(), bba.BranchRunningKeys, bba.BranchRunning, bba.DataQueue, "{---->3.0a")
			}
		} else { // 顶级时
			bba.BranchRunning = []map[string]interface{}{}
			bba.BRCkey = ""
			bba.BranchRunningKeys = []string{}
		}
		//bba.BRCIdx = bba.BRCIdx - 1
		//println(bba.BRCkey, key, value, "{->", tmpMap, bba.BranchRunning, BRCLen)
		bba.BRCIdx = bba.BRCIdx - 1
		if OnlyTestPutOut {
			fmt.Println("3).   ", BRCLen, bba.BRCIdx, bba.BRCkey, key, value, "{-->3（剪枝处理）", BRCLen)
			fmt.Println("       3.a-"+bba.Ini.File.GetLineString(), tmpMap, bba.BranchRunning, bba.BranchRunningKeys)
		}
		bba.PushQueue(key, value)
		isSuccess = true
	}

	return isSuccess
}
