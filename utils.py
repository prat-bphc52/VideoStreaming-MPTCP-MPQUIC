import numpy as np

def encodeNumPyArray(fr):
    nor = len(fr).to_bytes(2,'big') # number of rows
    noc = len(fr[0]).to_bytes(2, 'big') # number of cols
    temp = np.array([nor[0],nor[1],noc[0],noc[1]], np.uint8)

    return np.concatenate((temp, fr), axis=None).tobytes()
