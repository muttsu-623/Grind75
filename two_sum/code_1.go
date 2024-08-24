package 1_two_sum

func twoSum(nums []int, target int) []int {
    answers := make([]int, 0, 2)
    for i, a := range nums {
        for j := i + 1; j < len(nums); j++ {
            b := nums[j]
            if a + b == target {
                answers = append(answers, i, j)
                break
            }
            if len(answers) > 0 {
                break
            }
        }
    }

    return answers
}