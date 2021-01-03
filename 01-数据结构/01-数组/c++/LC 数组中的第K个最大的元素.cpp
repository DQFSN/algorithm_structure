#include <vector>
using namespace std;

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {

        return findKth(nums,0, nums.size()-1, k);

    }

    int findKth(vector<int>& nums, int l, int r, int k) {
        int n = partion(nums,l,r);
        if (n == nums.size() - k) {
            return nums[n];
        }else if (n < nums.size() - k){
            return findKth(nums,n+1,r,k);
        }else {
            return findKth(nums,l,n-1,k);
        }
    }

    int partion(vector<int>& nums, int l, int r) {
        int povt = nums[l];
        int povtIndex = l;
        int left = l;
        int right = r;

        while (l < r) {
            while (l < right && nums[l] <= povt) l++;
            while (left < r && nums[r] >= povt) r--;

            if (l < r) {
                int temp = nums[l];
                nums[l] = nums[r];
                nums[r] = temp;
                l++;
                r--;
            }
        }
        int temp = nums[r];
        nums[r] = nums[povtIndex];
        nums[povtIndex] = temp;
        // for (auto x : nums) {
        //     cout<<x<<" ";
        // }
        // cout<<"-> "<<povt<<endl;
        return r;
    }
};