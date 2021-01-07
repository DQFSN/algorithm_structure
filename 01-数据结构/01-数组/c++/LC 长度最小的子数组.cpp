#include<vector>

using namespace std;

class Solution {
public:
    int minSubArrayLen(int s, vector<int>& nums) {
        int l = 0;
        int r = 0;
        int min = nums.size() + 1;
        int sum = 0;
        while (r < nums.size() + 1) {
            if (sum < s) {
                if (r < nums.size()) {
                    sum += nums[r];
                }
                r++;
            }else {
                // cout<<l<<":"<<r<<"->"<<sum<<endl;
                if (r - l < min) {
                    min = r - l;
                }
                sum -= nums[l];
                l++;
            }
        }

        return min > nums.size()? 0 : min;

    }


    int minSubArrayLen2(int s, vector<int>& nums) {
        int l = 0;
        int r = 0;
        int min = nums.size() + 1;
        int sum = 0;
        while (l < nums.size()) {
            if (sum < s) {
                if (r == nums.size()) {
                    if (l == 0) {
                        return 0;
                    }else {
                        return min;
                    }
                }
                sum += nums[r];
                r++;
            }else {
                
                if (r - l < min) {
                    min = r - l;
                }
                // cout<<l<<":"<<r<<"->"<<sum<<" "<<min<<endl;

                sum -= nums[l];
                l++;
            }
        }

        return min > nums.size()? 0 : min;

    }
};