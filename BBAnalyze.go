/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月31日 星期二
 * @ini 基枝分析模型   Base Branch Analyze(BBAnalyze)
 */

package inigo

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
	MlValue           interface{}              // 多行抽象类型
	Ini               *Ini                     // Ini 对象
}

//生成基枝分析模型实例
func BBAnalyze(I *Ini) *BBA {
	bba := &BBA{
		DataQueue:         map[string]interface{}{},
		BRCIdx:            -1,
		BranchRunning:     []map[string]interface{}{},
		BranchRunningKeys: []string{},
		Ini:               I,
	}
	return bba
}
// 获取系统当前的键值以及配置数据
func (bba *BBA) getCurrentKeyMap() (string, map[string]interface{}) {
	key := ""
	settingMap := map[string]interface{}{}
	if bba.BRCIdx > -1{
		key = bba.BranchRunningKeys[bba.BRCIdx]
		settingMap = bba.BranchRunning[bba.BRCIdx]
	}
	return key, settingMap
}
// 更新基键值， 可能为map/array 键值
func (bba *BBA) UpdateBaseKey(bKey string) {
	// setting map
	sMap := map[string]interface{}{
		"isInit": false,
	}
	// 行分支不存在缓存时，先环境准备
	if -1 == bba.BRCIdx {
		// 分支
		bba.BranchRunning = []map[string]interface{}{sMap}
	} else { // 已经存在值时
		bba.BranchRunning = append(bba.BranchRunning, sMap)
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
	if -1 == bba.BRCIdx { // 推送数据到顶层数据中心
		bba.DataQueue[key] = value
		// 同步初始化处理
		bba.BRCIdx = -1
		bba.BranchRunningKeys = []string{}
		bba.BranchRunning = []map[string]interface{}{}
		//} else {
	} else if bba.BRCIdx > -1 {
		isUpdateMk := false
		sKey, sMap := bba.getCurrentKeyMap()
		if OnlyTestPutOut {
			//println(bba.BRCIdx, "<&>", bba.BranchRunning, bba.BranchRunningKeys)
			fmt.Println("2)  2a.", bba.BRCIdx, sKey, "<&>(推送值到数据中心/map)", value, key)
			fmt.Println("       2b.", bba.BranchRunning, bba.BranchRunningKeys)
		}
		//tmpMap := bba.BranchRunning[bba.BRCIdx]
		if OnlyTestPutOut {
			fmt.Println("       2c.", sMap, bba.BranchRunning, bba.BRCIdx, bba.BranchRunningKeys, ".------------>2", bba.Ini.File.GetLine(), value, key)
		}
		if sMap["isInit"] == false { // 如果还没有初始化并且，首先是map时
			sMap["map"] = map[string]interface{}{
				key: value,
			}
			sMap["type"] = "MAP"
			sMap["isInit"] = true
			isUpdateMk = true
		} else if sMap["isInit"] == true { // 已经初始化
			tMValue := map[string]interface{}{}
			if sMap["type"] == "MAP"{ // 检测到的全部为 map
				tMValue = sMap["map"].(map[string]interface{})
				tMValue[key] = value
				sMap["map"] = tMValue
			}else if sMap["type"] == "ARRAY" || sMap["type"] == "BOTH"{ // 检测到的为 array -> both
				tMValue = map[string]interface{}{
					key: value,
				}
				tMArray := []interface{}{}
				if "ARRAY" == sMap["type"]{
					tMArray = sMap["array"].([]interface{})
					sMap["type"] = "BOTH"
				}
				tMArray = append(tMArray, tMValue)
				sMap["array"] = tMValue
			}
			isUpdateMk = true
		}
		if isUpdateMk {
			bba.BranchRunning[bba.BRCIdx] = sMap
		}
	}
	return bba
}

// 分行数组值推送
func (bba *BBA) MultiLineToArray(value string) *BBA {
	// string,	-> string
	//value = strings.TrimSpace(value)
	if strings.LastIndex(value, IniParseSettings["limiter"]) == 0 {
		value = value[0 : len(value)-2]
	}
	if  bba.BRCIdx > -1 {
		isUpdateMk := false
		_ , sMap := bba.getCurrentKeyMap()
		if sMap["isInit"] == false { // 如果还没有初始化并且，首先是array时
			sMap["array"] = []interface{}{value}
			sMap["type"] = "ARRAY"
			sMap["isInit"] = true
			isUpdateMk = true
		} else if sMap["isInit"] == true { // 已经初始化
			//fmt.Println(sMap)
			tArrayValue := sMap["array"].([]interface{})
			if sMap["type"] == "ARRAY" || sMap["type"] == "BOTH" { // 检测到的全部为 array
				tArrayValue = append(tArrayValue, value)
			} else if sMap["type"] == "MAP" {
				sMap["type"] = "BOTH"
				// 将历史的map中的值推送带array 中
				for mKey, mValue := range sMap["map"].(map[string]interface{}){
					tArrayValue = append(tArrayValue, map[string]interface{}{
						mKey: mValue,
					})
				}
				tArrayValue = append(tArrayValue, value)
			}
			sMap["array"] = tArrayValue
			isUpdateMk = true
		}
		if isUpdateMk {
			bba.BranchRunning[bba.BRCIdx] = sMap
		}
	}
	return bba
}

// 跨行字符串解析
func (bba *BBA) MultiLineString(line string)  {
	_ , sMap := bba.getCurrentKeyMap()
	// 非字符串时跳过
	if sMap["type"] != "STRING"{
		return
	}
	if !sMap["isInit"].(bool){
		sMap["string"] = line
		sMap["isInit"] = true
		sMap["type"] = "STRING"
	}else {
		sMap["string"] = sMap["string"].(string) + line
	}
	bba.BranchRunning[bba.BRCIdx] = sMap
}

// 提交数据到外表 queue
// 当前的基枝遍历完成，去基枝
func (bba *BBA) CommitQueue() bool {
	isSuccess := false
	if bba.BRCIdx > -1 {
		key , sMap := bba.getCurrentKeyMap()
		//BRCLen := len(bba.BranchRunning) - 1
		//tmpMap := bba.BranchRunning[bba.BRCIdx]
		var value interface{}
		value = ""
		if sMap["isInit"] == false { // 空数组
			value = []string{}
		} else if sMap["isInit"] == true {
			switch sMap["type"] {
			case "ARRAY": // 纯数组
				value = sMap["array"]
			case "MAP": // 数组 map 类型
				value = sMap["map"]
			case "BOTH":
				value = sMap["array"]
				value = append(value.([]interface{}), sMap["map"])
			}
		}
		if bba.BRCIdx > 0 {
			if OnlyTestPutOut {
				fmt.Println("       3.a.a", bba.BranchRunningKeys, bba.BranchRunning, bba.DataQueue, "{---->3.0a")
			}
			bba.BranchRunning = bba.BranchRunning[0:bba.BRCIdx]
			bba.BranchRunningKeys = bba.BranchRunningKeys[0:bba.BRCIdx]
			if OnlyTestPutOut {
				fmt.Println("       3.a.b-"+bba.Ini.File.GetLineString(), bba.BranchRunningKeys, bba.BranchRunning, bba.DataQueue, "{---->3.0a")
			}
		} else { // 顶级时
			bba.BranchRunning = []map[string]interface{}{}
			bba.BranchRunningKeys = []string{}
		}
		//bba.BRCIdx = bba.BRCIdx - 1
		//println(bba.BRCkey, key, value, "{->", tmpMap, bba.BranchRunning, BRCLen)
		bba.BRCIdx = bba.BRCIdx - 1
		if OnlyTestPutOut {
			fmt.Println("3).   ", bba.BRCIdx, bba.BRCIdx, key, value, "{-->3（剪枝处理）", bba.BRCIdx)
			fmt.Println("       3.a-"+bba.Ini.File.GetLineString(), sMap, bba.BranchRunning, bba.BranchRunningKeys)
		}
		bba.PushQueue(key, value)
		isSuccess = true
	}

	return isSuccess
}
