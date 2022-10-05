using System.Collections.Generic;
using System.Linq;
using System;
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
    }
}