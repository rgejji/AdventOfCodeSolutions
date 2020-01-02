import sympy as sy

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


def readInput():
    inputStrSplit = inputStr.split("\n")
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
    M = sy.zeros(len(inputStrSplit),numItems)
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
#print((M.T).LUdecomposition())
print((M).rref())
