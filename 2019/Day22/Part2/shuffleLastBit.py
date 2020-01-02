inputStr = """deal into new stack
cut 7990
deal into new stack
cut -5698
deal with increment 29
cut 1503
deal with increment 65
cut -9095
deal with increment 56
cut 9104
deal into new stack
deal with increment 5
cut -7708
deal with increment 20
cut 4813
deal with increment 2
cut 4728
deal into new stack
cut -5429
deal with increment 47
cut 1739
deal with increment 63
cut 6707
deal with increment 29
cut 4293
deal with increment 44
cut 8873
deal with increment 53
cut 6046
deal into new stack
cut 8054
deal into new stack
deal with increment 14
cut 2426
deal with increment 11
cut 4006
deal with increment 49
cut -6277
deal with increment 3
cut 2231
deal with increment 45
cut -5059
deal with increment 7
cut 4251
deal with increment 16
cut -6081
deal with increment 25
cut -4067
deal with increment 29
cut 7656
deal into new stack
cut 5091
deal with increment 57
deal into new stack
deal with increment 63
cut 4047
deal with increment 24
cut -8596
deal with increment 13
cut 1946
deal with increment 16
cut -1656
deal into new stack
deal with increment 15
cut -6557
deal with increment 10
cut 2378
deal with increment 24
cut -2162
deal with increment 7
deal into new stack
deal with increment 37
cut -4310
deal into new stack
deal with increment 48
cut 6842
deal with increment 13
cut 2960
deal into new stack
cut 7128
deal with increment 30
cut -2529
deal with increment 31
cut -2500
deal with increment 28
deal into new stack
deal with increment 37
cut -8133
deal with increment 74
cut -7823
deal with increment 42
cut 2092
deal with increment 41
cut -6752
deal with increment 56
cut -9577
deal into new stack
cut -4736
deal with increment 8
cut -3584"""

def newStackPos(pos, L):
    return L-1-pos

def cutPos(pos, N, L):
    while N>L:
        N-=L
    return (pos+L-N)%L

def dealPos(pos, inc,L):
    return (pos*inc)%L

def doShuffle(pos, L):
    for row in inputStr.split("\n"):
        vals = row.split(" ")
        instruction = vals[0]
        modifier = vals[-1]

        if instruction == "deal":
            if modifier == "stack":
                pos = newStackPos(pos,L)
            else:
                pos = dealPos(pos, int(modifier), L)

        if instruction == "cut":
            pos = cutPos(pos, int(modifier), L)
    return pos





L= 119315717514047
x=2020
posV=[]
for i in range(3):
    pos = x*i
    pos = doShuffle(pos, L)
    posV.append(pos)
    print("Have position", pos)



xinv = 108506422313517
p0 = 72616846730143
p1 = 25259071599808
p2 = 97217013983520
a=88872671823520
b=72616846730143
numIter=101741582076661
am1inv=93300017469100

test = doShuffle(doShuffle(x,L),L)

print(test)
print("Test: ", test-(a*(a*x+b)+b)%L)


#print(x*xinv%L)
#print((a-1)*am1inv%L)

#Answer is solve 2020 = x*a^n+b(a^n-1)/(a-1)
#2020-b(a^n-1)/(a-1)=  x a^n
aninv = 38001627100167
an = pow(a,numIter,L)
print("Check: aninv*an=", an*aninv%L)
print("an=", an)
val = (-2020+b*(an-1)*am1inv)%L
print("-anx = ", val)
x=(-aninv*val)%L
print("x=", x)

print("Check: ", (x*pow(a,numIter,L)+b*(pow(a,numIter,L)-1)*am1inv-2020)%L)



'''
numIter=15
newTest = x
for i in range(numIter):
    newTest = doShuffle(newTest,L)
print(newTest)

an = pow(a,numIter,L)
val = (an*x+b*(an-1)*am1inv)%L

print(val)
print("Done", x)
'''
