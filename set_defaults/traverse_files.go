package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	models "github.com/SpirentOrion/gemini/middleware/go_models/models"
)

var DefaultValue map[string]reflect.Value

var EmptyObjMap map[string]interface{}

func SetDefaults(obj interface{}) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return
	}
	v := reflect.ValueOf(obj).Elem()
	/*str := fmt.Sprintf("%s", v.Type())
	fmt.Printf("Type:%s\n", str)
	dType := strings.Split(str, ".")[1]
	//fmt.Printf("Default Obj:%v\n", DefaultValue[dType])
	d, ok := DefaultValue[dType]
	if !ok {
		return
	}*/
	//fmt.Println("Check -11")
	setDefaultValue(v)
}
func checkPrimitiveTypePtr(v reflect.Value) bool {

	t := v.Type()
	switch t.Kind() {
	case reflect.String:
		return true
	}
	return false
}
func setDefaultValue(v reflect.Value) {
	//fmt.Printf("\n=======>setDefaultValue - Start<=======================\n")
	if !(v.Kind() == reflect.Struct || v.Kind() == reflect.Ptr) {
		return
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	//t := v.Type()
	str := fmt.Sprintf("%s", v.Type())
	//fmt.Printf("Type:%s\n", str)
	dType := strings.Split(str, ".")[1]
	d, ok := DefaultValue[dType]
	if !ok {
		return
	}
	dv := reflect.Indirect(d)
	//dt := dv.Type()
	for i := 0; i < v.NumField(); i++ {
		rf := v.Field(i)
		df := dv.Field(i)
		//rf_name := t.Field(i).Name
		//df_name := dt.Field(i).Name
		switch rf.Kind() {
		/*case reflect.String:
		//	fmt.Printf("String Check -12: i:%d\n", i)
			defaultVal := df.String()
			if rf.String() == "" && defaultVal != "" {
				rf.Set(reflect.ValueOf(defaultVal).Convert(rf.Type()))
			//	fmt.Printf("%s\n", defaultVal)
			}*/
		case reflect.Ptr:
			//		fmt.Println("ptrrrrrrrrrrrrrrrrr")
			//		fmt.Printf("rf:%s  df:%s\n", rf_name, df_name)
			//checkIfPrimitDataTypePointer
			if checkPrimitiveTypePtr(rf.Elem()) {
				rf.Set(df)
				fmt.Printf("yes nil")
			} else {
				setDefaultValue(rf)
			}
		case reflect.Slice:
			//		 fmt.Printf("Struct Check -123: i:%d\n", i)
			for j := 0; j < rf.Len(); j++ {
				//fmt.Printf("Struct Check -123: j:%d\n", j)
				setDefaultValue(rf.Index(j))
			}
		case reflect.Struct:
			//fmt.Printf("Struct Check -12: i:%d\n", i)
			setDefaultValue(rf)
		}
	}
	//fmt.Printf("\n=======>setDefaultValue - End<=======================\n")
}

func main() {
	path := "/work/rvallabhu/gemini/middleware/cmd/topology-middleware/topostore/default_data"
	EmptyObjMap = make(map[string]interface{})
	EmptyObjMap["GeminiSpirentTest"] = models.GeminiSpirentTest{}
	EmptyObjMap["TopologyNetworkdevice"] = models.TopologyNetworkdevice{}
	EmptyObjMap["TopologyTopology"] = models.TopologyTopology{}
	EmptyObjMap["TopologyEntitiesEthernetIfGroup"] = models.TopologyEntitiesEthernetIfGroup{}
	EmptyObjMap["TopologyEntitiesEthernetifgroupEthernet"] = models.TopologyEntitiesEthernetifgroupEthernet{}
	EmptyObjMap["TopologyEntitiesMacPatternGroup"] = models.TopologyEntitiesMacPatternGroup{}
	EmptyObjMap["TopologyEntitiesBasepatterngroupCounter"] = models.TopologyEntitiesBasepatterngroupCounter{}

	fileList := []string{}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fileInfo, _ := os.Stat(path)
			if !fileInfo.Mode().IsDir() {
				fileList = append(fileList, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	DefaultValue = make(map[string]reflect.Value)
	for _, file := range fileList {
		//fmt.Println("Check -1")
		dType := strings.Split(filepath.Base(file), ".")[0]
		if obj, ok := EmptyObjMap[dType]; ok {
			//	fmt.Println("Check -2")
			fHandle, _ := os.Open(file)
			v := reflect.ValueOf(obj)
			i := reflect.New(v.Type()).Interface()
			byteValue, _ := ioutil.ReadAll(fHandle)
			if err := json.Unmarshal(byteValue, i); err != nil {
				return
			}
			DefaultValue[dType] = reflect.ValueOf(i)
			//strB, _ := json.Marshal(i)
			//	fmt.Printf("Default Obj:%s\n", strB)
		}
	}
	/*obj := models.GeminiSpirentTest{}
	SetDefaults(&obj)
	strB, _ := json.Marshal(obj)
	fmt.Printf("\n\nTest:%s\n\n", strB)*/
	fh, _ := os.Open("./topo.json")
	bytes, _ := ioutil.ReadAll(fh)
	var topo models.TopologyTopology
	json.Unmarshal(bytes, &topo)
	//st := TopologyEntitiesStepGroup{Step:"00:00:00:00:00:01", Type:models.TopologyEntitiesTypeEnumerationPORT}
	abc := ""
	topo.SpirentHosts[0].Ethernet.Mac.Counter.Start = abc
	topo.SpirentHosts[0].Ethernet.Mac.Counter.Step = &abc
	strB, _ := json.Marshal(topo.SpirentHosts[0].Ethernet.Mac.Counter)
	fmt.Printf("Counter Orig:%s\n", strB)
	SetDefaults(&topo)
	strB, _ = json.Marshal(topo.SpirentHosts[0].Ethernet.Mac.Counter)
	//strB, _ = json.Marshal(topo)
	fmt.Printf("Set Default Counter:%s\n", strB)

}
