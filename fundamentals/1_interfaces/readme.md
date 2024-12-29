# UseCase 1
1. Unknown type unmarshall from data, assign as interface {}
```
Error: json: cannot unmarshal object into Go struct field Data.info of type string
```
2. To enable polymorphism by allowing different types (rectangle, circle) to be treated as Shape