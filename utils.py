def encodeNumPyArray(fr):
    mybytes = bytearray()
    nor = len(fr).to_bytes(2,'big') # number of rows
    noc = len(fr[0]).to_bytes(2, 'big') # number of cols
    mybytes.append(int(nor[0]))
    mybytes.append(int(nor[1]))
    mybytes.append(int(noc[0]))
    mybytes.append(int(noc[1]))

    for i in fr:
        for j in i:
            mybytes.append(int(j[0]))
            mybytes.append(int(j[1]))
            mybytes.append(int(j[2]))
    return mybytes