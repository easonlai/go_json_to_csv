import multiprocessing
import pandas as pd
import os
import time
from datetime import datetime

start_time = time.time()

cpuCount = 8
fileList = os.listdir('./data')
fileCount = len(fileList)
chunkSize = round(fileCount/8)

FinalDF = pd.DataFrame(columns = ['Sample1', 'Sample2', 'Sample3'])

def LoadDF(fileList):
    PartDF = pd.DataFrame(columns = ['Sample1', 'Sample2', 'Sample3'])
    for fileName in fileList:
        filePath = "./data/" + fileName
        df = pd.read_json(filePath)
        PartDF = PartDF.append(df)
    return PartDF

if __name__ == '__main__':
    pool = multiprocessing.Pool(processes=multiprocessing.cpu_count())
    FirstDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*0:chunkSize*1],))
    SecondDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*1:chunkSize*2],))
    ThirdDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*2:chunkSize*3],))
    ForthDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*3:chunkSize*4],))
    FifthDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*4:chunkSize*5],))
    SixDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*5:chunkSize*6],))
    SeventhDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*6:chunkSize*7],))
    EighthDFjob = pool.apply_async(LoadDF, (fileList[chunkSize*7:chunkSize*8],))
    FirstDF = FirstDFjob.get()
    SecondDF = SecondDFjob.get()
    ThirdDF = ThirdDFjob.get()
    ForthDF = ForthDFjob.get()
    FifthDF = FifthDFjob.get()
    SixDF = SixDFjob.get()
    SeventhDF = SeventhDFjob.get()
    EighthDF = EighthDFjob.get()
    FinalDF = FirstDF.append(SecondDF)
    FinalDF = FinalDF.append(ThirdDF)
    FinalDF = FinalDF.append(ForthDF)
    FinalDF = FinalDF.append(FifthDF)
    FinalDF = FinalDF.append(SixDF)
    FinalDF = FinalDF.append(SeventhDF)
    FinalDF = FinalDF.append(EighthDF)
    FinalDF.to_csv('from_json_py.csv', mode='a', index=False, header=False)

now = datetime.now()
dt_string = now.strftime("%d/%m/%Y %H:%M:%S")
duration = "{:.7f}".format(time.time() - start_time)
print(dt_string + " %ss" % (duration))