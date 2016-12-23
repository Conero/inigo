// mysql
package main

import (
	"fmt"
	"log"
)

import (
	"database/sql"
	//https://github.com/go-sql-driver/mysql/  参考地址
	_ "github.com/go-sql-driver/mysql"
)
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type mydb struct {
	db *sql.DB
}

// go struct 参考  http://blog.csdn.net/chuangrain/article/details/9335041
func (mw *MyMainWindow) mysql_Triggered() {
	mysqlMw := new(MyMainWindow)
	var sqlIpt, sqlOpt *walk.TextEdit
	var sqlOptTv *walk.TableView

	//dbset := new(DbSetting)
	//  conero:151009@/cro?charset=utf8
	//数据表单初始化
	dbset := &DbSetting{
		DbType:     "mysql",
		DbUser:     "conero",
		DbPassword: "151009",
		DbName:     "cro",
		DbChatset:  "urtf",
	}
	thisdb := new(mydb)

	(MainWindow{
		AssignTo: &mysqlMw.MainWindow,
		Title:    "MySQL控制台",
		MinSize:  Size{600, 460},
		Layout:   VBox{}, //开启VBox时图层才可见
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				Action{
					Text: "设置数据库",
					OnTriggered: func() {
						if cmd, err := RunAnimalDialog(mysqlMw, dbset); err != nil {
							sqlOpt.AppendText("/r/n" + err.Error())
						} else if cmd == walk.DlgCmdOK {
							sqlOpt.AppendText(fmt.Sprintf("%+v", dbset))
							//conero:151009@/cro?charset=utf8
							dirver := dbset.DbUser + ":" + dbset.DbPassword + "@/" + dbset.DbName + "?charset" + dbset.DbChatset
							db, err := sql.Open(dbset.DbType, dirver)
							thisdb.db = db
							if err != nil {
								sqlOpt.AppendText("/r/n" + dirver)
								sqlOpt.AppendText("/r/n" + err.Error())
							}
						}
					},
				},
			},
		},
		Children: []Widget{
			Composite{
				Layout: VBox{}, //开启VBox时图层才可见 自定义Composite也会出错，需要放到 Children下
				Children: []Widget{
					TextEdit{
						AssignTo: &sqlIpt,
						OnKeyDown: func(key walk.Key) { //sql语句输入框
							if key == walk.KeyReturn {
								sqlStr := sqlIpt.Text()
								sqlOpt.AppendText(sqlStr)
								db := thisdb.db
								rows, err := db.Query(sqlStr)
								if err != nil {
									sqlOpt.AppendText(err.Error())
								} else {
									fmt.Println(rows)
									col, err := rows.Columns()
									if err != nil {
										sqlOpt.AppendText("\r\n" + err.Error())
									} else {
										ret := ""
										for i := 0; i < len(col); i++ {
											ret += " " + col[i]
										}
										sqlOpt.AppendText("\r\n" + ret)
									}
								}
								sqlOptTv.Columns()
							}
						},
					},
				},
			},
			Composite{
				Layout: VBox{}, //开启VBox时图层才可见
				Children: []Widget{
					TableView{
						AssignTo: &sqlOptTv,
					},
					TextEdit{
						AssignTo: &sqlOpt,
					},
				},
			},
		},
	}).Run()
	//sqltest()
}

func sqltest() {
	db, err := sql.Open("mysql", "conero:151009@/cro?charset=utf8") //db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(db)
		rows, err := db.Query("SELECT finc_no FROM finc_set")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rows.Columns())
			defer rows.Close()
			for rows.Next() {
				var key string
				if err := rows.Scan(&key); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(key)
				}
			}
		}

	}
}

//"mysql", "conero:151009@/cro?charset=utf8"
type DbSetting struct {
	DbType     string
	DbUser     string
	DbPassword string
	DbName     string
	DbChatset  string
}

func RunAnimalDialog(owner walk.Form, dbsetting *DbSetting) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var ep walk.ErrorPresenter
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Animal Details",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			DataSource:     dbsetting,
			ErrorPresenter: ErrorPresenterRef{&ep},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "数据库类型",
					},
					LineEdit{
						Text: Bind("DbType"),
					},

					Label{
						Text: "数据库用户",
					},
					LineEdit{
						Text: Bind("DbUser"),
					},

					Label{
						Text: "数据库密码",
					},
					LineEdit{
						Text: Bind("DbPassword"),
					},

					Label{
						Text: "数据库名称",
					},
					LineEdit{
						Text: Bind("DbName"),
					},

					Label{
						Text: "数据库字符集",
					},
					LineEdit{
						Text: Bind("DbChatset"),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}

func mysql_main_guaqi() {
	fmt.Println("Hello World!")
}
