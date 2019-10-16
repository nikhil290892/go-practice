package main

func main() {
	arrayintro()
	slice()         //slice allows you to group together values of same type
	slicerange()    //how to print the range using for loop
	sliceslice()    //use of colon operator to slice an array
	sliceappend()   //use of append() to append unlimited values of type to an array
	slicedel()      //use append to delete a slice from an array and need to use ... to append the end values after the sliced values
	slicemake()     // use make to define the current length and the maximum capacity the array can hold the number of elements
	slicemultidim() //add the multi-dimensional slice/array
}
