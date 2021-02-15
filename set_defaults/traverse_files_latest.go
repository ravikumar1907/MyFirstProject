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

        uu "github.com/google/uuid"

	models "github.com/SpirentOrion/gemini/middleware/go_models/models"
)


var DefaultValue map[string]reflect.Value
var EmptyObjMap map[string]interface{}

//GenerateUUID - generate uuid
func GenerateUUID() *string {
        uuid := uu.New()
        uuidString := uuid.String()
        return &uuidString
}

func SetDefaults(obj interface{}) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return
	}
	v := reflect.ValueOf(obj).Elem()
	if v.Kind() != reflect.Struct {
		return
	}
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
	if !(v.Kind() == reflect.Struct || v.Kind() == reflect.Ptr) {
		return
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	str := fmt.Sprintf("%s", v.Type())
	dType := strings.Split(str, ".")[1]
	d, ok := DefaultValue[dType]
	if !ok {
		return
	}
	t := v.Type()
	dv := reflect.Indirect(d)
	for i := 0; i < v.NumField(); i++ {
		rf := v.Field(i)
		df := dv.Field(i)
		switch rf.Kind() {
		case reflect.String:
			rf_name := t.Field(i).Name
			if rf_name == "ID" && rf.String() == "" {
				uuid := GenerateUUID()
				rf.Set(reflect.ValueOf(*uuid).Convert(rf.Type()))
			}
		case reflect.Ptr:
			if checkPrimitiveTypePtr(rf.Elem()) {
				rf.Set(df)
			} else {
				setDefaultValue(rf)
			}
		case reflect.Slice:
			for j := 0; j < rf.Len(); j++ {
				setDefaultValue(rf.Index(j))
			}
		case reflect.Struct:
			setDefaultValue(rf)
		}
	}
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
		dType := strings.Split(filepath.Base(file), ".")[0]
		if obj, ok := EmptyObjMap[dType]; ok {
			fHandle, _ := os.Open(file)
			v := reflect.ValueOf(obj)
			i := reflect.New(v.Type()).Interface()
			byteValue, _ := ioutil.ReadAll(fHandle)
			if err := json.Unmarshal(byteValue, i); err != nil {
				return
			}
			DefaultValue[dType] = reflect.ValueOf(i)
		}
	}
	fh, _ := os.Open("./topo.json")
	bytes, _ := ioutil.ReadAll(fh)
	var topo models.TopologyTopology
	json.Unmarshal(bytes, &topo)
	abc := ""
	topo.SpirentHosts[0].Ethernet.Mac.Counter.Step = &abc
	strB, _ := json.Marshal(topo.SpirentHosts[0].Ethernet.Mac.Counter)
	fmt.Printf("Counter Orig:%s\n", strB)
	SetDefaults(&topo)
	strB, _ = json.Marshal(topo.SpirentHosts[0].Ethernet.Mac.Counter)
	//strB, _ = json.Marshal(topo)
	fmt.Printf("Set Default Counter:%s\n", strB)

}
