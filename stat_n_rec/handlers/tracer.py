import os
import itertools

class Model():
    def __init__(self, file, data):
        self.sums = []
        self.sets = dict()
        self.file = file
        # self.ress = dict()
        self.load(data)
        if file: self.open(file=self.file)

    
    def open(self, file=None, data=None):
        if file:
            data = open(file).readlines()
            self.sums = list(map(lambda x: int(x), data[0].strip().split(";")))
            for i in data[1:]:
                summ, d1, d2, d3, d4, res = i.split(";")
                if summ not in self.sums: self.sets[summ] = [{"data":[d1,d2,d3,d4], "res": res}]
                else: self.sets[summ].append({"data":[d1,d2,d3,d4], "res": res})

            # self.ress = {summ: res for summ, d1, d2, d3, d4, res in data[1:].split(";")}

    
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
                    result2[sum(j[:3])].append(j)
                except:
                    result2[sum(j[:3])] = [j]
        self.write(";".join([str(i) for i in result2.keys()]) + "\n", "w")
        c = 0
        stroke = ""
        for i in result2.keys():
            # a = [i]
            # a.extend(result2[i])
            # print(a)
            # self.write(";".join(list(map(lambda x: str(x), a))) + "\n", "a")
            for j in result2[i]:
                c += 1
                stroke+=str(i) + ";" + ";".join(list(map(lambda x: str(x), j))) + "\n"
                print(f"{c}/300400", "█" * round(c / 300400 * 20) + "▒" * (20 - round(c / 300400 * 20)), f"{round(c / 300400 * 100, 2)}%", end="\r", flush=True)
        self.write(stroke, "a")
        


def checker(arr, des_data):
    new_arr = dict()
    for i in arr:
        new_arr[i] = sum([1 for j in i if j in des_data["data"]])
    return max(new_arr, key=lambda key: new_arr[key])


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


# def configurator(data):
