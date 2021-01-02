#include<vector>

using namespace std;

class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] != 0) continue;
            for (int j = i+1; j < nums.size(); j++) {
                if (nums[j] == 0) continue;
                int temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
                break;
            }
        }
    }

    void moveZeroes2(vector<int>& nums) {
        int n = nums.size();
        int left = 0;
        int right = 0;

        while (right < n) {
            if (nums[right] != 0) {
                swap(nums[left], nums[right]);
                left++;
            }
            right++;
        }
    }
};