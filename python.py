def generateEquation(subArr1, subArr2):
    equation = ""
    if subArr2[1] == subArr1[1]:
        equation = f"y = {subArr1[1]}"
    elif subArr2[0] == subArr1[0]:
        equation = f"x = {subArr2[0]}"
    else:
        gradient = (subArr2[0] - subArr1[0]) / (subArr2[1] - subArr1[1])
        yIntercept = subArr2[1] - (gradient * subArr2[0])
        equation = f"y = {gradient}x + {yIntercept}"
    print(f"{equation} for A = {subArr1} and B = {subArr2}")
    return equation

def maxPoints(points):
    lLookUp = {}
    idx1 = 0
    idx2 = 1
    while idx1 < len(points) - 1:
        equation = generateEquation(points[idx1], points[idx2])
        if equation in lLookUp.keys():
            lLookUp[equation] += 0.5
        else:
            lLookUp[equation] = 1
        idx2 += 1
        if idx2 >= len(points):
            idx1 += 1
            idx2 = idx1 + 1
    highest = 0
    for k, v in lLookUp.items():
        if v > highest:
            highest = v
    return highest

if __name__ == "__main__":
    lArray = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
    print(maxPoints(lArray))
        