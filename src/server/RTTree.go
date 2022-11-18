package server

import (
	"fmt"
	"strings"
)

type RTLeaf struct{
	module IModule
	parent *RTLeaf
	leaves map[string]*RTLeaf
}

func GetNilLeaf() *RTLeaf{
	return &RTLeaf{nil,nil, map[string]*RTLeaf{}}
}

type RTTree struct{
	root *RTLeaf
}

func (t *RTTree)Init(){
	t.root = GetNilLeaf();
}

func (t *RTTree)Add(route string, module IModule) error {
	routes := strings.Split(strings.ToLower(route), "/")
	p := t.root
	for _, l := range routes{
		switch l {
			case "": continue
			case ".": continue
			case "..":
				if p.parent == nil {return fmt.Errorf("invalide path: %s", route)}
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
	if p.module != nil {return fmt.Errorf("module overlapping occurred at %s", route)}
	p.module = module
	return nil
}

func (t *RTTree)Run(route string, context *ModuleContext) error{
	routes := strings.Split(strings.ToLower(route), "/")
	p := t.root
	for _, l := range routes{
		switch l {
			case "": continue
			case ".": continue
			case "..":
				if p.parent == nil {return fmt.Errorf("invalide path: %s", route)}
					p = p.parent
					continue
		}
		if ch,ok := p.leaves[l]; ok{
			p=ch
		}else{
			break
		}
	}
	fallback := false
	context.Fallback = func(){
		fallback = true
	}
	for{
		if p==nil { // 404
			return fmt.Errorf("404 not found. Request URL: %s", route)
		}

		if p.module == nil {
			fallback = true
		}else{
			err:=p.module.Run(context)
			if err != nil {
				return err
			}
		}

		if fallback {
			fallback = false
			p = p.parent
		}else{
			break
		}
	}
	return nil
}