# A repository for fun little Go thingiez

## concurrencyExmaple

Creates an infinite counting funciton 

```go
func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
```
"0 sheep, 1 sheep, 2 sheep..."
Calls it in main() and then calls a second one 

```go
func main() {
	count("sheep")
	count("fish")
}
```
Since the first one will never finish, the program will never count fish
UNLESS!!
We use a go routine to send the first counting function to the background so that both infinite loops run in tandem 

```go
func main() {
	go count("sheep")
	count("fish")
}
```
"0 sheep, 0 fish, 1 sheep, 1 fish, 2 sheep..."

Concurrency!


## firstGoes

My favorite thing in this little playground here is the one about pointers and de-referencing

```go

func main() {
	/*--------
	 POINTERS
	---------*/

	//initiates an integer, i=7
	//passes the memory address of i's value to inc
	i := 7
	inc(&i)
	fmt.Println(i)
}




//increment function that accepts a pointer to an int
func inc(x *int) {
	//de-references the pointer
	*x++
}
```
