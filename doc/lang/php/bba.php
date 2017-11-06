<?php
/**
 * 2017年11月6日 星期一
 * php 快速实现 bba 模型分析
 */

class BBA{
    public $IsSuccess = false;
    public $lineCount = 0;  // 行统计
    public $msg;
    protected $file;
    protected $dataQueue = array(); // map
    protected $branchQueue = array();         // []map
    protected $branchKeys = array();          // []string
    protected $CBIdx = -1;                    // 索引    
    public function __construct($file){
        $this->file = $file;        
        if(is_file($file)){
            $this->IsSuccess = true;
            $this->parseFile();
        }else{
            $this->IsSuccess = false;
            $this->msg = "$file 文件不存在！";
        }
    }
    // 解析文件
    protected function parseFile(){
        if(!$this->IsSuccess) return;
        $fs = @fopen($this->file, 'r');
        $isComment = false;
        if($fs){
            while(($line = fgets($fs, 4096)) !== false){
                $this->lineCount += 1;
                $line = trim($line);
                // 空行
                if(empty($line)){continue;}
                // 注释
                if($isComment){
                    // 多行注释结束
                    if("'''" == $line){
                        $isComment = false;
                    }
                    continue;
                }
                // 多行注释开始
                if("'''" == $line){
                    $isComment = true;
                    continue;
                }                
                // 单行注释
                if(($fc = substr($line, 0, 1)) == ';' || $fc == '#'){
                    continue;
                }                
                $this->parseLine($line);
            }
            fclose($fs);
        }
    }
    // 获取当前的map配置以及键值
    protected function getCurrentMap(){
        $key = false;
        $map = false;
        if($this->CBIdx > -1){
            $key = $this->branchKeys[$this->CBIdx];
            $map = $this->branchQueue[$this->CBIdx];
        }
        return [$key, $map];
    }
    // 行解析
    protected function parseLine($line){
        // 等于符号解析, 有等于符号
        if(($idx = strpos($line, '='))>0){  //  key = any; key = { 或 {} ; key = tetss, yshhhshs, ; key = '/"; key = 'dddd'
            $key = trim(substr($line, 0, $idx));    // 键
            $value = trim(substr($line, $idx+1));   // 值
            // 作用域起始
            if('{' == $value){  // key = {
                $this->branchKeys[] = $key;
                $this->branchQueue[] = ['isInit'=> false];
                //print_r([$this->branchKeys, $this->branchQueue]);
                $this->CBIdx = count($this->branchKeys)-1;
                return false;
            }
            else if(preg_match('/^\{.*\}$/', $value)){   // key = {{{string}}}
                if(strpos($value, '=') === false){ // key = {dyhh,ff,f,,fffff,fff}
                    $value = explode(',', preg_replace('/\{|\}/', '', $value));
                }else{  // {{key = value}}
                    //
                }
            }else if(preg_match('/^["\']{1}.*["\']{1}$/', $value)){    // key = 'ddddd'
                $value = preg_replace('/"|\'/', '', $value);
            }else if(preg_match('/^["\']{1}[^"\']*/', $value)){         // key = '          字符串跨行
                $value = preg_replace('/"|\'/', '', $value);
                $this->branchKeys[] = $key;
                $this->branchQueue[] = ['isInit'=> true, 'type'=>'STRING', 'string'=> $value];
                $this->CBIdx = count($this->branchKeys)-1;
                return false;
            }
            /*
            else{   // key = string                   
            }
            */
            if(is_string($value) && substr_count($value, ',') > 0){
                $value = explode(',', $value);   // 数组       
            }
            $this->pushKeyValue($key, $value);
        }
        else if($line == '}'){  //  作用域结束
            if($this->CBIdx> -1){
                list($cKey, $map) = $this->getCurrentMap();
                //print_r([$cKey, $map]);
                if(isset($map['isInit']) && $map['isInit']){
                    $type = $map['type'];
                    $value = false;
                    if('ARRAY' == $type){
                        $value = $map['array'];
                    }else if('MAP' == $type){
                        $value = $map['map'];
                    }else if('BOTH' == $type){
                        $value = $map['both'];
                    }
                    array_pop($this->branchKeys);
                    array_pop($this->branchQueue);
                    $this->CBIdx = ($ctt = count($this->branchKeys))>0? $ctt-1:-1;
                    $this->pushKeyValue($cKey, $value);
                }
            } 
        }
        else if(preg_match('/^["\']*["\']$/', $line)){ // yetttehd" 或 "  字符串跨行
            $value = preg_replace('/"|\'/', '', $line);
            list($cKey, $map) = $this->getCurrentMap();
            $matched = false;
            if(isset($map['string'])){
                $value = $map['string']."\r\n".$value;
                array_pop($this->branchKeys);
                array_pop($this->branchQueue);
                $this->CBIdx = ($ctt = count($this->branchKeys))>0? $ctt-1:-1;
                $matched = true;
            }
            if($matched) $this->pushKeyValue($cKey, $value);
        }
        else{
            list($bKey, $bMap) = $this->getCurrentMap();
            if($bKey && $bMap){
                if($bMap['type'] == 'STRING'){  // 多行数组测试
                    $bMap['string'] .= "\r\n".$line;
                    $this->branchQueue[$this->CBIdx] = $bMap;
                }
            }
        }
    }
    // 推送值到 dataqueue
    protected function pushKeyValue($key, $value){
        if($this->CBIdx === -1){
            $this->dataQueue[$key] = $value;
        }else if($this->CBIdx > -1){
            //$cKey = $this->branchKeys[$this->CBIdx];
            $map = $this->branchQueue[$this->CBIdx];
            if($map['isInit'] === false){   // map
                $map['type'] = 'MAP';
                $map['map'] = [
                    $key => $value
                ];
                $map['isInit'] = true;
            }else{
                if($map['type'] == 'MAP'){
                    $map['map'][$key] = $value;
                    $map['isInit'] = true;
                }
                else if($map['type'] == 'BOTH'){
                    $map['both'][] = [
                        $key => $value
                    ];
                    $map['isInit'] = true;
                }
                else if($map['type'] == 'ARRAY'){
                    $map['both'] = $map['array'];
                    $map['type'] = 'BOTH';
                    $map['both'][] = [
                        $key => $value
                    ];
                    $map['isInit'] = true;
                }
            }
            $this->branchQueue[$this->CBIdx] = $map;
        } 
    }
    /**
     * 生成json字段字符串
     */
    public function JsonString(){
        return $this->IsSuccess? json_encode($this->dataQueue): false;
    }
}