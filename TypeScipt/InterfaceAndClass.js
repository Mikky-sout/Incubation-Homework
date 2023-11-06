var Employee = /** @class */ (function () {
    function Employee(name, age, role) {
        this.name = name;
        this.age = age;
        this.role = role;
    }
    Employee.prototype.getDetail = function () {
        return "Name: " + this.name + " Age: " + this.age + " Role: " + this.role;
    };
    return Employee;
}());
var em1 = new Employee("Paritat", 24, "Back-end");
em1.role = "DevOps";
console.log(em1);
console.log(em1.getDetail());
