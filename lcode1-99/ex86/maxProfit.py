class Solution(object):
    def maxProfit(self, k, prices):
        """
        :type k: int
        :type prices: List[int]
        :rtype: int
        """

        length = len(prices)
        dp = [[[-2**31 for _ in range(2)] for _ in range(k)]for _ in range(length)]
        # dp[i][k][0/1]
        
        # 第i天
        for i in range(0, length):
            # 交易次数
            for j in range(k):
                if i == 0:
                    dp[i][j][0] = 0
                    dp[i][j][1] = -prices[0]
                else:
                    if j == 0:
                    dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
                    dp[i][j][1] = max(dp[i][j][1], prices[i]+dp[i][j][0])

        return max(0, dp[length-1][k-1][1])
"""
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2:

输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""