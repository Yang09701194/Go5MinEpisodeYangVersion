package TestSqlDatabaseMssql

import (
	"flag"
	"fmt"
	"time"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB = nil

func GetDB() (*sql.DB, error) {

	if db != nil {
		return db, nil
	}

	var (
		userid   = flag.String("U", "sa", "login_id")
		password = flag.String("P", "1111", "password")
		server   = flag.String("S", "MSI\\SQL2K14", "server_name[\\instance_name]")
		database = flag.String("d", "FET2017", "db_name")
	)
	flag.Parse()

	dsn := "server=" + *server + ";user id=" + *userid + ";password=" + *password + ";database=" + *database
	//這句open還是印象深刻   帶領我進介面大門的鑰匙!!!
	var err error
	db, err = sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return new(sql.DB), err
	}

	//db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return new(sql.DB), err
	}

	return db, nil

	//這裡作用先定為   把db初始完成後回傳
	//db回傳後  再執行底下export版本的Exex

	//defer db.Close()
	// r := bufio.NewReader(os.Stdin)
	// for {
	// 	_, err = os.Stdout.Write([]byte("> "))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	cmd, err := r.ReadString('\n')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println()
	// 			return
	// 		}
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	err = exec(db, cmd)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

}

func ExecSql(db *sql.DB, cmd string) (string, error) {

	rows, err := db.Query(cmd)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return "", err
	}
	if cols == nil {
		return "", nil
	}

	var exeResult = ""

	var vals = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			exeResult += "\t"
			//fmt.Print("\t")
		}
		exeResult += cols[i]
		//fmt.Print(cols[i])
	}
	exeResult += "\r\n"
	//fmt.Println()
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			exeResult += err.Error()
			//fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				exeResult += "\t"
				//fmt.Print("\t")
			}
			exeResult += GetSqlValue(vals[i].(*interface{}))
			//getSqlValue(vals[i].(*interface{}))
		}
		exeResult += "\r\n"
		//fmt.Println()

	}
	if rows.Err() != nil {
		return "", rows.Err()
	}
	return exeResult, nil
}

func GetSqlValue(pval *interface{}) string {

	switch v := (*pval).(type) {
	case nil:
		return "NULL"
		//fmt.Print("NULL")
	case bool:
		if v {
			return "1"
			//fmt.Print("1")
		} else {
			return "0"
			//fmt.Print("0")
		}
	case []byte:
		return string(v)
		//fmt.Print(string(v))
	case time.Time:
		return v.Format("2006-01-02 15:04:05.999")
		//fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		return fmt.Sprint(v)
		//fmt.Print(v)
	}
}
