interface Person{
    name:string
    age:number
}

class Employee implements Person{
    name: string
    age: number
    role :string

    constructor(name:string,age:number,role:string) {
        this.name = name
        this.age = age
        this.role = role
    }

    getDetail(){
        return "Name: " + this.name + " Age: " + this.age+ " Role: "+this.role
    }
}

let em1 = new Employee("Paritat",24,"Back-end")
em1.role = "DevOps"
console.log(em1)
console.log(em1.getDetail())
