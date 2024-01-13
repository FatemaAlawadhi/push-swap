## push-swap 

## Objectives
Push-Swap is a  project that uses a Non-Comparative Sorting Algorithm. Using a list of int values, two stacks (a and b) and a set of instructions.

It consist of two programs:
- push-swap, which calculates and displays on the standard output the smallest program using push-swap instruction language that sorts integer arguments received.
- checker, which takes integer arguments and reads instructions on the standard input. Once read, checker executes them and displays OK if integers are sorted. Otherwise, it will display KO.

Goal is to sort the stack a, that will contain the int values received, in ascending order, using both stacks and a set of instructions.

These are the instructions that ycan be used to sort the stack :
- pa push the top first element of stack b to stack a
- pb push the top first element of stack a to stack b
- sa swap first 2 elements of stack a
- sb swap first 2 elements of stack b
- ss execute sa and sb
- ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
- rb rotate stack b
- rr execute ra and rb
- rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
- rrb reverse rotate b
- rrr execute rra and rrb 

## Push-swap Program
- It will receive as an argument the stack a formatted as a list of integers. The first integer should be at the top of the stack.
- The program must display the smallest list of instructions possible to sort the stack a, with the smallest number being at the top.
- Instructions must be separated by a \n and nothing else.
- The goal is to sort the stack with the minimum possible number of operations.
- In case of error, It must display Error followed by a \n on the standard error. Errors are understood as: some arguments are not integers and/or there are duplicates.
- If there are no arguments, the program does not display anything (0 instructions).


## checker program
- It will get as an argument the stack a formatted as a list of integers. The first argument should be at the top of the stack. If no argument is given, checker stops and displays nothing.
- Checker will then read instructions on the standard input, each instruction will be followed by \n. Once all the instructions have been read, checker will execute them on the stack received as an argument.
- If after executing those instructions, stack a is actually sorted and b is empty, then checker must display "OK" followed by a \n on the standard output. In every other case, checker must display "KO" followed by a \n on the standard output.
- In case of error, it must display Error followed by a \n on the standard error. Errors include for example: some arguments are not integers, there are duplicates, an instruction do not exist and/or is incorrectly formatted.
- In case of there are no arguments, the program displays nothing.

## Usage
Follow the Steps below:
1. Write on the terminal:
```
go build -o push-swap
```
```
cd Checker1
```
```
go build -o checker
```
```
cd ..
```
2. Pull the checker program from Checker folder to the root folder (push-swap)

3. To execute Push-swap Program write on the Git Bash terminal:
```
./push-swap "4 67 3 87 2"
```

```
ARG="4 67 3 87 23"; ./push-swap "$ARG"
```

4. To execute Checker Program write on the Git Bash terminal:
```
echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
```
```
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
```


## AUTHORS

* Fatema Alawadhi
* Asma Ahmed


## LICENSES
This program developed within the scope of Reboot.