using System.Text;
using System.Collections;

namespace TwoSum
{
    public static class ChallengeSolutions
    {
        public static int[] TwoSum(int[] nums, int target)
        {
            var map = new Dictionary<int, int>();
            for (var i = 0; i < nums.Length; i++)
            {
                if (map.ContainsKey(target - nums[i]))
                {
                    return new int[] { i, map[target - nums[i]] };
                }

                map.TryAdd(nums[i], i);
            }

            return new int[] { 0, 1 };
        }
        //https://leetcode.com/problems/palindrome-number/
        public static string LongestPalindrome(string s)
        {
            Dictionary<char, int> memo = new();
            var sA = new StringBuilder(s);
            var longest = -1;
            var lp = "" + sA[0];
            var i = 0;
            while (i < sA.Length)
            {
                if (!memo.ContainsKey(sA[i]))
                {
                    memo.Add(sA[i], i);
                    i++;
                }
                else
                {
                    var (start, stop) = (memo[sA[i]], i);

                    var len = stop - start + 1;
                    if (len >= 4)
                    {
                        memo[sA[i]] = stop;
                        i++;
                        continue;
                    }
                    else if (len == 2)
                    {
                        var j = i + 1;
                        while (j < sA.Length && sA[i] == sA[j])
                        {
                            stop++;
                            j++;
                        }
                    }

                    memo[sA[i]] = stop;

                    while (start - 1 > -1 && stop + 1 < sA.Length && sA[start - 1] == sA[stop + 1])
                    {
                        start--;
                        stop++;
                    }

                    len = stop - start + 1;
                    if (longest < len)
                    {
                        longest = len;
                        lp = sA.ToString(start, len);
                    }

                    i++;
                }
            }

            return lp;
        }
        // https://leetcode.com/problems/regular-expression-matching/
        public static bool IsMatch(string s, string p) {
            if(p == null || p.Length == 0) return s.Length == 0 || s == null;

            var dp = new bool[s.Length +1, p.Length +1];
            dp[0,0] = true;
            for(var j = 2; j<=p.Length; j++) {
                dp[0,j] = p[j-1] == '*' && dp[0,j-2];
            }
            for(var j = 1; j<= p.Length; j++) {
                for(var i=1; i<=s.Length; i++) {
                    if(p[j-1] == s[i-1] || p[j-1] == '.') {
                        dp[i,j] = dp[i-1, j-1];
                    } else if(p[j-1] == '*') {
                        dp[i,j] = dp[i, j-2] || ((p[j-2] == s[i-1] || p[j-2] == '.') && dp[i-1, j]);
                    }
                }
            }

            return dp[s.Length, p.Length];
        }
        // https://leetcode.com/problems/generate-parentheses/
        public static IList<string> GenerateParenthesis(int n) {
                var dp = new List<string>[n + 1];
                dp[0] = new List<string>() {""};
                dp[1] = new List<string> {"()"};

                if(n < 2) return dp[n];

                for(var i =2; i<=n;i++) {
                    dp[i] = new List<string>();
                    for(var j=0;j<i; j++) {
                      foreach(var x in dp[j]) {
                        foreach(var y in dp[i-j-1]) {
                            dp[i].Add($"({x}){y}");
                        }
                      }  
                    }
                }
                Console.WriteLine(dp[n].Count);
                return dp[n];
        }
        public static IList<string> GenerateParenthesis2(int n) {
                var dp = new List<string>[n + 1];
                dp[0] = new List<string>() {""};
                dp[1] = new List<string> {"()"};

                if(n < 2) return dp[n];
                dp[2] = new List<string>() {"(())", "()()"};
                for(var i =3; i<=n;i++) {
                    var seen = new Dictionary<string, bool> ();
                    for(var j = 0; j<dp[i-1].Count; j++) {
                        seen.TryAdd($"({dp[i-1][j]})", true);
                        seen.TryAdd(dp[i-1][j]+"()", true);
                        seen.TryAdd("()"+dp[i-1][j], true);
                    }
                    for(var k=2; k<i; k++) {
                        for(var j = 0; j<dp[k].Count;j++) {
                            for(var l = 0; l< dp[i-k].Count;l++) {
                            seen.TryAdd(dp[i-k][l]+dp[k][j], true);
                            seen.TryAdd(dp[k][j]+dp[i-k][l], true);
                            }
                        }
                    }
                    dp[i] = seen.Select(x => x.Key).ToList();
                }
                Console.WriteLine(dp[n].Count);
                return dp[n];
        }
    }

}