package model

type Community struct {
	Name string
}

type Communities []Community

type Person struct {
	Name        string
	Age         int
	Communities Communities
}

type Persons []Person
