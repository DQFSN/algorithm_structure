#include<vector>

using namespace std;

class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int n = nums.size();
        int left = 0;
        int right = 0;

        while(right < n) {
            if (nums[right] != val) {
                swap(nums[left], nums[right]);
                left++;
            }
            right++;
        }
        return left;
    }
};