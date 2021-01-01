#include<vector>
#include<map>
#include<iostream>

using namespace std;

class Solution {
public:

    int majorityElement(vector<int>& nums){

        map<int,int> dict;
        for (auto x : nums) {
            if (++dict[x] * 2 > nums.size()){
                return x;
            }
        }

        return -1;
    }

int majorityElement2(vector<int>& nums) {
        if (nums.size() < 1 ) {
            return -1;
        }

        int count = 1;
        int hero = nums[0];

        for (auto it = ++nums.begin(); it != nums.end(); it++) 
        {
            if (*it == hero) {
                count++;
            }else {
                if (count > 0) {
                    count--;
                }else {
                    hero = *it;
                    count++;
                }
            }
        }

        if (count) {
            count = 0;
            for (auto x : nums) {
                if (x == hero) count++;
            }

            count = count*2 > nums.size()?count:0;
        }
        // cout<<count<<"->"<<hero<<endl;

        return count?hero:-1;
    }
};


int main() {
    vector<int> input {1,2,3,4,5,6,7,7};

    int output = Solution().majorityElement(input);

    cout<<output<<endl;

    return 0;
}