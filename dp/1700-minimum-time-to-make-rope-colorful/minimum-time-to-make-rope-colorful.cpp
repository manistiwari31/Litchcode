class Solution {
public:
    int minCost(string colors, vector<int>& neededTime) {
        int n = colors.size();
        vector<int> dp(n,-1);
        return calc(n-1,colors, neededTime, dp, 'A', neededTime[n-1]);
    }
int calc(int i, string& colorSequence, vector<int>& timeRequired, vector<int>& memo, char previousColor, int previousTime) {
    if (i<0) {
        return 0;
    }
    if (memo[i]!=-1) {
        return memo[i];
    }
     if (colorSequence[i] == previousColor) {
            return memo[i] = calc(i - 1, colorSequence, timeRequired, memo, colorSequence[i], max(timeRequired[i], previousTime)) + min(timeRequired[i], previousTime);
        } else {
            return memo[i] = calc(i - 1, colorSequence, timeRequired, memo, colorSequence[i], timeRequired[i]);
        }
}
};