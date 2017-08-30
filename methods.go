// Go supports _methods_ defined on struct types.

package main

import "fmt"

type rect struct {
    width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
    return r.width * r.height
}


//you can call this method with instance created like this :  r := rect{width: 10, height: 5}

func area(r rect) int{
	 return r.width * r.height
}

//you can call this method with instance created like this :  r := &rect{width: 10, height: 5}

func area2(r *rect) int{
	 return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}


//you can call this method with instance created like this :  r := rect{width: 10, height: 5}
func  perim(r rect) int {
    return 2*r.width + 2*r.height
}


//you can call this method with instance created like this :  r := &rect{width: 10, height: 5}
func  perim2(r *rect) int {
    return 2*r.width + 2*r.height
}

func main() {
    r := &rect{width: 10, height: 5}

    // Here we call the 2 methods defined for our struct.
    fmt.Println("area: ", area2(r))
    fmt.Println("perim:", perim2(r))

    // Go automatically handles conversion between values
    // and pointers for method calls. You may want to use
    // a pointer receiver type to avoid copying on method
    // calls or to allow the method to mutate the
    // receiving struct.
    rp := r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
