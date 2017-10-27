# The Monty Hall Problem

## The problem

You pick one of three doors at random in an attempt to win a car. A non-winning door is then opened, and you are given the choice of switching your selection.

## The Math

Your initial selection has 1/3 probability of being correct- the unchosen doors
have a 2/3 chance of being correct. When a non-winning door is opened, it is
opened in the set that has 2/3 of being correct. What this means is that the set
that has the 2/3 chance of being correct still holds that value, with only one
choice to be made. Switching is always the correct choice. [Examples](https://en.wikipedia.org/wiki/Monty_Hall_problem)

## Runs

This is a solution by simulation

```shell
  Stay Won 31 of 100
Switch Won 69 of 100
Rutime: 2.172895ms
```

```shell
  Stay Won 353 of 1000
Switch Won 666 of 1000
Rutime: 20.532108ms
```

```shell
  Stay Won 334195 of 1000000
Switch Won 666549 of 1000000
Rutime: 18.500569021s
```
