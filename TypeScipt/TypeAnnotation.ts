
// Number
let n:number = 10
var no:number = 50
no = 60
console.log(n)
console.log(no)


// String
let a:string = "Paritat"
var b:string = "Wilaikum"
console.log("My name is " + a+" " +b)


// Constant
const pi:number = 3.14
//pi = 3.15
let r:number = 2
console.log("Area ="+(pi*r*r))

// Any
let c:any = 1
c = "Eiei"
console.log(typeof(c), c)

// Boolean
let check:boolean = true
if (check){
    console.log("It's true")
}

// Array
let testArr :string[] = ["a","b","c"]
let arr = ["a",2,"c"]
for (let i=0;i<testArr.length;i++){
    console.log(testArr[i])
}