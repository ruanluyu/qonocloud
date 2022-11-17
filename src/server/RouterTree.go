package server

import (
	"errors"
	"fmt"
	"strings"
)

type RTLeaf struct{
	module *IModule
	parent *RTLeaf
	leaves map[string]*RTLeaf
}

func GetNilLeaf() *RTLeaf{
	return &RTLeaf{nil,nil, map[string]*RTLeaf{}}
}

type RTTree struct{
	root *RTLeaf
}

func (this *RTTree)Init(){
	this.root = GetNilLeaf();
}

func (this *RTTree)Add(route string, module IModule) error {
	routes := strings.Split(route, "/")
	p := this.root
	for _, l := range routes{
		switch l {
			case "": continue
			case ".": continue
			case "..":
				if p.parent == nil {return errors.New(fmt.Sprintf("Invalide path: %s", route))}
				p = p.parent
				continue
			default: 
				if child, ok:=p.leaves[l]; ok{
					p = child
					continue
				}else{
					newl := GetNilLeaf()
					newl.parent = p
					p.leaves[l] = newl
					p = newl
					continue
				}
		}
	}
	if p.module != nil {return errors.New(fmt.Sprintf("Module overlapping occurred at %s", route))}
	p.module = &module
	return nil
}

func (this *RTTree)Run(route string, context ModuleContext) error{
	routes := strings.Split(route, "/")
	p := this.root
	
	return nil
}