<?php
/**
 * 2017年11月6日 星期一
 * php 快速实现 bba 模型分析
 */

require_once __DIR__.'/bba.php';

function microtime_float()
{
    list($usec, $sec) = explode(" ", microtime());
    return ((float)$usec + (float)$sec);
}
echo "  The PHP Version ".phpversion().".\r\n";
// print_r($argv);
$name = isset($argv[1])? trim($argv[1]):false;
if($name){
    $time1 = microtime_float();
    $bba = new BBA($name);
    if($bba->msg){
        echo $bba->msg.'\n';
    }else{
        $tt = pathinfo($name);
        file_put_contents(__DIR__.'/'.$tt['filename'].'.json', $bba->JsonString());
        // echo '  用时'.(time() - $time1).'s， 遍历行数：'.($bba->lineCount).'.\n';
        echo '  spend times '.(microtime_float() - $time1).'s, count lines: '.($bba->lineCount).".\n";
    }
}
