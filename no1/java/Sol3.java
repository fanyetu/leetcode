public class Sol3 {

    public int[] twoSum(int[] nums, int target) {
        int[] result = new int[]{};
        for (int i = 0; i< nums.length; i++) {
            int first = nums[i];
            int second = target - first;

            for(int j= 0; j < nums.length; j ++) {
                if (i != j && nums[j] == second) {
                    result = new int[]{i, j};
                    return result;
                }
            }
        }
        return result;
    }
}
