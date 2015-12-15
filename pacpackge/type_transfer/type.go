package main

import "fmt"

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	//普通类型向接口类型的转换是隐式的。
	//接口类型向普通类型转换需要类型断言

	// 普通类型向接口类型的转换是隐式的,是编译期确定的
	var val interface{} = "hello"
	P(val)
	val = []byte{'a', 'b', 2}
	P(val)

	// 接口类型向普通类型转换有两种方式：Comma-ok断言和switch测试

	// Comma-ok断言的语法是：value, ok := element.(T)。element必须是接口类型的变量，T是普通类型。
	type Html []interface{}
	html := make(Html, 4)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = 33

	for _, element := range html {
		if value, ok := element.(string); ok {
			P(value, "is string")
		} else if value, ok := element.([]byte); ok {
			P(value, "is []byte")
		}
	}

	// switch 测试:
	for _, element := range html {
		switch value := element.(type) {
		case string:
			P(value, "is string")
		case []byte:
			P(value, "is []byte")
		default:
			P("Unkown type")
		}
	}
	// Comma-ok断言还支持另一种简化使用的方式：value := element.(T)。但这种方式不建议使用，因为一旦element.(T)断言失败，则会产生运行时错误。
	tem := html[0].(string)
	P(tem)
	// tem2 will panic, 因为断言失败
	tem2 := html[0].(int)
	P(tem2)
}

/*

func (m *Manager) getCountByCode(tr trace.T, code string) (*db.ViewResult, int, error) {
	// glog.Infoln("code: ", code)
	deviceCode, err := db.View(tr, m.profileDB.DB(), "_design/popularize", "get_device_code", map[string]interface{}{})
	inviteAndDeviceCode, err := db.View(tr, m.profileDB.DB(), "_design/popularize", "get_invitationCode_and_deviceCode", map[string]interface{}{"key": code})

	// viewResult, err := db.View(tr, m.profileDB.DB(), "_design/popularize", "by_code", map[string]interface{}{"key": code})
	if err != nil {
		return nil, 0, err
	}
	keyList := make([]string, 0)
	for _, row := range deviceCode.Rows {
		tem, ok := row.Key.(string)
		if !ok {
			glog.Error("key is supposed to be string type but not")
			continue
		}
		keyList = append(keyList, strings.TrimSpace(tem))
	}
	repeatList := findRepeat(keyList)
	// glog.Infoln("repeatList: ", repeatList, "loop in repeatList: ")
	// for _, v := range repeatList {
	// 	glog.Infoln(v)
	// }
	// repeatList = common.FilterStringSlice(keyList, func)

	count := len(inviteAndDeviceCode.Rows)
	for _, v := range inviteAndDeviceCode.Rows {
		temString := string([]byte(v.Value))
		// glog.Infoln("device_id: ", temString, "\t222: ", temString[1:len(temString)-1])
		// glog.Infoln("judge: ", common.StringInSlice(temString[1:len(temString)-1], repeatList))
		if common.StringInSlice(temString[1:len(temString)-1], repeatList) {
			count = count - 1
		}
	}

	return inviteAndDeviceCode, count, nil
}
func findRepeat(input []string) []string {
	output := make([]string, 0)
	// glog.Infoln("test: ", []string{"abc", "sdg"})
	tem := strings.Join(input, ",")
	for _, i := range input {
		t := strings.TrimSpace(i)
		// glog.Infoln("device_id: ", t)
		if strings.Count(tem, t) > 1 {
			if common.StringInSlice(t, output) {
				continue
			}
			output = append(output, t)
		}
	}
	return output
}
*/
/*

func countByCode(tracer *lu.Tracer, param *lu.Param, m *Manager) lu.Replyer {
	tr, done := tracer.Trace("user.countByCode")
	defer done()
	var code string
	if err := param.Required("code", &code).Error(); err != nil {
		return lu.JSON(M{"error": "param: code not set"})
	}
	viewResult, count, err := m.getCountByCode(tr, code)
	if err != nil || len(viewResult.Rows) == 0 {
		return lu.JSON(M{"error": fmt.Sprintf("code: %s not found", code)})
	}
	// test will be rm when merge
	glog.Infoln("loop in viewResult.Rows: ")
	for _, row := range viewResult.Rows {
		glog.Infoln("id: ", row.Id, "\tkey: ", row.Key)
		glog.Infoln("value: ", string([]byte(row.Value)))
		glog.Infoln("doc: ", string([]byte(row.Doc)))
	}

	// var count interface{}
	// json.Unmarshal(viewResult.Rows[0].Value, &count)
	return lu.JSON(M{"count": count})
}
*/

/*
	designDocPopularize := map[string]interface{}{
		"_id":      "_design/popularize",
		"language": "coffeescript",
		"views": map[string]interface{}{
			"get_device_code": map[string]interface{}{
				"map": `(doc) ->
	 return if doc.doc_type isnt 'Profile'
	 emit(doc.app_info.device_identifier,null)`,
				// "reduce": `_count`,
			},
			"get_invitationCode_and_deviceCode": map[string]interface{}{
				"map": `(doc) ->
	 return if doc.doc_type isnt 'Profile'
	 emit(doc.app_info.invitation_code,doc.app_info.device_identifier)`,
				// "reduce": `_count`,
			},
		},
	}
	// designDocPopularize := map[string]interface{}{
	// 	"_id":      "_design/popularize",
	// 	"language": "coffeescript",
	// 	"views": map[string]interface{}{
	// 		"by_code": map[string]interface{}{
	// 			"map": `(doc) ->
	//  return if doc.doc_type isnt 'Profile'
	//  emit(doc.app_info.device_identifier,null)`,
	// 			// "reduce": `_count`,
	// 		},
	// 	},
	// }


// glog.Infoln("loop in deviceTokens.Rows: ", deviceTokens.Rows, "Rows==nil: ", deviceTokens.Rows == nil, "len(device.Rows): ", len(deviceTokens.Rows))
	// for _, row := range deviceTokens.Rows {
	// 	glog.Infoln("id: ", row.Id, "\tkey: ", row.Key)
	// 	glog.Infoln("value: ", string([]byte(row.Value)))
	// 	glog.Infoln("doc: ", string([]byte(row.Doc)))
	// }
*/

/*

// logic:
// we want the num of the given code, but if a "device_identifier" show up in more than one profile,
// thrat means this device register more than one acount, that can be a cheat, so, we don't count
// this profile
// logic we deal with cheat as following:
// sep1:
// get a set of invitation_code-device_identifier pair (key=code,value="device_identifier") -- `deviceTokens` by using view "get_device_tokens"
// spe2:
// deal with this set, and find repeat device_identifiers (not include "")
// get the sum of repeat device_identifiers -- `numRepeat` and put no-repeat device_identifier to a list `noRepeatDevicetokens`
// sep3
// pass the `noRepeatDevicetokens` as couchdb.Options to view `filter_repeat_device`
// get of set of device_identifier-count -- `tokenCounts`
// deal with `tokenCounts` to find out if any count > 1 (that means this device registered more than one account) and put this device_identifier to `repeatTokens`

// then, the count will like this :
// `numCode := len(deviceTokens.Rows) - numRepeat - len(repeatTokens)`
func (m *Manager) getCountByCode(tr trace.T, code string) (int, error) {
	deviceTokens, err := db.View(tr, m.profileDB.DB(), "_design/popularize", "get_device_tokens", map[string]interface{}{"key": code})
	if err != nil {
		glog.Error(err)
		return 0, err
	}

	glog.Infoln("loop in deviceTokens.Rows: ", deviceTokens.Rows, "Rows==nil: ", deviceTokens.Rows == nil, "len(device.Rows): ", len(deviceTokens.Rows))
	for _, row := range deviceTokens.Rows {
		glog.Infoln("id: ", row.Id, "\tkey: ", row.Key)
		glog.Infoln("value: ", string([]byte(row.Value)))
		glog.Infoln("doc: ", string([]byte(row.Doc)))
	}

	noRepeatDevicetokens, numRepeat := getNoRepeatValue(deviceTokens)

	glog.Infof("noRepeatDeviceTokens: %s, num of repeat: %d length of noRepeat: %d", noRepeatDevicetokens, numRepeat, len(noRepeatDevicetokens))

	glog.V(2).Infof("noRepeatDeviceTokens: %s, num of repeat: %d", noRepeatDevicetokens, numRepeat)

	if len(noRepeatDevicetokens) == 0 {
		return len(deviceTokens.Rows), nil
	}
	tokenCounts, err := db.View(tr, m.profileDB.DB(), "_design/popularize", "filter_repeat_device", map[string]interface{}{"keys": noRepeatDevicetokens, "group": true})

	if err != nil {
		glog.Error(err)
		return 0, err
	}

	glog.Infoln("loop in tokenCounts.Rows: ", tokenCounts.Rows, "Rows==nil: ", tokenCounts.Rows == nil, "len(device.Rows): ", len(tokenCounts.Rows))
	for _, row := range tokenCounts.Rows {
		glog.Infoln("id: ", row.Id, "\tkey: ", row.Key)
		glog.Infoln("value: ", string([]byte(row.Value)))
		glog.Infoln("doc: ", string([]byte(row.Doc)))
	}

	repeatTokens, err := filterRepeatToken(tokenCounts)

	glog.Infoln("loop repeatTokens: ")
	for _, v := range repeatTokens {
		glog.Infoln("token: ", v.Key)
	}

	if err != nil {
		glog.Error(err)
		return 0, err
	}
	numCode := len(deviceTokens.Rows) - numRepeat - len(repeatTokens)
	return numCode, nil
}

// return the list of device_identifier if its's count > 1
func filterRepeatToken(input *db.ViewResult) ([]db.ViewResultRow, error) {
	var output []db.ViewResultRow
	for _, v := range input.Rows {
		var count int
		if err := json.Unmarshal(v.Value, &count); err != nil {
			glog.Error(err)
			return nil, err
		}
		if count > 1 {
			output = append(output, v)
		}
	}
	return output, nil

}

// get value=device_identifier,get a list of not-repeat value,and the sum of repeat device_identifier
func getNoRepeatValue(input *db.ViewResult) ([]string, int) {
	var output []string
	var temList []string
	var numRepeat int
	for _, v := range input.Rows {
		temString := strings.TrimSpace(string([]byte(v.Value)))
		temList = append(temList, temString)
	}
	builtString := strings.Join(temList, ",")
	for _, v := range temList {
		// v = v[1:len(v)-1]
		glog.Infoln("v: ", v)
		if v[1:len(v)-1] == "" || v == "null" {
			// if !common.StringInSlice("", output) {
			// 	output = append(output, "")
			// 	continue
			// }
			continue
		}
		if strings.Count(builtString, v) > 1 {

			glog.Infoln("repeat string :", v)

			numRepeat = numRepeat + 1
			continue
		}
		output = append(output, v[1:len(v)-1])
	}
	return output, numRepeat
}

// // test
	// glog.Infoln("loop in rawResult.Rows: ", rawResult.Rows, "\nlen(device.Rows): ", len(rawResult.Rows), "\noffset: ", rawResult.Offset, "\ntotalRow: ", rawResult.TotalRows)
	// for _, row := range rawResult.Rows {
	// 	glog.Infoln("id: ", row.Id, "\tkey: ", row.Key)
	// 	glog.Infoln("value: ", string([]byte(row.Value)))
	// 	glog.Infoln("doc: ", string([]byte(row.Doc)), "~~~~")
	// }
	// // end test

*/
