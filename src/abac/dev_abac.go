package main

import ("fmt"
        "unicode"
        "strings"
        )

// Access Control Decision (ACD) function

// the return is (result, isSuccess)
func acd(policy string, att_map * map[string] []string) (bool , bool) {

    my_stack := &GenericStack{}
    parentheses_count := 0

    for i:=0 ; i < len(policy) && parentheses_count >= 0 ; i++ {

        fmt.Printf("(%d,%s)\n",i,string(policy[i]))
        if unicode.IsSpace(rune(policy[i])){
            continue
        }

        if policy[i] == '('{
            parentheses_count++
        }else if  policy[i] == ')' {

            parentheses_count--
            if parentheses_count < 0{
                println("Bad policy formatting. \n " +
                    "Policy should be of format (attr1=val1 & attr2=val2)\n" +
                    "There should matching ( and )")
                return false, false
            }

            for my_stack.stackSize() > 2  {
                right, ok1 := my_stack.pop().(bool)
                op, ok2 := my_stack.pop().(byte)
                left, ok3 := my_stack.pop().(bool)

                if ok1 && ok2 && ok3 {
                    result := evalBool(left, right, op)
                    my_stack.push(result)
                }else {
                    println("Bad policy formatting. Policy should be of format (attr1=val1 & attr2=val2)")
                    return false, false
                }

                if parentheses_count != 0{
                    break
                }

            }

        }else if policy[i] == '&' || policy[i] == '|' { // operator
            my_stack.push(policy[i])
        }else { // attributes

            j := i

            for i<len(policy) && policy[i]!= '(' && policy[i]!= ')' && policy[i]!= '&' && policy[i]!= '|' {
                i++
            }
            //attr_name_val := policy[j:i]
            att_name, att_val, success := splitAttributeNameValue(policy[j:i])

            if success{

                my_stack.push(attributeValExists(att_name, att_val, att_map))

            }else {
                fmt.Printf("Attribute value could not be splitted correctly %s\n",policy[j:i])
                return false, false
            }
            i--
        }
    }

    result, success := my_stack.pop().(bool)

    if parentheses_count == 0   && success && my_stack.stackSize() == 0{
        return result, true
    }else {
        if parentheses_count !=0 {
            fmt.Println("Bad formatting. Parenthesis mismatch!")
        }else if !success {
            fmt.Println("Bad formatting. Stack got empty")
        }else {
            fmt.Println("Bad formatting. More than one element left in the stack.")
        }
        return false, false
    }

}


func assertExample() {

}


func evalBool(left bool, right bool, opType byte ) bool{
    if opType == '&'{
        return left && right
    }else {
        return left || right
    }
}

func evalTNGate(t int, n int, arr []bool) bool{

    count:=0
    for _,elem := range arr {
        if elem{
            count++
        }
    }
    if t<=count {
        return true

    }
    return false
}

func attributeValExists(attr string, attr_val string, att_map * map[string] []string )  bool {

    if   values , found := (*att_map)[attr]; found {
        for _,value := range values  {
            if strings.TrimSpace(value) == attr_val{
                return true
            }
        }

    }
    return false
}

func splitAttributeNameValue(att_and_val string) (string, string, bool){
    name_val := strings.Split(att_and_val,"=")
    if len(name_val) != 2 {
        return "", "", false
    }

    return strings.TrimSpace(name_val[0]), strings.TrimSpace(name_val[1]), true
}

