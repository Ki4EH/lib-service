class Model():
    def __init__(self, data=None):
        self.sums = []
        self.sets = dict()
        # self.ress = dict()
        if data: self.open(file=data)

    
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
    

    def write(self, data):
        
        


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


    
    # def write(self, new_data):



# def tracing(dataset):
#     data = open("data.rat").readlines()
    