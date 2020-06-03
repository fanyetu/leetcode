public class Main {

    public static void main(String[] args) {
        int[] nums = new int[]{2, 7, 11, 15};
        int target = 9;

//        Sol1 solution = new Sol1();
//        Sol2 solution = new Sol2();
        Sol3 solution = new Sol3();
        int[] result = solution.twoSum(nums, target);

        ArrayUtils.printArray(result);
    }
}
