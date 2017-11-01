/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月31日 星期二
 * @ini 基枝分析模型   Base Branch Analyze(BBAnalyze)
 */

package ini

// CBaseKeys => CBaseValues
// 基枝模型
type BBA struct {
	BaseKey     string                   // 当前指向的基键
	CBaseKeys   []string                 // 枝基键列
	CBaseValues []map[string]interface{} // 枝基键值
	BranchQueue map[string]interface{}   // 分支队列
	DataQueue   map[string]interface{}   // 数据队列值
}

//生成基枝分析模型实例
func BBAnalyze() *BBA {
	bba := &BBA{
		BaseKey:     "",
		CBaseKeys:   []string{},
		CBaseValues: []map[string]interface{}{},
		BranchQueue: map[string]interface{}{},
		DataQueue:   map[string]interface{}{},
	}
	return bba
}

// 更新基键值
func (bba *BBA) UpdateBaseKey(bKey string) {
	bba.BaseKey = bKey
	bba.CBaseKeys = append(bba.CBaseKeys, bKey)
	bba.CBaseValues = append(bba.CBaseValues, map[string]interface{}{})
}

// 推值送到分支列队
func (bba *BBA) PushQueue(key string, value interface{}) *BBA {
	if bba.BaseKey == "" {
		bba.DataQueue[key] = value
	} else {
		bkLen := len(bba.CBaseKeys)
		bvLen := len(bba.CBaseValues)
		// CBaseKeys -> CBaseValues 长度一一对应时加入到当前的子节点
		if bkLen == bvLen {
			tmpV := bba.CBaseValues[bvLen-1]
			tmpV[key] = value
			bba.CBaseValues[bvLen-1] = tmpV
		}
	}
	return bba
}

// 提交数据到外表 queue
// 当前的基枝遍历完成，去基枝
func (bba *BBA) CommitQueue() bool {
	isSuccess := false
	bkLen := len(bba.CBaseKeys)
	// 存在多级时
	if bkLen > 1 {
		branchBv := bba.CBaseValues[bkLen-1]
		topBv := bba.CBaseValues[bkLen-2]
		tKey := bba.CBaseKeys[bkLen-1]
		topBv[tKey] = branchBv
		bba.CBaseValues = bba.CBaseValues[0 : bkLen-2]
		bba.BaseKey = bba.CBaseKeys[bkLen-2]
		bba.CBaseKeys = bba.CBaseKeys[0 : bkLen-2]
		isSuccess = true

	} else if bkLen == 1 {
		bba.BaseKey = ""
		bba.DataQueue[bba.CBaseKeys[0]] = bba.CBaseValues[0]
	}
	return isSuccess
}
