# CLEAN CRAFTSMANSHIP

This repo is intended for executing all the exercises that appear on the book called _clean craftsmanship_. Also the idea is learning `go` along the process.

## 1. Queue

Implement a FIFO queue of integer following TDD.

### FIFO Principle

The first element added to the queue is the first one to be removed.

### Two Ends

A queue has a distinct front (or head) for removals and a rear (or tail) for additions.

### Operations

- Enqueue: Adds an element to the rear of the queue.
- Dequeue: Removes and returns the element from the front of the queue.
- Peek: Returns the element at the front without removing it.
- isEmpty: Checks if the queue contains any elements.
- Size: Returns the number of elements currently in the queue.

## 2. The Word Wrap Problem

Imagine you're writing text, but you have a limited width - like an old terminal screen or a narrow column in a newspaper. The word wrap problem is: given a string of text and a maximum line width, how do you break that text into multiple lines so that:

- No line exceeds the maximum width
- You break at sensible places (ideally between words, not in the middle of words)
- You don't waste too much space

A Simple Example
Let's say you have this text:

```text
"the quick brown fox jumps"

And your maximum width is **10 characters**. 

Without word wrap, this would go off the edge:

the quick brown fox jumps  (25 characters - too long!)


With word wrap, you'd break it into lines:

the quick   (9 characters - fits!)
brown fox   (9 characters - fits!)
jumps       (5 characters - fits!)
```
