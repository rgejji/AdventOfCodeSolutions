import numpy as np
#import scipy as scipy
#import sympy as sy



#inputStr = """12345678"""
#inputStr = """03036732577212944063491565474664"""
inputStr = """59708372326282850478374632294363143285591907230244898069506559289353324363446827480040836943068215774680673708005813752468017892971245448103168634442773462686566173338029941559688604621181240586891859988614902179556407022792661948523370366667688937217081165148397649462617248164167011250975576380324668693910824497627133242485090976104918375531998433324622853428842410855024093891994449937031688743195134239353469076295752542683739823044981442437538627404276327027998857400463920633633578266795454389967583600019852126383407785643022367809199144154166725123539386550399024919155708875622641704428963905767166129198009532884347151391845112189952083025"""

#Observe that the offset is large and the matrix is triangular.
#Specifically 1 triangular matrix with all 1s for j>=i
#We construct a new x starting from offset and ending at signalSize and multiple against the triangular matrix
def readInput():
    size = len(inputStr)
    x = np.zeros((size), dtype=int)
    for i in range(size):
        x[i] = int(inputStr[i])
    return x

pattern = [0,1,0,-1]
numReps = 10000
x = readInput()
L = len(x)
signalSize = numReps*L

offset = 0
for i in range(7):
    offset=offset*10+x[i]

print("Start:", x)
print("Size:",signalSize)
print("Offset", offset)


#Shift x based on offset
subX = np.zeros((signalSize-offset), dtype=int)
for i in range(len(subX)):
    subX[i] = x[(offset+i)%L]
x=subX
print("Shifted x:", x)
print("w/ length:", len(x))

#Go through each phase
numPhases = 100
for p in range(numPhases):
    y = np.zeros((len(x)), dtype=int)
    #Sum goes to i=0
    y[0] = sum(x)
    for i in range(1,len(x)):
        y[i]=y[i-1]-x[i-1]

    for i in range(len(x)):
        x[i]=abs(y[i])%10
    print("> p=",p)

print(x[:8])
