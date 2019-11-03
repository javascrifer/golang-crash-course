package main

import (
	"github.com/nikas-lebedenko/golang-crash-course/sectionchannels"
	"github.com/nikas-lebedenko/golang-crash-course/sectionconcurrency"
	"github.com/nikas-lebedenko/golang-crash-course/sectionfunctions"
	"github.com/nikas-lebedenko/golang-crash-course/sectionjson"
	"github.com/nikas-lebedenko/golang-crash-course/sectionmaps"
	"github.com/nikas-lebedenko/golang-crash-course/sectionpointers"
	"github.com/nikas-lebedenko/golang-crash-course/sectionslices"
	"github.com/nikas-lebedenko/golang-crash-course/sectionstructs"
)

func sectionSlices() {
	sectionslices.Example()
}

func sectionMaps() {
	sectionmaps.Example()
}

func sectionStructs() {
	sectionstructs.Example()
}

func sectionFunctions() {
	sectionfunctions.ExampleFunctions()
	sectionfunctions.ExampleMethods()
}

func sectionPointers() {
	sectionpointers.Example()
}

func sectionJson() {
	sectionjson.Example()
}

func sectionConcurrency() {
	sectionconcurrency.Example()
}

func sectionChannels() {
	sectionchannels.Example()
}

func main() {
	//sectionSlices()
	//sectionMaps()
	//sectionStructs()
	//sectionFunctions()
	//sectionPointers()
	//sectionJson()
	//sectionConcurrency()
	sectionChannels()
}
