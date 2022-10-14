using System.Text;

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

        public static bool IsMatch(string s, string p)
        {
            var sb = new StringBuilder(s);
            var tokens = new List<string>();
            var cp = p.ToCharArray();
            for (var i = 0; i < cp.Length;)
            {
                if (i + 1 < cp.Length && cp[i + 1] == '*')
                {
                    tokens.Add(cp[i] + "" + cp[i + 1]);
                    i += 2;
                }
                else
                {
                    tokens.Add("" + cp[i]);
                    i++;
                }
            }
            var j = 0;
            foreach (var tok in tokens)
            {
                if (j >= sb.Length) return false;
                var star = tok.Length == 2 && tok[1] == '*' ? true : false;
                var cm = star ? "" + tok[0] : tok;
                if (!star)
                {
                    if (cm != "." && tok != sb.ToString(j, cm.Length))
                    {
                        return false;
                    }

                    j++;
                }
                else
                {
                    while (j < sb.Length && (cm == "." || sb[j] + "" == cm))
                    {
                        j++;
                    }
                }
            }

            if (j < sb.Length)
            {
                return false;
            }

            return true;
        }

        public static bool IsMatchDynamic(string s, string p)
        {
            var tokens = GetTokens(p);
            return IsMatchDynamicHelper(s, tokens, 0, 0);
        }
        private static bool IsMatchDynamicHelper(string s, List<string> tokens, int i, int j)
        {
            for (;i < s.Length || j<tokens.Count; i++)
            {
                if (j<tokens.Count)
                {
                    if(tokens[j].ElementAtOrDefault(1) ==  '*') {
                        if (i> s.Length || (tokens[j][0] != s.ElementAtOrDefault(i) && tokens[j][0] != '.')) return IsMatchDynamicHelper(s, tokens, i, j + 1);
                    } else {
                        if (tokens[j][0] == s.ElementAtOrDefault(i) || tokens[j][0] == '.') return IsMatchDynamicHelper(s, tokens, i + 1, j + 1);
                        else return false;
                    }
                } else {
                    return false;
                }
            }
            return true;
        }
        private static List<string> GetTokens(string p)
        {
            var tokens = new List<string>();
            var cp = p.ToCharArray();
            for (var i = 0; i < cp.Length;)
            {
                if (i + 1 < cp.Length && cp[i + 1] == '*')
                {
                    tokens.Add(cp[i] + "" + cp[i + 1]);
                    i += 2;
                }
                else
                {
                    tokens.Add("" + cp[i]);
                    i++;
                }
            }
            return tokens;
        }
    }
}