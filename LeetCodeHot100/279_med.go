package LeetCodeHot100

func numSquares(n int) int {
    dp := make([]int,n+1)
    dp[0]=0
    for i:=1;i<=n;i++ {
        dp[i]=dp[i-1]+1
        for j:=1;j*j<=i;j++ {
            dp[i]=min(dp[i],dp[i-j*j]+1)
        }
    }
    return dp[n]
}
