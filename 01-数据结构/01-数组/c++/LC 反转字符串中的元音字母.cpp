#include<vector>
#include<string>

using namespace std;

class Solution {
public:
    string reverseVowels(string s) {
        vector<char> chs;
        for (auto x : s) {
            if (isVowel(x)) {
                chs.push_back(x);
            }
        }

        string result = "";
        for (auto x : s) {
            if (isVowel(x)) {
                result += chs[chs.size()-1];
                chs.pop_back();
            }else {
                result += x;
            }
        }

        return result;
    }

    bool isVowel(char x) {
        if (x == 'a' || x == 'u' || x == 'o' || x == 'e' || x == 'i' 
        || x == 'A' || x == 'U' || x == 'O' || x == 'E' || x == 'I') {
            return true;
        }else {
            return false;
        }
    }
};







// 对撞指针，超出时间限制，最后两个用例没通过

class Solution {
public:
    string reverseVowels(string s) {
        vector<char> chs;
        for (auto x : s) {
            chs.push_back(x);
        }
       int l = 0;
       int r = chs.size()-1;
       while (l < r) {
           while (l < r && !isVowel(chs[l])) l++;
           while (l < r && !isVowel(chs[r])) r--;

           if (isVowel(chs[l]) && isVowel(chs[r])) {
               swap(chs[l],chs[r]);
               l++;
               r--;
           }
       }
        string result = "";
        for (auto x : chs) {
            result = result + x;
        }
       return result;
    }

    bool isVowel(char x) {
        if (x == 'a' || x == 'u' || x == 'o' || x == 'e' || x == 'i' 
        || x == 'A' || x == 'U' || x == 'O' || x == 'E' || x == 'I') {
            return true;
        }else {
            return false;
        } 
    }
};








// 对撞指针，拼接错误
class Solution {
public:
    string reverseVowels(string s) {
        if (s.size() < 2) {
            return s;
        }
        vector<char> chs;
        for (auto x : s) {
            chs.push_back(x);
        }
        string lStr = "";
        string rStr = "";
        int l = 0;
        int r = s.size()-1;
        while (l < r) {
            while (l < r && !isVowel(s[l])) {
                 lStr = lStr + s[l];
                 l++;
            }
            while (l < r && !isVowel(s[r])) {
                rStr = s[r] + rStr;
                r--;
            }

            if (isVowel(chs[l]) && isVowel(chs[r])) {
                rStr = s[l] + rStr;
                lStr = lStr + s[r];
               
                l++;
                r--;
            }

            if (l == r) {
                rStr = s[l] + rStr;
                break;
            }
            // cout<<lStr<<"->"<<rStr<<endl;
            // cout<<l<<"->"<<r<<endl;
        }
        // string result = "";
        // for (auto x : chs) {
        //     result = result + x;
        // }
        return lStr+rStr;
    }

    bool isVowel(char x) {
        if (x == 'a' || x == 'u' || x == 'o' || x == 'e' || x == 'i' 
        || x == 'A' || x == 'U' || x == 'O' || x == 'E' || x == 'I') {
            return true;
        }else {
            return false;
        } 
    }
};