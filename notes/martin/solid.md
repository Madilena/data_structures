## polymorphism

- polymorphism is about an application of pointers to functions
- Object Oriented (OO) allows plugin artchitecture to be used anywhere for anything
- OO is the ability through polymorphism to gain absolute control over every source code dependency

## different types of programming

- structured programming: discipline upon direct transfer of control (SQL)
- object oriented programming: discipline upon indirect transfer of control
- functional programming: discipline upon variable assignment

# [SOLID] prinicples

[SOLID](https://blog.cleancoder.com/uncle-bob/2020/10/18/Solid-Relevance.html)

## single responsibility principle

- Gather together the things that change for the same reasons. Separate things that change for different reasons.

* each software module should have only one reason to change

## open-closed principle

- open for extension, closed for modification
- for software systems to be easy to change, they must be designed to allow the behavior to be changed by adding new code, rather than existing code

## liskov substitution principle

- This is about sub-typing. All implementations of interfaces are subtypes of an interface. A program that uses an interface must not be confused by an implementation of that interface.

* this helps us prevent model hierarchies that don't conform to the Open/Closed principle.
* Functions that use pointers or references to base classes must be able to use objects of derived classes without knowing it.

## interface segregation principle

- Keep interfaces small so that users do not end up depending on things they do not need.

## dependency inversion principle

- Depend in the direction of abstraction. High level modules should not depend upon low level details.
