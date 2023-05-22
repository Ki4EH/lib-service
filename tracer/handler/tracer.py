import os
import itertools

class Model():
    def __init__(self, file, data):
        self.sums = []
        self.sets = dict()
        self.file = file
        self.load(data)
        if file: self.open(file=self.file)

    
    def open(self, file=None, data=None):
        if file:
            data = open(file).readlines()
            self.sums = list(map(lambda x: int(x), data[0].strip().split(";")))
            for i in data[1:]:
                summ, d1, d2, d3, d4, res = i.split(";")
                summ, d1, d2, d3, d4, res = int(summ), int(d1), int(d2), int(d3), int(d4), int(res)
                if summ not in self.sets: self.sets[summ] = [{"data":[d1,d2,d3,d4], "res": res}]
                else: self.sets[summ].append({"data":[d1,d2,d3,d4], "res": res})

    
    def search(self, des_data):
        data = self.sums[binary_search(self.sums, des_data["summ"])]
        stack = checker(self.sets[data], des_data)
        return stack
    

    def write(self, data, arg):
        with open(self.file, arg) as f:
            f.seek(0, os.SEEK_END)
            f.write(data)


    def load(self, data):
        result1 = dict()
        result2 = dict()
        for i in data:
            try:
                result1[i[1]].append(i[2])
            except:
                result1[i[1]] = [i[2]]
        for i in result1.keys():
            perm_set = list(itertools.permutations(result1[i],5))
            for j in perm_set:
                try:
                    result2[sum(j[:4])].append(j)
                except:
                    result2[sum(j[:4])] = [j]
        self.write(";".join([str(i) for i in sorted(result2.keys())]) + "\n", "w")
        c = 0
        stroke = ""
        for i in result2.keys():
            for j in result2[i]:
                c += 1
                stroke+=str(i) + ";" + ";".join(list(map(lambda x: str(x), j))) + "\n"
                print(f"{c}/300400", "█" * round(c / 300400 * 20) + "▒" * (20 - round(c / 300400 * 20)), f"{round(c / 300400 * 100, 2)}%", end="\r", flush=True)
        self.write(stroke, "a")


# def cleaner(array):


def checker(arr, des_data):
    new_arr = dict()
    for i in arr:
        try: new_arr[sum([1 for j in i["data"] if j in des_data["data"]])].append(arr.index(i))
        except: new_arr[sum([1 for j in i["data"] if j in des_data["data"]])] = [(arr.index(i))]
    returnable = new_arr[max(new_arr.keys())].copy()
    ret = [arr[i]["res"] for i in returnable]
    return ret[:5]


def binary_search(arr, x):
    lo, hi = 0, len(arr) - 1
    mid = 0
 
    while lo <= hi:
 
        mid = (hi + lo) // 2
 
        if arr[mid] < x:
            lo = mid + 1

        elif arr[mid] > x:
            hi = mid - 1

        else:
            return mid
    return mid
