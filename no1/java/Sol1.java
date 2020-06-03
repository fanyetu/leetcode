import java.util.HashMap;
import java.util.Map;

public class Sol1 {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int com = target - nums[i];
            if (map.containsKey(com) && map.get(com) != i) {
                return new int[]{map.get(com), i};
            }
            map.put(nums[i], i);
        }

        return new int[]{};
    }
}
