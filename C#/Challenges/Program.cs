// See https://aka.ms/new-console-template for more information

using System;
using System.Collections.Generic;
using Newtonsoft.Json;
using TwoSum;

// Console.WriteLine(JsonConvert.SerializeObject(ChallengeSolutions.TwoSum(new int[]{3,2,4}, 6)));
// Console.WriteLine(ChallengeSolutions.LongestPalindrome("aaabaaaa"));
// Console.WriteLine(ChallengeSolutions.LongestPalindrome("abacab"));
// Console.WriteLine(ChallengeSolutions.LongestPalindrome("abcdbbfcba"));
// Console.WriteLine(ChallengeSolutions.LongestPalindrome("ababababababa"));
Console.WriteLine(ChallengeSolutions.IsMatchDynamic("aabcadwdrgw", "a*b*"));
Console.WriteLine(ChallengeSolutions.IsMatchDynamic("mississippi", "mis*is*ip*."));
Console.WriteLine(ChallengeSolutions.IsMatchDynamic("ab", ".*c"));
Console.WriteLine(ChallengeSolutions.IsMatchDynamic("aaa", "a*a"));
