class Solution {
public:
    int knightDialer(int n) {
        const int MOD = 1000000007;
        vector<vector<int>> map = {{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9}, {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4}};
        vector<vector<int>> dp(10,vector<int>(n + 1, 0));

        // Initialize base case for n = 1
        for (int i = 0; i < 10; ++i) {
            dp[i][1] = 1;
        }

        // Build the solution bottom-up
        for (int step = 2; step <= n; ++step) {
            for (int i = 0; i < 10; ++i) {
                for (auto index : map[i]) {
                    dp[i][step] = (dp[i][step] + dp[index][step - 1]) % MOD;
                }
            }
        }

        // Calculate the total ways for n steps by summing up the possibilities for each starting digit
        int totalWays = 0;
        for (int i = 0; i < 10; ++i) {
            totalWays = (totalWays + dp[i][n]) % MOD;
        }

        return totalWays;
    }
};