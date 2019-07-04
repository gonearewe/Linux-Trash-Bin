package main

import (
	. "del/src/mypackage"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	flag_list, flag_list_detail, flag_help, flag_config bool
	trash_name                                          string
)

func init() {
	flag.BoolVar(&flag_list, "l", false, "list stuff in the trash bin")
	flag.BoolVar(&flag_list_detail, "ll", false, "list stuff in the trash bin in details")
	flag.BoolVar(&flag_help, "h", false, "show usage")
	flag.BoolVar(&flag_config, "c", false, "configure the config file")

	flag.StringVar(&trash_name, "d", "", "specify the trash (file or dictionary) name")

	// 改变默认的 Usage

	flag.Usage = usage

}
func main() {

	flag.Parse()
	if flag_help == true {
		usage()
	} else if flag_config == true {
		cmd := exec.Command("code", "-r", "/opt/myprogram/etc/del_config.json")
		_, err := cmd.CombinedOutput()
		ErrorCheck(err, "Unable to Open Config File with VS Code")

	} else if flag_list == true {
		listTrashBin(false)
	} else if flag_list_detail == true {
		listTrashBin(true)
	} else if trash_name != "" {
		if AddrExists(trash_name) {

			var conf config //instead of conf:=config{} !!!!!
			conf.readConfigFile()

			if AddrExists(conf.Trash_bin_path+"/TrashBin") == false {
				os.MkdirAll(conf.Trash_bin_path+"/TrashBin", 0755)
				fmt.Println("TrashBin Doesn't Exist and Was Created Automatically")
			}

			cmd := exec.Command("mv", trash_name, conf.Trash_bin_path+"/TrashBin")
			_, err := cmd.CombinedOutput()
			ErrorCheck(err, "Unable to Move the Trash to the Trash Bin")
			return
		} else {
			fmt.Println("Trash Doesn't Exist or Inaccess !!!")
			fmt.Println("Please Check : " + trash_name)
			return
		}
	}

}
func usage() {
	fmt.Fprintf(os.Stderr, `del version: 1.0
Usage: del [-hlc] [-d trash name] [-c ]

Options:
`)
	flag.PrintDefaults()
}

type config struct {
	Trash_bin_path string //instead of trash_bin_path,which is demanded by json.Unmarshal()
}

func (conf *config) readConfigFile() { //pass *config type so that you can actually modify it
	if contents, err := ioutil.ReadFile("/opt/myprogram/etc/del_config.json"); err != nil {
		fmt.Println("Unable to Open Config File")
		os.Exit(0)
	} else {
		err := json.Unmarshal(contents, conf)
		ErrorCheck(err, "Unable to Parse Config File,Please Check the Format")

	}
}

func listTrashBin(isDetail bool) {
	var conf config //instead of conf:=config{} !!!!!
	conf.readConfigFile()

	if AddrExists(conf.Trash_bin_path+"/TrashBin") == false {
		fmt.Println("TrashBin Doesn't Exist")
	} else {
		var cmd *exec.Cmd //instead of exec.*Cmd or just *Cmd
		if isDetail {
			cmd = exec.Command("ls", "-l", conf.Trash_bin_path+"/TrashBin") //you can't define "cmd" here,or
		} else {
			cmd = exec.Command("ls", conf.Trash_bin_path+"/TrashBin") //it will only work in this scope
		}
		out, err := cmd.CombinedOutput() //by default,no output,you need to print by yourself
		ErrorCheck(err, "Unable to Read the Trash Bin")
		fmt.Println(string(out)) //string()is needed
		return

	}
}

// contents, err := ioutil.ReadFile("/opt/myprogram/etc/del_config.json")
// var inter interface{}
// json.Unmarshal(contents,&inter)
// keymap, _ := inter.(map[string]interface{})
// for key,keyval:=range keymap{
// 	fmt.Println(key,keyval)
// }

// func  readConfigFile() *map[string] {
// 	if contents, err := ioutil.ReadFile("/opt/myprogram/etc/del_config.json"); err != nil {
// 		fmt.Println("Unable to Open Config File")
// 		os.Exit(1)
// 	} else {
// 		fmt.Println(string(contents))
// 		err := json.Unmarshal(contents, conf)
// 		fmt.Println(conf.Trash_bin_path)
// 		if nil != err {
// 			fmt.Println("Unable to Parse Config File,Please Check the Format")
// 			os.Exit(1)
// 		}
// 	}
// }
