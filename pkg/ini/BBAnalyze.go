/* @ini-go V1.x
 * @Joshua Conero
 * @2017年10月31日 星期二
 * @ini 基枝分析模型   Base Branch Analyze(BBAnalyze)
 */

package ini

type BBA struct {
	BaseKey  string                 // 当前指向的基键
	CBaseKeys   []string               // 枝基键列
	BranchQueue map[string]interface{} // 分支队列
	DataQueue	map[string]interface{} // 数据队列值
}

//生成基枝分析模型实例
func BBAnalyze() *BBA {
	bba := &BBA{
		BaseKey:     "",
		CBaseKeys:   []string{},
		BranchQueue: map[string]interface{}{},
		DataQueue: map[string]interface{}{},
	}
	return bba
}

// 更新基键值
func (bba *BBA) UpdateBaseKey(bKey string){
	bba.BaseKey = bKey
}

// 推送到分支列队
func (bba *BBA) PushQueue(key string, value interface{}) *BBA {
	if bba.BaseKey == ""{
		bba.DataQueue[key] = value
	}else{
		bba.BranchQueue[key] = value
	}
	return bba
}

// 提交数据到外表 queue
func (bba *BBA) CommitQueue(queue map[string]interface{}) map[string]interface{} {
	if bba.BaseKey != "" {
		if len(bba.BranchQueue) > 0 {
			queue[bba.BaseKey] = bba.BranchQueue
			bba.BranchQueue = map[string]interface{}{}
		}
		bba.BaseKey = ""
	}
	bba.CBaseKeys = []string{}
	return queue
}
