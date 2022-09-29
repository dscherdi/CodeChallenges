// See https://aka.ms/new-console-template for more information

using System;
using System.Collections.Generic;
using Newtonsoft.Json;

Console.WriteLine(JsonConvert.SerializeObject(TwoSum(new int[]{3,2,4}, 6)));

int[] TwoSum(int[] nums, int target)
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