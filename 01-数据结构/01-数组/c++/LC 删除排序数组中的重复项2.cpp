#include<vector>

using namespace std;

class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.size() < 1 ) return 0;
        int l = 0;
        int r = 0;
        int count = 0;

        while(r < nums.size()) {
            if (nums[r] != nums[l]) {
                swap(nums[++l],nums[r]);
                count = 1;
            }else {
                count++;
                if (count == 2) {
                    swap(nums[++l],nums[r]);
                }
            }
            r++;
        }
        return l+1;
    }
};