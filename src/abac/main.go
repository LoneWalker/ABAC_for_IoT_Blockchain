package main


import ("fmt"
    "encoding/json"
)


type Attribute struct {
    name string `json:"name"`
    values []interface{} `json:"values"`
}

func main() {

    //attr := map[string] []string{"designation":{"faculty", "chair"}, "dept":{"cs","eee"}}
    //acd("policy is made", &attr)

    attr_map := map[string][]string{"designation":{"faculty ", "chair"}, "dept":{"cs","eee"}, "security":{"clear","not_clear"}}
    //policy := "(((designation = chair & dept = cs))|(dept = civil & security = clear))|((designation = chair & dept = cs)|(dept = civil & security = clear))"
    //policy := "(designation = main_chair & dept = cs) | ((designation = chair & dept = eee) & (security = clear))"
    //policy := "(security = clear)"
    //policy := "(designation = chair & dept=eee & security = clear)"
    policy := "(designation = chair & dept=eee & security = clear)"

    result, success:= acd(policy, &attr_map)

    if success {
        if result {
            fmt.Println("Satisfied policy!!")
        }else {
            fmt.Println("Policy not satisfied!!")
        }
    }else {
        fmt.Println("Bad formating of policy")
    }

     /*
    var dat map[string]interface{}

    byt := []byte(`{"num":6.13,"strs":["a","b"], "nums":[12,13.10,14,15,"good"] }`)

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    num := dat["num"].(interface{})
    fmt.Println(num.(float64))

    strs := dat["strs"].([]interface{})

    fmt.Println(strs[0].(string))

    nums:= dat["nums"].([]interface{})

    fmt.Println(nums[4].(string))
    fmt.Println(nums[1].(float64))
    */
    testMarshalDataTable()
    testUnMarshalDataTable()
    testUnMarshalPolicy()

}


func testMarshalDataTable(){

    obj:= ObjectContainer{}
    obj.ObjectType =  TypeData
    //obj.Object = []Data{}

    //obj.Object = DataTable{}
    table := DataTable{}
    table.DataTableID = "tableid1232312"
    table.ReadPolicyID = "policyid26836273"
    table.WritePolicyID = "writepolicyid23827348"
    table.DeletePolicyID= "deletepolicyid6827343"

    map_data:= map[string][]byte{}
    map_data["item1"] = []byte(`{ "temp":123.4, "hum":34.6, "light":23.76, "infra": 62.90, "sound": 63.5, "gps": 62.8, "pressure":635, "altitude":524 }`)
    map_data["item2"] = []byte(`{ "temp":7623.4, "hum":34.6, "light":23.76, "infra": 62.90, "sound": 63.5, "gps": 62.8, "pressure":635, "altitude":555 }`)
    table.DataMap = map_data
    obj.Object = table

    data,_ := json.Marshal(obj)

    fmt.Println(string(data))
    /*
    if err:= json.Unmarshal(dat, &obj); err!=nil{
        fmt.Print(err)
    }

    fmt.Println(obj.ObjectType)

    var policy_interface = obj.Object.(interface{})

    var policy = policy_interface.(Policy)

    fmt.Println(policy.PolicyID)
    */
}
func testUnMarshalDataTable(){
    dat := []byte(`{"doc_type":3,"object":{"data_table_id":"tableid1232312","read_policy_id":"policyid26836273",
                    "write_policy_id":"writepolicyid23827348","delete_policy_id":"deletepolicyid6827343",
                    "data_map":{"item1":"eyAidGVtcCI6MTIzLjQsICJodW0iOjM0LjYsICJsaWdodCI6MjMuNzYsICJpbmZyYSI6IDYyLjkwLCAic291bmQiOiA2My41LCAiZ3BzIjogNjIuOCwgInByZXNzdXJlIjo2MzUsICJhbHRpdHVkZSI6NTI0IH0=",
                                "item2":"eyAidGVtcCI6NzYyMy40LCAiaHVtIjozNC42LCAibGlnaHQiOjIzLjc2LCAiaW5mcmEiOiA2Mi45MCwgInNvdW5kIjogNjMuNSwgImdwcyI6IDYyLjgsICJwcmVzc3VyZSI6NjM1LCAiYWx0aXR1ZGUiOjU1NSB9"}}}`)
    //dat:= []byte(`{"doc_type":1,"object":{"policy_id":"698712634b","payload":null}}`)
    obj:= ObjectContainer{}


    if err:= json.Unmarshal(dat, &obj); err!=nil{
        fmt.Print(err)
    }

    fmt.Println(obj.ObjectType)

    /***********if there is an interface inside an object, calling interface{} on that results in a map. for example, obj.Object.(interface{}) would convert
                the DataTable (assuming the obj.Object is a DataTable) into a map with its members (i.e. data_table_id, read_policy_id, write_policy_id, delete_policy_id, data_map )
                as the keys of that map

     */

    fmt.Println(obj.Object)

    // there are multiple types of objects as value in the map. So, the key would be string and the value type would be interface{}
    map_interface:= obj.Object.(map[string]interface{})

    data_table_id:=map_interface["data_table_id"]
    read_policy_id:=map_interface["read_policy_id"]
    fmt.Printf("data_table_id %s, read_policy_id: %s\n", data_table_id, read_policy_id)

    data_map:= map_interface["data_map"]

    fmt.Println(data_map)

    // there is []byte as value in the map. So, the key would be string and the value type would be []byte
    data_map_byte:= data_map.(map[string][]byte)

    fmt.Print(data_map_byte["item1"])

}


func testUnMarshalPolicy(){
    dat := []byte(`{"doc_type":1,"object": {"policy_id":"POL35426", "client_id":"628734iuwer", "policy":"(designation = chair & dept=eee & security = clear)", "meta_policy":"(designation = chair & dept=eee & security = clear)" }}`)
    //dat:= []byte(`{"doc_type":1,"object":{"policy_id":"698712634b","payload":null}}`)
    obj:= ObjectContainer{}

    if err:= json.Unmarshal(dat, &obj); err!=nil{
        fmt.Print(err)
    }

    fmt.Println(obj.ObjectType)

    //var policy_interface = obj.Object.(map[string]interface{})

    //var policy = policy_interface

    fmt.Println(obj.Object)
    //fmt.Println(policy_interface["policy"])

    map_interface:= obj.Object.(map[string]interface{})


    policy_id:=map_interface["policy_id"]
    policy:=map_interface["policy"]
    fmt.Printf("policy_id %s, policy: %s\n", policy_id, policy)

}


const (
    TypeClient=1
    TypePolicy=2
    TypeData=3
)

type ObjectContainer struct {
    ObjectType int `json:"doc_type"`
    Object interface{} `json:"object"`
    //Object map[string][]byte `json:"object"`
}

type FabClient struct {
    ClientID   string `json:"client_id"`
    ClientUsername   string `json:"client_username"`
    Org  string `json:"org"`
    SubjAttributes string `json:"subj_attributes"`
    ObjAttributes string `json:"obj_attributes"`
    EnvAttributes string `json:"env_attributes"`
    Policies []string `json:"policies"`
}

type Policy struct {
    PolicyID string `json:"policy_id"`
    ClientID string `json:"client_id"`
    PolicyRule string `json:"policy"`
    MetaPolicy string `json:"meta_policy"`
}


type DataTable struct {

    DataTableID string `json:"data_table_id"`
    ReadPolicyID string `json:"read_policy_id"`
    WritePolicyID string `json:"write_policy_id"`
    DeletePolicyID string `json:"delete_policy_id"`
    DataMap map[string][]byte `json:"data_map"`

}




