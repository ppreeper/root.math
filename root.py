#!/usr/bin/env python

def root(x,p):
    #print(x,p)
    y = float(x)/float(p)
    #print(y)
    for i in range(32):
        y = (x/(y**(p-1))+y*(p-1))/p
        #print(y,i)
    return y

if __name__ == '__main__':
    print(root(5,2))
