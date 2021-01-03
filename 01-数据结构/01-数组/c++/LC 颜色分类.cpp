#include <vector>
#include <map>

using namespace std;

class Solution {
public:
    void sortColors(vector<int>& nums) {
        map<int,int> dict;
        for (int i = 0; i < nums.size(); i++) {
            dict[nums[i]]++;
        }

        int c = 0;
        while (c < nums.size()) {
            if (c < dict[0]) {
                nums[c] = 0;
            }else if ((c - dict[0]) < dict[1]) {
                nums[c] = 1;
            }else {
                nums[c] = 2;
            }
            c++;
        }
    }
};