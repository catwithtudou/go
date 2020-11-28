package rpc

import (
	"log"
	"reflect"
	"testing"
)

/**
 *@Author tudou
 *@Date 2020/11/28
 **/

type Aa struct{
	Name string
	Nt   string
}

func (a *Aa)te()  {
	log.Println("hello")
}

func (a *Aa)Te(){
	log.Println("world")
}

func (a *Aa)Fe(d int){
	log.Println("world")
}

func (a *Aa)Re(d int,c int){
	log.Println("know")
}


func TestRegister(t *testing.T) {
	fun(&Aa{Name: "name"})
}

func fun(r interface{}){
	typ := reflect.TypeOf(r)
	rcvr := reflect.ValueOf(r)
	sname := reflect.Indirect(rcvr).Type().Name()
	log.Println(typ)
	log.Println(rcvr)
	log.Println(sname)
	methodNum:=typ.NumMethod()
	log.Println(methodNum)
	for i:=0;i<methodNum;i++{
		method:=typ.Method(i)
		mtype := method.Type
		mname := method.Name
		pkgPath:= method.PkgPath
		param:=mtype.NumIn()
		log.Println(mtype)
		log.Println(mname)
		log.Println(pkgPath)
		log.Println(param)
		for j:=0;j<param;j++{
			log.Println(mtype.In(j))
		}
	}
}