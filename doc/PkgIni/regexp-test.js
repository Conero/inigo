/**
 * 2017年11月4日 星期六
 * Joshua Conero
 * 字符处理正则测试
 */

// (1) 多行注释开始
// test = " 或 '
// test = 'xxxxxx  或 "xxxxxxx
(function(close){
    if(close) return false;
    // test = "dddddd
    // test = '
    // test = "
    var reg = /^[a-zA-Z0-9_-]+[=^"']*/;
    reg = /^[a-zA-Z0-9_-]*'?/;
    reg = /^[a-zA-Z0-9_-]+.*["']{1}.*/;
    reg = /['"]{1}/
    reg = /^[A-Za-z0-9_-]+[=\s]+("|').*[^"^']+$/;
    reg = /^[A-Za-z0-9_-]+[=\s]+("|'){1}[^"']*/;
    reg = /^[\w]+[=\s]+("|'){1}[^"']*$/;            // 符合开始符号
    //reg = /^[A-Za-z0-9_-]+[=\s]+("|'){1}.*/;
    // reg = /["']{1}/;
    var t = "";
    console.log(t = 'test = 1, 8, 9, 8, 5',reg.test(t));
    console.log(t = 'test = {1, 8, 9, 8, 5}',reg.test(t));
    console.log(t = 'test = "dddddd',reg.test(t));
    console.log(t = 'test = "',reg.test(t));
    console.log(t = `test = '`,reg.test(t));
    console.log(t = `test = "eee"`,reg.test(t));
    console.log(t = `test = 'eee'`,reg.test(t));
    console.log(t = `"eee"`,reg.test(t));
    console.log(t = `',`,reg.test(t));
    console.log(t = `',"`,reg.test(t));
    console.log(t = `','`,reg.test(t));
    console.log(t = `",`,reg.test(t));
})(true)

// (1.1) 多行注释开始  - 无键值， 用于 字符串切片处理
// " 或 '
// 'xxxxxx  或 "xxxxxxx
(function(close){
    if(close) return false;
    // test = "dddddd
    // test = '
    // test = "
    var reg = /^[a-zA-Z0-9_-]+[=^"']*/;
    reg = /["']{1}/;                // v1
    reg = /^["']{1}/;
    reg = /^["']{1}[^"'\,]+$/;      // 符合要求
    var t = "";
    console.log(t = `"`,reg.test(t));
    console.log(t = `'`,reg.test(t));
    console.log(t = 'test = 1, 8, 9, 8, 5',reg.test(t));
    console.log(t = 'format2 = "key1 _JC__EQUAL value1; 双引号"',reg.test(t));
    console.log(t = 'test = "dddddd',reg.test(t));
    console.log(t = 'test = "',reg.test(t));
    console.log(t = `test = '`,reg.test(t));
    console.log(t = `test = "eee"`,reg.test(t));
    console.log(t = `test = 'eee'`,reg.test(t));
    console.log(t = `"eee"`,reg.test(t));
    console.log(t = `',`,reg.test(t));
    console.log(t = `',"`,reg.test(t));
    console.log(t = `','`,reg.test(t));
    console.log(t = `",`,reg.test(t));
    console.log(t = `"eyehhdd //\\ddddd `,reg.test(t));
    console.log(t = `'字符串处理eeeee `,reg.test(t));
})(true)


// (2) 多行注释结束

(function(close){
    if(close) return false;
    // test = "dddddd
    // test = '
    // test = "
    var reg = /^[a-zA-Z0-9_-]+[=^"']*/;
    //reg = /^[A-Za-z0-9_-]+[=\s]+("|'){1}.*/;
    // reg = /["']{1}/;
    reg = /[^"']*['"]{1}$/;
    reg = /[^"']{0,}['"]{1}$/;
    reg = /^[^"'=]*['"]{1}$/;           // 未支持 , 分隔符号 其他匹配
    reg = /^[^"'=]*['"\,]+$/;
    var t = "";
    console.log(t = `字段结尾wee'`,reg.test(t));
    console.log(t = `字段结尾~~8555//wee"`,reg.test(t));
    console.log(t = `"`,reg.test(t));
    console.log(t = `'`,reg.test(t));
    console.log(t = 'test = "',reg.test(t));
    console.log(t = `test = 'eee'`,reg.test(t));
    console.log(t = `test = "eee"`,reg.test(t));
    console.log(t = `"eee"`,reg.test(t));
    console.log(t = `',`,reg.test(t));
    console.log(t = `',"`,reg.test(t));
    console.log(t = `','`,reg.test(t));
    console.log(t = `",`,reg.test(t));
})(true);