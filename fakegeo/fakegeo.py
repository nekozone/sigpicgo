import json


import json

with open("pca.json") as f:
    pca = json.load(f)
res=[]
# print(pca)
for province,p in pca.items():
    for city,c in p.items():
        for area in c:
            res.append({"province":province,"city":city,"area":area})
num=len(res)
out={"num":num,"data":res}
outstr=json.dumps(out,ensure_ascii=False)
with open("fakegeo.json", "w") as target:
    target.write(outstr)
    pass
