# def kthPalindrome( queries, l):
#     base = 10 ** ((l - 1) / 2)
#     print(base)
#     res = [q - 1 + base for q in queries]
#     for i,a in enumerate(res):
#         a = str(a) + str(a)[-1 - l % 2::-1]
#         res[i] = int(a) if len(a) == l else -1
#     return res

# print(kthPalindrome([1,2,3,4,5,90], 3))