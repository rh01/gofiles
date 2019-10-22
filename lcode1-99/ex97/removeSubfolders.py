class Solution:
    def removeSubfolders(self, folder):
        res = ['1']
        for i in sorted(folder):
            if not i.startswith(res[-1] + "/"):
                res.append(i)
        return res[1:]

