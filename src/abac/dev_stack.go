
package main


type GenericStack []interface{}

func (c *GenericStack) stackSize() int {
	return len(*c)
}

func (c *GenericStack) isEmpty() bool  {
	if len(*c)==0 {
		return true
	}
	return false
}

func (c *GenericStack) push(elem interface{}) {
	*c = append(*c, elem)
}

func (c *GenericStack) pop() interface{} {
	if (*c).isEmpty(){
		return nil
	}
	elem := (*c)[len(*c)-1]
	*c = (*c)[:len(*c)-1]
	return elem
}