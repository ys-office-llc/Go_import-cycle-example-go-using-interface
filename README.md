# Go import cycle example using interface

## Interface
Interfaces are considered separately from implementation and use.  
And the Go interface is also used to resolve "circular reference".  

### Circular reference
The Go language prohibits circular references to packages.  
A circular reference means that packages A and B refer to each other.  
The prohibition of circular references contributes to the realization of a layered architecture.  
Sometimes, the prohibition of circular references is too strict a restriction and becomes an obstacle to programming.  
Therefore, an interface may be used to separate the package from circular references and create a mechanism that allows circular references.

![image](https://user-images.githubusercontent.com/31363256/226831471-66451d96-7fcd-44e1-a429-253df5d3c408.png)

## Structure

```text
.
├── README.md
├── cat
│   ├── Cat.go
│   └── DogInterface.go
├── dog
│   ├── Dog.go
│   └── DogLoar.go
├── go.mod
└── main.go
```

## Point
Dog.go  
 
```Go
func (d *Dog) GetDog() string {
	cat := cat.Cat{DogIF: NewDogRoar()}
	// cat.Cat.GetCat.c.DogIF.GetDogRoar()
	return cat.GetCat()
}
```
note : In the GetDog function, the GetCat function is called, but the result is "Bow Wow !". 

## Code View

![image](https://user-images.githubusercontent.com/31363256/226886628-b01f0425-499c-4f74-a4bb-121bf9fed80f.png)

## Output Sample
```shell
go run main.go
```
```text   
Bow Wow !  
```

# Ref
- [Import Cycles in Golang: How To Deal With Them](https://jogendra.dev/import-cycles-in-golang-and-how-to-deal-with-them)
