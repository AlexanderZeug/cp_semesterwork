class Animal { // <1>
    private age: number; // <3>

    constructor(age: number) {
        this.age = age;
    }

    public setAge(age: number) { // <4>
        this.age = age;
    }

    public getAge(): number {
        return this.age;
    }
}

interface IMakeSound { // <5>
    makeSound(): void;
}

class Dog extends Animal implements IMakeSound{ // <2> <6>
    private colour: string;

    constructor(age: number, colour: string) {
        super(age);
        this.colour = colour;
    }

    public setColour(colour: string) {
        this.colour = colour;
    }

    public getColour(): string {
        return this.colour;
    }

    public makeSound(): void { // <7>
        console.log('Wuff');
    }
}

let laika = new Dog(13, "black"); // <8>
laika.makeSound();
console.log(laika.getAge());
console.log(laika.getColour());

/** 
 * output:
 *  Wuff
 *  13
 *  black
 */