import pandas as pd
import os
import time
from datetime import datetime

start_time = time.time()

fileList = os.listdir('./data')

for fileName in fileList:
       filePath = "./data/" + fileName
       df = pd.read_json(filePath)
       df.to_csv('from_json_py.csv', mode='a', index=False, header=False)

now = datetime.now()
dt_string = now.strftime("%d/%m/%Y %H:%M:%S")
duration = "{:.7f}".format(time.time() - start_time)
print(dt_string + " %ss" % (duration))