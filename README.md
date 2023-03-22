# Go Interface

## Interface
Interfaces are considered separately from implementation and use.  
And the Go interface is also used to resolve "circular reference".  

### Circular reference
The Go language prohibits circular references to packages.  
A circular reference means that packages A and B refer to each other.  
The prohibition of circular references contributes to the realization of a layered architecture.  
Sometimes, the prohibition of circular references is too strict a restriction and becomes an obstacle to programming.  
Therefore, an interface may be used to separate the package from circular references and create a mechanism that allows circular references.  

## Structure
dog/  
&ensp;&ensp;Dog.go  
&ensp;&ensp;DogLoar.go  
cat/  
&ensp;&ensp;Cat.go  
&ensp;&ensp;DogInterface.go  
main.go  

## Point
Dog.go  
 
```Go
func (d *Dog) GetDog() string {  
  cat := &cat.Cat{ CatObj: NewDogRoar() }  
  // GetCat -> DogInterface -> GetDogRoar
  return cat.GetCat()   
} 
```
note : In the GetDog function, the GetCat function is called, but the result is "Bow Wow !". 

## Output Sample
/Go_Interface $ go build  
/Go_Interface $ ./Go_Interface   
Bow Wow !  
