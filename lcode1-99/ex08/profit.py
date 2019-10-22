

def maxprofit(prices, fee):
    cash, hold = 0, -prices[0]
    for i in prices[1:len(prices)]:
        cash = max(cash, prices[i]+hold-fee)
        hold = max(hold, cash-prices[i])
    return cash

# print(maxprofit())