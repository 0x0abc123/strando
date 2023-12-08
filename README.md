# strando
Random string and phrase generator, when you need random values for IDs, usernames, hostnames, passwords, etc.

## usage:
```
strando [-c <complexity>] [-l <length>]
```
Length is ignored for complexity 0 (two-word phrase) and 4 (UUID).

Default length is 32 characters otherwise.

|Complexity|Range of Values|
|-|-|
|0|Generate a two word phrase "adjective-noun"|
|1|`[a-z0-9]`|
|2|`[a-zA-Z0-9]`|
|3|`[a-zA-Z0-9] and !@#$%^&*()-=_+:;{}[]\|,./?><~`|
|4|UUID version 4|

## examples:
```
$ strando
tq^W:Yusk2SdY{F]>gpl^1U:k5)$@3wC

$ strando -c 2
k8b2lLCbi01cKFzUkpqklW0QU2PTgmsF

$ strando -c 1
2107l0hsmo4hae03oh2v99bzhny0i8vy

$ strando -c 0
superhuman-resale

$ strando -c 3 -l 8
|x_3LCy<

$ strando -c 2 -l 16
odDYatWMMD9h3pZs

$ strando -c 4
b253d466-b5fb-4fd6-8e9d-8a613134209b
```

## build: 
```
GOOS=linux GOARCH=amd64 go build -o strando
```
