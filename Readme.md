# Use Golang as an object-oriented language
## Encapsulation
* Not keyword driven
* Package scoped (Only protected when accessing from outside the package)
* Letter-case driven (the first letter of the name of fields, functions, and structs)
  - Lowercase first letters encapsulate, uppercase first letters export (available outside package)
* Applies to methods at the package level
* Applies to structs at the package level

## Composition
* Simple
* Allows sharing logic
* Allows sharing data
* Useful for common patterns
### How to implement composition?
* Build a struct
* Include it in another struct as an attribute (anonymous field)
* Methods of the first struct appear as though they on the composed struct
## Using composition together with anonymous fields to implement inheritance
