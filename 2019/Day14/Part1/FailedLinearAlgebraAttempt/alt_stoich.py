#import sympy as sy
import numpy as num
#import scipy as scipy
from scipy.optimize import linprog


#The idea here is to use row reduction to solve for the value of ore needed to produce a fuel, unfortunately, this assumption operates on fractions of transactions

inputStr = """10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL"""


seenItems = {
        "ORE": 0,
        "FUEL": -1,
}


numEqns = 0

def readInput():
    inputStrSplit = inputStr.split("\n")
    global numEqns
    numEqns = len(inputStrSplit)
    print("numEqns is",numEqns)
    #Get number of items
    cnt = 1
    for row in inputStrSplit:
        sides = row.split(" => ")
        lhsParts = sides[0].split(", ")
        for part in lhsParts:
            element = part.split(" ")[1]
            cnt = checkItem(seenItems,  element, cnt)
        rhsItem = sides[1].split(" ")[1]
        cnt = checkItem(seenItems, rhsItem, cnt)
    numItems = len(seenItems)
    seenItems["FUEL"]=numItems-1
    print("Items", seenItems, numItems)


    #Make matrix Ore is first column, fuel is last column
    M = num.zeros((len(inputStrSplit),numItems), dtype=int)
    for i in range(len(inputStrSplit)):
        row = inputStrSplit[i]
        sides = row.split(" => ")
        lhsParts = sides[0].split(", ")
        for part in lhsParts:
            a, j = getValuesFromPart(part)
            M[i,j] = a  
        a,j = getValuesFromPart(sides[1])
        M[i,j] = -a
        cnt = checkItem(seenItems, rhsItem, cnt)
    return M


def checkItem(items, i, cnt):
    if i in items:
        return cnt
    items[i] = cnt
    return cnt+1


def getValuesFromPart(part):
    partSplit = part.split(" ")
    a = int(partSplit[0])
    j = seenItems[part.split(" ")[1]]
    return a,j

M = readInput()
print(M)
print()
b = num.zeros(len(seenItems),dtype=int)
#c = num.zeros((1,len(seenItems)),dtype=int)

print("numEqns is",numEqns)

c = num.zeros(numEqns,dtype=int)
c[0] = 1
bds = ((0,None),)*len(c)
res = linprog(c, A_ub=-M.transpose(), b_ub=b, bounds=bds, options={"disp": True})

#print((M.T).LUdecomposition())
