package leetcode

/*
* 1431. 拥有最多糖果的孩子
* https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
 */

/*
* 官方每日一题
* 做题时间：2020/6/1
* 儿童节快乐
 */

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	max := 0

	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	for i := 0; i < len(candies); i++ {
		ans[i] = candies[i]+extraCandies >= max
	}

	return ans
}
