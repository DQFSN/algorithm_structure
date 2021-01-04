#include <vector>
#include <string>

using namespace std;

class Solution {
public:
    bool isPalindrome(string s) {
        vector<char> chs;
        for (auto x : s) {
            if (x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z' || x >= '0' && x <= '9') {
                if (x >= 'A' && x <= 'Z') {
                    x = x - 'A' + 'a';
                }
                chs.push_back(x);
            }
        }

        
        int l = 0;
        int r = chs.size()-1;
        while (l < r) {
            if (chs[l] != chs[r]) {
                return false;
            }
            l++;
            r--;
        }
        return true;
    }
};