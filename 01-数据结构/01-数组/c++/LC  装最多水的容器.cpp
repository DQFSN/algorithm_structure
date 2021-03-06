#include<vector>

using namespace std;

class Solution {
public:
    int maxArea(vector<int>& height) {
        int max = 0;
        int l = 0;
        int r = height.size() - 1;
        while (l < r) {
            int w = r - l;
            int h = 0;
            if (height[l] < height[r]) {
                h = height[l];
                l++;
            }else {
                h = height[r];
                r--;
            }
            if (h * w > max) {
                max = h * w;
            }
        }

        return max;
    }
};